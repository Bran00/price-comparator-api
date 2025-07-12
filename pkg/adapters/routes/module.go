package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
  fx.Provide(

  ),
  fx.Invoke()
)

func startServer(r *gin.Engine) {
  r.Run(":8080")
  return nil
}

func newRouter() *gin.Engine {
  return gin.Default()
}
