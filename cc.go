/**
 * The following Golang code is and will stay very simple, readable and straightforward.
 * No 🏭 here.
 */
package main

import (
	"fmt"
	"github.com/mitchellh/colorstring"
	"github.com/strangebuzz/cc/structs"
	"github.com/strangebuzz/cc/symfony"
	"github.com/strangebuzz/cc/tools"
	"os"
	"path/filepath"
)

const version = "0.1.0"
const separator = "——————————————————————————————————————————————————————————————————————"
const welcomeStr = "  [green]Symfony CC [white]version [yellow]v%s[white] by [blue]COil - https://www.strangebuzz.com 🐝 [white]"
const aboutStr = "Symfony CC watches your config files (.env, yaml) and automatically refresh your application cache (CTRL+C to stop the process)."

func main() {
	var config structs.Config
	var err error

	welcome()

	// —— 1. Test arguments ————————————————————————————————————————————————————
	argsWithProg := os.Args
	if len(argsWithProg) < 2 {
		tools.PrintError(fmt.Errorf("you must provide the directory of the Symfony project to use as the first argument"))
		os.Exit(-1)
	}

	// —— 2. Test if project directory exists ——————————————————————————————————
	config.SymfonyProjectDir, err = getSymfonyProjectDir()
	if err != nil {
		tools.PrintError(err)
		os.Exit(-1)
	}
	_, _ = colorstring.Println(" > Project Directory: [green]" + config.SymfonyProjectDir)

	// —— 3. Test if it is a Symfony project ———————————————————————————————————
	config.SymfonyConsolePath, err = symfony.CheckSymfonyConsole(config)
	if err != nil {
		tools.PrintError(err)
		os.Exit(-1)
	}
	_, _ = colorstring.Println(" > Symfony console path: [green]" + config.SymfonyConsolePath)
	tools.PrettyPrint(config)
	//os.Exit(0)

	// —— 4. Test the Symfony console with the version command —————————————————
	//out, err := symfony.Version(consolePath)
	//if err != nil {
	//	tools.PrintError(err)
	//	os.Exit(-1)
	//}
	//_, _ = colorstring.Println(" > Symfony version: [green]" + strings.Trim(fmt.Sprintf("%s", out), "\n"))
	//
	//// —— 5. Watch files ———————————————————————————————————————————————————————
	//filesToWatch, err := symfony.GetWatchMap(symfonyProjectDir)
	//if err != nil {
	//	tools.PrintError(err)
	//	os.Exit(-1)
	//}
	//_, _ = colorstring.Println(fmt.Sprintf(" > [yellow]%d [white]file(s) watched in [yellow]%s[white] and ./[yellow]%s", len(filesToWatch), symfonyProjectDir, symfony.ConfigDirectory))
	//
	//// —— 6. Main process ——————————————————————————————————————————————————————
	//for {
	//	updatedFiles, _ := symfony.GetWatchMap(symfonyProjectDir)
	//	if !reflect.DeepEqual(filesToWatch, updatedFiles) {
	//		_, _ = colorstring.Println(" [yellow] ! Update detected[white] => refreshing cache...")
	//		_, _ = symfony.CacheWarmup(consolePath)
	//		_, _ = colorstring.Println("    > [green]Done!") // todo, diplay the time it took to refresh the cache
	//		filesToWatch = updatedFiles
	//	} else {
	//		time.Sleep(50 * time.Millisecond) // What time to use tp avoid overusing CPU?
	//	}
	//}
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
	_, _ = colorstring.Println(separator)
	_, _ = colorstring.Println(fmt.Sprintf(welcomeStr, version))
	_, _ = colorstring.Println(separator)
	fmt.Println(aboutStr)
}
