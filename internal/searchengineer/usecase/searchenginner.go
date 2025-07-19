// Package service has the logic of business
package service

import (
	"price-comparator-api/internal/searchengineer/domain"
	"price-comparator-api/internal/searchengineer/ports"
)

type PriceComparator struct {
	repo ports.PriceComparatorRepositoryImplemetation
}

func NewPriceComparator(repo ports.PriceComparatorRepositoryImplemetation) *PriceComparator {
	return &PriceComparator{repo: repo}
}

func (s *PriceComparator) FindPrices(productName string) ([]domain.ProductPrice, error) {
   product := domain.Product{
    Name: productName, 
   }

   foundProduct, err := s.repo.FindPrices(product)
   if err != nil {
     return nil, err
   }

   return foundProduct, nil
}
