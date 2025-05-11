package forge

import (
	"fmt"
	pkg "forge/pkg"

	"github.com/spf13/cobra"
)

// blankCmd represents the command to generate a blank Go project
var blankCmd = &cobra.Command{
	Use:     "blank [name]",
	Short:   "Generate a blank Go project with [name]",
	Long:    `Generate a blank Go project with the provided name, the world is your oyster.`,
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"b", "new"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting to forge....")
		pkg.Blank(args[0], git)
	},
}

func init() {
	rootCmd.AddCommand(blankCmd)
}
