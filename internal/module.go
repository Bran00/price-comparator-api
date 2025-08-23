package internal

import (
	"price-comparator-api/internal/product/usecase"

	"go.uber.org/fx"
)

var Module = fx.Module("internal",
  fx.Provide(usecase.NewProductUsecase),
)
