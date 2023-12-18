package data

import (
	"noah/apps/infra/data/repos"

	"go.uber.org/fx"
)

func FxOptions() fx.Option {
	return fx.Options(
		fx.Provide(
			repos.NewInfraLinuxRepo,
		),
	)
}
