package symfony_test

import (
	"github.com/strangebuzz/cache-watcher/structs"
	"github.com/strangebuzz/cache-watcher/symfony"
	"github.com/strangebuzz/cache-watcher/system"
	"path/filepath"
	"testing"
	"time"
)

// Test against the Symfony 5 fixtures
// @see fixtures/symfony4/.sfcw.yaml
func TestCheckCustomConfig(t *testing.T) {
	config := structs.Config{}
	config.Init()
	execDir, err := system.GetExecDir()
	if err != nil {
		t.Errorf("Error while testing getFilesToWatch(): '%s'", err)
	}

	config.SymfonyProjectDir = filepath.Clean(execDir + "/../fixtures/symfony4")
	config, err = symfony.CheckCustomConfig(config)
	if err != nil {
		t.Errorf("CheckCustomConfig problem: '%s'", err)
	}

	expected := "app/config"
	if config.SymfonyConfigDir != expected {
		t.Errorf("CheckCustomConfig problem, expected config dir: '%s', found '%s' ", expected, config.SymfonyConfigDir)
	}

	expected = "src/AppBundle/Resources/translations"
	if config.SymfonyTranslationsDir != expected {
		t.Errorf("CheckCustomConfig problem, expected translations dir: '%s', found '%s' ", expected, config.SymfonyTranslationsDir)
	}

	expected = "src/AppBundle/Resources/views"
	if config.SymfonyTemplatesDir != expected {
		t.Errorf("CheckCustomConfig problem, expected translations dir: '%s', found '%s' ", expected, config.SymfonyTemplatesDir)
	}

	expected = "src/AppBundle/Entity"
	if config.SymfonyEntitiesDir != expected {
		t.Errorf("CheckCustomConfig problem, expected entities dir: '%s', found '%s' ", expected, config.SymfonyEntitiesDir)
	}

	expected = "yml"
	if config.YamlExtension != expected {
		t.Errorf("CheckCustomConfig problem, expected YAML extension: '%s', found '%s' ", expected, config.YamlExtension)
	}

	expected = "php7"
	if config.PhpExtension != expected {
		t.Errorf("CheckCustomConfig problem, expected PHP extension: '%s', found '%s' ", expected, config.PhpExtension)
	}

	var expectedDuration time.Duration = 50000000
	if config.SleepTime != expectedDuration {
		t.Errorf("CheckCustomConfig problem, expected sleep time: '%v', found '%v' ", expectedDuration, config.SleepTime)
	}
}

// Test against the Symfony 4 fixtures
func TestGetFilesToWatchSymfony4(t *testing.T) {
	config := structs.Config{}
	config.Init()
	config.SymfonyConfigDir = "app/config"
	config.SymfonyTranslationsDir = "src/AppBundle/Resources/translations"
	config.SymfonyTemplatesDir = "src/AppBundle/Resources/views"
	config.YamlExtension = "yml"

	execDir, err := system.GetExecDir()
	if err != nil {
		t.Errorf("Error while testing getFilesToWatch(): '%s'", err)
	}

	config.SymfonyProjectDir = filepath.Clean(execDir + "/../fixtures/symfony4")
	filesToWatch, err := symfony.GetFilesToWatch(config)

	if err != nil {
		t.Errorf("Error while testing getFilesToWatch(): '%s'", err)
	}

	// app/config:                           3 files
	// src/AppBundle/Resources/translations: 2 file
	// src/AppBundle/Resources/view:         1 files
	// ========
	expected := 6
	if expected != len(filesToWatch) {
		t.Errorf("[Symfony 4] getFilesToWatch() problem, expected: '%d' files found '%d' ", expected, len(filesToWatch))
	}
}

// Test against the Symfony 5 fixtures
func TestGetFilesToWatchSymfony5(t *testing.T) {
	config := structs.Config{}
	config.Init()
	execDir, err := system.GetExecDir()
	if err != nil {
		t.Errorf("Error while testing getFilesToWatch(): '%s'", err)
	}

	config.SymfonyProjectDir = filepath.Clean(execDir + "/../fixtures/symfony5")
	filesToWatch, err := symfony.GetFilesToWatch(config)
	if err != nil {
		t.Errorf("Error while testing getFilesToWatch(): '%s'", err)
	}

	// root:         1 file
	// config:       2 files
	// templates:    1 file
	// translations: 2 files
	// ========
	expected := 6
	if expected != len(filesToWatch) {
		t.Errorf("[symfony 5] getFilesToWatch() problem, expected: '%d' files found '%d' ", expected, len(filesToWatch))
	}
}
