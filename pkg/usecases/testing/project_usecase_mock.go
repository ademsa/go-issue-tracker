package testing

import (
	"github.com/stretchr/testify/mock"
	"go-issue-tracker/pkg/domain"
)

// ProjectUseCaseMock is a mock of ProjectUseCase
type ProjectUseCaseMock struct {
	mock.Mock
}

// Add mock
func (m *ProjectUseCaseMock) Add(name string, description string) (*domain.Project, error) {
	args := m.Called(name, description)
	return args.Get(0).(*domain.Project), args.Error(1)
}

// Update mock
func (m *ProjectUseCaseMock) Update(id uint, name string, description string) (domain.Project, error) {
	args := m.Called(id, name, description)
	return args.Get(0).(domain.Project), args.Error(1)
}

// FindByID mock
func (m *ProjectUseCaseMock) FindByID(id uint) (domain.Project, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Project), args.Error(1)
}

// Find mock
func (m *ProjectUseCaseMock) Find(name string) ([]domain.Project, error) {
	args := m.Called(name)
	return args.Get(0).([]domain.Project), args.Error(1)
}

// FindAll mock
func (m *ProjectUseCaseMock) FindAll() ([]domain.Project, error) {
	args := m.Called()
	return args.Get(0).([]domain.Project), args.Error(1)
}

// Remove mock
func (m *ProjectUseCaseMock) Remove(id uint) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}
