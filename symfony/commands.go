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
func RunCommand(console string, argument string) (string, error) {
	out, err := exec.Command(console, argument).CombinedOutput()
	if err != nil {
		return "", nil
	}

	return fmt.Sprintf("%s", out), nil
}

/* ./bin/console --version */
func Version(consolePath string) (string, error) {
	return RunCommand(consolePath, versionOption)
}

/* ./bin/console cache:warmup */
func CacheWarmup(consolePath string) (string, error) {
	return RunCommand(consolePath, cacheWarmup)
}
