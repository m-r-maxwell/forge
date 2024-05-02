package forge

import (
	"fmt"
	muir "forge/pkg"

	"github.com/spf13/cobra"
)

// var useBlank bool
var blankCmd = &cobra.Command{
	Use:     "blank [name]",
	Short:   "Generate a blank Go project with an optional [name]",
	Long:    `Generate a blank Go project with the provided name or prompt for one.`,
	Args:    cobra.MaximumNArgs(1),
	Aliases: []string{"b", "--b", "new"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting to forge....")
		if len(args) > 0 {
			muir.Blank(args[0], git)
		} else {
			muir.Blank("", git)
		}
	},
}

func init() {
	rootCmd.AddCommand(blankCmd)
}
