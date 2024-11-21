package bootstrap

import (
	"noah/internal/dbmod"

	defaultHandler "noah/apps/default/handlers"

	apigwHandler "noah/apps/apigw/handlers"
	companyControllers "noah/apps/company/handlers"
	infraControllers "noah/apps/infra/controllers"
	securityControllers "noah/apps/security/controllers"
	sgwControllers "noah/apps/sgw/handlers"
	userHandlers "noah/apps/user/handlers"

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
		fx.Invoke(securityControllers.Register),
		fx.Invoke(infraControllers.Register),
		fx.Invoke(companyControllers.Register),
		fx.Invoke(userHandlers.Register),
	)
}
