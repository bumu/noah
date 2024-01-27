package bootstrap

import (
	"noah/modules/dbmod"

	apigwHandler "noah/apps/apigw/handlers"
	defaultHandler "noah/apps/default/handlers"
	infraControllers "noah/apps/infra/controllers"
	SecurityControllers "noah/apps/security/controllers"
	sgwControllers "noah/apps/sgw/handlers"

	"go.uber.org/fx"
)

func FxOptions() fx.Option {
	return fx.Options(
		fx.Provide(dbmod.NewConn),
		fx.Provide(
			New,
			NewMux,
		),

		// Register handlers.
		fx.Invoke(defaultHandler.Register),
		fx.Invoke(apigwHandler.Register),
		fx.Invoke(sgwControllers.Register),
		fx.Invoke(SecurityControllers.Register),
		fx.Invoke(infraControllers.Register),
	)
}
