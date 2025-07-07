package main

import (
	"fmt"
	"net/http"
	"price-comparator-api/internal/adapters/handler"
	"price-comparator-api/internal/adapters/repository"
	"price-comparator-api/internal/core/service"
)

func main() {
	repo := repository.NewSerpAPIRepository()
	priceComparatorService := service.NewPriceComparator(repo)
	httpHandler := handler.NewHTTPHandler(priceComparatorService)

	http.HandleFunc("/compare", httpHandler.Compare)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}