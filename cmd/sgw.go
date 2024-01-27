/*
Copyright © 2023 dean <dean@airdb.com>
More information, please vist https://airdb.team.
*/
package cmd

import (
	"fmt"
	"noah/cmd/sgw"

	"github.com/spf13/cobra"
)

// sgwCmd represents the sgw command
var sgwCmd = &cobra.Command{
	Use:   "sgw",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sgw called")
		sgw.Run()
	},
}

func init() {
	rootCmd.AddCommand(sgwCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sgwCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sgwCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
