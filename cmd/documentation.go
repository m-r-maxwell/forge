package forge

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "Generate a README.md file for the project",
	Long:  `Generate a basic README.md file for the project, including metadata.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating README.md...")

		// Detect project name dynamically if not provided
		if projectName == "" {
			cwd, err := os.Getwd()
			if err != nil {
				fmt.Println("Error detecting project name:", err)
				return
			}
			projectName = filepath.Base(cwd)
		}

		file, err := os.Create("README.md")
		if err != nil {
			fmt.Println("Error creating README.md:", err)
			return
		}
		defer file.Close()

		content := `# ` + projectName + `

This project was generated using the Forge CLI tool.

## Metadata

- Version: ` + version + `
- Git initialized: ` + fmt.Sprintf("%v", git) + `
- Generated on: ` + time.Now().Format("January 2, 2006") + `

`
		_, err = file.WriteString(content)
		if err != nil {
			fmt.Println("Error writing to README.md:", err)
			return
		}

		fmt.Println("README.md generated successfully!")
	},
}

func init() {
	docCmd.Flags().StringVarP(&projectName, "name", "n", "", "Specify a custom project name for the README")
	rootCmd.AddCommand(docCmd)
}
