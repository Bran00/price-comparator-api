// Package product has the connection to internal
package product

import (
	"encoding/json"
	"net/http"

	service "price-comparator-api/internal/searchengineer/usecase"
)

type HTTPHandler struct {
	service service.PriceComparator
}

func NewHTTPHandler(service service.PriceComparator) *HTTPHandler {
	return &HTTPHandler{service: service}
}

func (h *HTTPHandler) HistoryPriceProduct(w http.ResponseWriter, r *http.Request) {
	product := r.URL.Query().Get("product")
	if product == "" {
		http.Error(w, "Product parameter is missing", http.StatusBadRequest)
		return
	}

	prices, err := h.service.FindPrices(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prices)
}
