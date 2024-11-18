package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var buildsGetCmd = &cobra.Command{
	Use: "get",
	RunE: func(cmd *cobra.Command, args []string) error {
		locator, _ := cmd.Flags().GetString("locator")
		fields, _ := cmd.Flags().GetString("fields")

		resp, err := tcapp.GetBuild(locator, fields)
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
	buildsCmd.AddCommand(buildsGetCmd)

	buildsGetCmd.Flags().String("locator", "", "Locator for builds")
	buildsGetCmd.Flags().String("fields", "", "Fields to include in response")
}
