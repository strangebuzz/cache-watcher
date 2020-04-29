package structs

// Do we need this constants?
const EnvDefault = "dev"
const DebugDefault = true

/*
 * Main structure that contains all parameters we weed.
 *
 * @note Do we need the Symfony prefix?
 */
type Config struct {
	SymfonyProjectDir  string
	SymfonyConsolePath string
	SymfonyEnv         string
	SymfonyDebug       bool // default to false
}

func (obj *Config) Init() {
	obj.SymfonyEnv = EnvDefault
	obj.SymfonyDebug = DebugDefault
}
