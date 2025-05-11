package forge

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var taskCmd = &cobra.Command{
	Use:   "task [name]",
	Short: "Run tasks defined in Taskfile.yaml",
	Long:  `Run tasks defined in Taskfile.yaml. If no task name is provided, list all available tasks.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Listing available tasks...")
			out, err := exec.Command("task", "--list").Output()
			if err != nil {
				fmt.Println("Error listing tasks:", err)
				return
			}
			fmt.Println(string(out))
		} else {
			taskName := args[0]
			fmt.Printf("Running task: %s\n", taskName)
			cmd := exec.Command("task", taskName)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				fmt.Println("Error running task:", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(taskCmd)
}
