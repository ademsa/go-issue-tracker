package gql_test

import (
	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/assert"
	"go-issue-tracker/pkg/interfaces/gql"
	"testing"
)

func TestNewRequestManager(t *testing.T) {
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{})

	m := gql.NewRequestManager(schema)

	assert.NotNil(t, m)
}
