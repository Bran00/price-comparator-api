package main

import (
	"price-comparator-api/adapters/handler"
	"price-comparator-api/adapters/repository"
	"price-comparator-api/adapters/routes"
	"price-comparator-api/internal/service"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		handler.Module,
		repository.Module,
		routes.Module,
		service.Module,
	).Run()
}
