package domain

// ColorService interface
type ColorService interface {
	GetColor() (Color, error)
}

// colorService struct
type colorService struct {
	repository ColorRepository
}

// GetDefaultColorService alias to newColorService
var GetDefaultColorService = newColorService

// ResetDefaultColorService to reset GetDefaultColorService value
func ResetDefaultColorService() {
	GetDefaultColorService = newColorService
}

// newColorService to create new ColorService
func newColorService(repository ColorRepository) ColorService {
	return &colorService{
		repository: repository,
	}
}

// GetColor to get color
func (s *colorService) GetColor() (Color, error) {
	color, err := s.repository.GetColor()
	if err != nil {
		return color, err
	}
	return color, nil
}
