package routes

import (
	"net/http"

	"price-comparator-api/adapters/api/handler"
	"price-comparator-api/adapters/api/health"
)

func NewMux(h *handler.HTTPHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/compare", h.Compare)
	health.HealthRoutes(mux)
	return mux
}
