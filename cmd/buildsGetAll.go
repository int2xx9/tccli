package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var buildsGetAllCmd = &cobra.Command{
	Use: "getAll",
	RunE: func(cmd *cobra.Command, args []string) error {
		buildLocator, _ := cmd.Flags().GetString("build-locator")
		fields, _ := cmd.Flags().GetString("fields")

		resp, err := tcapp.GetAllBuilds(buildLocator, fields)
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

	buildsGetAllCmd.Flags().String("fields", "", "Fields to include in response")
}
