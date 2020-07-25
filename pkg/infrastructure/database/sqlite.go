package database

import (
	"github.com/jinzhu/gorm"
	"go-issue-tracker/pkg/domain"
	"go-issue-tracker/pkg/infrastructure/helpers"
	"path/filepath"
)

// GetDefaultSQLiteDBFilePath to get db default path
func GetDefaultSQLiteDBFilePath() (string, error) {
	path, err := helpers.GetProjectDirPath()
	if err != nil {
		return "", err
	}

	return filepath.Join(path, "data/db.sqlite3"), nil
}

// GetSQLiteDB to get DB
func GetSQLiteDB(path interface{}) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	db.AutoMigrate(&domain.Issue{})
	db.AutoMigrate(&domain.Label{})
	db.AutoMigrate(&domain.Project{})

	return db, nil
}
