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
	i := new(domain.Issue)
	i.Title = "test-title"
	i.Description = "test-description"
	i.Status = 1
	i.ProjectID = p.ID
	i.Project = p
	i.Labels = []domain.Label{
		domain.Label{
			ID:   1,
			Name: "test-name",
		},
	}

	ms := new(dTesting.IssueServiceMock)
	ms.On("Add", i).Return(i, nil)
	domain.GetDefaultIssueService = func(r domain.IssueRepository) domain.IssueService {
		return ms
	}
	defer domain.ResetDefaultIssueService()

	mr := new(dTesting.IssueRepositoryMock)

	uc := usecases.NewIssueUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.Add(i.Title, i.Description, i.Status, p, l)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, i, item)

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
	i := new(domain.Issue)
	i.Title = "test-title"
	i.Description = "test-description"
	i.Status = 1
	i.ProjectID = p.ID
	i.Project = p
	i.Labels = nil

	ms := new(dTesting.IssueServiceMock)
	ms.On("Add", i).Return(new(domain.Issue), errors.New("test error"))
	domain.GetDefaultIssueService = func(r domain.IssueRepository) domain.IssueService {
		return ms
	}
	defer domain.ResetDefaultIssueService()

	mr := new(dTesting.IssueRepositoryMock)

	uc := usecases.NewIssueUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.Add(i.Title, i.Description, i.Status, p, l)

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
	iff := domain.Issue{}
	iff.Title = "test-title"
	iff.Description = "test-description"
	iff.Status = 1
	iff.ProjectID = p.ID
	iff.Project = p
	iff.Labels = []domain.Label{}
	iu := iff
	iu.Labels = []domain.Label{
		domain.Label{
			ID:   1,
			Name: "test-name",
		},
	}

	ms := new(dTesting.IssueServiceMock)
	ms.On("FindByID", iff.ID).Return(iff, nil)
	ms.On("Update", iu).Return(iu, nil)
	domain.GetDefaultIssueService = func(r domain.IssueRepository) domain.IssueService {
		return ms
	}
	defer domain.ResetDefaultIssueService()

	mr := new(dTesting.IssueRepositoryMock)

	uc := usecases.NewIssueUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.Update(iff.ID, iff.Title, iff.Description, iff.Status, l)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.NotEqual(t, iff, item)
	assert.Equal(t, iu, item)

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
	iff := domain.Issue{}
	iff.Title = "test-title"
	iff.Description = "test-description"
	iff.Status = 1
	iff.ProjectID = p.ID
	iff.Project = p
	iff.Labels = []domain.Label{}
	iu := iff
	iu.Labels = []domain.Label{
		domain.Label{
			ID:   1,
			Name: "test-name",
		},
	}

	ms := new(dTesting.IssueServiceMock)
	ms.On("FindByID", iff.ID).Return(iff, nil)
	ms.On("Update", iu).Return(iu, errors.New("test error"))
	domain.GetDefaultIssueService = func(r domain.IssueRepository) domain.IssueService {
		return ms
	}
	defer domain.ResetDefaultIssueService()

	mr := new(dTesting.IssueRepositoryMock)

	uc := usecases.NewIssueUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.Update(iff.ID, iff.Title, iff.Description, iff.Status, l)

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.NotEqual(t, iff, item)
	assert.Equal(t, iu, item)

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
	i := domain.Issue{}
	i.Title = "test-title"
	i.Description = "test-description"
	i.Status = 1
	i.ProjectID = 1
	i.Labels = []domain.Label{}

	ms := new(dTesting.IssueServiceMock)
	ms.On("FindByID", i.ID).Return(i, nil)
	domain.GetDefaultIssueService = func(r domain.IssueRepository) domain.IssueService {
		return ms
	}
	defer domain.ResetDefaultIssueService()

	mr := new(dTesting.IssueRepositoryMock)

	uc := usecases.NewIssueUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.FindByID(i.ID)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, i, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseIssueFindByIDErr(t *testing.T) {
	i := domain.Issue{}
	i.Title = "test-title"
	i.Description = "test-description"
	i.Status = 1
	i.ProjectID = 1
	i.Labels = []domain.Label{}

	ms := new(dTesting.IssueServiceMock)
	ms.On("FindByID", i.ID).Return(i, errors.New("test error"))
	domain.GetDefaultIssueService = func(r domain.IssueRepository) domain.IssueService {
		return ms
	}
	defer domain.ResetDefaultIssueService()

	mr := new(dTesting.IssueRepositoryMock)

	uc := usecases.NewIssueUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.FindByID(i.ID)

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, i, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseIssueFind(t *testing.T) {
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
	ms.On("Find", "test", uint(1), []string{"test1", "test2"}).Return(issues, nil)
	domain.GetDefaultIssueService = func(r domain.IssueRepository) domain.IssueService {
		return ms
	}
	defer domain.ResetDefaultIssueService()

	mr := new(dTesting.IssueRepositoryMock)

	uc := usecases.NewIssueUseCase(mr)

	assert.NotNil(t, uc)

	items, err := uc.Find("test", uint(1), []string{"test1", "test2"})

	assert.Nil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, issues, items)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseIssueFindErr(t *testing.T) {
	issues := []domain.Issue{}

	ms := new(dTesting.IssueServiceMock)
	ms.On("Find", "test", uint(1), []string{"test1", "test2"}).Return(issues, errors.New("test error"))
	domain.GetDefaultIssueService = func(r domain.IssueRepository) domain.IssueService {
		return ms
	}
	defer domain.ResetDefaultIssueService()

	mr := new(dTesting.IssueRepositoryMock)

	uc := usecases.NewIssueUseCase(mr)

	assert.NotNil(t, uc)

	items, err := uc.Find("test", uint(1), []string{"test1", "test2"})

	assert.NotNil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, issues, items)

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
