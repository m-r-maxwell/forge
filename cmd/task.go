package forge

import (
	"fmt"
	pkg "forge/pkg"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var listTasks bool

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Create a default Taskfile.yaml for your project",
	Long:  `Creates a default Taskfile.yaml in the current directory if it does not exist.`,
	Run: func(cmd *cobra.Command, args []string) {
		taskfilePath := filepath.Join(".", "Taskfile.yaml")
		if listTasks {
			if _, err := os.Stat(taskfilePath); err != nil {
				fmt.Println("Taskfile.yaml does not exist. Use --init to create one.")
				return
			}
			out, err := exec.Command("task", "--list").Output()
			if err != nil {
				fmt.Println("Error listing tasks:", err)
				return
			}
			fmt.Println(string(out))
			return
		}
		if _, err := os.Stat(taskfilePath); err == nil {
			fmt.Println("Taskfile.yaml already exists.")
			return
		}
		pkg.Task()
		fmt.Println("Created Taskfile.yaml.")

	},
}

func init() {
	taskCmd.Flags().BoolVarP(&listTasks, "list", "l", false, "List tasks in Taskfile.yaml")
	rootCmd.AddCommand(taskCmd)
}
