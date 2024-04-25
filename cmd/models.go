package forge

import (
	muir "forge/pkg"

	"github.com/spf13/cobra"
)

var modelsCmd = &cobra.Command{
	Use:   "models",
	Short: "Generate a models folder and file",
	Run: func(cmd *cobra.Command, args []string) {
		muir.Models()
	},
}

func init() {
	rootCmd.AddCommand(modelsCmd)
}
