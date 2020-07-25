package testing

import (
	"github.com/stretchr/testify/mock"
	"go-issue-tracker/pkg/domain"
)

// LabelRepositoryMock is a mock of LabelRepository
type LabelRepositoryMock struct {
	mock.Mock
}

// Add mock
func (m *LabelRepositoryMock) Add(label *domain.Label) (*domain.Label, error) {
	args := m.Called(label)
	return args.Get(0).(*domain.Label), args.Error(1)
}

// Update mock
func (m *LabelRepositoryMock) Update(label domain.Label) (domain.Label, error) {
	args := m.Called(label)
	return args.Get(0).(domain.Label), args.Error(1)
}

// FindByID mock
func (m *LabelRepositoryMock) FindByID(id uint) (domain.Label, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Label), args.Error(1)
}

// FindByName mock
func (m *LabelRepositoryMock) FindByName(name string) (domain.Label, error) {
	args := m.Called(name)
	return args.Get(0).(domain.Label), args.Error(1)
}

// Find mock
func (m *LabelRepositoryMock) Find(name string) ([]domain.Label, error) {
	args := m.Called(name)
	return args.Get(0).([]domain.Label), args.Error(1)
}

// FindAll mock
func (m *LabelRepositoryMock) FindAll() ([]domain.Label, error) {
	args := m.Called()
	return args.Get(0).([]domain.Label), args.Error(1)
}

// Remove mock
func (m *LabelRepositoryMock) Remove(id uint) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}
