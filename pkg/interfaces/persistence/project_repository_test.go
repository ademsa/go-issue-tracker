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

func TestPersistenceProjectNewSQLiteProjectRepository(t *testing.T) {
	mockDB, _, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteProjectRepository(gormDB)

	assert.NotNil(t, r)
}

func TestPersistenceProjectAdd(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteProjectRepository(gormDB)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO \"projects\" (.+)$").WithArgs("test-name", "test-description", sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	p := new(domain.Project)
	p.Name = "test-name"
	p.Description = "test-description"

	item, err := r.Add(p)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, p.Name, item.Name)
	assert.Equal(t, p.Description, item.Description)
	assert.NotNil(t, item.CreatedAt)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceProjectAddErr(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteProjectRepository(gormDB)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO \"projects\" (.+)$").WithArgs("test-name", "test-description", sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(errors.New("test error"))
	mock.ExpectRollback()

	p := new(domain.Project)
	p.Name = "test-name"
	p.Description = "test-description"

	item, err := r.Add(p)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceProjectUpdate(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteProjectRepository(gormDB)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE \"projects\" SET (.+)$").WithArgs("test-name", "test-description", sqlmock.AnyArg(), 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	p := domain.Project{
		ID:          uint(1),
		Name:        "test-name",
		Description: "test-description",
	}

	item, err := r.Update(p)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, p.Name, item.Name)
	assert.Equal(t, p.Description, item.Description)
	assert.NotNil(t, item.UpdatedAt)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceProjectUpdateErr(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteProjectRepository(gormDB)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE \"projects\" SET (.+)$").WithArgs("test-name", "test-description", sqlmock.AnyArg(), 1).WillReturnError(errors.New("test error"))
	mock.ExpectRollback()

	p := domain.Project{
		ID:          uint(1),
		Name:        "test-name",
		Description: "test-description",
	}

	item, err := r.Update(p)

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, p.Name, item.Name)
	assert.Equal(t, p.Description, item.Description)
	assert.NotNil(t, item.UpdatedAt)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceProjectFindByID(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteProjectRepository(gormDB)

	projectData := sqlmock.NewRows([]string{
		"id", "name", "description",
	}).AddRow(uint(1), "test-name", "test-description")
	mock.ExpectQuery("SELECT (.+) FROM \"projects\" WHERE (.+)$").WithArgs(1).WillReturnRows(projectData)

	item, err := r.FindByID(uint(1))

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, uint(1), item.ID)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceProjectFindByIDErr(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteProjectRepository(gormDB)

	mock.ExpectQuery("SELECT (.+) FROM \"projects\" WHERE (.+)$").WithArgs(1).WillReturnError(errors.New("test error"))

	item, err := r.FindByID(uint(1))

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, uint(0), item.ID)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceProjectFind(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteProjectRepository(gormDB)

	projectData := sqlmock.NewRows([]string{
		"id", "name", "description",
	}).AddRow("1", "test-name-1", "test-description").AddRow("2", "test-name-2", "test-description")
	mock.ExpectQuery("SELECT (.+) FROM \"projects\"").WillReturnRows(projectData)

	items, err := r.Find("test")

	assert.Nil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, 2, len(items))

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceProjectFindErr(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteProjectRepository(gormDB)

	mock.ExpectQuery("SELECT (.+) FROM \"projects\"").WillReturnError(errors.New("test error"))

	items, err := r.Find("test")

	assert.NotNil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, 0, len(items))

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceProjectFindAll(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteProjectRepository(gormDB)

	projectData := sqlmock.NewRows([]string{
		"id", "name", "description",
	}).AddRow("1", "test-name-1", "test-description").AddRow("2", "test-name-2", "test-description")
	mock.ExpectQuery("SELECT (.+) FROM \"projects\"").WillReturnRows(projectData)

	items, err := r.FindAll()

	assert.Nil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, 2, len(items))

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceProjectFindAllErr(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteProjectRepository(gormDB)

	mock.ExpectQuery("SELECT (.+) FROM \"projects\"").WillReturnError(errors.New("test error"))

	items, err := r.FindAll()

	assert.NotNil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, 0, len(items))

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceProjectRemove(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteProjectRepository(gormDB)

	cdata := sqlmock.NewRows([]string{
		"count",
	}).AddRow(0)
	mock.ExpectQuery("SELECT count(.+) FROM \"issues\" (.+)$").WithArgs(1).WillReturnRows(cdata)

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM \"projects\" WHERE (.+)$").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	status, err := r.Remove(uint(1))

	assert.Nil(t, err)
	assert.True(t, status)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceProjectRemoveIssueExistErr(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteProjectRepository(gormDB)

	cdata := sqlmock.NewRows([]string{
		"count",
	}).AddRow(1)
	mock.ExpectQuery("SELECT count(.+) FROM \"issues\" (.+)$").WithArgs(1).WillReturnRows(cdata)

	status, err := r.Remove(uint(1))

	assert.Nil(t, err)
	assert.False(t, status)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceProjectRemoveErr(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteProjectRepository(gormDB)

	cdata := sqlmock.NewRows([]string{
		"count",
	}).AddRow(0)
	mock.ExpectQuery("SELECT count(.+) FROM \"issues\" (.+)$").WithArgs(1).WillReturnRows(cdata)

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM \"projects\" WHERE (.+)$").WithArgs(1).WillReturnError(errors.New("test error"))
	mock.ExpectRollback()

	status, err := r.Remove(uint(1))

	assert.NotNil(t, err)
	assert.False(t, status)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}
