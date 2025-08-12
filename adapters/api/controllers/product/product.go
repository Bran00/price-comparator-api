package product

import (
	"encoding/json"
	"net/http"
	"price-comparator-api/internal/product/usecase"
)

type Controller struct {
	usecase *usecase.Product
}

func NewProductController(usecase *usecase.Product) *Controller {
	return &Controller{
		usecase: usecase,
	}
}

func (c *Controller) HistoryProduct(w http.ResponseWriter, r *http.Request) {
	// TODO: implement controller
}

func (c *Controller) SuggestionOfProducts(w http.ResponseWriter, r *http.Request) {
	product := r.URL.Query().Get("product")
	if product == "" {
    http.Error(w, "product is required", http.StatusBadRequest)
		return
	}

	suggestions, err := c.usecase.SuggestionOfProduct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(suggestions)
}
