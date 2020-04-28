package main

import (
	"fmt"
	"github.com/mitchellh/colorstring"
	"github.com/strangebuzz/cc/tools"
	"os"
	"os/exec"
	"path/filepath"
)

const version = "0.1.0"
const welcomeStr = "[green]Symfony CC [white]version [yellow]v%s[white] by [blue]COil/Strangebuzz.com"
const aboutStr = "Symfony CC watches your config files (.env, yaml) and automatically refresh your application cache."

func main() {
	// —— Welcome users ————————————————————————————————————————————————————————
	welcome()

	// —— Test arguments ———————————————————————————————————————————————————————
	argsWithProg := os.Args
	if len(argsWithProg) < 2 {
		tools.PrintError(fmt.Errorf("you must provide the directory of the Symfony project to use"))
		os.Exit(-1)
	}

	// —— Test if project directory exists —————————————————————————————————————
	symfonyProjectDir, err := getSymfonyProjectDir()
	if err != nil {
		tools.PrintError(err)
		os.Exit(-1)
	}
	_, _ = colorstring.Println(" > Project Directory: [green]" + symfonyProjectDir)

	// —— Test if it's a Symfony project ———————————————————————————————————————
	console := symfonyProjectDir + "/bin/console"
	_, err = os.Stat(console)
	if err != nil {
		tools.PrintError(err)
		os.Exit(-1)
	}
	if os.IsNotExist(err) {
		tools.PrintError(err)
		os.Exit(-1)
	}
	_, _ = colorstring.Println(" > Symfony console: [green]" + console)

	getVersionCommand := console
	out, err := exec.Command(getVersionCommand, "--version").CombinedOutput()
	if err != nil {
		tools.PrintError(err)
		os.Exit(-1)
	}
	_, _ = colorstring.Println(" > Symfony version: [green]" + fmt.Sprintf("%s", out))
}

func getSymfonyProjectDir() (string, error) {
	argsWithProg := os.Args
	path := filepath.Clean(getExecDir() + "/" + argsWithProg[1])

	_, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	if os.IsNotExist(err) {
		return "", err
	}

	return path, err
}

func getExecDir() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return path
}

func welcome() {
	_, _ = colorstring.Println(fmt.Sprintf(welcomeStr, version))
	fmt.Println(aboutStr)
}
