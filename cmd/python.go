package forge

import (
	"fmt"
	pkg "forge/pkg"

	"github.com/spf13/cobra"
)

// pythonCmd represents the command to generate a Python project
var pythonCmd = &cobra.Command{
	Use:     "python [name]",
	Short:   "Generates a new Python project",
	Long:    "Generates a new Python project with a venv folder",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"p"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting to forge....")
		pkg.Python(args[0])
	},
}

func init() {
	rootCmd.AddCommand(pythonCmd)
}
