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

func TestPersistenceIssueNewSQLiteIssueRepository(t *testing.T) {
	mockDB, _, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteIssueRepository(gormDB)

	assert.NotNil(t, r)
}

func TestPersistenceIssueAdd(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteIssueRepository(gormDB)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO \"issues\" (.+)$").WithArgs("test-title", "test-description", 1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	p := new(domain.Issue)
	p.Title = "test-title"
	p.Description = "test-description"
	p.Status = 1
	p.ProjectID = 1

	item, err := r.Add(p)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, p, item)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceIssueAddErr(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteIssueRepository(gormDB)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO \"issues\" (.+)$").WithArgs("test-title", "test-description", 1, 1).WillReturnError(errors.New("test error"))
	mock.ExpectRollback()

	p := new(domain.Issue)
	p.Title = "test-title"
	p.Description = "test-description"
	p.Status = 1
	p.ProjectID = 1

	item, err := r.Add(p)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceIssueUpdate(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteIssueRepository(gormDB)

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM \"issues_labels\" WHERE (.+)$").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE \"issues\" SET (.+)$").WithArgs("test-title", "test-description", 1, 1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	p := domain.Issue{
		ID:          uint(1),
		Title:       "test-title",
		Description: "test-description",
		Status:      1,
		ProjectID:   1,
	}

	item, err := r.Update(p)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, p, item)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceIssueUpdateReplaceErr(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteIssueRepository(gormDB)

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM \"issues_labels\" WHERE (.+)$").WithArgs(1).WillReturnError(errors.New("test error"))
	mock.ExpectRollback()

	p := domain.Issue{
		ID:          uint(1),
		Title:       "test-title",
		Description: "test-description",
		Status:      1,
		ProjectID:   1,
	}

	item, err := r.Update(p)

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, p, item)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceIssueUpdateErr(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteIssueRepository(gormDB)

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM \"issues_labels\" WHERE (.+)$").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE \"issues\" SET (.+)$").WithArgs("test-title", "test-description", 1, 1, 1).WillReturnError(errors.New("test error"))
	mock.ExpectRollback()

	p := domain.Issue{
		ID:          uint(1),
		Title:       "test-title",
		Description: "test-description",
		Status:      1,
		ProjectID:   1,
	}

	item, err := r.Update(p)

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, p, item)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceIssueFindByID(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteIssueRepository(gormDB)
	issueData := sqlmock.NewRows([]string{
		"id", "title", "description", "status", "project_id",
	}).AddRow(uint(1), "test-title", "test-description", 1, 1)
	mock.ExpectQuery("SELECT (.+) FROM \"issues\" (.+)$").WithArgs(1).WillReturnRows(issueData)

	projectData := sqlmock.NewRows([]string{
		"id", "name", "description",
	}).AddRow(uint(1), "test-name", "test-description")
	mock.ExpectQuery("SELECT (.+) FROM \"projects\" (.+)$").WithArgs(1).WillReturnRows(projectData)
	labelData := sqlmock.NewRows([]string{
		"id", "name", "color_hex_code",
	}).AddRow(uint(1), "test-name", "FFFFFF")
	mock.ExpectQuery("SELECT (.+) FROM \"labels\" (.+)$").WithArgs(1).WillReturnRows(labelData)

	item, err := r.FindByID(uint(1))

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, uint(1), item.ID)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceIssueFindByIDErr(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteIssueRepository(gormDB)

	mock.ExpectQuery("SELECT (.+) FROM \"issues\" WHERE (.+)$").WithArgs(1).WillReturnError(errors.New("test error"))

	item, err := r.FindByID(uint(1))

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, uint(0), item.ID)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceIssueFindAll(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteIssueRepository(gormDB)

	mock.MatchExpectationsInOrder(false)

	issueData := sqlmock.NewRows([]string{
		"id", "title", "description", "status", "project_id",
	}).AddRow(uint(1), "test-title-1", "test-description", 1, 1)
	mock.ExpectQuery("SELECT (.+) FROM \"issues\"").WillReturnRows(issueData)

	projectData := sqlmock.NewRows([]string{
		"id", "name", "description",
	}).AddRow(uint(1), "test-name", "test-description")
	mock.ExpectQuery("SELECT (.+) FROM \"projects\"").WithArgs(1).WillReturnRows(projectData)

	labelData := sqlmock.NewRows([]string{
		"id", "name", "color_hex_code",
	}).AddRow(uint(1), "test-name", "FFFFFF")
	mock.ExpectQuery("SELECT (.+) FROM \"labels\"").WithArgs(1).WillReturnRows(labelData)

	items, err := r.FindAll()

	assert.Nil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, 1, len(items))

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceIssueFindAllErr(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteIssueRepository(gormDB)

	mock.ExpectQuery("SELECT (.+) FROM \"issues\"").WillReturnError(errors.New("test error"))

	items, err := r.FindAll()

	assert.NotNil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, 0, len(items))

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceIssueRemove(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteIssueRepository(gormDB)

	mock.ExpectExec("DELETE FROM \"issues_labels\" (.+)$").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM \"issues\" (.+)$").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	status, err := r.Remove(uint(1))

	assert.Nil(t, err)
	assert.True(t, status)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceIssueRemoveManyToManyDeleteErr(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteIssueRepository(gormDB)

	mock.ExpectExec("DELETE FROM \"issues_labels\" (.+)$").WithArgs(1).WillReturnError(errors.New("test error"))

	status, err := r.Remove(uint(1))

	assert.NotNil(t, err)
	assert.False(t, status)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}

func TestPersistenceIssueRemoveErr(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteIssueRepository(gormDB)

	mock.ExpectExec("DELETE FROM \"issues_labels\" (.+)$").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM \"issues\" (.+)$").WithArgs(1).WillReturnError(errors.New("test error"))
	mock.ExpectRollback()

	status, err := r.Remove(uint(1))

	assert.NotNil(t, err)
	assert.False(t, status)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met %s", err)
	}
}
