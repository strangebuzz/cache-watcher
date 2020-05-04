package symfony

import (
	"github.com/strangebuzz/cc/structs"
	"github.com/strangebuzz/cc/system"
	"path/filepath"
	"testing"
)

// Test against the Symfony 5 fixtures
func TestGetFilesToWatch(t *testing.T) {
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
	// translations: 2 files
	// ========
	expected := 5
	if expected != len(filesToWatch) {
		t.Errorf("[symfony 5] getFilesToWatch() problem, expected: '%d' files found '%d' ", expected, len(filesToWatch))
	}
}
