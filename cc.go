package main

import (
	"fmt"
	"github.com/mitchellh/colorstring"
	"os"
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

	fmt.Println("> Directory: " + argsWithProg[1])
}

func welcome() {
	_, err := colorstring.Println(fmt.Sprintf(welcomeStr, version))
	if err != nil {
		panic(err)
	}
	fmt.Println(aboutStr)
}
