package helpers

import (
	"os"
	"path/filepath"
)

// ExecutablePath is alias for os.Executable
var ExecutablePath = os.Executable

// GetProjectDirPath to get project's root dir path
func GetProjectDirPath() (string, error) {
	path, err := ExecutablePath()
	if err != nil {
		return "", err
	}
	return filepath.Dir(filepath.Dir(path)), nil
}
