package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"golang.org/x/net/context"
)

// IssueType graphql type
var IssueType *graphql.Object

// LabelType graphql type
var LabelType *graphql.Object

// ProjectType graphql type
var ProjectType *graphql.Object

// NodeDefinitions graphql node definitions
var NodeDefinitions *relay.NodeDefinitions

// SetTypesAndNodeDefinitions to set types and node definitions
func SetTypesAndNodeDefinitions(resolver Resolver) {
	NodeDefinitions = relay.NewNodeDefinitions(relay.NodeDefinitionsConfig{
		IDFetcher: func(id string, info graphql.ResolveInfo, context context.Context) (interface{}, error) {
			return resolver.ResolveNodeID(context, id, info)
		},
		TypeResolve: resolver.ResolveType,
	})

	LabelType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Label",
		Fields: graphql.Fields{
			"id":           relay.GlobalIDField("Label", nil),
			"name":         &graphql.Field{Type: graphql.String},
			"colorHexCode": &graphql.Field{Type: graphql.String},
			"createdAt":    &graphql.Field{Type: graphql.DateTime},
			"updatedAt":    &graphql.Field{Type: graphql.DateTime},
		},
		Interfaces: []*graphql.Interface{NodeDefinitions.NodeInterface},
	})

	ProjectType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Project",
		Fields: graphql.Fields{
			"id":          relay.GlobalIDField("Project", nil),
			"name":        &graphql.Field{Type: graphql.String},
			"description": &graphql.Field{Type: graphql.String},
			"createdAt":   &graphql.Field{Type: graphql.DateTime},
			"updatedAt":   &graphql.Field{Type: graphql.DateTime},
		},
		Interfaces: []*graphql.Interface{NodeDefinitions.NodeInterface},
	})

	labelConnectionDefinition := relay.ConnectionDefinitions(relay.ConnectionConfig{
		Name:     "Label",
		NodeType: LabelType,
	})

	IssueType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Issue",
		Fields: graphql.Fields{
			"id":          relay.GlobalIDField("Issue", nil),
			"title":       &graphql.Field{Type: graphql.String},
			"description": &graphql.Field{Type: graphql.String},
			"status":      &graphql.Field{Type: graphql.Int},
			"projectId":   &graphql.Field{Type: graphql.Int},
			"project":     &graphql.Field{Type: ProjectType},
			"labels": &graphql.Field{
				Type:    labelConnectionDefinition.ConnectionType,
				Args:    relay.ConnectionArgs,
				Resolve: resolver.ResolveFieldLabels,
			},
			"createdAt": &graphql.Field{Type: graphql.DateTime},
			"updatedAt": &graphql.Field{Type: graphql.DateTime},
		},
		Interfaces: []*graphql.Interface{NodeDefinitions.NodeInterface},
	})
}
