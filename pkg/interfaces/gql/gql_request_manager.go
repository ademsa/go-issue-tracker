package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
)

// RequestManager interface
type RequestManager interface {
	Handler(c echo.Context) error
	Execute(schema graphql.Schema, query string, variableValues map[string]interface{}) *graphql.Result
}

// requestManager contains base tooling
type requestManager struct {
	schema graphql.Schema
}

// NewRequestManager to init RequestManager
func NewRequestManager(schema graphql.Schema) RequestManager {
	return &requestManager{
		schema: schema,
	}
}
