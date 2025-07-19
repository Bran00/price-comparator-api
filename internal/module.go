// Package internal has all logic of the application
package internal

import "go.uber.org/fx"

var Module = fx.Module("internal",
  fx.Provide(constructors ...interface{})
)
