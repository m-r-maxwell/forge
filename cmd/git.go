package forge

import (
	muir "forge/pkg"

	"github.com/spf13/cobra"
)

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Generate a git project",
	Run: func(cmd *cobra.Command, args []string) {
		muir.Git()
	},
}

func init() {
	rootCmd.AddCommand(gitCmd)
}
