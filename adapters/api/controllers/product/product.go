package product

import (
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
	// TODO: implement controller
}
