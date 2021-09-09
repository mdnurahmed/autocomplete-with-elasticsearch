package services

import (
	"autocomplete/app/repositories"
)

type IAutocompleteService interface {
	Search(searchStringstring string) (result []string, err error)
	Insert(searchStringstring string) (err error)
	Delete() (err error)
}
type AutocompleteService struct {
	elasticSearchRepository repositories.IElasticSearchRepository
}

func NewInstanceOfAutocompleteService(elasticSearchRepository repositories.IElasticSearchRepository) AutocompleteService {
	return AutocompleteService{elasticSearchRepository: elasticSearchRepository}
}

func (a *AutocompleteService) Search(searchString string) ([]string, error) {
	result, err := a.elasticSearchRepository.Search(searchString)
	if err != nil {
		return []string{}, err
	}
	return result, nil
}

func (a *AutocompleteService) Insert(searchString string) error {
	err := a.elasticSearchRepository.Insert(searchString)
	return err
}

func (a *AutocompleteService) Delete() error {
	err := a.elasticSearchRepository.Delete()
	return err
}
