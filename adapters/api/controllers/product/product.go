package product

import (
	"net/http"

	"price-comparator-api/adapters/api/handler/product"
)

type Controller struct {
	httpHandler *product.HTTPHandler
}

func NewController(httpHandler *product.HTTPHandler) *Controller {
	return &Controller{httpHandler: httpHandler}
}

func (c *Controller) Compare(w http.ResponseWriter, r *http.Request) {
	c.httpHandler.Compare(w, r)
}
