package forge

import (
	"fmt"
	"os"

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
	rootCmd.PersistentFlags().BoolVarP(&git, "git", "g", false, "Initialize a git repository")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
