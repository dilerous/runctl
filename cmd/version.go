/*
Copyright © 2024 brad.soper@run.ai
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Set the version of the cnvrgctl cli
var Version = "v0.0.1"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the version of the runctl cli",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		InfoLogger.Println("version called")

		// Display the version
		displayVersion(Version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// displayVersion displays the version of the runctl cli
// v is the version to display
func displayVersion(v string) {
	InfoLogger.Println("runctl version " + v)
	fmt.Println("runctl version " + v)
}
