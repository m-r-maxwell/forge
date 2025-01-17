package forge

import (
	"fmt"
	pkg "forge/pkg"
	"os"

	"github.com/spf13/cobra"
)

var structureCmd = &cobra.Command{
	Use:     "structure [path]",
	Short:   "Generate a project structure at [path]",
	Long:    `Generate a project structure at the provided path, to make it easier to share.`,
	Args:    cobra.MaximumNArgs(1),
	Aliases: []string{"st", "struct"},
	Run: func(cmd *cobra.Command, args []string) {
		var path string
		if len(args) == 0 {
			var err error
			path, err = os.Getwd()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		} else {
			path = args[0]
		}
		pkg.PrintStructure(path, "")
	},
}

func init() {
	rootCmd.AddCommand(structureCmd)
}
