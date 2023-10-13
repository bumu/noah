package apigw

import (
	"apigw/apps/mall/data"

	"go.uber.org/fx"
)

func FxOptions() fx.Option {
	return fx.Options(
		data.FxOptions(),
		fx.Invoke(data.Migrate),
	)
}
