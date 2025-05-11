package helpers

import (
	"fmt"
	"os"
	"os/exec"
)

// GetInput reads a line of input from the user and returns it as a string.
// It validates that the input is not empty.
func GetInput() string {
	var input string
	_, err := fmt.Scanln(&input)
	CheckErrors(err)

	if input == "" {
		fmt.Println("Error: Input cannot be empty.")
		os.Exit(1)
	}

	return input
}

// CheckErrors checks if an error is not nil and exits the program if true.
// It prints the error message before exiting.
func CheckErrors(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

// GoFMT runs the `go fmt` command to format Go code in the current directory.
func GoFMT() {
	cmd := exec.Command("go", "fmt")
	err := cmd.Run()
	CheckErrors(err)
}

// GoModCobra installs the Cobra library by running `go get` for the Cobra package.
func GoModCobra() {
	cmd := exec.Command("go", "get", "github.com/spf13/cobra")
	err := cmd.Run()
	CheckErrors(err)
}
