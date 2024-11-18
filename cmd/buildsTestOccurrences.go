package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var buildsTestOccurrencesCmd = &cobra.Command{
	Use: "testOccurrences",
	RunE: func(cmd *cobra.Command, args []string) error {
		locator, err := cmd.Flags().GetString("locator")
		if err != nil {
			return err
		}

		resp, err := tcapp.GetBuildTestOccurrences(locator)
		if err != nil {
			return err
		}

		output, err := formatter.Format(resp)
		if err != nil {
			return err
		}

		fmt.Println(output)
		return nil
	},
}

func init() {
	buildsCmd.AddCommand(buildsTestOccurrencesCmd)

	buildsTestOccurrencesCmd.Flags().String("locator", "", "Locator for builds")
}
