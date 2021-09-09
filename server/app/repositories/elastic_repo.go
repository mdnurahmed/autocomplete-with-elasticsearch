package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	log "github.com/sirupsen/logrus"
)

type IElasticSearchRepository interface {
	Search(key string) (result []string, err error)
	Insert(searchString string) (err error)
	Delete() (err error)
	Bootstrap() error
}

type ElasticSearchRepository struct {
	es        *elasticsearch.Client
	Address   string
	IndexName string
	Size      int64
	Refresh   string
}

func (e *ElasticSearchRepository) Bootstrap() error {
	mapping := getMapping()
	createIndexRequest := esapi.IndicesCreateRequest{
		Index: e.IndexName,
		Body:  strings.NewReader(mapping),
	}
	_, err := createIndexRequest.Do(context.Background(), e.es)
	return err
}

func (e *ElasticSearchRepository) Insert(searchString string) error {
	insertItem := buildInsertItem(searchString)
	id := getDocumentId(searchString)
	updateRequest := esapi.UpdateRequest{
		Index:      e.IndexName,
		Body:       strings.NewReader(insertItem),
		DocumentID: id,
		Refresh:    e.Refresh,
	}
	_, err := updateRequest.Do(context.Background(), e.es)
	return err
}

func (e *ElasticSearchRepository) Search(searchString string) ([]string, error) {
	searchQuery := buildSearchQuery(searchString)
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(searchQuery); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	// Perform the search request.
	res, err := e.es.Search(
		e.es.Search.WithContext(context.Background()),
		e.es.Search.WithIndex(e.IndexName),
		e.es.Search.WithBody(&buf),
		e.es.Search.WithTrackTotalHits(true),
		e.es.Search.WithPretty(),
	)
	defer res.Body.Close()
	curatedResult := []string{}
	if err == nil && !res.IsError() {
		var errCurated error
		curatedResult, errCurated = getCuratedResult(res)
		if errCurated != nil {
			return []string{}, errCurated
		}
	}
	return curatedResult, err
}

func (e *ElasticSearchRepository) Delete() error {
	deleteAllDocumentByQuery := getDeleteAllDocumentByQuery()
	deleteRequest := esapi.DeleteByQueryRequest{
		Index: []string{e.IndexName},
		Body:  strings.NewReader(deleteAllDocumentByQuery),
	}
	_, err := deleteRequest.Do(context.Background(), e.es)
	return err
}

func NewInstanceOfElasticSearchRepository(
	address, indexName, refresh string, size int64,
) ElasticSearchRepository {

	cfg := elasticsearch.Config{
		Addresses: []string{
			address,
		},
	}

	// Instantiate a new Elasticsearch client object instance
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.WithFields(log.Fields{
			"erro_message": err.Error(),
		}).Fatal("Couldn't initialize elasticsearch client")
	}
	return ElasticSearchRepository{
		es:        es,
		Address:   address,
		IndexName: indexName,
		Size:      size,
		Refresh:   refresh,
	}
}
