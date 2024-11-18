package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var buildsGetAllCmd = &cobra.Command{
	Use: "getAll",
	RunE: func(cmd *cobra.Command, args []string) error {
		locator, err := cmd.Flags().GetString("locator")
		if err != nil {
			return err
		}

		resp, err := tcapp.GetAllBuilds(locator)
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
}
