package forge

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display metadata about the project",
	Long:  `Display metadata about the project, such as its name, type, and dependencies.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Project Metadata:")

		// Get the current working directory
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
			return
		}
		fmt.Println("- Project Path:", cwd)

		// Get the project name
		projectName := filepath.Base(cwd)
		fmt.Println("- Project Name:", projectName)

		// Display version
		fmt.Println("- Version:", version)

		// Check if git is initialized
		if _, err := os.Stat(filepath.Join(cwd, ".git")); err == nil {
			fmt.Println("- Git Initialized: Yes")
		} else {
			fmt.Println("- Git Initialized: No")
		}
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
