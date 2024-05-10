package forge

import (
	"fmt"
	muir "forge/pkg"

	"github.com/spf13/cobra"
)

var restCmd = &cobra.Command{
	Use:     "rest [name]",
	Short:   "Generate a RESTful API project with provided [name]",
	Args:    cobra.ExactArgs(1),
	Long:    `Generate a RESTful API project with the provided name with std lib net/http.`,
	Aliases: []string{"r"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting to forge....")
		if len(args) > 0 {
			muir.Rest(args[0], git)
		} else {
			fmt.Println("Please provide a name for your RESTful API project")
		}
	},
}

func init() {
	rootCmd.AddCommand(restCmd)
}
