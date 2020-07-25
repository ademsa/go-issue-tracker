package testing

import (
	"github.com/stretchr/testify/mock"
	"go-issue-tracker/pkg/domain"
)

// IssueServiceMock is a mock of IssueService
type IssueServiceMock struct {
	mock.Mock
}

// Add mock
func (m *IssueServiceMock) Add(issue *domain.Issue) (*domain.Issue, error) {
	args := m.Called(issue)
	return args.Get(0).(*domain.Issue), args.Error(1)
}

// Update mock
func (m *IssueServiceMock) Update(issue domain.Issue) (domain.Issue, error) {
	args := m.Called(issue)
	return args.Get(0).(domain.Issue), args.Error(1)
}

// FindByID mock
func (m *IssueServiceMock) FindByID(id uint) (domain.Issue, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Issue), args.Error(1)
}

// Find mock
func (m *IssueServiceMock) Find(title string, projectID uint, labels []string) ([]domain.Issue, error) {
	args := m.Called(title, projectID, labels)
	return args.Get(0).([]domain.Issue), args.Error(1)
}

// FindAll mock
func (m *IssueServiceMock) FindAll() ([]domain.Issue, error) {
	args := m.Called()
	return args.Get(0).([]domain.Issue), args.Error(1)
}

// Remove mock
func (m *IssueServiceMock) Remove(id uint) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}
