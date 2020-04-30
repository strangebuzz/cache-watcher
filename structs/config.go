package structs

import "time"

// Symfony parameters
const EnvDefault = "dev"
const DebugDefault = true
const ConfigDirectory = "config"
const TranslationsDirectory = "translations"

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
	SleepTime              time.Duration // Sleep time between two filesystem checks in milliseconds
}

func (obj *Config) Init() {
	obj.SymfonyEnv = EnvDefault
	obj.SymfonyDebug = DebugDefault // false by default
	obj.SymfonyConfigDir = ConfigDirectory
	obj.SymfonyTranslationsDir = TranslationsDirectory
	obj.SleepTime = SleepTime * time.Millisecond
}
