package usecases

import (
	"go-issue-tracker/pkg/domain"
)

// IssueUseCase interface
type IssueUseCase interface {
	Add(title string, description string, status int, project domain.Project, labels map[string]domain.Label) (*domain.Issue, error)
	Update(id uint, title string, description string, status int, labels map[string]domain.Label) (domain.Issue, error)
	FindByID(id uint) (domain.Issue, error)
	Find(title string, projectID uint, labels []string) ([]domain.Issue, error)
	FindAll() ([]domain.Issue, error)
	Remove(id uint) (bool, error)
}

// IssueUseCase struct
type issueUseCase struct {
	service domain.IssueService
}

// NewIssueUseCase to create new IssueUseCase
func NewIssueUseCase(repository domain.IssueRepository) IssueUseCase {
	return &issueUseCase{
		service: domain.GetDefaultIssueService(repository),
	}
}

// Add to add new issue
func (uc *issueUseCase) Add(title string, description string, status int, project domain.Project, labels map[string]domain.Label) (*domain.Issue, error) {
	item := new(domain.Issue)
	item.Title = title
	item.Description = description
	item.Status = status
	item.ProjectID = project.ID
	item.Project = project
	for _, label := range labels {
		item.Labels = append(item.Labels, label)
	}

	itemAdded, err := uc.service.Add(item)
	if err != nil {
		return nil, err
	}

	return itemAdded, nil
}

// Update to update issue
func (uc *issueUseCase) Update(id uint, title string, description string, status int, labels map[string]domain.Label) (domain.Issue, error) {
	item, err := uc.service.FindByID(id)
	if err != nil {
		return item, err
	}

	item.Title = title
	item.Description = description
	item.Status = status
	item.Labels = []domain.Label{}
	for _, label := range labels {
		item.Labels = append(item.Labels, label)
	}

	itemUpdated, err := uc.service.Update(item)
	if err != nil {
		return itemUpdated, err
	}

	return itemUpdated, nil
}

// FindByID to find issue by ID
func (uc *issueUseCase) FindByID(id uint) (domain.Issue, error) {
	item, err := uc.service.FindByID(id)
	if err != nil {
		return item, err
	}
	return item, nil
}

// Find to find issues
func (uc *issueUseCase) Find(title string, projectID uint, labels []string) ([]domain.Issue, error) {
	items, err := uc.service.Find(title, projectID, labels)
	if err != nil {
		return items, err
	}
	return items, nil
}

// FindAll to find all issues
func (uc *issueUseCase) FindAll() ([]domain.Issue, error) {
	items, err := uc.service.FindAll()
	if err != nil {
		return items, err
	}
	return items, nil
}

// Remove to remove issue
func (uc *issueUseCase) Remove(id uint) (bool, error) {
	status, err := uc.service.Remove(id)
	if err != nil {
		return status, err
	}
	return status, nil
}
