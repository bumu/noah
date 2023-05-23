package data

import (
	"apigw/apps/apigw/data/repos"

	"go.uber.org/fx"
)

func FxOptions() fx.Option {
	return fx.Options(
		fx.Provide(
			repos.NewUserRepo,
			repos.NewKeyRepo,
			repos.NewIpdbRepo,
			repos.NewIpdbV6Repo,
		),
	)
}
