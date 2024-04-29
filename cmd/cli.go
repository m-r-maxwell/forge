package forge

import (
	"fmt"
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
		fmt.Println("Starting to forge....")
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
