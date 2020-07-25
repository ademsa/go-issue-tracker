package testing

import (
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

// RequestManagerMock is a mock of RequestManager
type RequestManagerMock struct {
	mock.Mock
}

// Handler mock
func (m *RequestManagerMock) Handler(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// Execute mock
func (m *RequestManagerMock) Execute(schema graphql.Schema, query string, variableValues map[string]interface{}) *graphql.Result {
	args := m.Called(query, schema, variableValues)
	return args.Get(0).(*graphql.Result)
}
