
package ports

import "price-comparator-api/internal/core/domain"

type PriceRepository interface {
	FindPrices(productName string) ([]domain.ProductPrice, error)
}

type PriceComparatorService interface {
	Compare(productName string) ([]domain.ProductPrice, error)
}
