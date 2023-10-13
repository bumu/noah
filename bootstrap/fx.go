package bootstrap

import (
	"noah/apps/apigw/handlers"
	"noah/modules/dbmod"

	"go.uber.org/fx"
)

func FxOptions() fx.Option {
	return fx.Options(
		fx.Provide(dbmod.NewConn),
		fx.Provide(
			New,
			NewMux,
		),
		fx.Invoke(handlers.Register),
	)
}
