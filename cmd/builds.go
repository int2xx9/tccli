package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var buildsCmd = &cobra.Command{
	Use: "builds",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("builds called")
	},
}

func init() {
	rootCmd.AddCommand(buildsCmd)
}
