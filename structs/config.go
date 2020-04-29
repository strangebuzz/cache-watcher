package structs

/*
 * Main structure that contains all parameters we weed.
 *
 * @tothink Do we need the Symfony prefix?
 */
type Config struct {
	SymfonyProjectDir  string
	SymfonyConsolePath string
	SymfonyEnv         string
	SymfonyDebug       bool
}
