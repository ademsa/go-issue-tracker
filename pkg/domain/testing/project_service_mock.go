package testing

import (
	"github.com/stretchr/testify/mock"
	"go-issue-tracker/pkg/domain"
)

// ProjectServiceMock is a mock of ProjectService
type ProjectServiceMock struct {
	mock.Mock
}

// Add mock
func (m *ProjectServiceMock) Add(project *domain.Project) (*domain.Project, error) {
	args := m.Called(project)
	return args.Get(0).(*domain.Project), args.Error(1)
}

// Update mock
func (m *ProjectServiceMock) Update(project domain.Project) (domain.Project, error) {
	args := m.Called(project)
	return args.Get(0).(domain.Project), args.Error(1)
}

// FindByID mock
func (m *ProjectServiceMock) FindByID(id uint) (domain.Project, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Project), args.Error(1)
}

// Find mock
func (m *ProjectServiceMock) Find(name string) ([]domain.Project, error) {
	args := m.Called(name)
	return args.Get(0).([]domain.Project), args.Error(1)
}

// FindAll mock
func (m *ProjectServiceMock) FindAll() ([]domain.Project, error) {
	args := m.Called()
	return args.Get(0).([]domain.Project), args.Error(1)
}

// Remove mock
func (m *ProjectServiceMock) Remove(id uint) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}
