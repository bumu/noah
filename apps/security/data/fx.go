package data

import (
	"noah/apps/security/data/repos"

	"go.uber.org/fx"
)

func FxOptions() fx.Option {
	return fx.Options(
		fx.Provide(
			repos.NewMallShopRepo,
		),
	)
}
