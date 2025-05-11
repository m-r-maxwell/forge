package forge

import (
	"fmt"
	pkg "forge/pkg"

	"github.com/spf13/cobra"
)

// cliCmd represents the command to generate a CLI project
var cliCmd = &cobra.Command{
	Use:     "cli [name]",
	Short:   "Generate a CLI project with [name]",
	Long:    `Generate a Cobra CLI project with the provided name.`,
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"c"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting to forge....")
		pkg.Cli(args[0], git)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
}
