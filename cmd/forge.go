package forge

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"

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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
