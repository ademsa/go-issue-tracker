package testing

import (
	"github.com/stretchr/testify/mock"
	"go-issue-tracker/pkg/domain"
)

// ColorServiceMock is a mock of ColorService
type ColorServiceMock struct {
	mock.Mock
}

// GetColor mock
func (m *ColorServiceMock) GetColor() (domain.Color, error) {
	args := m.Called()
	return args.Get(0).(domain.Color), args.Error(1)
}
