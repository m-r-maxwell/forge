package forge

import (
	muir "forge/pkg"

	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Generate a service",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		muir.Service(args[0])
	},
}

func init() {
	rootCmd.AddCommand(serviceCmd)
}
