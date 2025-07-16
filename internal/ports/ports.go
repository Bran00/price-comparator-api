package ports

import "price-comparator-api/internal/domain"

type PriceComparator interface {
	FindPrices(productName string) ([]domain.ProductPrice, error)
	Compare(productName string) ([]domain.ProductPrice, error)
}