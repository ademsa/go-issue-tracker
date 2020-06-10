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
	v := new(domain.Project)
	v.Name = "test-name"
	v.Description = "test-description"

	m := new(dTesting.ProjectRepositoryMock)
	m.On("Add", v).Return(v, nil)

	s := domain.GetDefaultProjectService(m)

	item, err := s.Add(v)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, v, item)

	m.AssertExpectations(t)
}

func TestDomainProjectAddErr(t *testing.T) {
	v := new(domain.Project)

	m := new(dTesting.ProjectRepositoryMock)
	m.On("Add", v).Return(v, errors.New("test error"))

	s := domain.GetDefaultProjectService(m)

	item, err := s.Add(v)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	m.AssertExpectations(t)
}

func TestDomainProjectUpdate(t *testing.T) {
	v := domain.Project{
		ID:          1,
		Name:        "test-name",
		Description: "test-description",
	}

	m := new(dTesting.ProjectRepositoryMock)
	m.On("Update", v).Return(v, nil)

	s := domain.GetDefaultProjectService(m)

	item, err := s.Update(v)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, v, item)

	m.AssertExpectations(t)
}

func TestDomainProjectUpdateErr(t *testing.T) {
	v := domain.Project{}

	m := new(dTesting.ProjectRepositoryMock)
	m.On("Update", v).Return(v, errors.New("test error"))

	s := domain.GetDefaultProjectService(m)

	item, err := s.Update(v)

	assert.NotNil(t, err)
	assert.Equal(t, v, item)

	m.AssertExpectations(t)
}

func TestDomainProjectFindByID(t *testing.T) {
	v := domain.Project{
		ID:          1,
		Name:        "test-name",
		Description: "test-description",
	}

	m := new(dTesting.ProjectRepositoryMock)
	m.On("FindByID", v.ID).Return(v, nil)

	s := domain.GetDefaultProjectService(m)

	item, err := s.FindByID(v.ID)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, v, item)

	m.AssertExpectations(t)
}

func TestDomainProjectFindByIDErr(t *testing.T) {
	v := domain.Project{}

	m := new(dTesting.ProjectRepositoryMock)
	m.On("FindByID", uint(1)).Return(v, errors.New("test error"))

	s := domain.GetDefaultProjectService(m)

	item, err := s.FindByID(uint(1))

	assert.NotNil(t, err)
	assert.Equal(t, v, item)

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
