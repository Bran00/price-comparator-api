package main

import (
	"context"
	"fmt"
	"net/http"

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
		fx.Invoke(func(mux *http.ServeMux) {
			fmt.Println("Server is running on port 8080")
			http.ListenAndServe(":8080", mux)
		}),
	).Run()
}
