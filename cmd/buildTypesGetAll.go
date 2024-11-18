package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var buildTypesGetAllCmd = &cobra.Command{
	Use: "getAll",
	RunE: func(cmd *cobra.Command, args []string) error {
		buildTypeLocator, _ := cmd.Flags().GetString("build-type-locator")
		fields, _ := cmd.Flags().GetString("fields")

		resp, err := tcapp.GetAllBuildTypes(buildTypeLocator, fields)
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
	buildTypesCmd.AddCommand(buildTypesGetAllCmd)

	buildTypesGetAllCmd.Flags().String("fields", "", "Fields to include in response")
}
