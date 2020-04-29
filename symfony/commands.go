package symfony

import (
	"fmt"
	"github.com/strangebuzz/cc/structs"
	"os"
	"os/exec"
)

const versionOption = "--version"
const cacheWarmup = "cache:warmup"
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
func RunCommand(config structs.Config, argument string) (string, error) {
	var out []byte
	var err error
	envOption := "--env=" + config.SymfonyEnv

	// weird: if the debugOption if empty then the debug is set to false
	if config.SymfonyDebug == true {
		out, err = exec.Command(config.SymfonyConsolePath, argument, envOption).CombinedOutput()
	} else {
		out, err = exec.Command(config.SymfonyConsolePath, argument, envOption, "--no-debug").CombinedOutput()
	}

	if err != nil {
		return "", nil
	}

	return fmt.Sprintf("%s", out), nil
}

/* ./bin/console --version */
func Version(config structs.Config) (string, error) {
	return RunCommand(config, versionOption)
}

/* ./bin/console cache:warmup */
func CacheWarmup(config structs.Config) (string, error) {
	return RunCommand(config, cacheWarmup)
}
