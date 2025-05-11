package muir

import (
	"fmt"
	"forge/helpers"
	"forge/str"
)

func Blank(arg string, git bool) {
	str := str.BLANK

	err := createDir(arg)
	helpers.CheckErrors(err)
	err = changeDir(arg)
	helpers.CheckErrors(err)
	if git {
		Git()
	}
	err = runCommand("go", "mod", "init", arg)
	helpers.CheckErrors(err)

	err = createFileWithContent(MainGoFile, str)
	helpers.CheckErrors(err)
	fmt.Println("Project created successfully!")
	err = runCommand("go", "mod", "tidy")
	helpers.CheckErrors(err)
	helpers.GoFMT()
	helpers.CheckErrors(err)
	fmt.Println("You're good to Go :)")
}

func Rest(arg string, git bool) {
	st := str.WEB

	err := createDir(arg)
	helpers.CheckErrors(err)
	err = changeDir(arg)
	helpers.CheckErrors(err)
	if git {
		Git()
	}
	err = runCommand("go", "mod", "init", arg)
	helpers.CheckErrors(err)

	fmt.Println("Creating middleware folder...")
	err = createDir(MiddlewareFolder)
	helpers.CheckErrors(err)
	changeDir(MiddlewareFolder)
	fmwErr := createFileWithContent("middleware.go", str.MW)
	helpers.CheckErrors(fmwErr)

	err = changeDir("..")
	helpers.CheckErrors(err)
	// models
	fmt.Println("Creating models folder...")
	err = createDir(ModelsFolder)
	helpers.CheckErrors(err)
	changeDir(ModelsFolder)
	fmErr := createFileWithContent("models.go", str.MDL)
	helpers.CheckErrors(fmErr)

	err = changeDir("..")
	helpers.CheckErrors(err)
	// helpers
	fmt.Println("Creating helpers folder... delete if not needed")
	err = createDir(HelpersFolder)
	helpers.CheckErrors(err)

	// main file
	err = createFileWithContent(MainGoFile, st)
	helpers.CheckErrors(err)

	fmt.Println("Project created successfully!")
	err = runCommand("go", "mod", "tidy")
	helpers.CheckErrors(err)
	err = runCommand("go", "fmt")
	helpers.CheckErrors(err)
	fmt.Println("You're good to Go :)")
	helpers.GoFMT()
}
