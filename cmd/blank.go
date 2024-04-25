package forge

import (
	muir "forge/pkg"

	"github.com/spf13/cobra"
)

// var useBlank bool
var blankCmd = &cobra.Command{
	Use:   "blank",
	Short: "Generate a blank Go project",
	Run: func(cmd *cobra.Command, args []string) {
		//if useBlank {
		muir.Blank()
		//}
	},
}

func init() {
	//rootCmd.PersistentFlags().BoolVarP(&useBlank, "blank", "b", false, "Generate a blank Go project")
	rootCmd.AddCommand(blankCmd)
}
