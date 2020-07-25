package gql

import (
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"go-issue-tracker/pkg/domain"
	"go-issue-tracker/pkg/usecases"
	"golang.org/x/net/context"
	"strconv"
	"strings"
)

// Resolver interface
type Resolver interface {
	ResolveNodeID(context context.Context, id string, info graphql.ResolveInfo) (interface{}, error)
	ResolveType(p graphql.ResolveTypeParams) *graphql.Object
	ResolveFieldLabels(p graphql.ResolveParams) (interface{}, error)
	ResolveFindIssueByIDQuery(p graphql.ResolveParams) (interface{}, error)
	ResolveFindIssuesQuery(p graphql.ResolveParams) (interface{}, error)
	ResolveFindAllIssuesQuery(p graphql.ResolveParams) (interface{}, error)
	ResolveFindLabelByIDQuery(p graphql.ResolveParams) (interface{}, error)
	ResolveFindLabelsQuery(p graphql.ResolveParams) (interface{}, error)
	ResolveFindAllLabelsQuery(p graphql.ResolveParams) (interface{}, error)
	ResolveFindProjectByIDQuery(p graphql.ResolveParams) (interface{}, error)
	ResolveFindProjectsQuery(p graphql.ResolveParams) (interface{}, error)
	ResolveFindAllProjectsQuery(p graphql.ResolveParams) (interface{}, error)
	ResolveMutationOutputFieldItem(p graphql.ResolveParams) (interface{}, error)
	ResolveMutationOutputFieldItemID(p graphql.ResolveParams) (interface{}, error)
	ResolveMutationOutputFieldStatus(p graphql.ResolveParams) (interface{}, error)
	MutateAndGetPayloadForAddIssueMutation(ctx context.Context, inputMap map[string]interface{}, info graphql.ResolveInfo) (map[string]interface{}, error)
	MutateAndGetPayloadForUpdateIssueMutation(ctx context.Context, inputMap map[string]interface{}, info graphql.ResolveInfo) (map[string]interface{}, error)
	MutateAndGetPayloadForRemoveIssueMutation(ctx context.Context, inputMap map[string]interface{}, info graphql.ResolveInfo) (map[string]interface{}, error)
	MutateAndGetPayloadForAddProjectMutation(ctx context.Context, inputMap map[string]interface{}, info graphql.ResolveInfo) (map[string]interface{}, error)
	MutateAndGetPayloadForUpdateProjectMutation(ctx context.Context, inputMap map[string]interface{}, info graphql.ResolveInfo) (map[string]interface{}, error)
	MutateAndGetPayloadForRemoveProjectMutation(ctx context.Context, inputMap map[string]interface{}, info graphql.ResolveInfo) (map[string]interface{}, error)
	MutateAndGetPayloadForAddLabelMutation(ctx context.Context, inputMap map[string]interface{}, info graphql.ResolveInfo) (map[string]interface{}, error)
	MutateAndGetPayloadForUpdateLabelMutation(ctx context.Context, inputMap map[string]interface{}, info graphql.ResolveInfo) (map[string]interface{}, error)
	MutateAndGetPayloadForRemoveLabelMutation(ctx context.Context, inputMap map[string]interface{}, info graphql.ResolveInfo) (map[string]interface{}, error)
}

// resolver contains base tooling like GraphQL use case etc.
type resolver struct {
	iuc usecases.IssueUseCase
	luc usecases.LabelUseCase
	puc usecases.ProjectUseCase
	cuc usecases.ColorUseCase
}

// GetResolver to init Resolver
func GetResolver(iuc usecases.IssueUseCase, luc usecases.LabelUseCase, puc usecases.ProjectUseCase, cuc usecases.ColorUseCase) Resolver {
	return &resolver{
		iuc: iuc,
		luc: luc,
		puc: puc,
		cuc: cuc,
	}
}

