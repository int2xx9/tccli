package cmd

import (
	"github.com/spf13/cobra"
)

var buildsCmd = &cobra.Command{
	Use: "builds",
}

func init() {
	rootCmd.AddCommand(buildsCmd)

	buildsCmd.PersistentFlags().String("build-locator", "", "Locator for builds")
}
