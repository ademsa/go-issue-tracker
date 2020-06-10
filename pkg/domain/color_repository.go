package domain

// ColorRepository repository
type ColorRepository interface {
	GetColor() (Color, error)
}
