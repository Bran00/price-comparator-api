package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	userGroup := router.Group("compare")
	userGroup.GET("/", fmt.Println("Hola"))
}
