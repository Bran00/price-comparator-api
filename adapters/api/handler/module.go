// Package handler has the handlers to controllers
package handler

import (
	"price-comparator-api/adapters/api/handler/product"

	"go.uber.org/fx"
)

var Module = fx.Module("handler",
	fx.Provide(product.NewHTTPHandler),
)
