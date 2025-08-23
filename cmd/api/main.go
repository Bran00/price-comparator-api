package main

import (
	"price-comparator-api/adapters/api"
	"price-comparator-api/adapters/api/routes"
	"price-comparator-api/adapters/service"
	"price-comparator-api/internal"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		service.Module,
		internal.Module,
    api.Module,
	).Run()
}
