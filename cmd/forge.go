package forge

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.2"

var git bool

var rootCmd = &cobra.Command{
	Use:     "forge",
	Version: version,
	Short:   "Forge is a CLI tool for generating Go projects",
	Long:    `Forge is a CLI tool for generating Go projects of various types. Currently supported types are blank, web, or cli.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
	_______   ______   .______        _______  _______
	|   ____| /  __  \  |   _  \      /  _____||   ____|
	|  |__   |  |  |  | |  |_)  |    |  |  __  |  |__
	|   __|  |  |  |  | |      /     |  | |_ | |   __|
	|  |     |  '--'  | |  |\  \----.|  |__| | |  |____
	|__|      \______/  | _| '._____| \______| |_______|
	`)
		cmd.Help()
	},
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Print the version number of forge")
	rootCmd.Flags().BoolP("help", "h", false, "Print the help menu")
	//rootCmd.PersistentFlags().BoolVarP(&git, "git", "g", false, "Initialize a git repository")
	blankCmd.Flags().BoolVarP(&git, "git", "g", false, "Initialize a git repository")
	cliCmd.Flags().BoolVarP(&git, "git", "g", false, "Initialize a git repository")
	restCmd.Flags().BoolVarP(&git, "git", "g", false, "Initialize a git repository")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
