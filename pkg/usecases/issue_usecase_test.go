package usecases_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go-issue-tracker/pkg/domain"
	dTesting "go-issue-tracker/pkg/domain/testing"
	"go-issue-tracker/pkg/usecases"
	"testing"
)

func TestUseCaseIssueNewIssueUseCase(t *testing.T) {
	ms := new(dTesting.IssueServiceMock)
	domain.GetDefaultIssueService = func(r domain.IssueRepository) domain.IssueService {
		return ms
	}
	defer domain.ResetDefaultIssueService()

	mr := new(dTesting.IssueRepositoryMock)

	uc := usecases.NewIssueUseCase(mr)

	assert.NotNil(t, uc)
}

func TestUseCaseIssueAdd(t *testing.T) {
	p := domain.Project{
		ID:          1,
		Name:        "test-name",
		Description: "test-description",
	}
	l := map[string]domain.Label{
		"test-name": domain.Label{
			ID:   1,
			Name: "test-name",
		},
	}
	v := new(domain.Issue)
	v.Title = "test-title"
	v.Description = "test-description"
	v.Status = 1
	v.ProjectID = p.ID
	v.Project = p
	v.Labels = []domain.Label{
		domain.Label{
			ID:   1,
			Name: "test-name",
		},
	}

	ms := new(dTesting.IssueServiceMock)
	ms.On("Add", v).Return(v, nil)
	domain.GetDefaultIssueService = func(r domain.IssueRepository) domain.IssueService {
		return ms
	}
	defer domain.ResetDefaultIssueService()

	mr := new(dTesting.IssueRepositoryMock)

	uc := usecases.NewIssueUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.Add(v.Title, v.Description, v.Status, p, l)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, v, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseIssueAddErr(t *testing.T) {
	p := domain.Project{
		ID:          1,
		Name:        "test-name",
		Description: "test-description",
	}
	l := map[string]domain.Label{}
	v := new(domain.Issue)
	v.Title = "test-title"
	v.Description = "test-description"
	v.Status = 1
	v.ProjectID = p.ID
	v.Project = p
	v.Labels = nil

	ms := new(dTesting.IssueServiceMock)
	ms.On("Add", v).Return(new(domain.Issue), errors.New("test error"))
	domain.GetDefaultIssueService = func(r domain.IssueRepository) domain.IssueService {
		return ms
	}
	defer domain.ResetDefaultIssueService()

	mr := new(dTesting.IssueRepositoryMock)

	uc := usecases.NewIssueUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.Add(v.Title, v.Description, v.Status, p, l)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseIssueUpdate(t *testing.T) {
	p := domain.Project{
		ID:          1,
		Name:        "test-name",
		Description: "test-description",
	}
	l := map[string]domain.Label{
		"test-name": domain.Label{
			ID:   1,
			Name: "test-name",
		},
	}
	vf := domain.Issue{}
	vf.Title = "test-title"
	vf.Description = "test-description"
	vf.Status = 1
	vf.ProjectID = p.ID
	vf.Project = p
	vf.Labels = []domain.Label{}
	vu := vf
	vu.Labels = []domain.Label{
		domain.Label{
			ID:   1,
			Name: "test-name",
		},
	}

	ms := new(dTesting.IssueServiceMock)
	ms.On("FindByID", vf.ID).Return(vf, nil)
	ms.On("Update", vu).Return(vu, nil)
	domain.GetDefaultIssueService = func(r domain.IssueRepository) domain.IssueService {
		return ms
	}
	defer domain.ResetDefaultIssueService()

	mr := new(dTesting.IssueRepositoryMock)

	uc := usecases.NewIssueUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.Update(vf.ID, vf.Title, vf.Description, vf.Status, l)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.NotEqual(t, vf, item)
	assert.Equal(t, vu, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseIssueUpdateErr(t *testing.T) {
	p := domain.Project{
		ID:          1,
		Name:        "test-name",
		Description: "test-description",
	}
	l := map[string]domain.Label{
		"test-name": domain.Label{
			ID:   1,
			Name: "test-name",
		},
	}
	vf := domain.Issue{}
	vf.Title = "test-title"
	vf.Description = "test-description"
	vf.Status = 1
	vf.ProjectID = p.ID
	vf.Project = p
	vf.Labels = []domain.Label{}
	vu := vf
	vu.Labels = []domain.Label{
		domain.Label{
			ID:   1,
			Name: "test-name",
		},
	}

	ms := new(dTesting.IssueServiceMock)
	ms.On("FindByID", vf.ID).Return(vf, nil)
	ms.On("Update", vu).Return(vu, errors.New("test error"))
	domain.GetDefaultIssueService = func(r domain.IssueRepository) domain.IssueService {
		return ms
	}
	defer domain.ResetDefaultIssueService()

	mr := new(dTesting.IssueRepositoryMock)

	uc := usecases.NewIssueUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.Update(vf.ID, vf.Title, vf.Description, vf.Status, l)

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.NotEqual(t, vf, item)
	assert.Equal(t, vu, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseIssueUpdateFindByIDErr(t *testing.T) {
	l := map[string]domain.Label{
		"test-name": domain.Label{
			ID:   1,
			Name: "test-name",
		},
	}

	ms := new(dTesting.IssueServiceMock)
	ms.On("FindByID", uint(1)).Return(domain.Issue{}, errors.New("test error"))
	domain.GetDefaultIssueService = func(r domain.IssueRepository) domain.IssueService {
		return ms
	}
	defer domain.ResetDefaultIssueService()

	mr := new(dTesting.IssueRepositoryMock)

	uc := usecases.NewIssueUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.Update(1, "test-title", "test-description", 1, l)

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, domain.Issue{}, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseIssueFindByID(t *testing.T) {
	v := domain.Issue{}
	v.Title = "test-title"
	v.Description = "test-description"
	v.Status = 1
	v.ProjectID = 1
	v.Labels = []domain.Label{}

	ms := new(dTesting.IssueServiceMock)
	ms.On("FindByID", v.ID).Return(v, nil)
	domain.GetDefaultIssueService = func(r domain.IssueRepository) domain.IssueService {
		return ms
	}
	defer domain.ResetDefaultIssueService()

	mr := new(dTesting.IssueRepositoryMock)

	uc := usecases.NewIssueUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.FindByID(v.ID)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, v, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseIssueFindByIDErr(t *testing.T) {
	v := domain.Issue{}
	v.Title = "test-title"
	v.Description = "test-description"
	v.Status = 1
	v.ProjectID = 1
	v.Labels = []domain.Label{}

	ms := new(dTesting.IssueServiceMock)
	ms.On("FindByID", v.ID).Return(v, errors.New("test error"))
	domain.GetDefaultIssueService = func(r domain.IssueRepository) domain.IssueService {
		return ms
	}
	defer domain.ResetDefaultIssueService()

	mr := new(dTesting.IssueRepositoryMock)

	uc := usecases.NewIssueUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.FindByID(v.ID)

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, v, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseIssueFindAll(t *testing.T) {
	issues := []domain.Issue{
		domain.Issue{
			Title:       "test-title-1",
			Description: "test-description",
			Status:      1,
			ProjectID:   1,
		},
		domain.Issue{
			Title:       "test-title-2",
			Description: "test-description",
			Status:      1,
			ProjectID:   1,
		},
	}

	ms := new(dTesting.IssueServiceMock)
	ms.On("FindAll").Return(issues, nil)
	domain.GetDefaultIssueService = func(r domain.IssueRepository) domain.IssueService {
		return ms
	}
	defer domain.ResetDefaultIssueService()

	mr := new(dTesting.IssueRepositoryMock)

	uc := usecases.NewIssueUseCase(mr)

	assert.NotNil(t, uc)

	items, err := uc.FindAll()

	assert.Nil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, issues, items)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseIssueFindAllErr(t *testing.T) {
	issues := []domain.Issue{}

	ms := new(dTesting.IssueServiceMock)
	ms.On("FindAll").Return(issues, errors.New("test error"))
	domain.GetDefaultIssueService = func(r domain.IssueRepository) domain.IssueService {
		return ms
	}
	defer domain.ResetDefaultIssueService()

	mr := new(dTesting.IssueRepositoryMock)

	uc := usecases.NewIssueUseCase(mr)

	assert.NotNil(t, uc)

	items, err := uc.FindAll()

	assert.NotNil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, issues, items)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseIssueRemove(t *testing.T) {
	ms := new(dTesting.IssueServiceMock)
	ms.On("Remove", uint(1)).Return(true, nil)
	domain.GetDefaultIssueService = func(r domain.IssueRepository) domain.IssueService {
		return ms
	}
	defer domain.ResetDefaultIssueService()

	mr := new(dTesting.IssueRepositoryMock)

	uc := usecases.NewIssueUseCase(mr)

	assert.NotNil(t, uc)

	status, err := uc.Remove(uint(1))

	assert.Nil(t, err)
	assert.True(t, status)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseIssueRemoveErr(t *testing.T) {
	ms := new(dTesting.IssueServiceMock)
	ms.On("Remove", uint(1)).Return(false, errors.New("test error"))
	domain.GetDefaultIssueService = func(r domain.IssueRepository) domain.IssueService {
		return ms
	}
	defer domain.ResetDefaultIssueService()

	mr := new(dTesting.IssueRepositoryMock)

	uc := usecases.NewIssueUseCase(mr)

	assert.NotNil(t, uc)

	status, err := uc.Remove(uint(1))

	assert.NotNil(t, err)
	assert.False(t, status)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}
