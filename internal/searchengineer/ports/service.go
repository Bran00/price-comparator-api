package ports

import "price-comparator-api/internal/searchengineer/domain"

type PriceComparatorRepository interface {
	FindPrices(productName string) ([]domain.ProductPrice, error)
	Compare(productName string) ([]domain.ProductPrice, error)
}
