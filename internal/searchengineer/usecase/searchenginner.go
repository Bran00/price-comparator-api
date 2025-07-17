package service

import (
	"price-comparator-api/internal/searchengineer/domain"
	"price-comparator-api/internal/searchengineer/ports"
)

type PriceComparator struct {
	repo ports.PriceComparatorRepository
}

func NewPriceComparator(repo ports.PriceComparatorRepository) *PriceComparator {
	return &PriceComparator{repo: repo}
}

func (s *PriceComparator) Compare(productName string) ([]domain.ProductPrice, error) {
	return s.repo.FindPrices(productName)
}