// ResolveNodeID as IDFetcher func
func (r *resolver) ResolveNodeID(context context.Context, id string, info graphql.ResolveInfo) (interface{}, error) {
	resolvedID := relay.FromGlobalID(id)
	if resolvedID == nil {
		return nil, errors.New("provided id not valid")
	}
	intID, err := strconv.Atoi(resolvedID.ID)
	if err != nil {
		return nil, err
	}

	if resolvedID.Type == "Issue" {
		return r.iuc.FindByID(uint(intID))
	} else if resolvedID.Type == "Label" {
		return r.luc.FindByID(uint(intID))
	} else if resolvedID.Type == "Project" {
		return r.puc.FindByID(uint(intID))
	}

	return nil, errors.New("unknown type")
}

// ResolveType to resolve type
func (r *resolver) ResolveType(p graphql.ResolveTypeParams) *graphql.Object {
	switch p.Value.(type) {
	case domain.Issue:
		return IssueType
	case *domain.Issue:
		return IssueType
	case domain.Label:
		return LabelType
	case *domain.Label:
		return LabelType
	case domain.Project:
		return ProjectType
	case *domain.Project:
		return ProjectType
	}
	return nil
}

func (r *resolver) getLabelsConnectionData(labels []domain.Label, args relay.ConnectionArguments) (*relay.Connection, error) {
	data := make([]interface{}, len(labels))
	for i, v := range labels {
		data[i] = v
	}
	return relay.ConnectionFromArray(data, args), nil
}

// ResolveFieldLabels to get labels connection
func (r *resolver) ResolveFieldLabels(p graphql.ResolveParams) (interface{}, error) {
	args := relay.NewConnectionArguments(p.Args)
	if source, ok := p.Source.(domain.Issue); ok {
		return r.getLabelsConnectionData(source.Labels, args)
	}
	if source, ok := p.Source.(*domain.Issue); ok {
		return r.getLabelsConnectionData(source.Labels, args)
	}
	return nil, errors.New("no labels found")
}

// ResolveMutationOutputFieldItem to get item for output field
func (r *resolver) ResolveMutationOutputFieldItem(p graphql.ResolveParams) (interface{}, error) {
	if source, ok := p.Source.(map[string]interface{}); ok {
		return source["item"], nil
	}
	return nil, errors.New("item not resolved")
}

// ResolveMutationOutputFieldItemID to get item ID for output field
func (r *resolver) ResolveMutationOutputFieldItemID(p graphql.ResolveParams) (interface{}, error) {
	if source, ok := p.Source.(map[string]interface{}); ok {
		return source["id"], nil
	}
	return nil, errors.New("id not resolved")
}

// ResolveMutationOutputFieldStatus to get status for output field
func (r *resolver) ResolveMutationOutputFieldStatus(p graphql.ResolveParams) (interface{}, error) {
	if source, ok := p.Source.(map[string]interface{}); ok {
		return source["status"], nil
	}
	return nil, errors.New("status not resolved")
}

func (r *resolver) getLabels(inputMap map[string]interface{}) (map[string]domain.Label, error) {
	labelValues, labelValuesOK := inputMap["labels"].(string)
	if !labelValuesOK || labelValues == "" {
		return nil, errors.New("labels not provided")
	}
	labelStrings := strings.Split(strings.Trim(labelValues, " "), ",")
	labels := make(map[string]domain.Label)
	for _, ls := range labelStrings {
		if labels[ls].ID == 0 && ls != "" {
			lID := relay.FromGlobalID(ls)
			if lID == nil {
				return nil, errors.New("provided label id not valid")
			}
			lIDInt, err := strconv.Atoi(lID.ID)
			if err != nil {
				return nil, err
			}
			label, err := r.luc.FindByID(uint(lIDInt))
			if err != nil {
				return nil, fmt.Errorf("label %s is not valid", ls)
			}
			labels[ls] = label
		}
	}

	if len(labels) == 0 {
		return nil, errors.New("no labels provided")
	}

	return labels, nil
}

