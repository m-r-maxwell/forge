package forge

import (
	muir "forge/pkg"

	"github.com/spf13/cobra"
)

var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "Generate a RESTful API project",
	Run: func(cmd *cobra.Command, args []string) {
		muir.Rest()
	},
}

func init() {
	rootCmd.AddCommand(restCmd)
}
