package helpers_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go-issue-tracker/pkg/infrastructure/helpers"
	"os"
	"testing"
)

func TestGetProjectDirPath(t *testing.T) {
	path, err := helpers.GetProjectDirPath()

	assert.NotNil(t, path)
	assert.Nil(t, err)
}

func TestGetProjectDirPathErr(t *testing.T) {
	helpers.ExecutablePath = func() (string, error) {
		return "", errors.New("test error")
	}
	defer func() {
		helpers.ExecutablePath = os.Executable
	}()

	path, err := helpers.GetProjectDirPath()

	assert.Equal(t, "", path)
	assert.NotNil(t, err)
}
