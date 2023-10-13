/*
Copyright © 2023 dean <dean@airdb.com>
MIT
*/
package cmd

import (
	"context"
	"errors"
	"log"

	"noah/apps/apigw"
	"noah/bootstrap"
	"noah/pkg/configkit"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type invokeDeps struct {
	fx.In

	Server *bootstrap.Server
}

func run() {
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
