package usecase

import "price-comparator-api/internal/product/ports"

type ProductUseCase struct {
   repository ports.ProductHistoryRepositoryImplemetation
}

func NewProductUseCase(repository ports.ProductHistoryRepositoryImplemetation) *ProductUseCase {
  return &ProductUseCase{repository}
}

func (s *PriceComparator) HistoryProduct(productName string) (do, error) {
   product := domain.Product{
    Name: productName, 
   }

   foundProduct, err := s.repo.FindPrices(product)
   if err != nil {
     return nil, err
   }

   return foundProduct, nil
}
