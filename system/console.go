package system

import (
	"os"
	"path/filepath"
)

func GetSymfonyProjectDir() (string, error) {
	argsWithProg := os.Args
	execDir, err := getExecDir()
	path := filepath.Clean(execDir + "/" + argsWithProg[1])

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