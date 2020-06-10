package testing

import (
	"github.com/stretchr/testify/mock"
	"go-issue-tracker/pkg/domain"
)

// ColorRepositoryMock is a mock of ColorRepository
type ColorRepositoryMock struct {
	mock.Mock
}

// GetColor mock
func (m *ColorRepositoryMock) GetColor() (domain.Color, error) {
	args := m.Called()
	return args.Get(0).(domain.Color), args.Error(1)
}
