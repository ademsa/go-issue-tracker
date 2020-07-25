package gql_test

import (
	"github.com/stretchr/testify/assert"
	"go-issue-tracker/pkg/interfaces/gql"
	gqlTesting "go-issue-tracker/pkg/interfaces/gql/testing"
	"testing"
)

func TestGetMutation(t *testing.T) {
	rm := new(gqlTesting.ResolverMock)

	gql.SetTypesAndNodeDefinitions(rm)

	mutation := gql.GetMutation(rm)

	assert.NotNil(t, mutation)
}
