package symfony

import (
	"fmt"
	"github.com/strangebuzz/sfcw/structs"
	"os"
	"os/exec"
)

const versionOption = "--version"
const cacheWarmupArgument = "cache:warmup"

func CheckSymfonyConsole(config structs.Config) error {
	console := config.SymfonyProjectDir + "/" + config.SymfonyConsolePath
	_, err := os.Stat(console)
	if os.IsNotExist(err) {
		return err
	}

	return nil
}

/**
 * Run a Symfony command with a given argument or option.
 */
func RunCommand(config structs.Config, mainArgumentOrOption string) (string, error) {
	var out []byte
	var err error
	envOption := "--env=" + config.SymfonyEnv
	consoleFullPath := config.SymfonyProjectDir + "/" + config.SymfonyConsolePath

	// @fixme: if the debugOption (4th parameter) if empty then the debug is set to false, why?
	if config.SymfonyDebug == true {
		out, err = exec.Command(consoleFullPath, mainArgumentOrOption, envOption).CombinedOutput()
	} else {
		out, err = exec.Command(consoleFullPath, mainArgumentOrOption, envOption, "--no-debug").CombinedOutput()
	}

	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			return "", fmt.Errorf("The Symfony command didn't return a success code.")
		}
	}

	return fmt.Sprintf("%s", out), nil
}

/* ./bin/console --version --env=dev */
func Version(config structs.Config) (string, error) {
	return RunCommand(config, versionOption)
}

/* ./bin/console cache:warmup --env=dev */
func CacheWarmup(config structs.Config) (string, error) {
	return RunCommand(config, cacheWarmupArgument)
}
