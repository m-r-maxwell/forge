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

// pythonGuiCmd represents the command to generate a Python GUI project
var pythonGuiCmd = &cobra.Command{
	Use:   "python-gui [name]",
	Short: "Generates a new Python GUI project",
	Long:  "Generates a new Python project with a Tkinter GUI and a venv folder",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting to forge a Python GUI app....")
		pkg.PythonGUI(args[0])
	},
}

// pythonClassCmd represents the command to generate a Python class file
var pythonClassCmd = &cobra.Command{
	Use:   "python-class [dir] [ClassName]",
	Short: "Generates a new Python class file with the given class name",
	Long:  "Creates a Python file in the specified directory with a class of the given name.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		dir := args[0]
		className := args[1]
		fmt.Println("Creating Python class file...")
		pkg.PythonClass(dir, className)
	},
}

func init() {
	rootCmd.AddCommand(pythonCmd)
	rootCmd.AddCommand(pythonGuiCmd)
	rootCmd.AddCommand(pythonClassCmd)
}
