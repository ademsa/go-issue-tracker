package domain_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go-issue-tracker/pkg/domain"
	dTesting "go-issue-tracker/pkg/domain/testing"
	"testing"
)

var testLabels = []domain.Label{
	{
		ID: 1,
	},
	{
		ID: 2,
	},
	{
		ID: 3,
	},
	{
		ID: 4,
	},
	{
		ID: 5,
	},
	{
		ID: 6,
	},
	{
		ID: 7,
	},
	{
		ID: 8,
	},
	{
		ID: 9,
	},
	{
		ID: 10,
	},
	{
		ID: 11,
	},
}

func TestDomainIssueResetDefaultIssueService(t *testing.T) {
	assert.NotNil(t, domain.GetDefaultIssueService)

	domain.GetDefaultIssueService = nil
	defer domain.ResetDefaultIssueService()

	assert.Nil(t, domain.GetDefaultIssueService)

	domain.ResetDefaultIssueService()

	assert.NotNil(t, domain.GetDefaultIssueService)
}

func TestDomainIssueGetDefaultIssueService(t *testing.T) {
	m := new(dTesting.IssueRepositoryMock)

	s := domain.GetDefaultIssueService(m)

	assert.NotNil(t, s)
}

func TestDomainIssueAdd(t *testing.T) {
	i := new(domain.Issue)
	i.Title = "test-title"
	i.Description = "test-description"
	i.Status = 1
	i.ProjectID = 1

	m := new(dTesting.IssueRepositoryMock)
	m.On("Add", i).Return(i, nil)

	s := domain.GetDefaultIssueService(m)

	item, err := s.Add(i)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, i, item)

	m.AssertExpectations(t)
}

func TestDomainIssueAddErr(t *testing.T) {
	i := new(domain.Issue)

	m := new(dTesting.IssueRepositoryMock)
	m.On("Add", i).Return(i, errors.New("test error"))

	s := domain.GetDefaultIssueService(m)

	item, err := s.Add(i)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	m.AssertExpectations(t)
}

func TestDomainIssueAddValidateLabelsErr(t *testing.T) {
	l := testLabels
	i := new(domain.Issue)
	i.Title = "test-title"
	i.Description = "test-description"
	i.Status = 1
	i.ProjectID = 1
	i.Labels = l

	m := new(dTesting.IssueRepositoryMock)

	s := domain.GetDefaultIssueService(m)

	item, err := s.Add(i)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	m.AssertExpectations(t)
}

func TestDomainIssueUpdate(t *testing.T) {
	i := domain.Issue{
		ID:          1,
		Title:       "test-title",
		Description: "test-description",
		Status:      1,
		ProjectID:   1,
	}

	m := new(dTesting.IssueRepositoryMock)
	m.On("Update", i).Return(i, nil)

	s := domain.GetDefaultIssueService(m)

	item, err := s.Update(i)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, i, item)

	m.AssertExpectations(t)
}

func TestDomainIssueUpdateErr(t *testing.T) {
	i := domain.Issue{}

	m := new(dTesting.IssueRepositoryMock)
	m.On("Update", i).Return(i, errors.New("test error"))

	s := domain.GetDefaultIssueService(m)

	item, err := s.Update(i)

	assert.NotNil(t, err)
	assert.Equal(t, i, item)

	m.AssertExpectations(t)
}

func TestDomainIssueUpdateValidateLabelsErr(t *testing.T) {
	l := testLabels
	i := domain.Issue{
		Labels: l,
	}

	m := new(dTesting.IssueRepositoryMock)

	s := domain.GetDefaultIssueService(m)

	item, err := s.Update(i)

	assert.NotNil(t, err)
	assert.Equal(t, i, item)

	m.AssertExpectations(t)
}

func TestDomainIssueFindByID(t *testing.T) {
	i := domain.Issue{
		ID:          1,
		Title:       "test-title",
		Description: "test-description",
		Status:      1,
		ProjectID:   1,
	}

	m := new(dTesting.IssueRepositoryMock)
	m.On("FindByID", i.ID).Return(i, nil)

	s := domain.GetDefaultIssueService(m)

	item, err := s.FindByID(i.ID)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, i, item)

	m.AssertExpectations(t)
}

func TestDomainIssueFindByIDErr(t *testing.T) {
	i := domain.Issue{}

	m := new(dTesting.IssueRepositoryMock)
	m.On("FindByID", uint(1)).Return(i, errors.New("test error"))

	s := domain.GetDefaultIssueService(m)

	item, err := s.FindByID(uint(1))

	assert.NotNil(t, err)
	assert.Equal(t, i, item)

	m.AssertExpectations(t)
}

func TestDomainIssueFind(t *testing.T) {
	v := []domain.Issue{}

	m := new(dTesting.IssueRepositoryMock)
	m.On("Find", "test", uint(1), []string{"test1", "test2"}).Return(v, nil)

	s := domain.GetDefaultIssueService(m)

	items, err := s.Find("test", uint(1), []string{"test1", "test2"})

	assert.Nil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, v, items)

	m.AssertExpectations(t)
}

func TestDomainIssueFindErr(t *testing.T) {
	v := []domain.Issue{}

	m := new(dTesting.IssueRepositoryMock)
	m.On("Find", "test", uint(1), []string{"test1", "test2"}).Return(v, errors.New("test error"))

	s := domain.GetDefaultIssueService(m)

	items, err := s.Find("test", uint(1), []string{"test1", "test2"})

	assert.NotNil(t, err)
	assert.Equal(t, v, items)

	m.AssertExpectations(t)
}

func TestDomainIssueFindAll(t *testing.T) {
	v := []domain.Issue{}

	m := new(dTesting.IssueRepositoryMock)
	m.On("FindAll").Return(v, nil)

	s := domain.GetDefaultIssueService(m)

	items, err := s.FindAll()

	assert.Nil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, v, items)

	m.AssertExpectations(t)
}

func TestDomainIssueFindAllErr(t *testing.T) {
	v := []domain.Issue{}

	m := new(dTesting.IssueRepositoryMock)
	m.On("FindAll").Return(v, errors.New("test error"))

	s := domain.GetDefaultIssueService(m)

	items, err := s.FindAll()

	assert.NotNil(t, err)
	assert.Equal(t, v, items)

	m.AssertExpectations(t)
}

func TestDomainIssueRemove(t *testing.T) {
	m := new(dTesting.IssueRepositoryMock)
	m.On("Remove", uint(1)).Return(true, nil)

	s := domain.GetDefaultIssueService(m)

	status, err := s.Remove(uint(1))

	assert.Nil(t, err)
	assert.True(t, status)

	m.AssertExpectations(t)
}

func TestDomainIssueRemoveErr(t *testing.T) {
	m := new(dTesting.IssueRepositoryMock)
	m.On("Remove", uint(1)).Return(false, errors.New("test error"))

	s := domain.GetDefaultIssueService(m)

	status, err := s.Remove(uint(1))

	assert.NotNil(t, err)
	assert.False(t, status)

	m.AssertExpectations(t)
}
