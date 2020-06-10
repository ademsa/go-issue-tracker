package testing

import (
	"github.com/stretchr/testify/mock"
	"go-issue-tracker/pkg/domain"
)

// IssueUseCaseMock is a mock of IssueUseCase
type IssueUseCaseMock struct {
	mock.Mock
}

// Add mock
func (m *IssueUseCaseMock) Add(title string, description string, status int, project domain.Project, labels map[string]domain.Label) (*domain.Issue, error) {
	args := m.Called(title, description, status, project, labels)
	return args.Get(0).(*domain.Issue), args.Error(1)
}

// Update mock
func (m *IssueUseCaseMock) Update(id uint, title string, description string, status int, labels map[string]domain.Label) (domain.Issue, error) {
	args := m.Called(id, title, description, status, labels)
	return args.Get(0).(domain.Issue), args.Error(1)
}

// FindByID mock
func (m *IssueUseCaseMock) FindByID(id uint) (domain.Issue, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Issue), args.Error(1)
}

// FindAll mock
func (m *IssueUseCaseMock) FindAll() ([]domain.Issue, error) {
	args := m.Called()
	return args.Get(0).([]domain.Issue), args.Error(1)
}

// Remove mock
func (m *IssueUseCaseMock) Remove(id uint) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}
