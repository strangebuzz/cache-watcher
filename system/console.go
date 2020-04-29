package system

import (
	"os"
	"path/filepath"
)

func GetSymfonyProjectDir() (string, error) {
	execDir, err := getExecDir()
	path := os.Args[1]
	if !filepath.IsAbs(path) {
		path = filepath.Clean(execDir + "/" + path)
	}

	_, err = os.Stat(path)
	if err != nil {
		return "", err
	}

	if os.IsNotExist(err) {
		return "", err
	}

	return path, err
}

func getExecDir() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return path, nil
}
