package gql_test

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"github.com/stretchr/testify/assert"
	"go-issue-tracker/pkg/domain"
	"go-issue-tracker/pkg/interfaces/gql"
	ucTesting "go-issue-tracker/pkg/usecases/testing"
	"reflect"
	"testing"
)

func checkAssertions(t *testing.T, cucm *ucTesting.ColorUseCaseMock, iucm *ucTesting.IssueUseCaseMock, lucm *ucTesting.LabelUseCaseMock, pucm *ucTesting.ProjectUseCaseMock) {
	cucm.AssertExpectations(t)
	iucm.AssertExpectations(t)
	lucm.AssertExpectations(t)
	pucm.AssertExpectations(t)
}

func prepareMocksAndResolver() (*ucTesting.ColorUseCaseMock, *ucTesting.IssueUseCaseMock, *ucTesting.LabelUseCaseMock, *ucTesting.ProjectUseCaseMock, gql.Resolver) {
	cucm := new(ucTesting.ColorUseCaseMock)
	iucm := new(ucTesting.IssueUseCaseMock)
	lucm := new(ucTesting.LabelUseCaseMock)
	pucm := new(ucTesting.ProjectUseCaseMock)
	return cucm, iucm, lucm, pucm, gql.GetResolver(iucm, lucm, pucm, cucm)
}

func TestResolveNodeID(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	tests := []struct {
		id     string
		idType string
	}{
		{
			relay.ToGlobalID("Issue", "1"),
			"Issue",
		},
		{
			relay.ToGlobalID("Label", "1"),
			"Label",
		},
		{
			relay.ToGlobalID("Project", "1"),
			"Project",
		},
	}

	for _, ts := range tests {
		if ts.idType == "Issue" {
			iucm.On("FindByID", uint(1)).Return(domain.Issue{}, nil)
		} else if ts.idType == "Label" {
			lucm.On("FindByID", uint(1)).Return(domain.Label{}, nil)
		} else if ts.idType == "Project" {
			pucm.On("FindByID", uint(1)).Return(domain.Project{}, nil)
		}

		item, err := r.ResolveNodeID(nil, ts.id, graphql.ResolveInfo{})

		assert.Nil(t, err)
		assert.NotNil(t, item)

		checkAssertions(t, cucm, iucm, lucm, pucm)
	}
}

func TestResolveNodeIDArgErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	tests := []struct {
		id string
	}{
		{
			"",
		},
		{
			relay.ToGlobalID("Issue", "test"),
		},
	}

	for _, ts := range tests {
		item, err := r.ResolveNodeID(nil, ts.id, graphql.ResolveInfo{})

		assert.NotNil(t, err)
		assert.Nil(t, item)

		checkAssertions(t, cucm, iucm, lucm, pucm)
	}
}

