package cmd

import (
	"os"

	internalFormatter "github.com/int2xx9/tccli/internal/formatter"
	"github.com/int2xx9/tccli/teamcity"
	"github.com/spf13/cobra"
)

var tcapp *teamcity.Client
var formatter internalFormatter.Formatter = &internalFormatter.JsonFormatter{}

var rootCmd = &cobra.Command{
	Use: "tccli",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("profile", "", "Use a specific profile from your credential file")
	rootCmd.PersistentFlags().String("server", "", "Use a specific server")
	rootCmd.PersistentFlags().String("token", "", "Use a specific token")

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if cmd.Use == "config" {
			return nil
		}

		var err error
		tcapp, err = createTeamcityClient()
		return err
	}
}

func createTeamcityClient() (*teamcity.Client, error) {
	server, err := rootCmd.PersistentFlags().GetString("server")
	if err != nil {
		return nil, err
	}

	token, err := rootCmd.PersistentFlags().GetString("token")
	if err != nil {
		return nil, err
	}

	return teamcity.NewClient(server, token), nil
}
