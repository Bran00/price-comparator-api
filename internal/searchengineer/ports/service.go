// Package ports has the interface implemetation
package ports

import "price-comparator-api/internal/searchengineer/domain"

type PriceComparatorRepositoryImplemetation interface {
	FindPrices(productName string) ([]domain.ProductPrice, error)
	Compare(productName string) ([]domain.ProductPrice, error)
}
