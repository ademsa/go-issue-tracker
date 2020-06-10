package testing

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"testing"
)

// GetMockedDB to get mocked database
func GetMockedDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *gorm.DB) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("mock db error %s", err)
	}

	gormDB, err := gorm.Open("sqlite3", mockDB)
	if err != nil {
		t.Fatalf("gorm mock db error %s", err)
	}

	gormDB.LogMode(true)

	return mockDB, mock, gormDB
}
