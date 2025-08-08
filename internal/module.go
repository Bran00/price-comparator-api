// Package internal has the logics from business
package internal

import (
	"price-comparator-api/internal/product/usecase"

	"go.uber.org/fx"
)

var Module = fx.Module(
	"internal",
	usecase.Module,
)
