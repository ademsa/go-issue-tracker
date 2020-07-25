package database_test

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go-issue-tracker/pkg/infrastructure/database"
	"go-issue-tracker/pkg/infrastructure/helpers"
	"os"
	"testing"
)

func TestGetDefaultSQLiteDBFilePath(t *testing.T) {
	path, err := database.GetDefaultSQLiteDBFilePath()

	assert.Nil(t, err)
	assert.NotNil(t, path)
}

func TestGetDefaultSQLiteDBFilePathErr(t *testing.T) {
	helpers.ExecutablePath = func() (string, error) {
		return "", errors.New("test error")
	}
	defer func() {
		helpers.ExecutablePath = os.Executable
	}()

	path, err := database.GetDefaultSQLiteDBFilePath()

	assert.NotNil(t, err)
	assert.Equal(t, "", path)
}

func TestGetSQLiteDB(t *testing.T) {
	mockDB, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("mock db error %s", err)
	}

	db, err := database.GetSQLiteDB(mockDB)

	assert.Nil(t, err)
	assert.NotNil(t, db)
}

func TestGetSQLiteDBErr(t *testing.T) {
	mock := new(interface{})

	db, err := database.GetSQLiteDB(mock)

	assert.NotNil(t, err)
	assert.Nil(t, db)
}
