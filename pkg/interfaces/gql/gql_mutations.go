package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"golang.org/x/net/context"
)

// GetMutation to get mutation
func GetMutation(resolver Resolver) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"addIssue": relay.MutationWithClientMutationID(relay.MutationConfig{
				Name: "AddIssue",
				InputFields: graphql.InputObjectConfigFieldMap{
					"title":       &graphql.InputObjectFieldConfig{Type: graphql.String},
					"description": &graphql.InputObjectFieldConfig{Type: graphql.String},
					"status":      &graphql.InputObjectFieldConfig{Type: graphql.Int},
					"projectId":   &graphql.InputObjectFieldConfig{Type: graphql.String},
					"labels":      &graphql.InputObjectFieldConfig{Type: graphql.String},
				},
				OutputFields: graphql.Fields{
					"issue": &graphql.Field{
						Type:    IssueType,
						Resolve: resolver.ResolveMutationOutputFieldItem,
					},
				},
				MutateAndGetPayload: func(inputMap map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
					return resolver.MutateAndGetPayloadForAddIssueMutation(ctx, inputMap, info)
				},
			}),
			"updateIssue": relay.MutationWithClientMutationID(relay.MutationConfig{
				Name: "UpdateIssue",
				InputFields: graphql.InputObjectConfigFieldMap{
					"id":          &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.ID)},
					"title":       &graphql.InputObjectFieldConfig{Type: graphql.String},
					"description": &graphql.InputObjectFieldConfig{Type: graphql.String},
					"status":      &graphql.InputObjectFieldConfig{Type: graphql.Int},
					"labels":      &graphql.InputObjectFieldConfig{Type: graphql.String},
				},
				OutputFields: graphql.Fields{
					"issue": &graphql.Field{
						Type:    IssueType,
						Resolve: resolver.ResolveMutationOutputFieldItem,
					},
				},
				MutateAndGetPayload: func(inputMap map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
					return resolver.MutateAndGetPayloadForUpdateIssueMutation(ctx, inputMap, info)
				},
			}),
			"removeIssue": relay.MutationWithClientMutationID(relay.MutationConfig{
				Name: "RemoveIssue",
				InputFields: graphql.InputObjectConfigFieldMap{
					"id": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.ID)},
				},
				OutputFields: graphql.Fields{
					"issueId": &graphql.Field{
						Type:    graphql.NewNonNull(graphql.ID),
						Resolve: resolver.ResolveMutationOutputFieldItemID,
					},
					"status": &graphql.Field{
						Type:    graphql.Boolean,
						Resolve: resolver.ResolveMutationOutputFieldStatus,
					},
				},
				MutateAndGetPayload: func(inputMap map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
					return resolver.MutateAndGetPayloadForRemoveIssueMutation(ctx, inputMap, info)
				},
			}),
			"addLabel": relay.MutationWithClientMutationID(relay.MutationConfig{
				Name: "AddLabel",
				InputFields: graphql.InputObjectConfigFieldMap{
					"name":         &graphql.InputObjectFieldConfig{Type: graphql.String},
					"colorHexCode": &graphql.InputObjectFieldConfig{Type: graphql.String},
				},
				OutputFields: graphql.Fields{
					"label": &graphql.Field{
						Type:    LabelType,
						Resolve: resolver.ResolveMutationOutputFieldItem,
					},
				},
				MutateAndGetPayload: func(inputMap map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
					return resolver.MutateAndGetPayloadForAddLabelMutation(ctx, inputMap, info)
				},
			}),
			"updateLabel": relay.MutationWithClientMutationID(relay.MutationConfig{
				Name: "UpdateLabel",
				InputFields: graphql.InputObjectConfigFieldMap{
					"id":           &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.ID)},
					"name":         &graphql.InputObjectFieldConfig{Type: graphql.String},
					"colorHexCode": &graphql.InputObjectFieldConfig{Type: graphql.String},
				},
				OutputFields: graphql.Fields{
					"label": &graphql.Field{
						Type:    LabelType,
						Resolve: resolver.ResolveMutationOutputFieldItem,
					},
				},
				MutateAndGetPayload: func(inputMap map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
					return resolver.MutateAndGetPayloadForUpdateLabelMutation(ctx, inputMap, info)
				},
			}),
			"removeLabel": relay.MutationWithClientMutationID(relay.MutationConfig{
				Name: "RemoveLabel",
				InputFields: graphql.InputObjectConfigFieldMap{
					"id": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.ID)},
				},
				OutputFields: graphql.Fields{
					"labelId": &graphql.Field{
						Type:    graphql.NewNonNull(graphql.ID),
						Resolve: resolver.ResolveMutationOutputFieldItemID,
					},
					"status": &graphql.Field{
						Type:    graphql.Boolean,
						Resolve: resolver.ResolveMutationOutputFieldStatus,
					},
				},
				MutateAndGetPayload: func(inputMap map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
					return resolver.MutateAndGetPayloadForRemoveLabelMutation(ctx, inputMap, info)
				},
			}),
			"addProject": relay.MutationWithClientMutationID(relay.MutationConfig{
				Name: "AddProject",
				InputFields: graphql.InputObjectConfigFieldMap{
					"name":        &graphql.InputObjectFieldConfig{Type: graphql.String},
					"description": &graphql.InputObjectFieldConfig{Type: graphql.String},
				},
				OutputFields: graphql.Fields{
					"project": &graphql.Field{
						Type:    ProjectType,
						Resolve: resolver.ResolveMutationOutputFieldItem,
					},
				},
				MutateAndGetPayload: func(inputMap map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
					return resolver.MutateAndGetPayloadForAddProjectMutation(ctx, inputMap, info)
				},
			}),
			"updateProject": relay.MutationWithClientMutationID(relay.MutationConfig{
				Name: "UpdateProject",
				InputFields: graphql.InputObjectConfigFieldMap{
					"id":          &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.ID)},
					"name":        &graphql.InputObjectFieldConfig{Type: graphql.String},
					"description": &graphql.InputObjectFieldConfig{Type: graphql.String},
				},
				OutputFields: graphql.Fields{
					"project": &graphql.Field{
						Type:    ProjectType,
						Resolve: resolver.ResolveMutationOutputFieldItem,
					},
				},
				MutateAndGetPayload: func(inputMap map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
					return resolver.MutateAndGetPayloadForUpdateProjectMutation(ctx, inputMap, info)
				},
			}),
			"removeProject": relay.MutationWithClientMutationID(relay.MutationConfig{
				Name: "RemoveProject",
				InputFields: graphql.InputObjectConfigFieldMap{
					"id": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.ID)},
				},
				OutputFields: graphql.Fields{
					"projectId": &graphql.Field{
						Type:    graphql.NewNonNull(graphql.ID),
						Resolve: resolver.ResolveMutationOutputFieldItemID,
					},
					"status": &graphql.Field{
						Type:    graphql.Boolean,
						Resolve: resolver.ResolveMutationOutputFieldStatus,
					},
				},
				MutateAndGetPayload: func(inputMap map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
					return resolver.MutateAndGetPayloadForRemoveProjectMutation(ctx, inputMap, info)
				},
			}),
		},
	})
}
