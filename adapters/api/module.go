package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Module("api", 
  fx.Provide(

  ),
  fx.Invoke(),
)

func startServer(lc fx.Lifecycle, router *gin.Engine) {
  server := &http.Server{
    Addr: fmt.Sprintf(":%s", 5000),
    Handler: router,
  }

  lc.Append(fx.Hook)
}
