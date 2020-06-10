package testing

import (
	"github.com/stretchr/testify/mock"
	"net/http"
)

// HTTPClientMock is a mock of HTTPClient
type HTTPClientMock struct {
	mock.Mock
}

// Get mock
func (m *HTTPClientMock) Get(url string) (resp *http.Response, err error) {
	args := m.Called(url)
	return args.Get(0).(*http.Response), args.Error(1)
}
