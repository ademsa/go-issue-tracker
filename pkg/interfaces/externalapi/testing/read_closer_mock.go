package testing

import (
	"github.com/stretchr/testify/mock"
)

// ReadCloserMock is a mock of io.ReadCloser
type ReadCloserMock struct {
	mock.Mock
}

// Read mock
func (m *ReadCloserMock) Read(p []byte) (n int, err error) {
	args := m.Called(p)
	return args.Int(0), args.Error(1)
}

// Close mock
func (m *ReadCloserMock) Close() error {
	args := m.Called()
	return args.Error(0)
}
