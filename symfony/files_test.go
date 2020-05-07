package symfony

import (
	"github.com/strangebuzz/cc/structs"
	"github.com/strangebuzz/cc/system"
	"path/filepath"
	"testing"
)

// Test against the Symfony 5 fixtures
func TestGetFilesToWatchSymfony5(t *testing.T) {
	config := structs.Config{}
	config.Init()
	execDir, err := system.GetExecDir()
	if err != nil {
		t.Errorf("Error while testing getFilesToWatch(): '%s'", err)
	}

	config.SymfonyProjectDir = filepath.Clean(execDir + "/../fixtures/symfony5")
	filesToWatch, err := getFilesToWatch(config)
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
	filesToWatch, err := getFilesToWatch(config)

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
