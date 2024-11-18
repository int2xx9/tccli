package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var buildTypesGetCmd = &cobra.Command{
	Use: "get",
	RunE: func(cmd *cobra.Command, args []string) error {
		buildTypeLocator, err := cmd.Flags().GetString("build-type-locator")
		if err != nil {
			return err
		}

		resp, err := tcapp.GetBuildType(buildTypeLocator)
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
	buildTypesCmd.AddCommand(buildTypesGetCmd)
}
