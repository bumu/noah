/*
Copyright © 2023 dean <dean@airdb.com>
MIT
*/
package cmd

import (
	"noah/tools/loader"

	"github.com/spf13/cobra"
)

// appsCmd represents the apps command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load or import data",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		loader.LoadCompany()
	},
}

func init() {
	rootCmd.AddCommand(loadCmd)
}
