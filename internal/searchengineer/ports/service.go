// Package ports has the interface implemetation
package ports

import "price-comparator-api/internal/searchengineer/domain"

type PriceComparatorRepositoryImplemetation interface {
	FindPrices(product domain.Product) ([]domain.ProductPrice, error)
}
