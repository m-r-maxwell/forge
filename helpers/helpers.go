package helpers

import (
	"fmt"
	"os"
	"os/exec"
)

func GetInput() string {
	var input string
	_, err := fmt.Scanln(&input)
	CheckErrors(err)
	return input

}

func CheckErrors(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

func GoFMT() {
	cmd := exec.Command("go", "fmt")
	err := cmd.Run()
	CheckErrors(err)
}

func GoModCobra() {
	cmd := exec.Command("go", "get", "github.com/spf13/cobra")
	err := cmd.Run()
	CheckErrors(err)
}

func ChangeDirSinceDone(name string) {
	err := os.Chdir(name)
	CheckErrors(err)
}
