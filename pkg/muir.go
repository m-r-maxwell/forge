package muir

import (
	"fmt"
	"forge/helpers"
	"forge/str"
	"os"
	"os/exec"
)

func Blank(arg string) {
	str := str.BLANK

	if arg == "" {
		fmt.Println("What would you like to name your project?")
		fmt.Println("This will be the name of the folder and the Go module.")
		name := helpers.GetInput()
		err := os.Mkdir(name, 0755)
		helpers.CheckErrors(err)
		helpers.CheckErrors(err)
		err = os.Chdir(name)
		helpers.CheckErrors(err)
		fmt.Println("Running Go mod init...")
		cmd := exec.Command("go", "mod", "init", name)
		err = cmd.Run()
		helpers.CheckErrors(err)
	} else {
		err := os.Mkdir(arg, 0755)
		helpers.CheckErrors(err)
		err = os.Chdir(arg)
		helpers.CheckErrors(err)
		fmt.Println("Running Go mod init...")
		cmd := exec.Command("go", "mod", "init", arg)
		err = cmd.Run()
		helpers.CheckErrors(err)
	}
	fmt.Println("Go mod init successful!")

	file, err := os.Create("main.go")
	helpers.CheckErrors(err)
	defer file.Close()
	_, err = file.WriteString(str)
	helpers.CheckErrors(err)
	fmt.Println("Project created successfully!")

	fmt.Println("Running Go mod tidy...")
	cmd := exec.Command("go", "mod", "tidy")
	err = cmd.Run()
	helpers.CheckErrors(err)
	fmt.Println("Go mod tidy successful!")
	fmt.Println("Running Go fmt...")
	helpers.GoFMT()
	helpers.CheckErrors(err)
	fmt.Println("Go fmt successful!")
	fmt.Println("You're good to Go :)")
}

func Rest() {
	st := str.WEB
	fmt.Println("What would you like to name your project?")
	fmt.Println("This will be the name of the folder and the Go module.")
	name := helpers.GetInput()
	err := os.Mkdir(name, 0755)
	helpers.CheckErrors(err)
	err = os.Chdir(name)
	helpers.CheckErrors(err)

	fmt.Println("Creating middleware folder...")
	err = os.Mkdir("middleware", 0755)
	helpers.CheckErrors(err)
	os.Chdir("middleware")
	fmw, err := os.Create("middleware.go")
	helpers.CheckErrors(err)
	fmt.Println("Creating middleware.go file...")
	_, err = fmw.WriteString(str.MW)
	defer fmw.Close()
	helpers.CheckErrors(err)

	err = os.Chdir("..")
	helpers.CheckErrors(err)
	// models
	fmt.Println("Creating models folder...")
	err = os.Mkdir("models", 0755)
	helpers.CheckErrors(err)
	os.Chdir("models")
	fmt.Println("Creating models.go file...")
	fm, err := os.Create("models.go")
	helpers.CheckErrors(err)
	_, err = fm.WriteString(str.MDL)
	defer fm.Close()
	helpers.CheckErrors(err)

	err = os.Chdir("..")
	helpers.CheckErrors(err)
	// helpers
	fmt.Println("Creating helpers folder... delete if not needed")
	err = os.Mkdir("helpers", 0755)
	helpers.CheckErrors(err)

	// main file
	file, err := os.Create("main.go")
	helpers.CheckErrors(err)
	defer file.Close()
	_, err = file.WriteString(st)
	helpers.CheckErrors(err)

	fmt.Println("Project created successfully!")
	fmt.Println("Running Go mod init...")
	cmd := exec.Command("go", "mod", "init", name)
	err = cmd.Run()
	helpers.CheckErrors(err)
	fmt.Println("Go mod init successful!")
	fmt.Println("Running Go mod tidy...")
	cmd = exec.Command("go", "mod", "tidy")
	err = cmd.Run()
	helpers.CheckErrors(err)
	fmt.Println("Go mod tidy successful!")
	fmt.Println("Running Go fmt...")
	cmd = exec.Command("go", "fmt")
	err = cmd.Run()
	helpers.CheckErrors(err)
	fmt.Println("Go fmt successful!")
	fmt.Println("You're good to Go :)")
	helpers.GoFMT()
}

