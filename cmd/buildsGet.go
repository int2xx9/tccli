package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var buildsGetCmd = &cobra.Command{
	Use: "get",
	RunE: func(cmd *cobra.Command, args []string) error {
		buildLocator, _ := cmd.Flags().GetString("build-locator")
		fields, _ := cmd.Flags().GetString("fields")

		resp, err := tcapp.GetBuild(buildLocator, fields)
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

	buildsGetCmd.Flags().String("fields", "", "Fields to include in response")
}