// MutateAndGetPayloadForAddIssueMutation func
func (r *resolver) MutateAndGetPayloadForAddIssueMutation(ctx context.Context, inputMap map[string]interface{}, info graphql.ResolveInfo) (map[string]interface{}, error) {
	errResponse := map[string]interface{}{
		"item": nil,
	}

	title, titleOK := inputMap["title"].(string)
	if !titleOK || title == "" {
		return errResponse, errors.New("title not provided")
	}
	description, descriptionOK := inputMap["description"].(string)
	if !descriptionOK || description == "" {
		return errResponse, errors.New("description not provided")
	}
	status, statusOK := inputMap["status"].(int)
	if !statusOK || status == 0 {
		return errResponse, errors.New("status not provided")
	}
	projectID, projectIDOK := inputMap["projectId"].(string)
	if !projectIDOK || projectID == "" {
		return errResponse, errors.New("project id not provided")
	}
	resolvedID := relay.FromGlobalID(projectID)
	if resolvedID == nil {
		return errResponse, errors.New("provided project id not valid")
	}
	projectIDInt, err := strconv.Atoi(resolvedID.ID)
	if err != nil {
		return errResponse, err
	}
	project, err := r.puc.FindByID(uint(projectIDInt))
	if err != nil {
		return errResponse, errors.New("provided project id not valid")
	}
	labels, err := r.getLabels(inputMap)
	if err != nil {
		return errResponse, err
	}

	item, err := r.iuc.Add(title, description, status, project, labels)
	if err != nil {
		return map[string]interface{}{
			"item": nil,
		}, err
	}

	return map[string]interface{}{
		"item": item,
	}, nil
}

func (r *resolver) getIDFromMutationData(inputMap map[string]interface{}) (uint, error) {
	id, idOK := inputMap["id"].(string)
	if !idOK {
		return uint(0), errors.New("id not provided")
	}
	resolvedID := relay.FromGlobalID(id)
	if resolvedID == nil {
		return uint(0), errors.New("provided id not valid")
	}
	intID, err := strconv.Atoi(resolvedID.ID)
	if err != nil {
		return uint(0), err
	}
	return uint(intID), nil
}

// MutateAndGetPayloadForUpdateIssueMutation func
func (r *resolver) MutateAndGetPayloadForUpdateIssueMutation(ctx context.Context, inputMap map[string]interface{}, info graphql.ResolveInfo) (map[string]interface{}, error) {
	errResponse := map[string]interface{}{
		"item": nil,
	}

	id, err := r.getIDFromMutationData(inputMap)
	if err != nil {
		return errResponse, err
	}
	title, titleOK := inputMap["title"].(string)
	if !titleOK || title == "" {
		return errResponse, errors.New("title not provided")
	}
	description, descriptionOK := inputMap["description"].(string)
	if !descriptionOK || description == "" {
		return errResponse, errors.New("description not provided")
	}
	status, statusOK := inputMap["status"].(int)
	if !statusOK || status == 0 {
		return errResponse, errors.New("status not provided")
	}
	labels, err := r.getLabels(inputMap)
	if err != nil {
		return errResponse, err
	}

	item, err := r.iuc.Update(id, title, description, status, labels)
	if err != nil {
		return errResponse, err
	}

	return map[string]interface{}{
		"item": item,
	}, nil
}

// MutateAndGetPayloadForRemoveIssueMutation func
func (r *resolver) MutateAndGetPayloadForRemoveIssueMutation(ctx context.Context, inputMap map[string]interface{}, info graphql.ResolveInfo) (map[string]interface{}, error) {
	id, err := r.getIDFromMutationData(inputMap)
	if err != nil {
		return map[string]interface{}{
			"id":     id,
			"status": false,
		}, err
	}

	status, err := r.iuc.Remove(id)
	if err != nil {
		return map[string]interface{}{
			"id":     id,
			"status": status,
		}, err
	}

	return map[string]interface{}{
		"id":     id,
		"status": status,
	}, nil
}

