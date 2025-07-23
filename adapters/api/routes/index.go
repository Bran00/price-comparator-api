//Package routes has the routes to api
package routes

import (
	"net/http"

	"price-comparator-api/adapters/api/controllers/product"
)

func NewMux(c *product.Controller) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/compare", c.Compare)
	return mux
}
