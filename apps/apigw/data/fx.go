package data

import (
	"noah/apps/apigw/data/repos"

	"go.uber.org/fx"
)

func FxOptions() fx.Option {
	return fx.Options(
		fx.Provide(
			repos.NewUserRepo,
			repos.NewKeyRepo,
			repos.NewIpdbRepo,
			repos.NewIpdbV6Repo,
			repos.NewSensitiveRepo,
			repos.NewIpdbDataCenterRepo,
			repos.NewIpdbClientIpRepo,
			repos.NewUseragentRepo,
			repos.NewUseragentOSRepo,
		),
	)
}
