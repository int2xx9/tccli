package cmd

import (
	"errors"
	"os"

	internalFormatter "github.com/int2xx9/tccli/internal/formatter"
	"github.com/int2xx9/tccli/teamcity"
	"github.com/spf13/cobra"
)

var tcapp *teamcity.Client
var formatter internalFormatter.Formatter

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
	rootCmd.PersistentFlags().String("server", "", "Use a specific server (required)")
	rootCmd.PersistentFlags().String("token", "", "Use a specific token (required)")

	rootCmd.PersistentFlags().String("output", "json", "Output format (json or csv, default is json)")
	rootCmd.PersistentFlags().String("csv-format", "", "CSV format (required if output is csv)")

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if cmd.Use == "config" {
			return nil
		}

		var err error

		output, _ := rootCmd.Flags().GetString("output")
		switch output {
		case "csv":
			csvFormat, _ := rootCmd.Flags().GetString("csv-format")
			if csvFormat == "" {
				return errors.New("csv-format is required if output is csv")
			}

			formatter, err = internalFormatter.NewCsvFormatter(csvFormat, true)
			if err != nil {
				return err
			}
		case "json":
			formatter = &internalFormatter.JsonFormatter{}
		default:
			return errors.New("output must be csv or json")
		}

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
