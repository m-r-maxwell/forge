package forge

import (
	muir "forge/pkg"

	"github.com/spf13/cobra"
)

var restCmd = &cobra.Command{
	Use:   "rest [name]",
	Short: "Generate a RESTful API project with optional [name]",
	Args:  cobra.MaximumNArgs(1),
	Long:  `Generate a RESTful API project with the provided name or prompt for one.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			muir.Rest(args[0])
		} else {
			muir.Rest("")
		}
	},
}

func init() {
	rootCmd.AddCommand(restCmd)
}
