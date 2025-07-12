package main

import (
	"context"
	"fmt"
	"net/http"

	"price-comparator-api/internal/domain"
	"price-comparator-api/internal/ports"
	"price-comparator-api/internal/service"
	"price-comparator-api/adapters/handler"
	"price-comparator-api/adapters/repository"
	"price-comparator-api/adapters/routes"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			repository.NewSerpAPIRepository,
			service.NewPriceComparator,
			handler.NewHTTPHandler,
			routes.NewMux,
		),
		fx.Invoke(func(mux *http.ServeMux) {
			fmt.Println("Server is running on port 8080")
			http.ListenAndServe(":8080", mux)
		}),
	).Run()
}
