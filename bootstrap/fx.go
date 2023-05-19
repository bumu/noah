package bootstrap

import (
	"apigw/modules/dbmod"

	"apigw/apps/apigw/handlers"

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
