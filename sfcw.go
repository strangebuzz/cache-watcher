/**
 * The following Golang code must stay simple, readable and straightforward. No 🏭.
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

// pas besoin de constantes pour ça
const welcomeStr = "  [green]Symfony CC [white]version [yellow]v%s[white] by [blue]COil - https://www.strangebuzz.com 🐝 [white]"
const aboutStr = "Symfony CC watches your config files (.env, .yaml) and automatically refreshes your application cache."
const aboutStr2 = "(CTRL+C to stop watching)."

func main() {
	var config structs.Config
	var err error
	config.Init()

	// —— 1. Hello world! ——————————————————————————————————————————————————————
	welcome()

	// —— 2. No argument then display the help and exit ————————————————————————
	if len(os.Args) == 1 {
		help()
		os.Exit(0)
	}

	// —— 2.1. Otherwise at least an argument has been specified...

	// —— 3. Test if the specified directory exists ————————————————————————————
	config.SymfonyProjectDir, err = system.GetSymfonyProjectDir()
	if err != nil {
		tools.PrintError(fmt.Errorf("Project directory not found."))
		tools.PrintError(err)
		os.Exit(1)
	}
	_, _ = colorstring.Println(" > Project directory: [green]" + config.SymfonyProjectDir)

	// —— 4. Test if it is a Symfony project ———————————————————————————————————
	config.SymfonyConsolePath, err = symfony.CheckSymfonyConsole(config)
	if err != nil {
		tools.PrintError(fmt.Errorf("Symfony console not found."))
		tools.PrintError(err)
		os.Exit(1)
	}
	_, _ = colorstring.Println(" > Symfony console path: [green]" + config.SymfonyConsolePath)

	// —— 5. Test the Symfony console with the version command —————————————————
	out, err := symfony.Version(config)
	if err != nil {
		tools.PrintError(fmt.Errorf("Error while running the Symfony version command (rights problems?)."))
		tools.PrintError(err)
		os.Exit(1)
	}
	_, _ = colorstring.Println(" > Symfony env: [green]" + strings.Trim(fmt.Sprintf("%s", out), "\n"))

	// —— 6. Get the files to watch ————————————————————————————————————————————
	start := time.Now()
	filesToWatch, err := symfony.GetWatchMap(config)
	end := time.Now()
	elapsed := end.Sub(start)
	if err != nil {
		tools.PrintError(fmt.Errorf("Error while analysing the files to watch."))
		tools.PrintError(err)
		os.Exit(1)
	}
	_, _ = colorstring.Println(fmt.Sprintf(" > [yellow]%d [white]file(s) watched in [yellow]%s[white] in [yellow]%d[white] millisecond(s).", len(filesToWatch), config.SymfonyProjectDir, elapsed.Milliseconds()))

	// —— 6. Main loop —————————————————————————————————————————————————————————
	for {
		updatedFiles, _ := symfony.GetWatchMap(config)
		if !reflect.DeepEqual(filesToWatch, updatedFiles) {
			start := time.Now()
			_, _ = symfony.CacheWarmup(config) // handle errors
			end := time.Now()
			elapsed := end.Sub(start)
			_, _ = colorstring.Println(" [yellow] ⬇ Update detected[white] > refreshing cache...")
			_, _ = colorstring.Println(fmt.Sprintf("  [green]✅  Done![white] in [yellow]%.2f[white] second(s).", elapsed.Seconds()))
			filesToWatch = updatedFiles
		} else {
			time.Sleep(config.SleepTime) // What time to use to avoid overusing CPU?
		}
	}
}

func welcome() {
	fmt.Println(separator)
	_, _ = colorstring.Println(fmt.Sprintf(welcomeStr, version))
	fmt.Println(separator)
	fmt.Println(aboutStr)
	fmt.Println(aboutStr2)
	fmt.Println(separator)
}

func help() {
	fmt.Println("Call sfcw with the path of your Symfony project as the first argument.")
	_, _ = colorstring.Println("Example: \"[green]cc [yellow]../strangebuzz.com[white]\"")
	_, _ = colorstring.Println("Or: \"[green]cc [yellow].[white]\" if you call it from the root of your Symfony project directory.")
	_, _ = colorstring.Println("[yellow][💡][white] Add it to your $PATH if not already done.\n")
}
