package main

import (
	"fmt"
	"github.com/mitchellh/colorstring"
	"os"
	"path/filepath"
)

const version = "0.1.0"
const welcomeStr = "[green]Symfony CC [white]version [yellow]v%s[white] by [blue]COil/Strangebuzz.com"
const aboutStr = "Symfony CC watches your config files and automatically refresh your application cache."

func main() {
	welcome()
	argsWithProg := os.Args
	if len(argsWithProg) < 2 {
		panic(fmt.Errorf("you must provide the directory of the Symfony project to use"))
	}
	_, _ = colorstring.Println(" > Project Directory: [green]" + getSymfonyProjectDir())
}

func welcome() {
	_, _ = colorstring.Println(fmt.Sprintf(welcomeStr, version))
	fmt.Println(aboutStr)
}

func getSymfonyProjectDir() string {
	argsWithProg := os.Args
	path := filepath.Clean(getExecDir() + "/" + argsWithProg[1])

	_, err := os.Stat(path)
	if err != nil {
		_, _ = colorstring.Println(fmt.Sprintf("[red]/!\\ %s.", path))
		os.Exit(-1)
	}

	if os.IsNotExist(err) {
		_, _ = colorstring.Println(fmt.Sprintf("[red]/!\\ %s.", err))
		os.Exit(-1)
	}

	return path
}

func getExecDir() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return path
}
