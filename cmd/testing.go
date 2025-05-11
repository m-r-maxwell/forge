package forge

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Generate boilerplate test files for Go projects",
	Long:  `Generate a *_test.go file for each .go file in the project that doesn't already have a test file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating test files...")
		err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if strings.HasSuffix(info.Name(), ".go") && !strings.HasSuffix(info.Name(), "_test.go") {
				testFile := strings.TrimSuffix(path, ".go") + "_test.go"
				if _, err := os.Stat(testFile); os.IsNotExist(err) {
					file, err := os.Create(testFile)
					if err != nil {
						return err
					}
					defer file.Close()
					file.WriteString("package " + filepath.Base(filepath.Dir(path)) + "\n\n")
					file.WriteString("import \"testing\"\n\n")
					file.WriteString("func TestPlaceholder(t *testing.T) {\n\t// Add your tests here\n}\n")
					fmt.Println("Created:", testFile)
				}
			}
			return nil
		})

		if err != nil {
			fmt.Println("Error generating test files:", err)
		} else {
			fmt.Println("Test files generated successfully!")
		}
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
