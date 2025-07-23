// Package internal has the logics from business
package internal
import (
	service "price-comparator-api/internal/searchengineer/usecase"

	"go.uber.org/fx"
)

var Module = fx.Module(
	"internal",
	fx.Provide(service.NewPriceComparator),
)
