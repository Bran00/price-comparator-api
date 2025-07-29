package ports

import "price-comparator-api/internal/searchengineer/domain"

type PriceComparatorRepositoryImplemetation interface {
	FindPrices(product domain.Product) ([]domain.Product, error)
}
