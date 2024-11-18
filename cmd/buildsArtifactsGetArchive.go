package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

var buildsArtifactsGetArchiveCmd = &cobra.Command{
	Use: "getArchive",
	RunE: func(cmd *cobra.Command, args []string) error {
		locator, _ := cmd.Flags().GetString("locator")
		path, _ := cmd.Flags().GetString("path")
		archiveLocator, _ := cmd.Flags().GetString("archive-locator")

		resp, err := tcapp.GetBuildArtifactArchive(locator, path, archiveLocator)
		if err != nil {
			return err
		}
		defer resp.Close()

		io.Copy(os.Stdout, resp)
		return nil
	},
}

func init() {
	buildsArtifactsCmd.AddCommand(buildsArtifactsGetArchiveCmd)

	buildsArtifactsGetArchiveCmd.Flags().String("locator", "", "Locator for builds")
	buildsArtifactsGetArchiveCmd.Flags().String("path", "", "Path to the artifact")
	buildsArtifactsGetArchiveCmd.Flags().String("archive-locator", "", "Locator for the archive")
}
