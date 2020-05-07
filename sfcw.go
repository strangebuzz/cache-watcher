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

const accronym = "Sfcw"
const binary = "sfcw"
const version = "0.3.0"
const separator = "——————————————————————————————————————————————————————————————————————"
const repository = "https://github.com/strangebuzz/cc"

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

	// —— 4. Test if we have a custom config file and load it ——————————————————
	config, err = symfony.CheckCustomConfig(config)
	if err != nil {
		tools.PrintError(fmt.Errorf("Error while checking the custom config."))
		tools.PrintError(err)
		os.Exit(1)
	}

	// —— 5. Test if it is a Symfony project ———————————————————————————————————
	err = symfony.CheckSymfonyConsole(config)
	if err != nil {
		tools.PrintError(fmt.Errorf("Symfony console not found."))
		tools.PrintError(err)
		os.Exit(1)
	}
	_, _ = colorstring.Println(" > Symfony console path: [green]" + config.SymfonyConsolePath)

	// —— 6. Test the Symfony console with the version command —————————————————
	out, err := symfony.Version(config)
	if err != nil {
		tools.PrintError(fmt.Errorf("Error while running the Symfony version command."))
		tools.PrintError(err)
		os.Exit(1)
	}
	_, _ = colorstring.Println(" > Symfony env: [green]" + strings.Trim(fmt.Sprintf("%s", out), "\n"))

	// —— 7. Get the files to watch ————————————————————————————————————————————
	start := time.Now()
	filesToWatch, _ := symfony.GetWatchMap(config)
	//tools.PrettyPrint(filesToWatch)
	end := time.Now()
	elapsed := end.Sub(start)

	// —— 7.1 No error, but no file was found ——————————————————————————————————
	if len(filesToWatch) == 0 {
		errorNothingtoWatch()
	}
	_, _ = colorstring.Println(fmt.Sprintf(" > [yellow]%d [white]file(s) watched in [yellow]%s[white] in [yellow]%d[white] millisecond(s).", len(filesToWatch), config.SymfonyProjectDir, elapsed.Milliseconds()))
	_, _ = colorstring.Println(fmt.Sprintf(" > [green]CTRL+C[white] to stop watching or run [green]kill -9 [yellow]%d[white].", os.Getpid()))

	// —— 6. Main loop —————————————————————————————————————————————————————————
	mainLoop(config, filesToWatch)
}

func mainLoop(config structs.Config, filesToWatch map[string]string) {
	for {
		updatedFiles, _ := symfony.GetWatchMap(config)
		if !reflect.DeepEqual(filesToWatch, updatedFiles) {
			start := time.Now()
			_, _ = colorstring.Println(fmt.Sprintf(" [yellow] ⬇ Update detected[white] at [green]%s[white] > refreshing cache...", start.Format("15:04:05")))
			_, _ = symfony.CacheWarmup(config) // handle errors
			end := time.Now()
			elapsed := end.Sub(start)
			_, _ = colorstring.Println(fmt.Sprintf("  [green]✅  Done![white] in [yellow]%.2f[white] second(s).", elapsed.Seconds()))
			filesToWatch = updatedFiles
		} else {
			time.Sleep(config.SleepTime)
		}
	}
}

func errorNothingtoWatch() {
	tools.PrintError(fmt.Errorf("No file to watch found."))
	_, _ = colorstring.Println("[yellow][💡][white] If you are using an \"old\" Symfony project directory structure")
	_, _ = colorstring.Println("     you have to customize the watched directories with a [yellow].sfcw.yaml [white]file")
	_, _ = colorstring.Println(fmt.Sprintf("     at the root of your Symfony project. Check out the doc: [magenta]%s", repository))
	os.Exit(0)
}

func welcome() {
	fmt.Println(separator)
	_, _ = colorstring.Println(fmt.Sprintf("  [bold][green]%s [white]version [yellow]v%s[white] by [blue]COil - https://www.strangebuzz.com 🐝 [white]", accronym, version))
	fmt.Println(separator)
	_, _ = colorstring.Println(fmt.Sprintf("[green]%s[white] watches your files (.env, YAML, Twig) and automatically refreshes your application cache.", accronym))
	fmt.Println(separator)
}

func help() {
	_, _ = colorstring.Println(fmt.Sprintf("Call [green]%s[white] with the path of your Symfony project as the first argument.", binary))
	_, _ = colorstring.Println(fmt.Sprintf("Example: \"[green]%s [yellow]../strangebuzz.com[white]\"", binary))
	_, _ = colorstring.Println(fmt.Sprintf("Or even: \"[green]%s [yellow].[white]\" if you call it from the root of your Symfony project directory.", binary))
	_, _ = colorstring.Println("[yellow][💡][white] Add it to your $PATH if not done already.")
}
