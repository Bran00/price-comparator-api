
package handler

import (
	"encoding/json"
	"net/http"
	"price-comparator-api/internal/ports"
)

type HTTPHandler struct {
	service ports.PriceComparatorService
}

func NewHTTPHandler(service ports.PriceComparatorService) *HTTPHandler {
	return &HTTPHandler{service: service}
}

func (h *HTTPHandler) Compare(w http.ResponseWriter, r *http.Request) {
	product := r.URL.Query().Get("product")
	if product == "" {
		http.Error(w, "Product parameter is missing", http.StatusBadRequest)
		return
	}

	prices, err := h.service.Compare(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prices)
}
