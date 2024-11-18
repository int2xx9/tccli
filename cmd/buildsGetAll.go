package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var buildsGetAllCmd = &cobra.Command{
	Use: "getAll",
	RunE: func(cmd *cobra.Command, args []string) error {
		locator, _ := cmd.Flags().GetString("locator")
		fields, _ := cmd.Flags().GetString("fields")

		resp, err := tcapp.GetAllBuilds(locator, fields)
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
	buildsCmd.AddCommand(buildsGetAllCmd)

	buildsGetAllCmd.Flags().String("locator", "", "Locator for builds")
	buildsGetAllCmd.Flags().String("fields", "", "Fields to include in response")
}
