package routes

import (
	"net/http"

	"price-comparator-api/adapters/handler"
)

func NewMux(h *handler.HTTPHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/compare", h.Compare)
	return mux
}