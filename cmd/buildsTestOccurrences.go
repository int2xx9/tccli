package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var buildsTestOccurrencesCmd = &cobra.Command{
	Use: "testOccurrences",
	RunE: func(cmd *cobra.Command, args []string) error {
		buildLocator, _ := cmd.Flags().GetString("build-locator")
		fields, _ := cmd.Flags().GetString("fields")

		resp, err := tcapp.GetBuildTestOccurrences(buildLocator, fields)
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

	buildsTestOccurrencesCmd.Flags().String("fields", "", "Fields to include in response")
}
