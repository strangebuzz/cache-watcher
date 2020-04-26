package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args
	if len(argsWithProg) < 2 {
		panic(fmt.Errorf("you must provide the directory of the Symfony project to use"))
	}

	fmt.Println("Directoty: " + argsWithProg[1])
}
