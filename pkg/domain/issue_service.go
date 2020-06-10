package domain

import (
	"errors"
)

// IssueService interface
type IssueService interface {
	Add(issue *Issue) (*Issue, error)
	Update(issue Issue) (Issue, error)
	FindByID(id uint) (Issue, error)
	FindAll() ([]Issue, error)
	Remove(id uint) (bool, error)
}

// issueService struct
type issueService struct {
	repository IssueRepository
}

// GetDefaultIssueService alias to newIssueService
var GetDefaultIssueService = newIssueService

// ResetDefaultIssueService to reset GetDefaultIssueService value
func ResetDefaultIssueService() {
	GetDefaultIssueService = newIssueService
}

// newIssueService to create new IssueService
func newIssueService(repository IssueRepository) IssueService {
	return &issueService{
		repository: repository,
	}
}

// validateLabels validates if there are too many labels
func (s *issueService) validateLabels(labels []Label) error {
	if len(labels) > 10 {
		return errors.New("max. 10 labels can be assigned to issue")
	}
	return nil
}

// Add to add new issue
func (s *issueService) Add(issue *Issue) (*Issue, error) {
	if err := s.validateLabels(issue.Labels); err != nil {
		return nil, err
	}

	item, err := s.repository.Add(issue)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// Update to update issue
func (s *issueService) Update(issue Issue) (Issue, error) {
	if err := s.validateLabels(issue.Labels); err != nil {
		return issue, err
	}

	item, err := s.repository.Update(issue)
	if err != nil {
		return item, err
	}
	return item, nil
}

// FindByID to find issue by ID
func (s *issueService) FindByID(id uint) (Issue, error) {
	item, err := s.repository.FindByID(id)
	if err != nil {
		return item, err
	}
	return item, nil
}

// FindAll to find all issues
func (s *issueService) FindAll() ([]Issue, error) {
	items, err := s.repository.FindAll()
	if err != nil {
		return items, err
	}
	return items, nil
}

// Remove to remove issue
func (s *issueService) Remove(id uint) (bool, error) {
	status, err := s.repository.Remove(id)
	if err != nil {
		return status, err
	}
	return status, nil
}
