package structs

import "time"

// Symfony parameters, these are the best defaults for Symfony 5/Flex.
const EnvDefault = "dev"
const DebugDefault = true
const ConfigDirectory = "config"
const TranslationsDirectory = "translations"
const TemplatesDir = "templates"
const TemplatesExtension = "twig"
const YamlExtension = "yaml"

// Watcher process parameters
const SleepTime = 30

/*
 * Main application structure that contains all the parameters we weed.
 *
 * @see Init()
 */
type Config struct {
	SymfonyProjectDir      string        // The main Symfony project directory
	SymfonyConsolePath     string        // Full path to "bin/console"
	SymfonyEnv             string        // This is the APP_ENV parameter of the Symfony application
	SymfonyDebug           bool          // This is the APP_DEBUG parameter of the Symfony application
	SymfonyConfigDir       string        // Relative directory where are stored the configuration files of the Symfony application
	SymfonyTranslationsDir string        // Relative directory where are stored the translations files of the Symfony application
	SymfonyTemplatesDir    string        // Relative directory where are stored the templates files of the Symfony application
	TemplatesExtension     string        // Default extension for templates files
	YamlExtension          string        // Default extension for YAML files, we consider it must be consistent within an application
	SleepTime              time.Duration // Sleep time between two filesystem checks in milliseconds
}

func (obj *Config) Init() {
	obj.SymfonyEnv = EnvDefault
	obj.SymfonyDebug = DebugDefault // false by default
	obj.SymfonyConfigDir = ConfigDirectory
	obj.SymfonyTranslationsDir = TranslationsDirectory
	obj.SymfonyTemplatesDir = TemplatesDir
	obj.TemplatesExtension = TemplatesExtension
	obj.YamlExtension = YamlExtension
	obj.SleepTime = SleepTime * time.Millisecond
}
