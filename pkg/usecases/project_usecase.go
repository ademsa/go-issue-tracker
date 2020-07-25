package usecases

import (
	"go-issue-tracker/pkg/domain"
)

// ProjectUseCase interface
type ProjectUseCase interface {
	Add(name string, description string) (*domain.Project, error)
	Update(id uint, name string, description string) (domain.Project, error)
	FindByID(id uint) (domain.Project, error)
	Find(name string) ([]domain.Project, error)
	FindAll() ([]domain.Project, error)
	Remove(id uint) (bool, error)
}

// ProjectUseCase struct
type projectUseCase struct {
	service domain.ProjectService
}

// NewProjectUseCase to create new ProjectUseCase
func NewProjectUseCase(repository domain.ProjectRepository) ProjectUseCase {
	return &projectUseCase{
		service: domain.GetDefaultProjectService(repository),
	}
}

// Add to add new project
func (uc *projectUseCase) Add(name string, description string) (*domain.Project, error) {
	item := new(domain.Project)
	item.Name = name
	item.Description = description
	itemAdded, err := uc.service.Add(item)
	if err != nil {
		return nil, err
	}
	return itemAdded, nil
}

// Update to update project
func (uc *projectUseCase) Update(id uint, name string, description string) (domain.Project, error) {
	item, err := uc.service.FindByID(id)
	if err != nil {
		return item, err
	}

	item.Name = name
	item.Description = description
	itemUpdated, err := uc.service.Update(item)
	if err != nil {
		return itemUpdated, err
	}
	return itemUpdated, nil
}

// FindByID to find project by ID
func (uc *projectUseCase) FindByID(id uint) (domain.Project, error) {
	item, err := uc.service.FindByID(id)
	if err != nil {
		return item, err
	}
	return item, nil
}

// Find to find project by name
func (uc *projectUseCase) Find(name string) ([]domain.Project, error) {
	items, err := uc.service.Find(name)
	if err != nil {
		return items, err
	}
	return items, nil
}

// FindAll to find all projects
func (uc *projectUseCase) FindAll() ([]domain.Project, error) {
	items, err := uc.service.FindAll()
	if err != nil {
		return items, err
	}
	return items, nil
}

// Remove to remove label
func (uc *projectUseCase) Remove(id uint) (bool, error) {
	status, err := uc.service.Remove(id)
	if err != nil {
		return status, err
	}
	return status, nil
}
