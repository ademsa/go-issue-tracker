package testing

import (
	"github.com/stretchr/testify/mock"
	"go-issue-tracker/pkg/domain"
)

// IssueRepositoryMock is a mock of IssueRepository
type IssueRepositoryMock struct {
	mock.Mock
}

// Add mock
func (m *IssueRepositoryMock) Add(issue *domain.Issue) (*domain.Issue, error) {
	args := m.Called(issue)
	return args.Get(0).(*domain.Issue), args.Error(1)
}

// Update mock
func (m *IssueRepositoryMock) Update(issue domain.Issue) (domain.Issue, error) {
	args := m.Called(issue)
	return args.Get(0).(domain.Issue), args.Error(1)
}

// FindByID mock
func (m *IssueRepositoryMock) FindByID(id uint) (domain.Issue, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Issue), args.Error(1)
}

// FindAll mock
func (m *IssueRepositoryMock) FindAll() ([]domain.Issue, error) {
	args := m.Called()
	return args.Get(0).([]domain.Issue), args.Error(1)
}

// Remove mock
func (m *IssueRepositoryMock) Remove(id uint) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}
