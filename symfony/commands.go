package symfony

import (
	"fmt"
	"github.com/strangebuzz/cc/structs"
	"os"
	"os/exec"
)

const versionOption = "--version"
const cacheWarmupArgument = "cache:warmup"
const consoleRelativePath = "bin/console"

func CheckSymfonyConsole(config structs.Config) (string, error) {
	console := config.SymfonyProjectDir + "/" + consoleRelativePath
	_, err := os.Stat(console)
	if err != nil {
		return "", err
	}
	if os.IsNotExist(err) {
		return "", err
	}

	return console, nil
}

/**
 * @todo permetre de passer un environnement.
 */
func RunCommand(config structs.Config, mainArgumentOrOption string) (string, error) {
	var out []byte
	var err error
	envOption := "--env=" + config.SymfonyEnv

	// @fixme: if the debugOption if empty then the debug is set to false
	if config.SymfonyDebug == true {
		out, err = exec.Command(config.SymfonyConsolePath, mainArgumentOrOption, envOption).CombinedOutput()
	} else {
		out, err = exec.Command(config.SymfonyConsolePath, mainArgumentOrOption, envOption, "--no-debug").CombinedOutput()
	}

	if err != nil {
		return "", nil
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
