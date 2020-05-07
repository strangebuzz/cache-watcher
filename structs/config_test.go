package structs_test

import (
	"github.com/strangebuzz/cc/structs"
	"testing"
)

// Test default values for Symfony 5
func TestConfig(t *testing.T) {
	var config structs.Config
	config.Init()

	expected := "dev"
	if expected != config.SymfonyEnv {
		t.Errorf("SymfonyEnv default value is not OK, expected '%s', found '%s'", expected, config.SymfonyEnv)
	}

	if true != config.SymfonyDebug {
		t.Errorf("SymfonyDebug default value is not OK, expected '%t', found '%t'", true, config.SymfonyDebug)
	}

	expected = "config"
	if expected != config.SymfonyConfigDir {
		t.Errorf("SymfonyConfigDir default value is not OK, expected '%s', found '%s'", expected, config.SymfonyConfigDir)
	}

	expected = "translations"
	if expected != config.SymfonyTranslationsDir {
		t.Errorf("TranslationsDirectory default value is not OK, expected '%s', found '%s'", expected, config.SymfonyTranslationsDir)
	}

	expected = "templates"
	if expected != config.SymfonyTemplatesDir {
		t.Errorf("SymfonyTemplatesDir default value is not OK, expected '%s', found '%s'", expected, config.SymfonyTemplatesDir)
	}

	expected = "twig"
	if expected != config.TemplatesExtension {
		t.Errorf("TemplatesExtension default value is not OK, expected '%s', found '%s'", expected, config.TemplatesExtension)
	}

	expected = "yaml"
	if expected != config.YamlExtension {
		t.Errorf("YamlExtension default value is not OK, expected '%s', found '%s'", expected, config.YamlExtension)
	}
}
