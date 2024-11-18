package cmd

import (
	"github.com/spf13/cobra"
)

var buildsArtifactsCmd = &cobra.Command{
	Use: "artifacts",
}

func init() {
	buildsCmd.AddCommand(buildsArtifactsCmd)
}
