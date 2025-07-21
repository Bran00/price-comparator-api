package controllers

import (
	"price-comparator-api/adapters/api/controllers/product"

	"go.uber.org/fx"
)

var Module = fx.Module("controllers",
	fx.Provide(product.NewController),
)
