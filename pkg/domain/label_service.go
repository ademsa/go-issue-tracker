package domain

import (
	"fmt"
)

// LabelService interface
type LabelService interface {
	Add(label *Label) (*Label, error)
	Update(label Label) (Label, error)
	FindByID(id uint) (Label, error)
	FindByName(name string) (Label, error)
	FindAll() ([]Label, error)
	Remove(id uint) (bool, error)
}

// labelService struct
type labelService struct {
	repository LabelRepository
}

// GetDefaultLabelService alias to newLabelService
var GetDefaultLabelService = newLabelService

// ResetDefaultLabelService to reset GetDefaultIssueService value
func ResetDefaultLabelService() {
	GetDefaultLabelService = newLabelService
}

// newLabelService to create new LabelService
func newLabelService(repository LabelRepository) LabelService {
	return &labelService{
		repository: repository,
	}
}

// alreadyExists checks if label already exists
func (s *labelService) alreadyExists(name string) error {
	item, err := s.repository.FindByName(name)
	if item.ID != 0 {
		return fmt.Errorf("%s label already exists", name)
	}
	if err != nil && err.Error() != "record not found" {
		return err
	}
	return nil
}

// Add to add new label
func (s *labelService) Add(label *Label) (*Label, error) {
	if err := s.alreadyExists(label.Name); err != nil {
		return nil, err
	}

	item, err := s.repository.Add(label)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// Update to update label
func (s *labelService) Update(label Label) (Label, error) {
	item, err := s.repository.Update(label)
	if err != nil {
		return item, err
	}
	return item, nil
}

// FindByID to find label by ID
func (s *labelService) FindByID(id uint) (Label, error) {
	item, err := s.repository.FindByID(id)
	if err != nil {
		return item, err
	}
	return item, nil
}

// FindByName to find label by name
func (s *labelService) FindByName(name string) (Label, error) {
	item, err := s.repository.FindByName(name)
	if err != nil {
		return item, err
	}
	return item, nil
}

// FindAll to find all labels
func (s *labelService) FindAll() ([]Label, error) {
	items, err := s.repository.FindAll()
	if err != nil {
		return items, err
	}
	return items, nil
}

// Remove to remove label
func (s *labelService) Remove(id uint) (bool, error) {
	status, err := s.repository.Remove(id)
	if err != nil {
		return status, err
	}
	return status, nil
}
