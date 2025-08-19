package api

import (
	"context"
	"fmt"
	"log"
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

  lc.Append(fx.Hook{
    OnStart: func(ctx context.Context) error {
      go func() {
          err := server.ListenAndServe()
          if err != nil && err != http.ErrServerClosed{
            log.Fatal("error starting server")
          }
      } ()
      
      log.Println("application running in server 5000!")
      return nil
    },
  })
}
