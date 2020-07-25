package testing

import (
	"github.com/stretchr/testify/mock"
	"go-issue-tracker/pkg/domain"
)

// LabelUseCaseMock is a mock of LabelUseCase
type LabelUseCaseMock struct {
	mock.Mock
}

// Add mock
func (m *LabelUseCaseMock) Add(name string, colorHexCode string) (*domain.Label, error) {
	args := m.Called(name, colorHexCode)
	return args.Get(0).(*domain.Label), args.Error(1)
}

// Update mock
func (m *LabelUseCaseMock) Update(id uint, name string, colorHexCode string) (domain.Label, error) {
	args := m.Called(id, name, colorHexCode)
	return args.Get(0).(domain.Label), args.Error(1)
}

// FindByID mock
func (m *LabelUseCaseMock) FindByID(id uint) (domain.Label, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Label), args.Error(1)
}

// FindByName mock
func (m *LabelUseCaseMock) FindByName(name string) (domain.Label, error) {
	args := m.Called(name)
	return args.Get(0).(domain.Label), args.Error(1)
}

// Find mock
func (m *LabelUseCaseMock) Find(name string) ([]domain.Label, error) {
	args := m.Called(name)
	return args.Get(0).([]domain.Label), args.Error(1)
}

// FindAll mock
func (m *LabelUseCaseMock) FindAll() ([]domain.Label, error) {
	args := m.Called()
	return args.Get(0).([]domain.Label), args.Error(1)
}

// Remove mock
func (m *LabelUseCaseMock) Remove(id uint) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}
