package forge

import (
	"fmt"
	muir "forge/pkg"

	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:     "service [name]",
	Short:   "Generate a service",
	Long:    `Generate a service with the provided name.`,
	Aliases: []string{"s"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting to forge....")
		muir.Service(args[0])
	},
}

func init() {
	rootCmd.AddCommand(serviceCmd)
}
