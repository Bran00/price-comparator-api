package controllers

import (
	"log"
	"net/http"

	"price-comparator-api/adapters/transport/request"
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
  var isoReq request.RequestProductParams

  if err := ctx.ShouldBindQuery(&isoReq); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
  }

  log.Println(isoReq)

  //TODO: tem que verificar como ta chegando o valor aqui
	suggestions, err := c.usecaseProduct.SuggestionOfProduct(isoReq.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

  ctx.JSON(http.StatusOK, suggestions)
}
