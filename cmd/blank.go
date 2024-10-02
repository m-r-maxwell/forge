package forge

import (
	"fmt"
	pkg "forge/pkg"

	"github.com/spf13/cobra"
)

// var useBlank bool
var blankCmd = &cobra.Command{
	Use:     "blank [name]",
	Short:   "Generate a blank Go project with [name]",
	Long:    `Generate a blank Go project with the provided name, the world is your oyster.`,
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"b", "new"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting to forge....")
		if len(args) > 0 {
			pkg.Blank(args[0], git)
		} else {
			fmt.Println("Please provide a name for your blank project")
		}
	},
}

func init() {
	rootCmd.AddCommand(blankCmd)
}