// MutateAndGetPayloadForAddLabelMutation func
func (r *resolver) MutateAndGetPayloadForAddLabelMutation(ctx context.Context, inputMap map[string]interface{}, info graphql.ResolveInfo) (map[string]interface{}, error) {
	errResponse := map[string]interface{}{
		"item": nil,
	}

	name, nameOK := inputMap["name"].(string)
	if !nameOK || name == "" {
		return errResponse, errors.New("name not provided")
	}

	colorHexCode := inputMap["colorHexCode"].(string)
	if colorHexCode == "" {
		cl, err := r.cuc.GetColor()
		if err != nil {
			return errResponse, err
		}
		colorHexCode = cl.HexCode
	}

	item, err := r.luc.Add(name, colorHexCode)
	if err != nil {
		return map[string]interface{}{
			"item": nil,
		}, err
	}

	return map[string]interface{}{
		"item": item,
	}, nil
}

// MutateAndGetPayloadForUpdateLabelMutation func
func (r *resolver) MutateAndGetPayloadForUpdateLabelMutation(ctx context.Context, inputMap map[string]interface{}, info graphql.ResolveInfo) (map[string]interface{}, error) {
	errResponse := map[string]interface{}{
		"item": nil,
	}

	id, err := r.getIDFromMutationData(inputMap)
	if err != nil {
		return errResponse, err
	}

	name, nameOK := inputMap["name"].(string)
	if !nameOK || name == "" {
		return errResponse, errors.New("name not provided")
	}

	colorHexCode := inputMap["colorHexCode"].(string)
	if colorHexCode == "" {
		cl, err := r.cuc.GetColor()
		if err != nil {
			return errResponse, err
		}
		colorHexCode = cl.HexCode
	}

	item, err := r.luc.Update(id, name, colorHexCode)
	if err != nil {
		return map[string]interface{}{
			"item": item,
		}, err
	}

	return map[string]interface{}{
		"item": item,
	}, nil
}

// MutateAndGetPayloadForRemoveLabelMutation func
func (r *resolver) MutateAndGetPayloadForRemoveLabelMutation(ctx context.Context, inputMap map[string]interface{}, info graphql.ResolveInfo) (map[string]interface{}, error) {
	id, err := r.getIDFromMutationData(inputMap)
	if err != nil {
		return map[string]interface{}{
			"id":     id,
			"status": false,
		}, err
	}

	status, err := r.luc.Remove(id)
	if err != nil {
		return map[string]interface{}{
			"id":     id,
			"status": status,
		}, err
	}

	return map[string]interface{}{
		"id":     id,
		"status": status,
	}, nil
}

// MutateAndGetPayloadForAddProjectMutation func
func (r *resolver) MutateAndGetPayloadForAddProjectMutation(ctx context.Context, inputMap map[string]interface{}, info graphql.ResolveInfo) (map[string]interface{}, error) {
	errResponse := map[string]interface{}{
		"item": nil,
	}

	name, nameOK := inputMap["name"].(string)
	if !nameOK || name == "" {
		return errResponse, errors.New("name not provided")
	}

	description := inputMap["description"].(string)

	item, err := r.puc.Add(name, description)
	if err != nil {
		return errResponse, err
	}

	return map[string]interface{}{
		"item": item,
	}, nil
}

// MutateAndGetPayloadForUpdateProjectMutation func
func (r *resolver) MutateAndGetPayloadForUpdateProjectMutation(ctx context.Context, inputMap map[string]interface{}, info graphql.ResolveInfo) (map[string]interface{}, error) {
	errResponse := map[string]interface{}{
		"item": nil,
	}

	id, err := r.getIDFromMutationData(inputMap)
	if err != nil {
		return errResponse, err
	}

	name, nameOK := inputMap["name"].(string)
	if !nameOK || name == "" {
		return errResponse, errors.New("name not provided")
	}

	description := inputMap["description"].(string)

	item, err := r.puc.Update(id, name, description)
	if err != nil {
		return map[string]interface{}{
			"item": item,
		}, err
	}

	return map[string]interface{}{
		"item": item,
	}, nil
}

