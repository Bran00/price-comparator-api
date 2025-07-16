package routes

import (
	"fmt"
	"net/http"

	"go.uber.org/fx"
)

var Module = fx.Module("routes",
	fx.Provide(NewMux),
	fx.Invoke(StartServe),
)

func StartServe(mux *http.Server) {
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", mux)
}
