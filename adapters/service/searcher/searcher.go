package searcher

import (
	"price-comparator-api/internal/product/domain"
	"price-comparator-api/internal/product/ports"
)

type Service struct{}

func NewSearcherService() ports.ProductHistoryRepositoryImplemetation {
	return &Service{}
}

func (s *Service) ProductHistory(product string) (domain.Product, error) {
	// TODO: implement searcher
	return domain.Product{}, nil
}