// MutateAndGetPayloadForRemoveProjectMutation func
func (r *resolver) MutateAndGetPayloadForRemoveProjectMutation(ctx context.Context, inputMap map[string]interface{}, info graphql.ResolveInfo) (map[string]interface{}, error) {
	id, err := r.getIDFromMutationData(inputMap)
	if err != nil {
		return map[string]interface{}{
			"id":     id,
			"status": false,
		}, err
	}

	status, err := r.puc.Remove(id)
	if err != nil {
		return map[string]interface{}{
			"id":     id,
			"status": status,
		}, err
	}

	return map[string]interface{}{
		"id":     id,
		"status": status,
	}, nil
}

func (r *resolver) getIDFromQueryData(p graphql.ResolveParams) (uint, error) {
	id, idOK := p.Args["id"].(string)
	if !idOK {
		return uint(0), errors.New("id not provided")
	}

	resolvedID := relay.FromGlobalID(id)
	if resolvedID == nil {
		return uint(0), errors.New("provided id not valid")
	}
	intID, err := strconv.Atoi(resolvedID.ID)
	if err != nil {
		return uint(0), err
	}

	return uint(intID), err
}

func (r *resolver) ResolveFindIssueByIDQuery(p graphql.ResolveParams) (interface{}, error) {
	id, err := r.getIDFromQueryData(p)
	if err != nil {
		return nil, err
	}

	item, err := r.iuc.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *resolver) ResolveFindIssuesQuery(p graphql.ResolveParams) (interface{}, error) {
	title := p.Args["title"].(string)
	projectID, projectIDOK := p.Args["projectId"].(string)
	projectIDInt := 0
	if projectIDOK && projectID != "" && projectID != "0" {
		resolvedID := relay.FromGlobalID(projectID)
		if resolvedID == nil {
			return nil, errors.New("provided project id not valid")
		}
		var err error
		projectIDInt, err = strconv.Atoi(resolvedID.ID)
		if err != nil {
			return nil, err
		}
	}
	labelValues := p.Args["labels"].(string)
	labels := []string{}
	if labelValues != "" {
		labelStrings := strings.Split(strings.Trim(labelValues, " "), ",")
		for _, ls := range labelStrings {
			if ls != "" {
				resolvedID := relay.FromGlobalID(ls)
				if resolvedID == nil {
					return nil, errors.New("provided label id not valid")
				}
				labels = append(labels, resolvedID.ID)
			}
		}
	}

	items, err := r.iuc.Find(title, uint(projectIDInt), labels)
	if err != nil {
		return items, err
	}
	return items, nil
}

func (r *resolver) ResolveFindAllIssuesQuery(p graphql.ResolveParams) (interface{}, error) {
	items, err := r.iuc.FindAll()
	if err != nil {
		return items, err
	}
	return items, nil
}

func (r *resolver) ResolveFindLabelByIDQuery(p graphql.ResolveParams) (interface{}, error) {
	id, err := r.getIDFromQueryData(p)
	if err != nil {
		return nil, err
	}

	item, err := r.luc.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *resolver) ResolveFindLabelsQuery(p graphql.ResolveParams) (interface{}, error) {
	name := p.Args["name"].(string)

	items, err := r.luc.Find(name)
	if err != nil {
		return nil, err
	}

	return &items, nil
}

func (r *resolver) ResolveFindAllLabelsQuery(p graphql.ResolveParams) (interface{}, error) {
	items, err := r.luc.FindAll()
	if err != nil {
		return items, err
	}
	return items, nil
}

func (r *resolver) ResolveFindProjectByIDQuery(p graphql.ResolveParams) (interface{}, error) {
	id, err := r.getIDFromQueryData(p)
	if err != nil {
		return nil, err
	}

	item, err := r.puc.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *resolver) ResolveFindProjectsQuery(p graphql.ResolveParams) (interface{}, error) {
	name := p.Args["name"].(string)

	items, err := r.puc.Find(name)
	if err != nil {
		return items, err
	}

	return &items, nil
}

func (r *resolver) ResolveFindAllProjectsQuery(p graphql.ResolveParams) (interface{}, error) {
	items, err := r.puc.FindAll()
	if err != nil {
		return items, err
	}
	return items, nil
}
