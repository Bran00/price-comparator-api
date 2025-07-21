package routes

import (
	"net/http"

	"price-comparator-api/adapters/api/controllers/product"
	"price-comparator-api/adapters/api/health"
)

func NewMux(c *product.Controller) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/compare", c.Compare)
	health.HealthRoutes(mux)
	return mux
}
