package forge

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

// forge.go defines the root command for the Forge CLI tool

const asciiArt = `
_______   ______   .______        _______  _______
|   ____| /  __  \  |   _  \      /  _____||   ____|
|  |__   |  |  |  | |  |_)  |    |  |  __  |  |__
|   __|  |  |  |  | |      /     |  | |_ | |   __|
|  |     |  '--'  | |  |\  \----.|  |__| | |  |____
|__|      \______/  | _| '._____| \______| |_______|
`

var version = "0.0.6"

var git bool
var generateReadme bool
var projectName string

var rootCmd = &cobra.Command{
	Use:     "forge",
	Version: version,
	Short:   "Forge is a CLI tool for generating Go projects",
	Long:    `Forge is a CLI tool for generating Go projects of various types. Currently supported types are blank, web, or cli.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(asciiArt)
		cmd.Help()
	},
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Print the version number of forge")
	rootCmd.PersistentFlags().BoolVarP(&generateReadme, "readme", "r", false, "Generate a README.md file after executing the command")
	rootCmd.PersistentFlags().BoolVarP(&git, "git", "g", false, "Initialize a git repository")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Check if the --readme flag is set
	if generateReadme {
		fmt.Println("--readme flag detected. Generating README.md...")
		generateReadmeFile()
	}
}

// Reuse the logic from the `forge doc` command to generate the README.md
func generateReadmeFile() {
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
}
