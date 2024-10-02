package muir

import (
	"fmt"
	"forge/helpers"
	"forge/str"
	"os"
	"os/exec"
)

func Blank(arg string, git bool) {
	str := str.BLANK

	err := os.Mkdir(arg, 0755)
	helpers.CheckErrors(err)
	err = os.Chdir(arg)
	helpers.CheckErrors(err)
	if git {
		Git()
	}
	cmd := exec.Command("go", "mod", "init", arg)
	err = cmd.Run()
	helpers.CheckErrors(err)

	file, err := os.Create("main.go")
	helpers.CheckErrors(err)
	defer file.Close()
	_, err = file.WriteString(str)
	helpers.CheckErrors(err)
	fmt.Println("Project created successfully!")
	cmd = exec.Command("go", "mod", "tidy")
	err = cmd.Run()
	helpers.CheckErrors(err)
	helpers.GoFMT()
	helpers.CheckErrors(err)
	fmt.Println("You're good to Go :)")
}

func Rest(arg string, git bool) {
	st := str.WEB

	err := os.Mkdir(arg, 0755)
	helpers.CheckErrors(err)
	err = os.Chdir(arg)
	helpers.CheckErrors(err)
	if git {
		Git()
	}
	cmd := exec.Command("go", "mod", "init", arg)
	err = cmd.Run()
	helpers.CheckErrors(err)

	fmt.Println("Creating middleware folder...")
	err = os.Mkdir("middleware", 0755)
	helpers.CheckErrors(err)
	os.Chdir("middleware")
	fmw, err := os.Create("middleware.go")
	helpers.CheckErrors(err)
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
	cmd = exec.Command("go", "mod", "tidy")
	err = cmd.Run()
	helpers.CheckErrors(err)
	cmd = exec.Command("go", "fmt")
	err = cmd.Run()
	helpers.CheckErrors(err)
	fmt.Println("You're good to Go :)")
	helpers.GoFMT()
}

func Git() {
	fmt.Println("Running Git init...")
	cmd := exec.Command("git", "init")
	err := cmd.Run()
	helpers.CheckErrors(err)
}

func Service(arg1 string) {
	fmt.Println("Creating: ", arg1, " service in the "+arg1+" folder")
	err := os.Mkdir(arg1, 0755)
	helpers.CheckErrors(err)
	err = os.Chdir(arg1)
	helpers.CheckErrors(err)
	pf := fmt.Sprintf("package %s", arg1)
	tf := fmt.Sprintf("type %sService struct {}", arg1)
	str := pf + "\n" + tf + "\n" + "// Add your service code here\n"
	str += fmt.Sprintf("func %sService(...interface{}) *%sService\n", arg1, arg1)
	str += "{\n"
	str += fmt.Sprintf("return &%sService{}\n", arg1)
	str += "}\n"
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
	file, err := os.Create("models.go")
	helpers.CheckErrors(err)
	defer file.Close()
	_, err = file.WriteString(str.MDL)
	helpers.CheckErrors(err)
	fmt.Println("You're good to Go :)")
	helpers.GoFMT()
}

func Cli(arg string, git bool) {
	st := str.CLI

	err := os.Mkdir(arg, 0755)
	helpers.CheckErrors(err)
	err = os.Chdir(arg)
	helpers.CheckErrors(err)
	if git {
		Git()
	}
	cmd := exec.Command("go", "mod", "init", arg)
	err = cmd.Run()
	helpers.CheckErrors(err)

	fmt.Println("Creating cmd folder...")
	err = os.Mkdir("cmd", 0755)
	helpers.CheckErrors(err)
	os.Chdir("cmd")
	file, err := os.Create("main.go")
	helpers.CheckErrors(err)
	defer file.Close()
	_, err = file.WriteString(str.CLIROOT)
	helpers.CheckErrors(err)

	os.Chdir("..")
	fmt.Println("Creating pkg folder...")
	os.Mkdir("pkg", 0755)

	file, err = os.Create("main.go")
	helpers.CheckErrors(err)
	defer file.Close()
	_, err = file.WriteString(st)
	helpers.CheckErrors(err)

	fmt.Println("Project created successfully!")
	fmt.Println("Installing Cobra...")
	helpers.GoModCobra()
	cmd = exec.Command("go", "mod", "tidy")
	err = cmd.Run()
	helpers.CheckErrors(err)
	cmd = exec.Command("go", "fmt")
	err = cmd.Run()
	helpers.CheckErrors(err)
	fmt.Println("You're good to Go :)")
	helpers.GoFMT()
}

func Python(arg string) {
	str := str.PY

	err := os.Mkdir(arg, 0755)
	helpers.CheckErrors(err)
	err = os.Chdir(arg)
	helpers.CheckErrors(err)
	cmd := exec.Command("python3", "-m", "venv", "venv")
	err = cmd.Run()
	helpers.CheckErrors(err)
	file, err := os.Create("app.py")
	helpers.CheckErrors(err)
	defer file.Close()
	_, err = file.WriteString(str)
	helpers.CheckErrors(err)
	fmt.Println("Project created successfully!")
}
