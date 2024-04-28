package forge

import (
	muir "forge/pkg"

	"github.com/spf13/cobra"
)

// var useBlank bool
var blankCmd = &cobra.Command{
	Use:   "blank",
	Short: "Generate a blank Go project",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			muir.Blank(args[0])
		} else {
			muir.Blank("")
		}
	},
}

func init() {
	//rootCmd.PersistentFlags().BoolVarP(&useBlank, "blank", "b", false, "Generate a blank Go project")
	rootCmd.AddCommand(blankCmd)
}
