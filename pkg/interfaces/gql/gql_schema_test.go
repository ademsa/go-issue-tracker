package gql_test

import (
	"github.com/stretchr/testify/assert"
	"go-issue-tracker/pkg/interfaces/gql"
	gqlTesting "go-issue-tracker/pkg/interfaces/gql/testing"
	"testing"
)

func TestSchema(t *testing.T) {
	rm := new(gqlTesting.ResolverMock)

	gql.SetTypesAndNodeDefinitions(rm)

	query := gql.GetQuery(rm)

	mutation := gql.GetMutation(rm)

	schema, err := gql.GetSchema(query, mutation)

	assert.Nil(t, err)
	assert.NotNil(t, schema)
}
