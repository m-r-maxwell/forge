package forge

import (
	"fmt"
	pkg "forge/pkg"

	"github.com/spf13/cobra"
)

// restCmd represents the command to generate a RESTful API project
var restCmd = &cobra.Command{
	Use:     "rest [name]",
	Short:   "Generate a RESTful API project with provided [name]",
	Args:    cobra.ExactArgs(1),
	Long:    `Generate a RESTful API project with the provided name using the standard library net/http.`,
	Aliases: []string{"r"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting to forge....")
		pkg.Rest(args[0], git)
	},
}

func init() {
	rootCmd.AddCommand(restCmd)
}
