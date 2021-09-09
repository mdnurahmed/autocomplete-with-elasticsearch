package repositories

import (
	"crypto/sha1"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

func buildInsertItem(searchString string) string {
	return fmt.Sprintf(`{"script" : {"source": "ctx._source.frequency++","lang": "painless"}, "upsert": {"search-text" : "%s","frequency":1}}`, searchString)
}

func getDocumentId(searchString string) string {
	document := fmt.Sprintf(`{"search-text" : "%s","frequency":1}`, searchString)
	h := sha1.New()
	h.Write([]byte(document))
	bs := h.Sum(nil)
	sEnc := b64.URLEncoding.EncodeToString(bs)
	return sEnc
}

func buildSearchQuery(searchString string) map[string]interface{} {
	query := map[string]interface{}{
		"size": 5,
		"sort": []map[string]interface{}{{"frequency": "desc"}},
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query": searchString,
				"type":  "bool_prefix",
				"fields": []string{
					"search-text",
					"search-text._2gram",
					"search-text._3gram",
				},
			},
		},
	}
	return query
}

func getMapping() string {
	return `{"settings":{"number_of_shards":1,"number_of_replicas":0},"mappings":{"properties":{"search-text":{"type":"search_as_you_type"},"frequency":{"type":"long"}}}}`
}

func getCuratedResult(response *esapi.Response) ([]string, error) {
	var result []string
	r := make(map[string]interface{})
	if err := json.NewDecoder(response.Body).Decode(&r); err != nil {
		return []string{}, err
	}

	for _, res := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		result = append(result, res.(map[string]interface{})["_source"].(map[string]interface{})["search-text"].(string))
	}
	return result, nil
}

func getDeleteAllDocumentByQuery() string {
	return `{"query":{"match_all": {}}}`
}
