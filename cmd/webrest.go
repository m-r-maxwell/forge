package forge

import (
	"fmt"
	muir "forge/pkg"

	"github.com/spf13/cobra"
)

var restCmd = &cobra.Command{
	Use:     "rest [name]",
	Short:   "Generate a RESTful API project with optional [name]",
	Args:    cobra.MaximumNArgs(1),
	Long:    `Generate a RESTful API project with the provided name or prompt for one.`,
	Aliases: []string{"r"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting to forge....")
		if len(args) > 0 {
			muir.Rest(args[0], git)
		} else {
			muir.Rest("", git)
		}
	},
}

func init() {
	rootCmd.AddCommand(restCmd)
}
