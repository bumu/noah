package bootstrap

import (
	"api/modules/ip2locationmod"
	"api/modules/ipipmod"

	"go.uber.org/fx"
)

func FxOptions() fx.Option {
	return fx.Options(
		fx.Provide(ip2locationmod.NewIPdb),
		fx.Provide(ipipmod.NewDB),
		fx.Provide(NewServe),
	)
}
