package forge

import (
	muir "forge/pkg"

	"github.com/spf13/cobra"
)

var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Generate a cli project",
	Run: func(cmd *cobra.Command, args []string) {
		muir.Cli()
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
}
