package forge

import (
	muir "forge/pkg"

	"github.com/spf13/cobra"
)

var cliCmd = &cobra.Command{
	Use:     "cli [name]",
	Short:   "Generate a cli project with an optional [name]",
	Long:    `Generate a cli project with the provided name or prompt for one.`,
	Args:    cobra.MaximumNArgs(1),
	Aliases: []string{"c", "--c"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			muir.Cli(args[0])
		} else {
			muir.Cli("")
		}
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
}
