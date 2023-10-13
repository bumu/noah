package data

import (
	"apigw/apps/mall/data/repos"

	"go.uber.org/fx"
)

func FxOptions() fx.Option {
	return fx.Options(
		fx.Provide(
			repos.NewMallShopRepo,
		),
	)
}
