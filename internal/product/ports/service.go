package ports

import "price-comparator-api/internal/product/domain"

type ProductHistoryRepositoryImplemetation interface {
	ProductHistory(product string) (domain.Product, error)
  SuggestionOfProduct(name string) (domain.Product, error)
}
