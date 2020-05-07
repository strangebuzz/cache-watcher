package symfony

import (
	"fmt"
	"github.com/strangebuzz/cc/structs"
	"github.com/strangebuzz/cc/tools"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

const customConfigFilename = ".sfcw.yaml"

/**
 * Get and load the custom config if it exists.
 */
func CheckCustomConfig(config structs.Config) (structs.Config, error) {
	customConfigFilepath := config.SymfonyProjectDir + "/" + customConfigFilename
	_, err := os.Stat(customConfigFilepath)
	if os.IsNotExist(err) {
		fmt.Println("Custom config file not found.")
		return config, nil
	}

	fmt.Println("Custom config file found!")

	yamlFile, err := ioutil.ReadFile(customConfigFilepath)
	if err != nil {
		fmt.Print(fmt.Errorf("Error when trying to load the %s file: #%v ", customConfigFilename, err))
		os.Exit(1)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Print(fmt.Errorf("Error when reading the custom file values: #%v ", err))
		os.Exit(1)
	}

	// Convert milliseconds of user setting to the duration unit
	if config.SleepTime != structs.SleepTime {
		config.SleepTime = config.SleepTime * time.Millisecond
	}

	return config, nil
}

/**
 * Get all the files of the Symfony application which have to be watched.
 *
 * @todo handle nesting level? Add a parameter for that?
 */
func getFilesToWatch(config structs.Config) ([]string, error) {
	var filesToWatch []string

	// 1) ".env" files at the root of the project
	filesToWatch = append(filesToWatch, getFilesFromPath(config, ".env*")...)

	// 2) Yaml files in "config/"...
	filesToWatch = append(filesToWatch, getFilesFromPath(config, fmt.Sprintf("%s/*.%s", config.SymfonyConfigDir, config.YamlExtension))...)
	filesToWatch = append(filesToWatch, getFilesFromPath(config, fmt.Sprintf("%s/*/*.%s", config.SymfonyConfigDir, config.YamlExtension))...)
	filesToWatch = append(filesToWatch, getFilesFromPath(config, fmt.Sprintf("%s/*/*/*.%s", config.SymfonyConfigDir, config.YamlExtension))...)

	// 3) Yaml files in "translations/"...
	filesToWatch = append(filesToWatch, getFilesFromPath(config, fmt.Sprintf("%s/*.%s", config.SymfonyTranslationsDir, config.YamlExtension))...)
	filesToWatch = append(filesToWatch, getFilesFromPath(config, fmt.Sprintf("%s/*/*.%s", config.SymfonyTranslationsDir, config.YamlExtension))...)

	// 4) Twig files in "templates/"...
	filesToWatch = append(filesToWatch, getFilesFromPath(config, fmt.Sprintf("%s/*.%s", config.SymfonyTemplatesDir, config.TemplatesExtension))...)
	filesToWatch = append(filesToWatch, getFilesFromPath(config, fmt.Sprintf("%s/*/*.%s", config.SymfonyTemplatesDir, config.TemplatesExtension))...)
	filesToWatch = append(filesToWatch, getFilesFromPath(config, fmt.Sprintf("%s/*/*/*.%s", config.SymfonyTemplatesDir, config.TemplatesExtension))...)
	filesToWatch = append(filesToWatch, getFilesFromPath(config, fmt.Sprintf("%s/*/*/*/*.%s", config.SymfonyTemplatesDir, config.TemplatesExtension))...)

	return filesToWatch, nil
}

/**
 * Get all files corresponding to a glob pattern. To simplify the code we don't
 * return an error. The process will stop if an error is encountered.
 */
func getFilesFromPath(config structs.Config, glob string) []string {
	files, err := filepath.Glob(fmt.Sprintf("%s/%s", config.SymfonyProjectDir, glob))
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	return files
}

/**
 * Get the watched files as a simple map: "path" => "datetime"
 *
 * @example "/Users/coil/Sites/strangebuzz.com/config/services.yaml": "2020-04-28 16:33:44.73756727 +0200 CEST"
 */
func GetWatchMap(config structs.Config) (map[string]string, error) {
	watchMap := map[string]string{}
	filesTowatch, _ := getFilesToWatch(config)

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
