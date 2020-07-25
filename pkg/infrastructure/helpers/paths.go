package helpers

import (
	"os"
	"path/filepath"
)

// ExecutablePath is alias for os.Executable
var ExecutablePath = os.Executable

// Stat alias
var Stat = os.Stat

// GetProjectDirPath to get project's root dir path
func GetProjectDirPath() (string, error) {
	path, err := ExecutablePath()
	if err != nil {
		return "", err
	}
	return filepath.Dir(filepath.Dir(path)), nil
}

// PathExists to check if path exists
func PathExists(path string) bool {
	if _, err := Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
