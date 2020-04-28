package main

import (
	"fmt"
	"os"
)

const version = "0.1.0"

func main() {
	welcome()

	argsWithProg := os.Args
	if len(argsWithProg) < 2 {
		panic(fmt.Errorf("you must provide the directory of the Symfony project to use"))
	}

	fmt.Println("> Directory: " + argsWithProg[1])
}

func welcome() {
	fmt.Println("Symfony cc v " + version)
}