func Git() {
	fmt.Println("Running Git init...")
	cmd := exec.Command("git", "init")
	err := cmd.Run()
	helpers.CheckErrors(err)
	fmt.Println("Git init successful!")
}

func Service(arg1 string) {
	fmt.Println("Creating: ", arg1, " service in the "+arg1+" folder")
	err := os.Mkdir(arg1, 0755)
	helpers.CheckErrors(err)
	err = os.Chdir(arg1)
	helpers.CheckErrors(err)
	pf := fmt.Sprintf("package %s", arg1)
	tf := fmt.Sprintf("type %sService struct {}", arg1)
	str := pf + "\n" + tf + "\n" + `
	// Add your service code here
	// Such as func NewService(...interface{}) *Service
	//{
		// return &Service{}
	//}
	`
	file, err := os.Create(arg1 + ".service.go")
	helpers.CheckErrors(err)
	defer file.Close()
	_, err = file.WriteString(str)
	helpers.CheckErrors(err)
	helpers.GoFMT()
	fmt.Println("Service created successfully!")
	fmt.Println("You're good to Go :)")
}

func Models() {
	fmt.Println("Creating models folder...")
	err := os.Mkdir("models", 0755)
	helpers.CheckErrors(err)
	os.Chdir("models")
	fmt.Println("Creating models.go file...")
	file, err := os.Create("models.go")
	helpers.CheckErrors(err)
	defer file.Close()
	_, err = file.WriteString(str.MDL)
	helpers.CheckErrors(err)
	fmt.Println("Models created successfully!")
	fmt.Println("You're good to Go :)")
	helpers.GoFMT()
}

func Cli() {
	st := str.CLI
	fmt.Println("What would you like to name your project?")
	fmt.Println("This will be the name of the folder and the Go module.")
	name := helpers.GetInput()
	err := os.Mkdir(name, 0755)
	helpers.CheckErrors(err)
	err = os.Chdir(name)
	helpers.CheckErrors(err)

	fmt.Println("Creating cmd folder...")
	err = os.Mkdir("cmd", 0755)
	helpers.CheckErrors(err)
	os.Chdir("cmd")
	fmt.Println("Creating cmd main.go file...")
	file, err := os.Create("main.go")
	helpers.CheckErrors(err)
	defer file.Close()
	_, err = file.WriteString(str.CLIROOT)
	helpers.CheckErrors(err)

	os.Chdir("..")
	fmt.Println("Creating pkg folder...")
	os.Mkdir("pkg", 0755)

	fmt.Println("Creating main.go file...")
	file, err = os.Create("main.go")
	helpers.CheckErrors(err)
	defer file.Close()
	_, err = file.WriteString(st)
	helpers.CheckErrors(err)

	fmt.Println("Project created successfully!")
	fmt.Println("Running Go mod init...")
	cmd := exec.Command("go", "mod", "init", name)
	err = cmd.Run()
	helpers.CheckErrors(err)
	fmt.Println("Go mod init successful!")
	fmt.Println("Installing Cobra...")
	helpers.GoModCobra()
	fmt.Println("Running Go mod tidy...")
	cmd = exec.Command("go", "mod", "tidy")
	err = cmd.Run()
	helpers.CheckErrors(err)
	fmt.Println("Go mod tidy successful!")
	fmt.Println("Running Go fmt...")
	cmd = exec.Command("go", "fmt")
	err = cmd.Run()
	helpers.CheckErrors(err)
	fmt.Println("Go fmt successful!")
	fmt.Println("You're good to Go :)")
	helpers.GoFMT()
}
