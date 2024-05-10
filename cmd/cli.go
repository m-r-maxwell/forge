package forge

import (
	"fmt"
	muir "forge/pkg"

	"github.com/spf13/cobra"
)

var cliCmd = &cobra.Command{
	Use:     "cli [name]",
	Short:   "Generate a cli project with [name]",
	Long:    `Generate a cobra cli project with the provided name.`,
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"c"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting to forge....")
		if len(args) > 0 {
			muir.Cli(args[0], git)
		} else {
			fmt.Println("Please provide a name for your cli project")
		}
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
}
