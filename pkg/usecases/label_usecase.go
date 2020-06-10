package usecases

import (
	"go-issue-tracker/pkg/domain"
)

// LabelUseCase interface
type LabelUseCase interface {
	Add(name string, colorHexCode string) (*domain.Label, error)
	Update(id uint, name string, colorHexCode string) (domain.Label, error)
	FindByID(id uint) (domain.Label, error)
	FindByName(name string) (domain.Label, error)
	FindAll() ([]domain.Label, error)
	Remove(id uint) (bool, error)
}

// LabelUseCase struct
type labelUseCase struct {
	service domain.LabelService
}

// NewLabelUseCase to create new LabelUseCase
func NewLabelUseCase(repository domain.LabelRepository) LabelUseCase {
	return &labelUseCase{
		service: domain.GetDefaultLabelService(repository),
	}
}

// Add to add new label
func (uc *labelUseCase) Add(name string, colorHexCode string) (*domain.Label, error) {
	var item = new(domain.Label)
	item.Name = name
	item.ColorHexCode = colorHexCode
	itemAdded, err := uc.service.Add(item)
	if err != nil {
		return nil, err
	}
	return itemAdded, nil
}

// Update to update label
func (uc *labelUseCase) Update(id uint, name string, colorHexCode string) (domain.Label, error) {
	item, err := uc.service.FindByID(id)
	if err != nil {
		return item, err
	}

	item.Name = name
	item.ColorHexCode = colorHexCode
	itemUpdated, err := uc.service.Update(item)
	if err != nil {
		return itemUpdated, err
	}
	return itemUpdated, nil
}

// FindByID to find label by ID
func (uc *labelUseCase) FindByID(id uint) (domain.Label, error) {
	item, err := uc.service.FindByID(id)
	if err != nil {
		return item, err
	}
	return item, nil
}

// FindByName to find label by name
func (uc *labelUseCase) FindByName(name string) (domain.Label, error) {
	item, err := uc.service.FindByName(name)
	if err != nil {
		return item, err
	}
	return item, nil
}

// FindAll to find all labels
func (uc *labelUseCase) FindAll() ([]domain.Label, error) {
	items, err := uc.service.FindAll()
	if err != nil {
		return items, err
	}
	return items, nil
}

// Remove to remove label
func (uc *labelUseCase) Remove(id uint) (bool, error) {
	status, err := uc.service.Remove(id)
	if err != nil {
		return status, err
	}
	return status, nil
}
