package usecases

import (
	"go-issue-tracker/pkg/domain"
)

// ColorUseCase interface
type ColorUseCase interface {
	GetColor() (domain.Color, error)
}

// colorUseCase struct
type colorUseCase struct {
	service domain.ColorService
}

// NewColorUseCase to create new ColorUseCase
func NewColorUseCase(repository domain.ColorRepository) ColorUseCase {
	return &colorUseCase{
		service: domain.GetDefaultColorService(repository),
	}
}

// GetColor to get color
func (uc *colorUseCase) GetColor() (domain.Color, error) {
	color, err := uc.service.GetColor()
	if err != nil {
		return color, err
	}
	return color, nil
}
