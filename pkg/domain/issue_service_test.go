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
	v := new(domain.Issue)
	v.Title = "test-title"
	v.Description = "test-description"
	v.Status = 1
	v.ProjectID = 1

	m := new(dTesting.IssueRepositoryMock)
	m.On("Add", v).Return(v, nil)

	s := domain.GetDefaultIssueService(m)

	item, err := s.Add(v)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, v, item)

	m.AssertExpectations(t)
}

func TestDomainIssueAddErr(t *testing.T) {
	v := new(domain.Issue)

	m := new(dTesting.IssueRepositoryMock)
	m.On("Add", v).Return(v, errors.New("test error"))

	s := domain.GetDefaultIssueService(m)

	item, err := s.Add(v)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	m.AssertExpectations(t)
}

func TestDomainIssueAddValidateLabelsErr(t *testing.T) {
	l := testLabels
	v := new(domain.Issue)
	v.Title = "test-title"
	v.Description = "test-description"
	v.Status = 1
	v.ProjectID = 1
	v.Labels = l

	m := new(dTesting.IssueRepositoryMock)

	s := domain.GetDefaultIssueService(m)

	item, err := s.Add(v)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	m.AssertExpectations(t)
}

func TestDomainIssueUpdate(t *testing.T) {
	v := domain.Issue{
		ID:          1,
		Title:       "test-title",
		Description: "test-description",
		Status:      1,
		ProjectID:   1,
	}

	m := new(dTesting.IssueRepositoryMock)
	m.On("Update", v).Return(v, nil)

	s := domain.GetDefaultIssueService(m)

	item, err := s.Update(v)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, v, item)

	m.AssertExpectations(t)
}

func TestDomainIssueUpdateErr(t *testing.T) {
	v := domain.Issue{}

	m := new(dTesting.IssueRepositoryMock)
	m.On("Update", v).Return(v, errors.New("test error"))

	s := domain.GetDefaultIssueService(m)

	item, err := s.Update(v)

	assert.NotNil(t, err)
	assert.Equal(t, v, item)

	m.AssertExpectations(t)
}

func TestDomainIssueUpdateValidateLabelsErr(t *testing.T) {
	l := testLabels
	v := domain.Issue{
		Labels: l,
	}

	m := new(dTesting.IssueRepositoryMock)

	s := domain.GetDefaultIssueService(m)

	item, err := s.Update(v)

	assert.NotNil(t, err)
	assert.Equal(t, v, item)

	m.AssertExpectations(t)
}

func TestDomainIssueFindByID(t *testing.T) {
	v := domain.Issue{
		ID:          1,
		Title:       "test-title",
		Description: "test-description",
		Status:      1,
		ProjectID:   1,
	}

	m := new(dTesting.IssueRepositoryMock)
	m.On("FindByID", v.ID).Return(v, nil)

	s := domain.GetDefaultIssueService(m)

	item, err := s.FindByID(v.ID)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, v, item)

	m.AssertExpectations(t)
}

func TestDomainIssueFindByIDErr(t *testing.T) {
	v := domain.Issue{}

	m := new(dTesting.IssueRepositoryMock)
	m.On("FindByID", uint(1)).Return(v, errors.New("test error"))

	s := domain.GetDefaultIssueService(m)

	item, err := s.FindByID(uint(1))

	assert.NotNil(t, err)
	assert.Equal(t, v, item)

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
