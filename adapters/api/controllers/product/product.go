package controllers

import (
	"net/http"

	"price-comparator-api/internal/product/usecase"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	client            *http.Client
  usecaseProduct    *usecase.Product
}

func NewProductController(usecase *usecase.Product) *ProductController {
	return &ProductController{
    client:         http.DefaultClient,
		usecaseProduct: usecase,
	}
}

func (c *ProductController) HistoryProduct(w http.ResponseWriter, r *http.Request) {
	// TODO: implement controller
}

func (c *ProductController) SuggestionOfProducts(ctx *gin.Context) {
  var isoReq request.RequestIsochroneParams
	err := validate.BindQuery(ctx.Request.URL.Query(), &isoReq) 
}
