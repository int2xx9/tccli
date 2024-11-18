package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

var buildsArtifactsGetArchiveCmd = &cobra.Command{
	Use: "getArchive",
	RunE: func(cmd *cobra.Command, args []string) error {
		buildLocator, _ := cmd.Flags().GetString("build-locator")
		path, _ := cmd.Flags().GetString("path")
		archiveLocator, _ := cmd.Flags().GetString("archive-locator")

		resp, err := tcapp.GetBuildArtifactArchive(buildLocator, path, archiveLocator)
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

	buildsArtifactsGetArchiveCmd.Flags().String("path", "", "Path to the artifact")
	buildsArtifactsGetArchiveCmd.Flags().String("archive-locator", "", "Locator for the archive")
}