func TestResolveNodeIDUnknownType(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	item, err := r.ResolveNodeID(nil, relay.ToGlobalID("Unknown", "1"), graphql.ResolveInfo{})

	assert.NotNil(t, err)
	assert.Nil(t, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveType(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	tests := []struct {
		value interface{}
	}{
		{
			domain.Issue{},
		},
		{
			new(domain.Issue),
		},
		{
			domain.Label{},
		},
		{
			new(domain.Label),
		},
		{
			domain.Project{},
		},
		{
			new(domain.Project),
		},
	}

	bType := reflect.TypeOf(new(graphql.Object))

	for _, ts := range tests {
		rtp := graphql.ResolveTypeParams{
			Value: ts.value,
		}

		fType := r.ResolveType(rtp)

		assert.Equal(t, bType, reflect.TypeOf(fType))

		checkAssertions(t, cucm, iucm, lucm, pucm)
	}
}

func TestResolveTypeUnknown(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	rtp := graphql.ResolveTypeParams{
		Value: map[string]interface{}{
			"example": "value",
		},
	}

	fType := r.ResolveType(rtp)

	assert.Nil(t, fType)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveFieldLabels(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	i := new(domain.Issue)
	i.Labels = []domain.Label{
		domain.Label{},
		domain.Label{},
	}

	tests := []struct {
		source interface{}
	}{
		{
			domain.Issue{
				Labels: i.Labels,
			},
		},
		{
			i,
		},
	}

	for _, ts := range tests {
		rp := graphql.ResolveParams{
			Source: ts.source,
			Args:   map[string]interface{}{},
		}

		connectionData, err := r.ResolveFieldLabels(rp)

		assert.Nil(t, err)
		assert.NotNil(t, connectionData)

		checkAssertions(t, cucm, iucm, lucm, pucm)
	}
}

func TestResolveFieldLabelsErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	rp := graphql.ResolveParams{
		Source: domain.Project{},
		Args:   map[string]interface{}{},
	}

	connectionData, err := r.ResolveFieldLabels(rp)

	assert.NotNil(t, err)
	assert.Nil(t, connectionData)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveMutationOutputFieldItem(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	rp := graphql.ResolveParams{
		Source: map[string]interface{}{
			"item": domain.Project{},
		},
	}

	item, err := r.ResolveMutationOutputFieldItem(rp)

	assert.Nil(t, err)
	assert.Equal(t, domain.Project{}, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveMutationOutputFieldItemErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	rp := graphql.ResolveParams{
		Source: nil,
	}

	item, err := r.ResolveMutationOutputFieldItem(rp)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveMutationOutputFieldItemID(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	rp := graphql.ResolveParams{
		Source: map[string]interface{}{
			"id": 1,
		},
	}

	item, err := r.ResolveMutationOutputFieldItemID(rp)

	assert.Nil(t, err)
	assert.Equal(t, 1, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveMutationOutputFieldItemIDErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	rp := graphql.ResolveParams{
		Source: nil,
	}

	item, err := r.ResolveMutationOutputFieldItemID(rp)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveMutationOutputFieldStatus(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	rp := graphql.ResolveParams{
		Source: map[string]interface{}{
			"status": false,
		},
	}

	item, err := r.ResolveMutationOutputFieldStatus(rp)

	assert.Nil(t, err)
	assert.Equal(t, false, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveMutationOutputFieldStatusErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	rp := graphql.ResolveParams{
		Source: nil,
	}

	item, err := r.ResolveMutationOutputFieldStatus(rp)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestMutateAndGetPayloadForAddIssueMutation(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	p := domain.Project{
		ID:          1,
		Name:        "test-name",
		Description: "test-description",
	}
	pucm.On("FindByID", uint(1)).Return(p, nil)

	l := domain.Label{
		ID:   1,
		Name: "test-name",
	}
	lucm.On("FindByID", uint(1)).Return(l, nil)

	i := new(domain.Issue)
	i.Title = "test-title"
	i.Description = "test-description"
	i.Status = 1
	i.ProjectID = p.ID
	i.Project = p
	i.Labels = []domain.Label{l}
	iucm.On("Add", i.Title, i.Description, i.Status, p, map[string]domain.Label{
		relay.ToGlobalID("Label", "1"): l,
	}).Return(i, nil)

	inputMap := map[string]interface{}{
		"title":       i.Title,
		"description": i.Description,
		"status":      i.Status,
		"projectId":   relay.ToGlobalID("Project", "1"),
		"labels":      relay.ToGlobalID("Label", "1"),
	}

	result, err := r.MutateAndGetPayloadForAddIssueMutation(nil, inputMap, graphql.ResolveInfo{})

	assert.Nil(t, err)
	assert.NotNil(t, result)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestMutateAndGetPayloadForAddIssueMutationArgErr(t *testing.T) {
	tests := []struct {
		title                  string
		description            string
		status                 interface{}
		projectID              string
		labels                 string
		mockProjectFindByID    bool
		mockProjectFindByIDErr bool
		mockLabelFindByID      bool
		mockLabelFindByIDErr   bool
	}{
		{
			"",
			"test-description",
			1,
			relay.ToGlobalID("Project", "1"),
			relay.ToGlobalID("Label", "1"),
			false,
			false,
			false,
			false,
		},
		{
			"test-title",
			"",
			1,
			relay.ToGlobalID("Project", "1"),
			relay.ToGlobalID("Label", "1"),
			false,
			false,
			false,
			false,
		},
		{
			"test-title",
			"test-description",
			"test",
			relay.ToGlobalID("Project", "1"),
			relay.ToGlobalID("Label", "1"),
			false,
			false,
			false,
			false,
		},
		{
			"test-title",
			"test-description",
			1,
			"",
			relay.ToGlobalID("Label", "1"),
			false,
			false,
			false,
			false,
		},
		{
			"test-title",
			"test-description",
			1,
			"test",
			relay.ToGlobalID("Label", "1"),
			false,
			false,
			false,
			false,
		},
		{
			"test-title",
			"test-description",
			1,
			relay.ToGlobalID("Project", "test"),
			relay.ToGlobalID("Label", "1"),
			false,
			false,
			false,
			false,
		},
		{
			"test-title",
			"test-description",
			1,
			relay.ToGlobalID("Project", "1"),
			relay.ToGlobalID("Label", "1"),
			false,
			true,
			false,
			false,
		},
		{
			"test-title",
			"test-description",
			1,
			relay.ToGlobalID("Project", "1"),
			"",
			true,
			false,
			false,
			false,
		},
		{
			"test-title",
			"test-description",
			1,
			relay.ToGlobalID("Project", "1"),
			",",
			true,
			false,
			false,
			false,
		},
		{
			"test-title",
			"test-description",
			1,
			relay.ToGlobalID("Project", "1"),
			"test",
			true,
			false,
			false,
			false,
		},
		{
			"test-title",
			"test-description",
			1,
			relay.ToGlobalID("Project", "1"),
			relay.ToGlobalID("Label", "test"),
			true,
			false,
			false,
			false,
		},
		{
			"test-title",
			"test-description",
			1,
			relay.ToGlobalID("Project", "1"),
			relay.ToGlobalID("Label", "1"),
			true,
			false,
			false,
			true,
		},
	}

	for _, ts := range tests {
		cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

		if ts.mockProjectFindByID {
			p := domain.Project{
				ID:          1,
				Name:        "test-name",
				Description: "test-description",
			}
			pucm.On("FindByID", uint(1)).Return(p, nil)
		}

		if ts.mockProjectFindByIDErr {
			p := domain.Project{
				ID:          1,
				Name:        "test-name",
				Description: "test-description",
			}
			pucm.On("FindByID", uint(1)).Return(p, errors.New("test error"))
		}

		if ts.mockLabelFindByID {
			l := domain.Label{
				ID:   1,
				Name: "test-name",
			}
			lucm.On("FindByID", uint(1)).Return(l, nil)
		}

		if ts.mockLabelFindByIDErr {
			l := domain.Label{
				ID:   1,
				Name: "test-name",
			}
			lucm.On("FindByID", uint(1)).Return(l, errors.New("test error"))
		}

		inputMap := map[string]interface{}{
			"title":       ts.title,
			"description": ts.description,
			"status":      ts.status,
			"projectId":   ts.projectID,
			"labels":      ts.labels,
		}

		result, err := r.MutateAndGetPayloadForAddIssueMutation(nil, inputMap, graphql.ResolveInfo{})

		assert.NotNil(t, err)
		assert.NotNil(t, result)

		checkAssertions(t, cucm, iucm, lucm, pucm)
	}
}

func TestMutateAndGetPayloadForAddIssueMutationErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	p := domain.Project{
		ID:          1,
		Name:        "test-name",
		Description: "test-description",
	}
	pucm.On("FindByID", uint(1)).Return(p, nil)

	l := domain.Label{
		ID:   1,
		Name: "test-name",
	}
	lucm.On("FindByID", uint(1)).Return(l, nil)

	i := new(domain.Issue)
	i.Title = "test-title"
	i.Description = "test-description"
	i.Status = 1
	i.ProjectID = p.ID
	i.Project = p
	i.Labels = []domain.Label{l}
	iucm.On("Add", i.Title, i.Description, i.Status, p, map[string]domain.Label{
		relay.ToGlobalID("Label", "1"): l,
	}).Return(i, errors.New("test error"))

	inputMap := map[string]interface{}{
		"title":       i.Title,
		"description": i.Description,
		"status":      i.Status,
		"projectId":   relay.ToGlobalID("Project", "1"),
		"labels":      relay.ToGlobalID("Label", "1"),
	}

	result, err := r.MutateAndGetPayloadForAddIssueMutation(nil, inputMap, graphql.ResolveInfo{})

	assert.NotNil(t, err)
	assert.NotNil(t, result)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestMutateAndGetPayloadForUpdateIssueMutation(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	p := domain.Project{
		ID:          1,
		Name:        "test-name",
		Description: "test-description",
	}

	l := domain.Label{
		ID:   1,
		Name: "test-name",
	}
	lucm.On("FindByID", uint(1)).Return(l, nil)

	i := domain.Issue{
		Title:       "test-title",
		Description: "test-description",
		Status:      1,
		ProjectID:   p.ID,
		Project:     p,
		Labels:      []domain.Label{l},
	}
	iucm.On("Update", uint(1), i.Title, i.Description, i.Status, map[string]domain.Label{
		relay.ToGlobalID("Label", "1"): l,
	}).Return(i, nil)

	inputMap := map[string]interface{}{
		"id":          relay.ToGlobalID("Issue", "1"),
		"title":       i.Title,
		"description": i.Description,
		"status":      i.Status,
		"labels":      relay.ToGlobalID("Label", "1"),
	}

	result, err := r.MutateAndGetPayloadForUpdateIssueMutation(nil, inputMap, graphql.ResolveInfo{})

	assert.Nil(t, err)
	assert.NotNil(t, result)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestMutateAndGetPayloadForUpdateIssueMutationArgErr(t *testing.T) {
	tests := []struct {
		id                   string
		title                string
		description          string
		status               interface{}
		labels               string
		mockLabelFindByID    bool
		mockLabelFindByIDErr bool
	}{
		{
			"",
			"test-title",
			"test-description",
			1,
			relay.ToGlobalID("Label", "1"),
			false,
			false,
		},
		{
			relay.ToGlobalID("Issue", "1"),
			"",
			"test-description",
			1,
			relay.ToGlobalID("Label", "1"),
			false,
			false,
		},
		{
			relay.ToGlobalID("Issue", "1"),
			"test-title",
			"",
			1,
			relay.ToGlobalID("Label", "1"),
			false,
			false,
		},
		{
			relay.ToGlobalID("Issue", "1"),
			"test-title",
			"test-description",
			"test",
			relay.ToGlobalID("Label", "1"),
			false,
			false,
		},
		{
			relay.ToGlobalID("Issue", "1"),
			"test-title",
			"test-description",
			1,
			"",
			false,
			false,
		},
		{
			relay.ToGlobalID("Issue", "1"),
			"test-title",
			"test-description",
			1,
			",",
			false,
			false,
		},
		{
			relay.ToGlobalID("Issue", "1"),
			"test-title",
			"test-description",
			1,
			"test",
			false,
			false,
		},
		{
			relay.ToGlobalID("Issue", "1"),
			"test-title",
			"test-description",
			1,
			relay.ToGlobalID("Label", "test"),
			false,
			false,
		},
		{
			relay.ToGlobalID("Issue", "1"),
			"test-title",
			"test-description",
			1,
			relay.ToGlobalID("Label", "1"),
			false,
			true,
		},
	}

	for _, ts := range tests {
		cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

		if ts.mockLabelFindByID {
			l := domain.Label{
				ID:   1,
				Name: "test-name",
			}
			lucm.On("FindByID", uint(1)).Return(l, nil)
		}

		if ts.mockLabelFindByIDErr {
			l := domain.Label{
				ID:   1,
				Name: "test-name",
			}
			lucm.On("FindByID", uint(1)).Return(l, errors.New("test error"))
		}

		inputMap := map[string]interface{}{
			"id":          ts.id,
			"title":       ts.title,
			"description": ts.description,
			"status":      ts.status,
			"labels":      ts.labels,
		}

		result, err := r.MutateAndGetPayloadForUpdateIssueMutation(nil, inputMap, graphql.ResolveInfo{})

		assert.NotNil(t, err)
		assert.NotNil(t, result)

		checkAssertions(t, cucm, iucm, lucm, pucm)
	}
}

func TestMutateAndGetPayloadForUpdateIssueMutationErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	p := domain.Project{
		ID:          1,
		Name:        "test-name",
		Description: "test-description",
	}

	l := domain.Label{
		ID:   1,
		Name: "test-name",
	}
	lucm.On("FindByID", uint(1)).Return(l, nil)

	i := domain.Issue{
		Title:       "test-title",
		Description: "test-description",
		Status:      1,
		ProjectID:   p.ID,
		Project:     p,
		Labels:      []domain.Label{l},
	}
	iucm.On("Update", uint(1), i.Title, i.Description, i.Status, map[string]domain.Label{
		relay.ToGlobalID("Label", "1"): l,
	}).Return(i, errors.New("test error"))

	inputMap := map[string]interface{}{
		"id":          relay.ToGlobalID("Issue", "1"),
		"title":       i.Title,
		"description": i.Description,
		"status":      i.Status,
		"labels":      relay.ToGlobalID("Label", "1"),
	}

	result, err := r.MutateAndGetPayloadForUpdateIssueMutation(nil, inputMap, graphql.ResolveInfo{})

	assert.NotNil(t, err)
	assert.NotNil(t, result)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestMutateAndGetPayloadForRemoveIssueMutation(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	iucm.On("Remove", uint(1)).Return(true, nil)

	inputMap := map[string]interface{}{
		"id": relay.ToGlobalID("Issue", "1"),
	}

	result, err := r.MutateAndGetPayloadForRemoveIssueMutation(nil, inputMap, graphql.ResolveInfo{})

	assert.Nil(t, err)
	assert.NotNil(t, result)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestMutateAndGetPayloadForRemoveIssueMutationArgErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	tests := []struct {
		id interface{}
	}{
		{
			nil,
		},
		{
			"",
		},
		{
			relay.ToGlobalID("Issue", "test"),
		},
	}

	for _, ts := range tests {
		inputMap := map[string]interface{}{}
		if ts.id != nil {
			inputMap = map[string]interface{}{
				"id": ts.id,
			}
		}
		result, err := r.MutateAndGetPayloadForRemoveIssueMutation(nil, inputMap, graphql.ResolveInfo{})

		assert.NotNil(t, err)
		assert.NotNil(t, result)

		checkAssertions(t, cucm, iucm, lucm, pucm)
	}
}

func TestMutateAndGetPayloadForRemoveIssueMutationErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	iucm.On("Remove", uint(1)).Return(false, errors.New("test error"))

	inputMap := map[string]interface{}{
		"id": relay.ToGlobalID("Issue", "1"),
	}

	result, err := r.MutateAndGetPayloadForRemoveIssueMutation(nil, inputMap, graphql.ResolveInfo{})

	assert.NotNil(t, err)
	assert.NotNil(t, result)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestMutateAndGetPayloadForAddLabelMutation(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	l := new(domain.Label)
	l.Name = "test-name"
	l.ColorHexCode = "FFFFFF"

	lucm.On("Add", "test-name", "FFFFFF").Return(l, nil)

	inputMap := map[string]interface{}{
		"name":         l.Name,
		"colorHexCode": l.ColorHexCode,
	}

	result, err := r.MutateAndGetPayloadForAddLabelMutation(nil, inputMap, graphql.ResolveInfo{})

	assert.Nil(t, err)
	assert.NotNil(t, result)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestMutateAndGetPayloadForAddLabelMutationColorHexCodeNotProvided(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	c := domain.Color{
		HexCode: "FFFFFF",
	}
	cucm.On("GetColor").Return(c, nil)

	l := new(domain.Label)
	l.Name = "test-name"
	l.ColorHexCode = "FFFFFF"

	lucm.On("Add", "test-name", "FFFFFF").Return(l, nil)

	inputMap := map[string]interface{}{
		"name":         l.Name,
		"colorHexCode": "",
	}

	result, err := r.MutateAndGetPayloadForAddLabelMutation(nil, inputMap, graphql.ResolveInfo{})

	assert.Nil(t, err)
	assert.NotNil(t, result)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestMutateAndGetPayloadForAddLabelMutationColorHexCodeNotProvidedErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	c := domain.Color{
		HexCode: "FFFFFF",
	}
	cucm.On("GetColor").Return(c, errors.New("test error"))

	inputMap := map[string]interface{}{
		"name":         "test-name",
		"colorHexCode": "",
	}

	result, err := r.MutateAndGetPayloadForAddLabelMutation(nil, inputMap, graphql.ResolveInfo{})

	assert.NotNil(t, err)
	assert.NotNil(t, result)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestMutateAndGetPayloadForAddLabelMutationArgErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	inputMap := map[string]interface{}{
		"name": "",
	}

	result, err := r.MutateAndGetPayloadForAddLabelMutation(nil, inputMap, graphql.ResolveInfo{})

	assert.NotNil(t, err)
	assert.NotNil(t, result)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestMutateAndGetPayloadForAddLabelMutationErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	l := new(domain.Label)
	l.Name = "test-name"
	l.ColorHexCode = "FFFFFF"

	lucm.On("Add", "test-name", "FFFFFF").Return(l, errors.New("test error"))

	inputMap := map[string]interface{}{
		"name":         l.Name,
		"colorHexCode": l.ColorHexCode,
	}

	result, err := r.MutateAndGetPayloadForAddLabelMutation(nil, inputMap, graphql.ResolveInfo{})

	assert.NotNil(t, err)
	assert.NotNil(t, result)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestMutateAndGetPayloadForUpdateLabelMutation(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	l := domain.Label{
		ID:           1,
		Name:         "test-name",
		ColorHexCode: "FFFFFF",
	}

	lucm.On("Update", uint(1), "test-name", "FFFFFF").Return(l, nil)

	inputMap := map[string]interface{}{
		"id":           relay.ToGlobalID("Label", "1"),
		"name":         l.Name,
		"colorHexCode": l.ColorHexCode,
	}

	result, err := r.MutateAndGetPayloadForUpdateLabelMutation(nil, inputMap, graphql.ResolveInfo{})

	assert.Nil(t, err)
	assert.NotNil(t, result)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestMutateAndGetPayloadForUpdateLabelMutationColorHexCodeNotProvided(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	c := domain.Color{
		HexCode: "FFFFFF",
	}
	cucm.On("GetColor").Return(c, nil)

	l := domain.Label{
		ID:           1,
		Name:         "test-name",
		ColorHexCode: "FFFFFF",
	}

	lucm.On("Update", uint(1), "test-name", "FFFFFF").Return(l, nil)

	inputMap := map[string]interface{}{
		"id":           relay.ToGlobalID("Label", "1"),
		"name":         l.Name,
		"colorHexCode": "",
	}

	result, err := r.MutateAndGetPayloadForUpdateLabelMutation(nil, inputMap, graphql.ResolveInfo{})

	assert.Nil(t, err)
	assert.NotNil(t, result)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestMutateAndGetPayloadForUpdateLabelMutationColorHexCodeNotProvidedErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	c := domain.Color{
		HexCode: "FFFFFF",
	}
	cucm.On("GetColor").Return(c, errors.New("test error"))

	inputMap := map[string]interface{}{
		"id":           relay.ToGlobalID("Label", "1"),
		"name":         "test-name",
		"colorHexCode": "",
	}

	result, err := r.MutateAndGetPayloadForUpdateLabelMutation(nil, inputMap, graphql.ResolveInfo{})

	assert.NotNil(t, err)
	assert.NotNil(t, result)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestMutateAndGetPayloadForUpdateLabelMutationArgErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	tests := []struct {
		id interface{}
	}{
		{
			nil,
		},
		{
			"",
		},
		{
			relay.ToGlobalID("Label", "test"),
		},
	}

	for _, ts := range tests {
		inputMap := map[string]interface{}{}
		if ts.id != nil {
			inputMap = map[string]interface{}{
				"id": ts.id,
			}
		}
		result, err := r.MutateAndGetPayloadForUpdateLabelMutation(nil, inputMap, graphql.ResolveInfo{})

		assert.NotNil(t, err)
		assert.NotNil(t, result)

		checkAssertions(t, cucm, iucm, lucm, pucm)
	}
}

func TestMutateAndGetPayloadForUpdateLabelMutationOtherArgErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	inputMap := map[string]interface{}{
		"id":   relay.ToGlobalID("Label", "1"),
		"name": "",
	}

	result, err := r.MutateAndGetPayloadForUpdateLabelMutation(nil, inputMap, graphql.ResolveInfo{})

	assert.NotNil(t, err)
	assert.NotNil(t, result)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestMutateAndGetPayloadForUpdateLabelMutationErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	l := domain.Label{
		ID:           1,
		Name:         "test-name",
		ColorHexCode: "FFFFFF",
	}

	lucm.On("Update", uint(1), "test-name", "FFFFFF").Return(l, errors.New("test error"))

	inputMap := map[string]interface{}{
		"id":           relay.ToGlobalID("Label", "1"),
		"name":         l.Name,
		"colorHexCode": l.ColorHexCode,
	}

	result, err := r.MutateAndGetPayloadForUpdateLabelMutation(nil, inputMap, graphql.ResolveInfo{})

	assert.NotNil(t, err)
	assert.NotNil(t, result)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestMutateAndGetPayloadForRemoveLabelMutation(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	lucm.On("Remove", uint(1)).Return(true, nil)

	inputMap := map[string]interface{}{
		"id": relay.ToGlobalID("Label", "1"),
	}

	result, err := r.MutateAndGetPayloadForRemoveLabelMutation(nil, inputMap, graphql.ResolveInfo{})

	assert.Nil(t, err)
	assert.NotNil(t, result)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestMutateAndGetPayloadForRemoveLabelMutationArgErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	tests := []struct {
		id interface{}
	}{
		{
			nil,
		},
		{
			"",
		},
		{
			relay.ToGlobalID("Label", "test"),
		},
	}

	for _, ts := range tests {
		inputMap := map[string]interface{}{}
		if ts.id != nil {
			inputMap = map[string]interface{}{
				"id": ts.id,
			}
		}
		result, err := r.MutateAndGetPayloadForRemoveLabelMutation(nil, inputMap, graphql.ResolveInfo{})

		assert.NotNil(t, err)
		assert.NotNil(t, result)

		checkAssertions(t, cucm, iucm, lucm, pucm)
	}
}

func TestMutateAndGetPayloadForRemoveLabelMutationErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	lucm.On("Remove", uint(1)).Return(false, errors.New("test error"))

	inputMap := map[string]interface{}{
		"id": relay.ToGlobalID("Label", "1"),
	}

	result, err := r.MutateAndGetPayloadForRemoveLabelMutation(nil, inputMap, graphql.ResolveInfo{})

	assert.NotNil(t, err)
	assert.NotNil(t, result)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestMutateAndGetPayloadForAddProjectMutation(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	p := new(domain.Project)
	p.Name = "test-name"
	p.Description = "test-description"

	pucm.On("Add", "test-name", "test-description").Return(p, nil)

	inputMap := map[string]interface{}{
		"name":        p.Name,
		"description": p.Description,
	}

	result, err := r.MutateAndGetPayloadForAddProjectMutation(nil, inputMap, graphql.ResolveInfo{})

	assert.Nil(t, err)
	assert.NotNil(t, result)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestMutateAndGetPayloadForAddProjectMutationArgErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	inputMap := map[string]interface{}{
		"name": "",
	}

	result, err := r.MutateAndGetPayloadForAddProjectMutation(nil, inputMap, graphql.ResolveInfo{})

	assert.NotNil(t, err)
	assert.NotNil(t, result)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestMutateAndGetPayloadForAddProjectMutationErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	p := new(domain.Project)
	p.Name = "test-name"
	p.Description = "test-description"

	pucm.On("Add", "test-name", "test-description").Return(p, errors.New("test error"))

	inputMap := map[string]interface{}{
		"name":        p.Name,
		"description": p.Description,
	}

	result, err := r.MutateAndGetPayloadForAddProjectMutation(nil, inputMap, graphql.ResolveInfo{})

	assert.NotNil(t, err)
	assert.NotNil(t, result)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestMutateAndGetPayloadForUpdateProjectMutation(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	p := domain.Project{
		ID:          1,
		Name:        "test-name",
		Description: "test-description",
	}

	pucm.On("Update", uint(1), "test-name", "test-description").Return(p, nil)

	inputMap := map[string]interface{}{
		"id":          relay.ToGlobalID("Project", "1"),
		"name":        p.Name,
		"description": p.Description,
	}

	result, err := r.MutateAndGetPayloadForUpdateProjectMutation(nil, inputMap, graphql.ResolveInfo{})

	assert.Nil(t, err)
	assert.NotNil(t, result)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestMutateAndGetPayloadForUpdateProjectMutationArgErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	tests := []struct {
		id interface{}
	}{
		{
			nil,
		},
		{
			"",
		},
		{
			relay.ToGlobalID("Project", "test"),
		},
	}

	for _, ts := range tests {
		inputMap := map[string]interface{}{}
		if ts.id != nil {
			inputMap = map[string]interface{}{
				"id": ts.id,
			}
		}
		result, err := r.MutateAndGetPayloadForUpdateProjectMutation(nil, inputMap, graphql.ResolveInfo{})

		assert.NotNil(t, err)
		assert.NotNil(t, result)

		checkAssertions(t, cucm, iucm, lucm, pucm)
	}
}

func TestMutateAndGetPayloadForUpdateProjectMutationOtherArgErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	inputMap := map[string]interface{}{
		"id":   relay.ToGlobalID("Project", "1"),
		"name": "",
	}

	result, err := r.MutateAndGetPayloadForUpdateProjectMutation(nil, inputMap, graphql.ResolveInfo{})

	assert.NotNil(t, err)
	assert.NotNil(t, result)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestMutateAndGetPayloadForUpdateProjectMutationErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	p := domain.Project{
		ID:          1,
		Name:        "test-name",
		Description: "test-description",
	}

	pucm.On("Update", uint(1), "test-name", "test-description").Return(p, errors.New("test error"))

	inputMap := map[string]interface{}{
		"id":          relay.ToGlobalID("Project", "1"),
		"name":        p.Name,
		"description": p.Description,
	}

	result, err := r.MutateAndGetPayloadForUpdateProjectMutation(nil, inputMap, graphql.ResolveInfo{})

	assert.NotNil(t, err)
	assert.NotNil(t, result)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestMutateAndGetPayloadForRemoveProjectMutation(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	pucm.On("Remove", uint(1)).Return(true, nil)

	inputMap := map[string]interface{}{
		"id": relay.ToGlobalID("Project", "1"),
	}

	result, err := r.MutateAndGetPayloadForRemoveProjectMutation(nil, inputMap, graphql.ResolveInfo{})

	assert.Nil(t, err)
	assert.NotNil(t, result)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestMutateAndGetPayloadForRemoveProjectMutationArgErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	tests := []struct {
		id interface{}
	}{
		{
			nil,
		},
		{
			"",
		},
		{
			relay.ToGlobalID("Project", "test"),
		},
	}

	for _, ts := range tests {
		inputMap := map[string]interface{}{}
		if ts.id != nil {
			inputMap = map[string]interface{}{
				"id": ts.id,
			}
		}
		result, err := r.MutateAndGetPayloadForRemoveProjectMutation(nil, inputMap, graphql.ResolveInfo{})

		assert.NotNil(t, err)
		assert.NotNil(t, result)

		checkAssertions(t, cucm, iucm, lucm, pucm)
	}
}

func TestMutateAndGetPayloadForRemoveProjectMutationErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	pucm.On("Remove", uint(1)).Return(false, errors.New("test error"))

	inputMap := map[string]interface{}{
		"id": relay.ToGlobalID("Project", "1"),
	}

	result, err := r.MutateAndGetPayloadForRemoveProjectMutation(nil, inputMap, graphql.ResolveInfo{})

	assert.NotNil(t, err)
	assert.NotNil(t, result)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveFindIssueByIDQuery(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	i := domain.Issue{
		ID:          1,
		Title:       "test-title",
		Description: "test-description",
		Status:      1,
		ProjectID:   1,
	}

	iucm.On("FindByID", i.ID).Return(i, nil)

	rp := graphql.ResolveParams{
		Args: map[string]interface{}{
			"id": relay.ToGlobalID("Issue", "1"),
		},
	}

	item, err := r.ResolveFindIssueByIDQuery(rp)

	assert.Nil(t, err)
	assert.NotNil(t, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveFindIssueByIDQueryArgNotValid(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	rp := graphql.ResolveParams{
		Args: map[string]interface{}{
			"id": relay.ToGlobalID("Issue", "test"),
		},
	}

	item, err := r.ResolveFindIssueByIDQuery(rp)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveFindIssueByIDQueryArgErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	rp := graphql.ResolveParams{
		Args: map[string]interface{}{
			"idtest": relay.ToGlobalID("Issue", "1"),
		},
	}

	item, err := r.ResolveFindIssueByIDQuery(rp)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveFindIssueByIDQueryNotFoundErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	i := domain.Issue{
		ID:          1,
		Title:       "test-title",
		Description: "test-description",
		Status:      1,
		ProjectID:   1,
	}

	iucm.On("FindByID", i.ID).Return(i, errors.New("record not found"))

	rp := graphql.ResolveParams{
		Args: map[string]interface{}{
			"id": relay.ToGlobalID("Issue", "1"),
		},
	}

	item, err := r.ResolveFindIssueByIDQuery(rp)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveFindIssueByIDQueryErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	i := domain.Issue{
		ID:          1,
		Title:       "test-title",
		Description: "test-description",
		Status:      1,
		ProjectID:   1,
	}

	iucm.On("FindByID", i.ID).Return(i, errors.New("test error"))

	rp := graphql.ResolveParams{
		Args: map[string]interface{}{
			"id": relay.ToGlobalID("Issue", "1"),
		},
	}

	item, err := r.ResolveFindIssueByIDQuery(rp)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveFindIssuesQuery(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	i := []domain.Issue{}

	iucm.On("Find", "test-title", uint(1), []string{"1"}).Return(i, nil)

	rp := graphql.ResolveParams{
		Args: map[string]interface{}{
			"title":     "test-title",
			"projectId": relay.ToGlobalID("Project", "1"),
			"labels":    relay.ToGlobalID("Label", "1"),
		}}

	item, err := r.ResolveFindIssuesQuery(rp)

	assert.Nil(t, err)
	assert.NotNil(t, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveFindIssuesQueryArgErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	tests := []struct {
		projectID interface{}
		labels    string
	}{
		{
			"test",
			relay.ToGlobalID("Label", "1"),
		},
		{
			relay.ToGlobalID("Project", "test"),
			relay.ToGlobalID("Label", "1"),
		},
		{
			relay.ToGlobalID("Project", "1"),
			"test",
		},
	}

	for _, ts := range tests {
		rp := graphql.ResolveParams{
			Args: map[string]interface{}{
				"title":     "test-title",
				"projectId": ts.projectID,
				"labels":    ts.labels,
			}}

		item, err := r.ResolveFindIssuesQuery(rp)

		assert.NotNil(t, err)
		assert.Nil(t, item)

		checkAssertions(t, cucm, iucm, lucm, pucm)
	}
}

func TestResolveFindIssuesQueryErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	i := []domain.Issue{}

	iucm.On("Find", "test-title", uint(1), []string{"1"}).Return(i, errors.New("test error"))

	rp := graphql.ResolveParams{
		Args: map[string]interface{}{
			"title":     "test-title",
			"projectId": relay.ToGlobalID("Project", "1"),
			"labels":    relay.ToGlobalID("Label", "1"),
		},
	}

	item, err := r.ResolveFindIssuesQuery(rp)

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, []domain.Issue{}, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveFindAllIssuesQuery(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	i := []domain.Issue{}

	iucm.On("FindAll").Return(i, nil)

	rp := graphql.ResolveParams{
		Args: map[string]interface{}{},
	}

	item, err := r.ResolveFindAllIssuesQuery(rp)

	assert.Nil(t, err)
	assert.NotNil(t, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveFindAllIssuesQueryErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	i := []domain.Issue{}

	iucm.On("FindAll").Return(i, errors.New("test error"))

	rp := graphql.ResolveParams{
		Args: map[string]interface{}{},
	}

	item, err := r.ResolveFindAllIssuesQuery(rp)

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, []domain.Issue{}, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveFindLabelByIDQuery(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	l := domain.Label{
		ID:           1,
		Name:         "test-name",
		ColorHexCode: "FFFFFF",
	}

	lucm.On("FindByID", l.ID).Return(l, nil)

	rp := graphql.ResolveParams{
		Args: map[string]interface{}{
			"id": relay.ToGlobalID("Label", "1"),
		},
	}

	item, err := r.ResolveFindLabelByIDQuery(rp)

	assert.Nil(t, err)
	assert.NotNil(t, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveFindLabelByIDQueryArgErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	tests := []struct {
		id interface{}
	}{
		{
			nil,
		},
		{
			"",
		},
		{
			relay.ToGlobalID("Label", "test"),
		},
	}

	for _, ts := range tests {
		rp := graphql.ResolveParams{
			Args: map[string]interface{}{},
		}
		if ts.id != nil {
			rp = graphql.ResolveParams{
				Args: map[string]interface{}{
					"id": ts.id,
				},
			}
		}

		item, err := r.ResolveFindLabelByIDQuery(rp)

		assert.NotNil(t, err)
		assert.Nil(t, item)

		checkAssertions(t, cucm, iucm, lucm, pucm)
	}
}

func TestResolveFindLabelByIDQueryNotFoundErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	l := domain.Label{
		ID:           1,
		Name:         "test-name",
		ColorHexCode: "FFFFFF",
	}

	lucm.On("FindByID", l.ID).Return(l, errors.New("record not found"))

	rp := graphql.ResolveParams{
		Args: map[string]interface{}{
			"id": relay.ToGlobalID("Label", "1"),
		},
	}

	item, err := r.ResolveFindLabelByIDQuery(rp)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveFindLabelByIDQueryErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	l := domain.Label{
		ID:           1,
		Name:         "test-name",
		ColorHexCode: "FFFFFF",
	}

	lucm.On("FindByID", l.ID).Return(l, errors.New("test error"))

	rp := graphql.ResolveParams{
		Args: map[string]interface{}{
			"id": relay.ToGlobalID("Label", "1"),
		},
	}

	item, err := r.ResolveFindLabelByIDQuery(rp)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveFindLabelsQuery(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	l := []domain.Label{
		domain.Label{
			ID:           1,
			Name:         "test-name-1",
			ColorHexCode: "FFFFFF",
		},
		domain.Label{
			ID:           2,
			Name:         "test-name-2",
			ColorHexCode: "FFFFFF",
		},
	}

	lucm.On("Find", "test").Return(l, nil)

	rp := graphql.ResolveParams{
		Args: map[string]interface{}{
			"name": "test",
		},
	}

	item, err := r.ResolveFindLabelsQuery(rp)

	assert.Nil(t, err)
	assert.NotNil(t, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveFindLabelsQueryErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	l := []domain.Label{
		domain.Label{
			ID:           1,
			Name:         "test-name-1",
			ColorHexCode: "FFFFFF",
		},
		domain.Label{
			ID:           2,
			Name:         "test-name-2",
			ColorHexCode: "FFFFFF",
		},
	}

	lucm.On("Find", "test").Return(l, errors.New("test error"))

	rp := graphql.ResolveParams{
		Args: map[string]interface{}{
			"name": "test",
		},
	}

	item, err := r.ResolveFindLabelsQuery(rp)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveFindAllLabelsQuery(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	l := []domain.Label{
		domain.Label{
			ID:           1,
			Name:         "test-name",
			ColorHexCode: "FFFFFF",
		},
	}

	lucm.On("FindAll").Return(l, nil)

	rp := graphql.ResolveParams{
		Args: map[string]interface{}{},
	}

	item, err := r.ResolveFindAllLabelsQuery(rp)

	assert.Nil(t, err)
	assert.NotNil(t, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveFindAllLabelsQueryErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	lucm.On("FindAll").Return([]domain.Label{}, errors.New("test error"))

	rp := graphql.ResolveParams{
		Args: map[string]interface{}{},
	}

	item, err := r.ResolveFindAllLabelsQuery(rp)

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, []domain.Label{}, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveFindProjectByIDQuery(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	p := domain.Project{
		ID:          1,
		Name:        "test-name",
		Description: "test-description",
	}

	pucm.On("FindByID", p.ID).Return(p, nil)

	rp := graphql.ResolveParams{
		Args: map[string]interface{}{
			"id": relay.ToGlobalID("Project", "1"),
		},
	}

	item, err := r.ResolveFindProjectByIDQuery(rp)

	assert.Nil(t, err)
	assert.NotNil(t, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveFindProjectByIDQueryArgErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	rp := graphql.ResolveParams{
		Args: map[string]interface{}{
			"idtest": relay.ToGlobalID("Project", "1"),
		},
	}

	item, err := r.ResolveFindProjectByIDQuery(rp)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveFindProjectByIDQueryNotFoundErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	p := domain.Project{
		ID:          1,
		Name:        "test-name",
		Description: "test-description",
	}

	pucm.On("FindByID", p.ID).Return(p, errors.New("record not found"))

	rp := graphql.ResolveParams{
		Args: map[string]interface{}{
			"id": relay.ToGlobalID("Project", "1"),
		},
	}

	item, err := r.ResolveFindProjectByIDQuery(rp)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveFindProjectByIDQueryErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	p := domain.Project{
		ID:          1,
		Name:        "test-name",
		Description: "test-description",
	}

	pucm.On("FindByID", p.ID).Return(p, errors.New("test error"))

	rp := graphql.ResolveParams{
		Args: map[string]interface{}{
			"id": relay.ToGlobalID("Project", "1"),
		},
	}

	item, err := r.ResolveFindProjectByIDQuery(rp)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveFindProjectsQuery(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	p := []domain.Project{
		domain.Project{
			ID:          1,
			Name:        "test-name",
			Description: "test-description",
		},
	}

	pucm.On("Find", "test-name").Return(p, nil)

	rp := graphql.ResolveParams{
		Args: map[string]interface{}{
			"name": "test-name",
		},
	}

	item, err := r.ResolveFindProjectsQuery(rp)

	assert.Nil(t, err)
	assert.NotNil(t, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveFindProjectsQueryErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	pucm.On("Find", "test-name").Return([]domain.Project{}, errors.New("test error"))

	rp := graphql.ResolveParams{
		Args: map[string]interface{}{
			"name": "test-name",
		},
	}

	items, err := r.ResolveFindProjectsQuery(rp)

	assert.NotNil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, []domain.Project{}, items)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveFindAllProjectsQuery(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	p := []domain.Project{
		domain.Project{
			ID:          1,
			Name:        "test-name",
			Description: "test-description",
		},
	}

	pucm.On("FindAll").Return(p, nil)

	rp := graphql.ResolveParams{
		Args: map[string]interface{}{},
	}

	item, err := r.ResolveFindAllProjectsQuery(rp)

	assert.Nil(t, err)
	assert.NotNil(t, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestResolveFindAllProjectsQueryErr(t *testing.T) {
	cucm, iucm, lucm, pucm, r := prepareMocksAndResolver()

	pucm.On("FindAll").Return([]domain.Project{}, errors.New("test error"))

	rp := graphql.ResolveParams{
		Args: map[string]interface{}{},
	}

	item, err := r.ResolveFindAllProjectsQuery(rp)

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, []domain.Project{}, item)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}
