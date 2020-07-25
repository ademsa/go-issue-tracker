package testing

import (
	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
)

// ResolverMock is a mock of Resolver
type ResolverMock struct {
	mock.Mock
}

// ResolveNodeID mock
func (m *ResolverMock) ResolveNodeID(context context.Context, id string, info graphql.ResolveInfo) (interface{}, error) {
	args := m.Called(context, id, info)
	return args.Get(0).(interface{}), args.Error(1)
}

// ResolveType mock
func (m *ResolverMock) ResolveType(p graphql.ResolveTypeParams) *graphql.Object {
	args := m.Called(p)
	return args.Get(0).(*graphql.Object)
}

// ResolveFieldLabels mock
func (m *ResolverMock) ResolveFieldLabels(p graphql.ResolveParams) (interface{}, error) {
	args := m.Called(p)
	return args.Get(0).(interface{}), args.Error(1)
}

// ResolveAddIssueQuery mock
func (m *ResolverMock) ResolveAddIssueQuery(p graphql.ResolveParams) (interface{}, error) {
	args := m.Called(p)
	return args.Get(0).(interface{}), args.Error(1)
}

// ResolveUpdateIssueQuery mock
func (m *ResolverMock) ResolveUpdateIssueQuery(p graphql.ResolveParams) (interface{}, error) {
	args := m.Called(p)
	return args.Get(0).(interface{}), args.Error(1)
}

// ResolveFindIssueByIDQuery mock
func (m *ResolverMock) ResolveFindIssueByIDQuery(p graphql.ResolveParams) (interface{}, error) {
	args := m.Called(p)
	return args.Get(0).(interface{}), args.Error(1)
}

// ResolveFindIssuesQuery mock
func (m *ResolverMock) ResolveFindIssuesQuery(p graphql.ResolveParams) (interface{}, error) {
	args := m.Called(p)
	return args.Get(0).(interface{}), args.Error(1)
}

// ResolveFindAllIssuesQuery mock
func (m *ResolverMock) ResolveFindAllIssuesQuery(p graphql.ResolveParams) (interface{}, error) {
	args := m.Called(p)
	return args.Get(0).(interface{}), args.Error(1)
}

// ResolveRemoveIssueQuery mock
func (m *ResolverMock) ResolveRemoveIssueQuery(p graphql.ResolveParams) (interface{}, error) {
	args := m.Called(p)
	return args.Get(0).(interface{}), args.Error(1)
}

// ResolveAddLabelQuery mock
func (m *ResolverMock) ResolveAddLabelQuery(p graphql.ResolveParams) (interface{}, error) {
	args := m.Called(p)
	return args.Get(0).(interface{}), args.Error(1)
}

// ResolveUpdateLabelQuery mock
func (m *ResolverMock) ResolveUpdateLabelQuery(p graphql.ResolveParams) (interface{}, error) {
	args := m.Called(p)
	return args.Get(0).(interface{}), args.Error(1)
}

// ResolveFindLabelByIDQuery mock
func (m *ResolverMock) ResolveFindLabelByIDQuery(p graphql.ResolveParams) (interface{}, error) {
	args := m.Called(p)
	return args.Get(0).(interface{}), args.Error(1)
}

// ResolveFindLabelsQuery mock
func (m *ResolverMock) ResolveFindLabelsQuery(p graphql.ResolveParams) (interface{}, error) {
	args := m.Called(p)
	return args.Get(0).(interface{}), args.Error(1)
}

// ResolveFindAllLabelsQuery mock
func (m *ResolverMock) ResolveFindAllLabelsQuery(p graphql.ResolveParams) (interface{}, error) {
	args := m.Called(p)
	return args.Get(0).(interface{}), args.Error(1)
}

// ResolveRemoveLabelQuery mock
func (m *ResolverMock) ResolveRemoveLabelQuery(p graphql.ResolveParams) (interface{}, error) {
	args := m.Called(p)
	return args.Get(0).(interface{}), args.Error(1)
}

// ResolveAddProjectQuery mock
func (m *ResolverMock) ResolveAddProjectQuery(p graphql.ResolveParams) (interface{}, error) {
	args := m.Called(p)
	return args.Get(0).(interface{}), args.Error(1)
}

// ResolveUpdateProjectQuery mock
func (m *ResolverMock) ResolveUpdateProjectQuery(p graphql.ResolveParams) (interface{}, error) {
	args := m.Called(p)
	return args.Get(0).(interface{}), args.Error(1)
}

// ResolveFindProjectByIDQuery mock
func (m *ResolverMock) ResolveFindProjectByIDQuery(p graphql.ResolveParams) (interface{}, error) {
	args := m.Called(p)
	return args.Get(0).(interface{}), args.Error(1)
}

// ResolveFindProjectsQuery mock
func (m *ResolverMock) ResolveFindProjectsQuery(p graphql.ResolveParams) (interface{}, error) {
	args := m.Called(p)
	return args.Get(0).(interface{}), args.Error(1)
}

