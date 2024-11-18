package cmd

import (
	"github.com/spf13/cobra"
)

var buildTypesCmd = &cobra.Command{
	Use: "buildTypes",
}

func init() {
	rootCmd.AddCommand(buildTypesCmd)

	buildTypesCmd.PersistentFlags().String("build-type-locator", "", "Locator for build types")
}
