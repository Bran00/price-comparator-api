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
      ginEngineProvider,
  ),
  fx.Invoke(startServer),
)

func startServer(lc fx.Lifecycle, router *gin.Engine) {
  server := &http.Server{
    Addr: fmt.Sprintf(":%d", 5000),
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
    OnStop: func(ctx context.Context) error {
      return server.Shutdown(ctx)
    },
  })
}

func ginEngineProvider() *gin.Engine {
  app := gin.New()

  return app
}
