package domain_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go-issue-tracker/pkg/domain"
	dTesting "go-issue-tracker/pkg/domain/testing"
	"testing"
)

func TestDomainProjectResetDefaultProjectService(t *testing.T) {
	assert.NotNil(t, domain.GetDefaultProjectService)

	domain.GetDefaultProjectService = nil
	defer domain.ResetDefaultProjectService()

	assert.Nil(t, domain.GetDefaultProjectService)

	domain.ResetDefaultProjectService()

	assert.NotNil(t, domain.GetDefaultProjectService)
}

func TestDomainProjectGetDefaultProjectService(t *testing.T) {
	m := new(dTesting.ProjectRepositoryMock)

	s := domain.GetDefaultProjectService(m)

	assert.NotNil(t, s)
}

func TestDomainProjectAdd(t *testing.T) {
	p := new(domain.Project)
	p.Name = "test-name"
	p.Description = "test-description"

	m := new(dTesting.ProjectRepositoryMock)
	m.On("Add", p).Return(p, nil)

	s := domain.GetDefaultProjectService(m)

	item, err := s.Add(p)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, p, item)

	m.AssertExpectations(t)
}

func TestDomainProjectAddErr(t *testing.T) {
	p := new(domain.Project)

	m := new(dTesting.ProjectRepositoryMock)
	m.On("Add", p).Return(p, errors.New("test error"))

	s := domain.GetDefaultProjectService(m)

	item, err := s.Add(p)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	m.AssertExpectations(t)
}

func TestDomainProjectUpdate(t *testing.T) {
	p := domain.Project{
		ID:          1,
		Name:        "test-name",
		Description: "test-description",
	}

	m := new(dTesting.ProjectRepositoryMock)
	m.On("Update", p).Return(p, nil)

	s := domain.GetDefaultProjectService(m)

	item, err := s.Update(p)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, p, item)

	m.AssertExpectations(t)
}

func TestDomainProjectUpdateErr(t *testing.T) {
	p := domain.Project{}

	m := new(dTesting.ProjectRepositoryMock)
	m.On("Update", p).Return(p, errors.New("test error"))

	s := domain.GetDefaultProjectService(m)

	item, err := s.Update(p)

	assert.NotNil(t, err)
	assert.Equal(t, p, item)

	m.AssertExpectations(t)
}

func TestDomainProjectFindByID(t *testing.T) {
	p := domain.Project{
		ID:          1,
		Name:        "test-name",
		Description: "test-description",
	}

	m := new(dTesting.ProjectRepositoryMock)
	m.On("FindByID", p.ID).Return(p, nil)

	s := domain.GetDefaultProjectService(m)

	item, err := s.FindByID(p.ID)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, p, item)

	m.AssertExpectations(t)
}

func TestDomainProjectFindByIDErr(t *testing.T) {
	p := domain.Project{}

	m := new(dTesting.ProjectRepositoryMock)
	m.On("FindByID", uint(1)).Return(p, errors.New("test error"))

	s := domain.GetDefaultProjectService(m)

	item, err := s.FindByID(uint(1))

	assert.NotNil(t, err)
	assert.Equal(t, p, item)

	m.AssertExpectations(t)
}

func TestDomainProjectFind(t *testing.T) {
	v := []domain.Project{}

	m := new(dTesting.ProjectRepositoryMock)
	m.On("Find", "test").Return(v, nil)

	s := domain.GetDefaultProjectService(m)

	items, err := s.Find("test")

	assert.Nil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, v, items)

	m.AssertExpectations(t)
}

func TestDomainProjectFindErr(t *testing.T) {
	v := []domain.Project{}

	m := new(dTesting.ProjectRepositoryMock)
	m.On("Find", "test").Return(v, errors.New("test error"))

	s := domain.GetDefaultProjectService(m)

	items, err := s.Find("test")

	assert.NotNil(t, err)
	assert.Equal(t, v, items)

	m.AssertExpectations(t)
}

func TestDomainProjectFindAll(t *testing.T) {
	v := []domain.Project{}

	m := new(dTesting.ProjectRepositoryMock)
	m.On("FindAll").Return(v, nil)

	s := domain.GetDefaultProjectService(m)

	items, err := s.FindAll()

	assert.Nil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, v, items)

	m.AssertExpectations(t)
}

func TestDomainProjectFindAllErr(t *testing.T) {
	v := []domain.Project{}

	m := new(dTesting.ProjectRepositoryMock)
	m.On("FindAll").Return(v, errors.New("test error"))

	s := domain.GetDefaultProjectService(m)

	items, err := s.FindAll()

	assert.NotNil(t, err)
	assert.Equal(t, v, items)

	m.AssertExpectations(t)
}

func TestDomainProjectRemove(t *testing.T) {
	m := new(dTesting.ProjectRepositoryMock)
	m.On("Remove", uint(1)).Return(true, nil)

	s := domain.GetDefaultProjectService(m)

	status, err := s.Remove(uint(1))

	assert.Nil(t, err)
	assert.True(t, status)

	m.AssertExpectations(t)
}

func TestDomainProjectRemoveErr(t *testing.T) {
	m := new(dTesting.ProjectRepositoryMock)
	m.On("Remove", uint(1)).Return(false, errors.New("test error"))

	s := domain.GetDefaultProjectService(m)

	status, err := s.Remove(uint(1))

	assert.NotNil(t, err)
	assert.False(t, status)

	m.AssertExpectations(t)
}
