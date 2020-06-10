package testing

import (
	"github.com/stretchr/testify/mock"
	"go-issue-tracker/pkg/domain"
)

// ColorUseCaseMock is a mock of ColorUseCase
type ColorUseCaseMock struct {
	mock.Mock
}

// GetColor mock
func (m *ColorUseCaseMock) GetColor() (domain.Color, error) {
	args := m.Called()
	return args.Get(0).(domain.Color), args.Error(1)
}
