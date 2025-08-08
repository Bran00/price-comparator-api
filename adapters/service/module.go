package service

import (
	"price-comparator-api/adapters/service/searcher"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		searcher.NewSearcherService,
	),
)
