// Package routes has the routes to api
package routes

import (
	"net/http"
	controllers "price-comparator-api/adapters/api/controllers/product"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(app *gin.Engine, crtl *controllers.ProductController) {
  
}
