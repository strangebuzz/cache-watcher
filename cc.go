package main

import (
	"fmt"
	"github.com/mitchellh/colorstring"
	"github.com/strangebuzz/cc/tools"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const version = "0.1.0"
const welcomeStr = "[green]Symfony CC [white]version [yellow]v%s[white] by [blue]COil - https://www.strangebuzz.com 🐝"
const aboutStr = "Symfony CC watches your config files (.env, yaml) and automatically refresh your application cache."

// —— Now comes the constants we will be able to transform into parameters later
const consolePath = "bin/console"
const versionArgument = "--version"

func main() {
	welcome()

	// —— 1. Test arguments ————————————————————————————————————————————————————
	argsWithProg := os.Args
	if len(argsWithProg) < 2 {
		tools.PrintError(fmt.Errorf("you must provide the directory of the Symfony project to use"))
		os.Exit(-1)
	}

	// —— 2. Test if project directory exists ——————————————————————————————————
	symfonyProjectDir, err := getSymfonyProjectDir()
	if err != nil {
		tools.PrintError(err)
		os.Exit(-1)
	}
	_, _ = colorstring.Println(" > Project Directory: [green]" + symfonyProjectDir)

	// —— 3. Test if it is a Symfony project ———————————————————————————————————
	console, err := checkSymfonyConsole(symfonyProjectDir)
	if err != nil {
		tools.PrintError(err)
		os.Exit(-1)
	}
	_, _ = colorstring.Println(" > Symfony console: [green]" + console)

	// —— 4. Test the Symfony console with the version command —————————————————
	out, err := exec.Command(console, versionArgument).CombinedOutput()
	if err != nil {
		tools.PrintError(err)
		os.Exit(-1)
	}
	_, _ = colorstring.Println(" > Symfony version: [green]" + strings.Trim(fmt.Sprintf("%s", out), ""))

	// OK everything seems, ok now lets grab the files to watch
	md5, err := tools.HashFileMd5(console)
	_, _ = colorstring.Println(" > Symfony console md5: [green]" + md5)
}

func checkSymfonyConsole(symfonyProjectDir string) (string, error) {
	console := symfonyProjectDir + "/" + consolePath
	_, err := os.Stat(console)
	if err != nil {
		return "", err
	}
	if os.IsNotExist(err) {
		return "", err
	}

	return console, nil
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
