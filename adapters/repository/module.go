//Package repository has the conections
package repository

import "go.uber.org/fx"

var Module = fx.Module("repository",
	fx.Provide(NewSerpAPIRepository),
)
