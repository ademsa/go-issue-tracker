package testing

import (
	"github.com/stretchr/testify/mock"
	"go-issue-tracker/pkg/domain"
)

// ProjectRepositoryMock is a mock of ProjectRepository
type ProjectRepositoryMock struct {
	mock.Mock
}

// Add mock
func (m *ProjectRepositoryMock) Add(project *domain.Project) (*domain.Project, error) {
	args := m.Called(project)
	return args.Get(0).(*domain.Project), args.Error(1)
}

// Update mock
func (m *ProjectRepositoryMock) Update(project domain.Project) (domain.Project, error) {
	args := m.Called(project)
	return args.Get(0).(domain.Project), args.Error(1)
}

// FindByID mock
func (m *ProjectRepositoryMock) FindByID(id uint) (domain.Project, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Project), args.Error(1)
}

// FindAll mock
func (m *ProjectRepositoryMock) FindAll() ([]domain.Project, error) {
	args := m.Called()
	return args.Get(0).([]domain.Project), args.Error(1)
}

// Remove mock
func (m *ProjectRepositoryMock) Remove(id uint) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}
