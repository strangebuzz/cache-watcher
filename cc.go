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
	"github.com/strangebuzz/cc/system"
	"github.com/strangebuzz/cc/tools"
	"os"
	"reflect"
	"strings"
	"time"
)

const version = "0.1.0"
const separator = "——————————————————————————————————————————————————————————————————————"
const welcomeStr = "  [green]Symfony CC [white]version [yellow]v%s[white] by [blue]COil - https://www.strangebuzz.com 🐝 [white]"
const aboutStr = "Symfony CC watches your config files (.env, yaml) and automatically refreshes your application cache (CTRL+C to stop watching)."

func main() {
	var config structs.Config
	var err error
	config.Init()

	// —— 1. Hello world! ——————————————————————————————————————————————————————
	welcome()

	// —— 2. Test arguments ————————————————————————————————————————————————————
	argsWithProg := os.Args
	if len(argsWithProg) < 2 {
		tools.PrintError(fmt.Errorf("you must provide the directory of the Symfony project to use as the first argument"))
		tools.PrintError(fmt.Errorf("example: cc ../strangebuzz.com"))
		os.Exit(1)
	}

	// —— 3. Test if project directory exists ——————————————————————————————————
	config.SymfonyProjectDir, err = system.GetSymfonyProjectDir()
	if err != nil {
		tools.PrintError(err)
		os.Exit(1)
	}
	_, _ = colorstring.Println(" > Project Directory: [green]" + config.SymfonyProjectDir)

	// —— 4. Test if it is a Symfony project ———————————————————————————————————
	config.SymfonyConsolePath, err = symfony.CheckSymfonyConsole(config)
	if err != nil {
		tools.PrintError(err)
		os.Exit(1)
	}
	_, _ = colorstring.Println(" > Symfony console path: [green]" + config.SymfonyConsolePath)

	// —— 5. Test the Symfony console with the version command —————————————————
	out, err := symfony.Version(config)
	if err != nil {
		tools.PrintError(err)
		os.Exit(1)
	}
	_, _ = colorstring.Println(" > Symfony env: [green]" + strings.Trim(fmt.Sprintf("%s", out), "\n"))

	//// —— 6. Watch files —————————————————————————————————————————————————————
	filesToWatch, err := symfony.GetWatchMap(config)
	if err != nil {
		tools.PrintError(err)
		os.Exit(1)
	}
	_, _ = colorstring.Println(fmt.Sprintf(" > [yellow]%d [white]file(s) watched in [yellow]%s[white] and [yellow](.)/%s", len(filesToWatch), config.SymfonyProjectDir, symfony.ConfigDirectory))

	// —— 6. Main loop —————————————————————————————————————————————————————————
	for {
		updatedFiles, _ := symfony.GetWatchMap(config)
		if !reflect.DeepEqual(filesToWatch, updatedFiles) {
			_, _ = symfony.CacheWarmup(config)
			_, _ = colorstring.Println(" [yellow] ⬇ Update detected[white] > refreshing cache...")
			_, _ = colorstring.Println("  [green]✅  Done!")
			filesToWatch = updatedFiles
		} else {
			time.Sleep(config.SleepTime) // What time to use to avoid overusing CPU?
		}
	}
}

func welcome() {
	_, _ = colorstring.Println(separator)
	_, _ = colorstring.Println(fmt.Sprintf(welcomeStr, version))
	_, _ = colorstring.Println(separator)
	fmt.Println(aboutStr)
	_, _ = colorstring.Println(separator)
}

func help() {
}
