package routes

import (
	"net/http"

	"price-comparator-api/adapters/api/handler/product"
	"price-comparator-api/adapters/api/health"
)

func NewMux(h *product.HTTPHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/compare", h.Compare)
	health.HealthRoutes(mux)
	return mux
}
