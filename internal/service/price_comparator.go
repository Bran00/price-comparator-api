package service

import (
	"price-comparator-api/internal/domain"
	"price-comparator-api/internal/ports"
)

type PriceComparator struct {
	repo ports.PriceRepository
}

func NewPriceComparator(repo ports.PriceRepository) *PriceComparator {
	return &PriceComparator{repo: repo}
}

func (s *PriceComparator) Compare(productName string) ([]domain.ProductPrice, error) {
	return s.repo.FindPrices(productName)
}
