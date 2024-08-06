package forge

import (
	"fmt"
	muir "forge/pkg"

	"github.com/spf13/cobra"
)

var pythonCmd = &cobra.Command{
	Use:     "python [name]",
	Short:   "Generates a new python project",
	Long:    "Generates a new pythong project with a venv folder",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"p"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting to forge....")
		if len(args) > 0 {
			muir.Python(args[0])

		} else {
			fmt.Println("Please provide a name for your project")
		}
	},
}

func init() {
	rootCmd.AddCommand(pythonCmd)
}
