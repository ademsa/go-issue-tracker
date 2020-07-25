package persistence_test

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go-issue-tracker/pkg/domain"
	"go-issue-tracker/pkg/interfaces/persistence"
	pTesting "go-issue-tracker/pkg/interfaces/persistence/testing"
	"testing"
)

func TestPersistenceLabelNewSQLiteLabelRepository(t *testing.T) {
	mockDB, _, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteLabelRepository(gormDB)

	assert.NotNil(t, r)
}

func TestPersistenceLabelAdd(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteLabelRepository(gormDB)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO \"labels\" (.+)$").WithArgs("test-name", "FFFFFF", sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	l := new(domain.Label)
	l.Name = "test-name"
	l.ColorHexCode = "FFFFFF"

	item, err := r.Add(l)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, l.Name, item.Name)
	assert.Equal(t, l.ColorHexCode, item.ColorHexCode)
	assert.NotNil(t, item.CreatedAt)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceLabelAddErr(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteLabelRepository(gormDB)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO \"labels\" (.+)$").WithArgs("test-name", "FFFFFF", sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(errors.New("test error"))
	mock.ExpectRollback()

	l := new(domain.Label)
	l.Name = "test-name"
	l.ColorHexCode = "FFFFFF"

	item, err := r.Add(l)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceLabelUpdate(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteLabelRepository(gormDB)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE \"labels\" SET (.+)$").WithArgs("test-name", "FFFFFF", sqlmock.AnyArg(), 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	l := domain.Label{
		ID:           uint(1),
		Name:         "test-name",
		ColorHexCode: "FFFFFF",
	}

	item, err := r.Update(l)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, l.Name, item.Name)
	assert.Equal(t, l.ColorHexCode, item.ColorHexCode)
	assert.NotNil(t, item.UpdatedAt)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceLabelUpdateErr(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteLabelRepository(gormDB)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE \"labels\" SET (.+)$").WithArgs("test-name", "FFFFFF", sqlmock.AnyArg(), 1).WillReturnError(errors.New("test error"))
	mock.ExpectRollback()

	l := domain.Label{
		ID:           uint(1),
		Name:         "test-name",
		ColorHexCode: "FFFFFF",
	}

	item, err := r.Update(l)

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, l.Name, item.Name)
	assert.Equal(t, l.ColorHexCode, item.ColorHexCode)
	assert.NotNil(t, item.UpdatedAt)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceLabelFindByID(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteLabelRepository(gormDB)

	labelData := sqlmock.NewRows([]string{
		"id", "name", "color_hex_code",
	}).AddRow(uint(1), "test-name", "FFFFFF")
	mock.ExpectQuery("SELECT (.+) FROM \"labels\" WHERE (.+)$").WithArgs(1).WillReturnRows(labelData)

	item, err := r.FindByID(uint(1))

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, uint(1), item.ID)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceLabelFindByIDErr(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteLabelRepository(gormDB)

	mock.ExpectQuery("SELECT (.+) FROM \"labels\" WHERE (.+)$").WithArgs(1).WillReturnError(errors.New("test error"))

	item, err := r.FindByID(uint(1))

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, uint(0), item.ID)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceLabelFindByName(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteLabelRepository(gormDB)

	labelData := sqlmock.NewRows([]string{
		"id", "name", "color_hex_code",
	}).AddRow(uint(1), "test-name", "FFFFFF")
	mock.ExpectQuery("SELECT (.+) FROM \"labels\" WHERE (.+)$").WithArgs("test-name").WillReturnRows(labelData)

	item, err := r.FindByName("test-name")

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, uint(1), item.ID)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceLabelFindByNameErr(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteLabelRepository(gormDB)

	mock.ExpectQuery("SELECT (.+) FROM \"labels\" WHERE (.+)$").WithArgs("test-name").WillReturnError(errors.New("test error"))

	item, err := r.FindByName("test-name")

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, uint(0), item.ID)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceLabelFind(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteLabelRepository(gormDB)

	labelData := sqlmock.NewRows([]string{
		"id", "name", "color_hex_code",
	}).AddRow("1", "test-name-1", "FFFFFF").AddRow("2", "test-name-2", "FFFFFF")
	mock.ExpectQuery("SELECT (.+) FROM \"labels\"").WillReturnRows(labelData)

	items, err := r.Find("test")

	assert.Nil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, 2, len(items))

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceLabelFindErr(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteLabelRepository(gormDB)

	mock.ExpectQuery("SELECT (.+) FROM \"labels\"").WillReturnError(errors.New("test error"))

	items, err := r.Find("test")

	assert.NotNil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, 0, len(items))

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceLabelFindAll(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteLabelRepository(gormDB)

	labelData := sqlmock.NewRows([]string{
		"id", "name", "color_hex_code",
	}).AddRow("1", "test-name-1", "FFFFFF").AddRow("2", "test-name-2", "FFFFFF")
	mock.ExpectQuery("SELECT (.+) FROM \"labels\"").WillReturnRows(labelData)

	items, err := r.FindAll()

	assert.Nil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, 2, len(items))

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceLabelFindAllErr(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteLabelRepository(gormDB)

	mock.ExpectQuery("SELECT (.+) FROM \"labels\"").WillReturnError(errors.New("test error"))

	items, err := r.FindAll()

	assert.NotNil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, 0, len(items))

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceLabelRemove(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteLabelRepository(gormDB)

	cdata := sqlmock.NewRows([]string{
		"count",
	}).AddRow(0)
	mock.ExpectQuery("SELECT count(.+) FROM \"issues_labels\" (.+)$").WithArgs(1).WillReturnRows(cdata)

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM \"labels\" WHERE (.+)$").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	status, err := r.Remove(uint(1))

	assert.Nil(t, err)
	assert.True(t, status)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceLabelRemoveIssueExistErr(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteLabelRepository(gormDB)

	cdata := sqlmock.NewRows([]string{
		"count",
	}).AddRow(1)
	mock.ExpectQuery("SELECT count(.+) FROM \"issues_labels\" (.+)$").WithArgs(1).WillReturnRows(cdata)

	status, err := r.Remove(uint(1))

	assert.Nil(t, err)
	assert.False(t, status)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceLabelRemoveErr(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteLabelRepository(gormDB)

	cdata := sqlmock.NewRows([]string{
		"count",
	}).AddRow(0)
	mock.ExpectQuery("SELECT count(.+) FROM \"issues_labels\" (.+)$").WithArgs(1).WillReturnRows(cdata)

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM \"labels\" WHERE (.+)$").WithArgs(1).WillReturnError(errors.New("test error"))
	mock.ExpectRollback()

	status, err := r.Remove(uint(1))

	assert.NotNil(t, err)
	assert.False(t, status)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}
