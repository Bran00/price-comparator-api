// Package product has the logic of business
package product

import (
	"net/http"

	"price-comparator-api/adapters/api/handler/product"
)

type ProductController struct {
	httpHandler *product.HTTPHandler
}

func NewController(httpHandler *product.HTTPHandler) *Controller {
	return &Controller{httpHandler: httpHandler}
}

func (c *Controller) History(w http.ResponseWriter, r *http.Request) {
	c.httpHandler.Compare(w, r)
}
