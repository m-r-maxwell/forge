package muir

import (
	"fmt"
	"forge/helpers"
)

// Git initializes a git repository
func Git() {
	fmt.Println("Running Git init...")
	err := runCommand("git", "init")
	helpers.CheckErrors(err)
}
