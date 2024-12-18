/*
Copyright © 2023 dean <dean@airdb.com>
MIT
*/
package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"

	"noah/apps/apigw"
	"noah/apps/company"
	"noah/apps/infra"
	"noah/apps/notification"
	"noah/apps/security"
	"noah/apps/sgw"
	"noah/apps/user"
	"noah/internal/bootstrap"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run noah apps",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

var runApp string

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	runCmd.PersistentFlags().StringVarP(&runApp, "app", "a", "all", "app name is default")
}

type invokeDeps struct {
	fx.In

	Server *bootstrap.Server
}

func run() {
	// Init config when run cobra on OnInitialize().
	// configkit.Init()

	opts := []fx.Option{
		bootstrap.FxOptions(),
		fx.Provide(
		// Add your dependency providers here
		// For example:
		// database.NewConnection,
		// repositories.NewUserRepository,
		),
	}

	// Init or create database.
	switch runApp {
	case "apigw":
		opts = append(opts, apigw.FxOptions())
	case "sgw":
		opts = append(opts, sgw.FxOptions())
	case "notification":
		opts = append(opts, notification.FxOptions())
	case "all":
		opts = append(opts,
			apigw.FxOptions(),
			sgw.FxOptions(),
			security.FxOptions(),
			notification.FxOptions(),
			infra.FxOptions(),
			company.FxOptions(),
			user.FxOptions(),
		)
	default:
		log.Println("no app database init")
	}

	fmt.Println(opts)

	// Add invoke func to opts.
	opts = append(opts,
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
		}))

	app := fx.New(opts...)
	app.Run()
}
