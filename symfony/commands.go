package symfony

import (
	"fmt"
	"os/exec"
)

const versionOption = "--version"
const cacheWarmup = "cache:warmup"

/**
 * @todo permetrre de passer un environnement.
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
