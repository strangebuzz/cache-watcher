/**
 * A renommer an files.go
 */

package symfony

import (
	"fmt"
	"github.com/strangebuzz/cc/structs"
	"github.com/strangebuzz/cc/tools"
	"os"
	"path/filepath"
)

const ConfigDirectory = "config"

/**
 * Get all the files who has to be watch.
 *
 * @todo Loop on a list of Glob patterns.
 * @todo Allow to set the max nesting level to handle.
 */
func getFilesToWatch(symfonyProjectDir string) ([]string, error) {
	var fileToWatch []string

	// Level 0 : .env files
	fileToWatch, err := filepath.Glob(fmt.Sprintf("%s/.env*", symfonyProjectDir))
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	// Level 1
	level1, err := filepath.Glob(fmt.Sprintf("%s/%s/*.yaml", symfonyProjectDir, ConfigDirectory))
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fileToWatch = append(fileToWatch, level1...)

	// Level 2
	level2, err := filepath.Glob(fmt.Sprintf("%s/%s/*/*.yaml", symfonyProjectDir, ConfigDirectory))
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fileToWatch = append(fileToWatch, level2...)

	// Level 3
	level3, err := filepath.Glob(fmt.Sprintf("%s/%s/*/*/*.yaml", symfonyProjectDir, ConfigDirectory))
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fileToWatch = append(fileToWatch, level3...)

	return fileToWatch, nil
}

/**
 * Get the watched files as a simple map: "path" => "datetime"
 *
 * @example "/Users/coil/Sites/strangebuzz.com/config/services.yaml": "2020-04-28 16:33:44.73756727 +0200 CEST"
 */
func GetWatchMap(config structs.Config) (map[string]string, error) {
	watchMap := map[string]string{}
	filesTowatch, _ := getFilesToWatch(config.SymfonyProjectDir)

	for _, file := range filesTowatch {
		stats, err := os.Stat(file)
		if err != nil {
			tools.PrintError(fmt.Errorf("Can't get stats for the \"%s\" file, check the project permissions.", err))
			os.Exit(1)
		}
		watchMap[file] = fmt.Sprintf("%s", stats.ModTime())
	}

	return watchMap, nil
}
