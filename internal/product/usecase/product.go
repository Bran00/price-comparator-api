package usecase

import (
	"price-comparator-api/internal/product/domain"
	"price-comparator-api/internal/product/ports"
)

type Product struct {
	repository ports.ProductHistoryRepositoryImplemetation
}

func NewProductUsecase(repository ports.ProductHistoryRepositoryImplemetation) *Product {
	return &Product{
		repository: repository,
	}
}

func (p *Product) ProductHistory(product string) (domain.Product, error) {
	return p.repository.ProductHistory(product)
}