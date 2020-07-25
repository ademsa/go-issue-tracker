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
	mock.ExpectExec("INSERT INTO \"issues\" (.+)$").WithArgs("test-title", "test-description", 1, 1, sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	i := new(domain.Issue)
	i.Title = "test-title"
	i.Description = "test-description"
	i.Status = 1
	i.ProjectID = 1

	item, err := r.Add(i)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, i.Title, item.Title)
	assert.Equal(t, i.Description, item.Description)
	assert.Equal(t, i.Status, item.Status)
	assert.Equal(t, i.ProjectID, item.ProjectID)
	assert.NotNil(t, item.CreatedAt)

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
	mock.ExpectExec("INSERT INTO \"issues\" (.+)$").WithArgs("test-title", "test-description", 1, 1, sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(errors.New("test error"))
	mock.ExpectRollback()

	i := new(domain.Issue)
	i.Title = "test-title"
	i.Description = "test-description"
	i.Status = 1
	i.ProjectID = 1

	item, err := r.Add(i)

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
	mock.ExpectExec("UPDATE \"issues\" SET (.+)$").WithArgs("test-title", "test-description", 1, 1, sqlmock.AnyArg(), 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	i := domain.Issue{
		ID:          uint(1),
		Title:       "test-title",
		Description: "test-description",
		Status:      1,
		ProjectID:   1,
	}

	item, err := r.Update(i)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, i.Title, item.Title)
	assert.Equal(t, i.Description, item.Description)
	assert.Equal(t, i.Status, item.Status)
	assert.Equal(t, i.ProjectID, item.ProjectID)
	assert.NotNil(t, item.UpdatedAt)

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

	i := domain.Issue{
		ID:          uint(1),
		Title:       "test-title",
		Description: "test-description",
		Status:      1,
		ProjectID:   1,
	}

	item, err := r.Update(i)

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, i.Title, item.Title)
	assert.Equal(t, i.Description, item.Description)
	assert.Equal(t, i.Status, item.Status)
	assert.Equal(t, i.ProjectID, item.ProjectID)
	assert.NotNil(t, item.UpdatedAt)

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
	mock.ExpectExec("UPDATE \"issues\" SET (.+)$").WithArgs("test-title", "test-description", 1, 1, sqlmock.AnyArg(), 1).WillReturnError(errors.New("test error"))
	mock.ExpectRollback()

	i := domain.Issue{
		ID:          uint(1),
		Title:       "test-title",
		Description: "test-description",
		Status:      1,
		ProjectID:   1,
	}

	item, err := r.Update(i)

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, i.Title, item.Title)
	assert.Equal(t, i.Description, item.Description)
	assert.Equal(t, i.Status, item.Status)
	assert.Equal(t, i.ProjectID, item.ProjectID)
	assert.NotNil(t, item.UpdatedAt)

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

func TestPersistenceIssueFind(t *testing.T) {
	tests := []struct {
		title     string
		projectID uint
		labels    []string
	}{
		{
			"test-title-1",
			uint(1),
			[]string{"test-name-1"},
		},
		{
			"test-title-1",
			uint(1),
			[]string{},
		},
		{
			"",
			uint(1),
			[]string{},
		},
		{
			"test-title-1",
			uint(0),
			[]string{},
		},
		{
			"",
			uint(0),
			[]string{"test-name-1"},
		},
		{
			"",
			uint(0),
			[]string{},
		},
	}

	for _, ts := range tests {
		mockDB, mock, gormDB := pTesting.GetMockedDB(t)
		defer mockDB.Close()
		defer gormDB.Close()

		r := persistence.NewSQLiteIssueRepository(gormDB)

		mock.MatchExpectationsInOrder(false)

		issueData := sqlmock.NewRows([]string{
			"id", "title", "description", "status", "project_id",
		}).AddRow(uint(1), "test-title-1", "test-description", 1, 1)
		mock.ExpectQuery("SELECT (.+) FROM \"issues\"").WillReturnRows(issueData)

		labelData := sqlmock.NewRows([]string{
			"id", "name", "color_hex_code",
		}).AddRow(uint(1), "test-name-1", "FFFFFF")
		mock.ExpectQuery("SELECT (.+) FROM \"labels\"").WithArgs(1).WillReturnRows(labelData)

		projectData := sqlmock.NewRows([]string{
			"id", "name", "description",
		}).AddRow(uint(1), "test-name-1", "test-description")
		mock.ExpectQuery("SELECT (.+) FROM \"projects\"").WillReturnRows(projectData)

		items, err := r.Find(ts.title, ts.projectID, ts.labels)

		assert.Nil(t, err)
		assert.NotNil(t, items)
		assert.Equal(t, 1, len(items))

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("expectations were not met %s", err)
		}
	}
}

func TestPersistenceIssueFindErr(t *testing.T) {
	mockDB, mock, gormDB := pTesting.GetMockedDB(t)
	defer mockDB.Close()
	defer gormDB.Close()

	r := persistence.NewSQLiteIssueRepository(gormDB)

	mock.ExpectQuery("SELECT (.+) FROM \"issues\"").WillReturnError(errors.New("test error"))

	tests := []struct {
		title     string
		projectID uint
		labels    []string
	}{
		{
			"test-title",
			uint(1),
			[]string{"test-name"},
		},
		{
			"",
			uint(0),
			[]string{},
		},
		{
			"test-title",
			uint(1),
			[]string{},
		},
	}

	for _, ts := range tests {
		items, err := r.Find(ts.title, ts.projectID, ts.labels)

		assert.NotNil(t, err)
		assert.NotNil(t, items)
		assert.Equal(t, 0, len(items))

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("expectations were not met %s", err)
		}
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
