package forge

import (
	"fmt"
	pkg "forge/pkg"

	"github.com/spf13/cobra"
)

// modelsCmd represents the command to generate a models folder and file
var modelsCmd = &cobra.Command{
	Use:     "models",
	Short:   "Generate a models folder and file",
	Aliases: []string{"m"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting to forge....")
		pkg.Models()
	},
}

func init() {
	rootCmd.AddCommand(modelsCmd)
}
