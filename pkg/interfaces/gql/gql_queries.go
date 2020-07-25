package gql

import (
	"github.com/graphql-go/graphql"
)

// GetQuery to get query
func GetQuery(resolver Resolver) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"issue": &graphql.Field{
				Type:        IssueType,
				Description: "Find Issue by ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
				},
				Resolve: resolver.ResolveFindIssueByIDQuery,
			},
			"issues": &graphql.Field{
				Type:        graphql.NewList(IssueType),
				Description: "Find Issues",
				Args: graphql.FieldConfigArgument{
					"title":     &graphql.ArgumentConfig{Type: graphql.String},
					"projectId": &graphql.ArgumentConfig{Type: graphql.String},
					"labels":    &graphql.ArgumentConfig{Type: graphql.String},
				},
				Resolve: resolver.ResolveFindIssuesQuery,
			},
			"allIssues": &graphql.Field{
				Type:        graphql.NewList(IssueType),
				Description: "Find All Issues",
				Resolve:     resolver.ResolveFindAllIssuesQuery,
			},
			"label": &graphql.Field{
				Type:        LabelType,
				Description: "Find Label by ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
				},
				Resolve: resolver.ResolveFindLabelByIDQuery,
			},
			"labels": &graphql.Field{
				Type:        graphql.NewList(LabelType),
				Description: "Find Labels",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{Type: graphql.String},
				},
				Resolve: resolver.ResolveFindLabelsQuery,
			},
			"allLabels": &graphql.Field{
				Type:        graphql.NewList(LabelType),
				Description: "Find All Labels",
				Resolve:     resolver.ResolveFindAllLabelsQuery,
			},
			"project": &graphql.Field{
				Type:        ProjectType,
				Description: "Find Project by ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
				},
				Resolve: resolver.ResolveFindProjectByIDQuery,
			},
			"projects": &graphql.Field{
				Type:        graphql.NewList(ProjectType),
				Description: "Find Projects",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{Type: graphql.String},
				},
				Resolve: resolver.ResolveFindProjectsQuery,
			},
			"allProjects": &graphql.Field{
				Type:        graphql.NewList(ProjectType),
				Description: "Find All Projects",
				Resolve:     resolver.ResolveFindAllProjectsQuery,
			},
			"node": NodeDefinitions.NodeField,
		},
	})
}
