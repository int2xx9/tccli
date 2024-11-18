package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var buildsTestOccurrencesCmd = &cobra.Command{
	Use: "testOccurrences",
	RunE: func(cmd *cobra.Command, args []string) error {
		locator, _ := cmd.Flags().GetString("locator")
		fields, _ := cmd.Flags().GetString("fields")

		resp, err := tcapp.GetBuildTestOccurrences(locator, fields)
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
	buildsTestOccurrencesCmd.Flags().String("fields", "", "Fields to include in response")
}
