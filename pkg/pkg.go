package muir

import (
	"fmt"
	"forge/helpers"
	"forge/str"
	"os"
	"os/exec"
	"path/filepath"
)

// Define constants for repeated strings
const (
	MainGoFile       = "main.go"
	ModelsFolder     = "models"
	HelpersFolder    = "helpers"
	MiddlewareFolder = "middleware"
	VenvFolder       = "venv"
	PycacheFolder    = "__pycache__"
)

// Helper function to create a directory
func createDir(path string) error {
	return os.Mkdir(path, 0755)
}

// Helper function to change the current working directory
func changeDir(path string) error {
	return os.Chdir(path)
}

// Helper function to run a command
func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	return cmd.Run()
}

// Helper function to create a file and write content
func createFileWithContent(path, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	return err
}

func Service(arg1 string) {
	fmt.Println("Creating: ", arg1, " service in the "+arg1+" folder")
	err := createDir(arg1)
	helpers.CheckErrors(err)
	err = changeDir(arg1)
	helpers.CheckErrors(err)
	pf := fmt.Sprintf("package %s", arg1)
	tf := fmt.Sprintf("type %sService struct {}", arg1)
	str := pf + "\n" + tf + "\n" + "// Add your service code here\n"
	str += fmt.Sprintf("func %sService(...interface{}) *%sService\n", arg1, arg1)
	str += "{\n"
	str += fmt.Sprintf("return &%sService{}\n", arg1)
	str += "}\n"
	err = createFileWithContent(arg1+".service.go", str)
	helpers.CheckErrors(err)
	helpers.GoFMT()
	fmt.Println("Service created successfully!")
	fmt.Println("You're good to Go :)")
}

func Models() {
	fmt.Println("Creating models folder...")
	err := createDir(ModelsFolder)
	helpers.CheckErrors(err)
	changeDir(ModelsFolder)
	err = createFileWithContent("models.go", str.MDL)
	helpers.CheckErrors(err)
	fmt.Println("You're good to Go :)")
	helpers.GoFMT()
}

func Cli(arg string, git bool) {
	st := str.CLI

	err := createDir(arg)
	helpers.CheckErrors(err)
	err = changeDir(arg)
	helpers.CheckErrors(err)
	if git {
		Git()
	}
	err = runCommand("go", "mod", "init", arg)
	helpers.CheckErrors(err)

	fmt.Println("Creating cmd folder...")
	err = createDir("cmd")
	helpers.CheckErrors(err)
	changeDir("cmd")
	err = createFileWithContent(MainGoFile, str.CLIROOT)
	helpers.CheckErrors(err)

	changeDir("..")
	fmt.Println("Creating pkg folder...")
	createDir("pkg")

	err = createFileWithContent(MainGoFile, st)
	helpers.CheckErrors(err)

	fmt.Println("Project created successfully!")
	fmt.Println("Installing Cobra...")
	helpers.GoModCobra()
	err = runCommand("go", "mod", "tidy")
	helpers.CheckErrors(err)
	err = runCommand("go", "fmt")
	helpers.CheckErrors(err)
	fmt.Println("You're good to Go :)")
	helpers.GoFMT()
}

func Python(arg string) {
	str := str.PY

	err := createDir(arg)
	helpers.CheckErrors(err)
	err = changeDir(arg)
	helpers.CheckErrors(err)
	err = runCommand("python3", "-m", "venv", VenvFolder)
	helpers.CheckErrors(err)
	err = createFileWithContent("app.py", str)
	helpers.CheckErrors(err)
	fmt.Println("Project created successfully!")
}

// PrintStructure prints the folder structure of a given path
// path: the root directory to start printing from
// prefix: the prefix to use for formatting the output
func PrintStructure(path string, prefix string) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if !fileInfo.IsDir() {
		// Print file with "└──" (or other symbols depending on context)
		fmt.Println(prefix + "└── " + fileInfo.Name())
		return
	}

	base := filepath.Base(path)
	// Skip hidden directories or specific directories like "venv" and "__pycache__"
	if base[0] == '.' || base == VenvFolder || base == PycacheFolder {
		return
	}

	// Print the directory name
	fmt.Println(prefix + base + "/")

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Iterate through directory contents
	for i, file := range files {
		isLast := i == len(files)-1
		newPrefix := prefix
		if isLast {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}

		if file.IsDir() {
			PrintStructure(filepath.Join(path, file.Name()), newPrefix)
		} else {
			fmt.Println(newPrefix + "└── " + file.Name())
		}
	}
}
