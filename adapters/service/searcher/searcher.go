package searcher

import (
	"net/http"
	"price-comparator-api/internal/product/domain"
	"price-comparator-api/internal/product/ports"
	zoom "price-comparator-api/pkg/services/zoom/product"
)

type Service struct {
	client *http.Client
}

func NewSearcherService() ports.ProductHistoryRepositoryImplemetation {
	return &Service{
		client: &http.Client{},
	}
}

func (s *Service) ProductHistory(product string) (domain.Product, error) {
	// TODO: implement searcher
	return domain.Product{}, nil
}

func (s *Service) SuggestionOfProduct(product string) (domain.Suggestions, error) {
	suggestions, err := zoom.SuggestionsOfProducts(product, s.client)
	if err != nil {
		return domain.Suggestions{}, err
	}

	return domain.Suggestions{
		Products: suggestions.Query,
	}, nil
}