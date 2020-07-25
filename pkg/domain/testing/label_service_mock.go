package testing

import (
	"github.com/stretchr/testify/mock"
	"go-issue-tracker/pkg/domain"
)

// LabelServiceMock is a mock of LabelService
type LabelServiceMock struct {
	mock.Mock
}

// alreadyExists mock
func (m *LabelServiceMock) alreadyExists(name string) error {
	args := m.Called(name)
	return args.Error(0)
}

// Add mock
func (m *LabelServiceMock) Add(label *domain.Label) (*domain.Label, error) {
	args := m.Called(label)
	return args.Get(0).(*domain.Label), args.Error(1)
}

// Update mock
func (m *LabelServiceMock) Update(label domain.Label) (domain.Label, error) {
	args := m.Called(label)
	return args.Get(0).(domain.Label), args.Error(1)
}

// FindByID mock
func (m *LabelServiceMock) FindByID(id uint) (domain.Label, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Label), args.Error(1)
}

// FindByName mock
func (m *LabelServiceMock) FindByName(name string) (domain.Label, error) {
	args := m.Called(name)
	return args.Get(0).(domain.Label), args.Error(1)
}

// Find mock
func (m *LabelServiceMock) Find(name string) ([]domain.Label, error) {
	args := m.Called(name)
	return args.Get(0).([]domain.Label), args.Error(1)
}

// FindAll mock
func (m *LabelServiceMock) FindAll() ([]domain.Label, error) {
	args := m.Called()
	return args.Get(0).([]domain.Label), args.Error(1)
}

// Remove mock
func (m *LabelServiceMock) Remove(id uint) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}
