package forge

import (
	"fmt"
	muir "forge/pkg"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.2"

// var git bool

var rootCmd = &cobra.Command{
	Use:     "forge",
	Version: version,
	Short:   "Forge is a CLI tool for generating Go projects",
	Long:    `Forge is a CLI tool for generating Go projects of various types. Blank, web, or cli.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
	_______   ______   .______        _______  _______
	|   ____| /  __  \  |   _  \      /  _____||   ____|
	|  |__   |  |  |  | |  |_)  |    |  |  __  |  |__
	|   __|  |  |  |  | |      /     |  | |_ | |   __|
	|  |     |  '--'  | |  |\  \----.|  |__| | |  |____
	|__|      \______/  | _| '._____| \______| |_______|
	`)
	},
}

// func init() {
//rootCmd.PersistentFlags().BoolVarP(&git, "git", "g", false, "Initialize a git repository")
// blankCmd.PersistentFlags().BoolVarP(&git, "git", "g", false, "Initialize a git repository")
//restCmd.Flags().BoolP("rest", "r", false, "Generate a RESTful API project")

//}

func Execute() {
	if git, _ := rootCmd.Flags().GetBool("git"); git {
		muir.Git()
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
