package symfony

import (
	"fmt"
	"github.com/strangebuzz/cc/tools"
	"os"
)

/**
 * Get all the files who has to be watch.
 */
//func filesToWatch() {
//}

/**
 * Get the watched files as a simple map path => datetime.
 *
 * @example "/Users/coil/Sites/strangebuzz.com/config/services.yaml": "2020-04-28 16:33:44.73756727 +0200 CEST"
 */
func GetWatchMap(symfonyProjectDir string) (map[string]string, error) {
	watchMap := map[string]string{}

	// Example with one file
	serviceYamlFile := symfonyProjectDir + "/config/services.yaml"
	stats, err := os.Stat(serviceYamlFile)
	if err != nil {
		tools.PrintError(fmt.Errorf("can't get stats for the \"%s\" file, check the project permissions", err))
		os.Exit(-1)
	}

	watchMap[serviceYamlFile] = fmt.Sprintf("%s", stats.ModTime())

	return watchMap, nil
}
