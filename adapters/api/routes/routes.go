package routes

import (
	controllers "price-comparator-api/adapters/api/controllers/product"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(app *gin.Engine, crtl *controllers.ProductController) {
  getProductRouter := app.Group("api/v1")
  getProductRouter.GET("/getSuggestionOfProducts", crtl.SuggestionOfProducts)
  getProductRouter.GET("/getProduct", crtl.GetProductsNearest)
}
