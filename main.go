package main

import (
	"context"
	"errors"
	"log"

	"apigw/apps/apigw"

	"apigw/bootstrap"
	"apigw/pkg/configkit"

	"go.uber.org/fx"
)

type invokeDeps struct {
	fx.In

	Server *bootstrap.Server
}

func main() {
	configkit.Init()

	app := fx.New(
		bootstrap.FxOptions(),
		apigw.FxOptions(),
		// sensitivemod.FxOptions(),
		fx.Invoke(func(lc fx.Lifecycle, deps invokeDeps) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					go deps.Server.Start()
					log.Println("Press Ctrl+C to exit")
					return nil
				},
				OnStop: func(ctx context.Context) error {
					deps.Server.Stop()
					return errors.Join(deps.Server.Stop())
				},
			})
		}),
	)

	app.Run()
}
