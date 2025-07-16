
package ports

type PriceRepository interface {
	FindPrices(productName string) ([]domain.ProductPrice, error)
}

type PriceComparatorService interface {
	Compare(productName string) ([]domain.ProductPrice, error)
}
