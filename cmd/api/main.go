package main

import (
	"fmt"
	"net/http"
	"price-comparator-api/adapters/api/controllers"
	"price-comparator-api/adapters/api/routes"
	"price-comparator-api/adapters/service"
	"price-comparator-api/internal"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		controllers.Module,
		routes.Module,
		service.Module,
		internal.Module,
	).Run()
}