// ResolveFindAllProjectsQuery mock
func (m *ResolverMock) ResolveFindAllProjectsQuery(p graphql.ResolveParams) (interface{}, error) {
	args := m.Called(p)
	return args.Get(0).(interface{}), args.Error(1)
}

// ResolveRemoveProjectQuery mock
func (m *ResolverMock) ResolveRemoveProjectQuery(p graphql.ResolveParams) (interface{}, error) {
	args := m.Called(p)
	return args.Get(0).(interface{}), args.Error(1)
}

// ResolveMutationOutputFieldItem mock
func (m *ResolverMock) ResolveMutationOutputFieldItem(p graphql.ResolveParams) (interface{}, error) {
	args := m.Called(p)
	return args.Get(0).(interface{}), args.Error(1)
}

// ResolveMutationOutputFieldItemID mock
func (m *ResolverMock) ResolveMutationOutputFieldItemID(p graphql.ResolveParams) (interface{}, error) {
	args := m.Called(p)
	return args.Get(0).(interface{}), args.Error(1)
}

// ResolveMutationOutputFieldStatus mock
func (m *ResolverMock) ResolveMutationOutputFieldStatus(p graphql.ResolveParams) (interface{}, error) {
	args := m.Called(p)
	return args.Get(0).(interface{}), args.Error(1)
}

// MutateAndGetPayloadForAddIssueMutation mock
func (m *ResolverMock) MutateAndGetPayloadForAddIssueMutation(ctx context.Context, inputMap map[string]interface{}, info graphql.ResolveInfo) (map[string]interface{}, error) {
	args := m.Called(ctx, inputMap, info)
	return args.Get(0).(map[string]interface{}), args.Error(1)
}

// MutateAndGetPayloadForUpdateIssueMutation mock
func (m *ResolverMock) MutateAndGetPayloadForUpdateIssueMutation(ctx context.Context, inputMap map[string]interface{}, info graphql.ResolveInfo) (map[string]interface{}, error) {
	args := m.Called(ctx, inputMap, info)
	return args.Get(0).(map[string]interface{}), args.Error(1)
}

// MutateAndGetPayloadForRemoveIssueMutation mock
func (m *ResolverMock) MutateAndGetPayloadForRemoveIssueMutation(ctx context.Context, inputMap map[string]interface{}, info graphql.ResolveInfo) (map[string]interface{}, error) {
	args := m.Called(ctx, inputMap, info)
	return args.Get(0).(map[string]interface{}), args.Error(1)
}

// MutateAndGetPayloadForAddProjectMutation mock
func (m *ResolverMock) MutateAndGetPayloadForAddProjectMutation(ctx context.Context, inputMap map[string]interface{}, info graphql.ResolveInfo) (map[string]interface{}, error) {
	args := m.Called(ctx, inputMap, info)
	return args.Get(0).(map[string]interface{}), args.Error(1)
}

// MutateAndGetPayloadForUpdateProjectMutation mock
func (m *ResolverMock) MutateAndGetPayloadForUpdateProjectMutation(ctx context.Context, inputMap map[string]interface{}, info graphql.ResolveInfo) (map[string]interface{}, error) {
	args := m.Called(ctx, inputMap, info)
	return args.Get(0).(map[string]interface{}), args.Error(1)
}

// MutateAndGetPayloadForRemoveProjectMutation mock
func (m *ResolverMock) MutateAndGetPayloadForRemoveProjectMutation(ctx context.Context, inputMap map[string]interface{}, info graphql.ResolveInfo) (map[string]interface{}, error) {
	args := m.Called(ctx, inputMap, info)
	return args.Get(0).(map[string]interface{}), args.Error(1)
}

// MutateAndGetPayloadForAddLabelMutation mock
func (m *ResolverMock) MutateAndGetPayloadForAddLabelMutation(ctx context.Context, inputMap map[string]interface{}, info graphql.ResolveInfo) (map[string]interface{}, error) {
	args := m.Called(ctx, inputMap, info)
	return args.Get(0).(map[string]interface{}), args.Error(1)
}

// MutateAndGetPayloadForUpdateLabelMutation mock
func (m *ResolverMock) MutateAndGetPayloadForUpdateLabelMutation(ctx context.Context, inputMap map[string]interface{}, info graphql.ResolveInfo) (map[string]interface{}, error) {
	args := m.Called(ctx, inputMap, info)
	return args.Get(0).(map[string]interface{}), args.Error(1)
}

// MutateAndGetPayloadForRemoveLabelMutation mock
func (m *ResolverMock) MutateAndGetPayloadForRemoveLabelMutation(ctx context.Context, inputMap map[string]interface{}, info graphql.ResolveInfo) (map[string]interface{}, error) {
	args := m.Called(ctx, inputMap, info)
	return args.Get(0).(map[string]interface{}), args.Error(1)
}
