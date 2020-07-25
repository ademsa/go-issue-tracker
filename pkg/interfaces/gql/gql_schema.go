package gql

import (
	"github.com/graphql-go/graphql"
)

// GetSchema to get GraphQL schema
func GetSchema(query *graphql.Object, mutation *graphql.Object) (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    query,
		Mutation: mutation,
	})
}
