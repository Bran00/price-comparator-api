package main

import (
	"fmt"
	"net/http"
	"price-comparator-api/adapters/api/controllers"
	"price-comparator-api/adapters/api/handler"
	"price-comparator-api/adapters/repository"
	"price-comparator-api/adapters/api/routes"
	"price-comparator-api/internal"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		controllers.Module,
		handler.Module,
		repository.Module,
		routes.Module,
		internal.Module,
		fx.Invoke(StartServe),
	).Run()
}

func StartServe(mux *http.ServeMux) {
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", mux)
}
