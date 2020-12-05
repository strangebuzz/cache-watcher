package structs

import "time"

// Symfony parameters, these are the best defaults for Symfony 5/Flex.
const EnvDefault = "dev"
const ConsolePath = "bin/console"
const DebugDefault = true
const ConfigDirectory = "config"
const TranslationsDirectory = "translations"
const TemplatesDir = "templates"
const EntitiesDir = "src/Entity"
const TemplatesExtension = "twig"
const YamlExtension = "yaml"
const PhpExtension = "php"

// Watcher process parameters
const SleepTime = 30

/*
 * Main application structure that contains all the parameters we weed. The YAML
 * meta data are the keys of the Symfony custom config file that will be used to
 * override the config default data.
 */
type Config struct {
	SymfonyProjectDir      string        `yaml:"project_dir"`         // The main Symfony project directory
	SymfonyConsolePath     string        `yaml:"console_path"`        // Relative path to the Symfony console
	SymfonyEnv             string        `yaml:"env"`                 // This is the APP_ENV parameter
	SymfonyDebug           bool          `yaml:"debug"`               // This is the APP_DEBUG parameter
	SymfonyConfigDir       string        `yaml:"config_dir"`          // Relative directory where are stored the configuration files
	SymfonyTranslationsDir string        `yaml:"translations_dir"`    // Relative directory where are stored the translations files
	SymfonyTemplatesDir    string        `yaml:"templates_dir"`       // Relative directory where are stored the templates files
	SymfonyEntitiesDir     string        `yaml:"entities_dir"`        // Relative directory where are stored the Doctrine entites
	TemplatesExtension     string        `yaml:"templates_extension"` // Default extension for templates files
	YamlExtension          string        `yaml:"yaml_extension"`      // Default extension for YAML files, we consider it must be consistent within an application
	PhpExtension           string        `yaml:"php_extension"`       // Default extension for PHP files (no be consistent with other parameters)
	SleepTime              time.Duration `yaml:"sleep_time"`          // Sleep time between two filesystem checks in milliseconds
}

func (obj *Config) Init() {
	obj.SymfonyConsolePath = ConsolePath
	obj.SymfonyEnv = EnvDefault
	obj.SymfonyDebug = DebugDefault
	obj.SymfonyConfigDir = ConfigDirectory
	obj.SymfonyTranslationsDir = TranslationsDirectory
	obj.SymfonyTemplatesDir = TemplatesDir
	obj.SymfonyEntitiesDir = EntitiesDir
	obj.TemplatesExtension = TemplatesExtension
	obj.YamlExtension = YamlExtension
	obj.PhpExtension = PhpExtension
	obj.SleepTime = SleepTime * time.Millisecond
}
