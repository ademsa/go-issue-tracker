package usecases_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go-issue-tracker/pkg/domain"
	dTesting "go-issue-tracker/pkg/domain/testing"
	"go-issue-tracker/pkg/usecases"
	"testing"
)

func TestUseCaseProjectNewProjectUseCase(t *testing.T) {
	ms := new(dTesting.ProjectServiceMock)
	domain.GetDefaultProjectService = func(r domain.ProjectRepository) domain.ProjectService {
		return ms
	}
	defer domain.ResetDefaultProjectService()

	mr := new(dTesting.ProjectRepositoryMock)

	uc := usecases.NewProjectUseCase(mr)

	assert.NotNil(t, uc)
}

func TestUseCaseProjectAdd(t *testing.T) {
	p := new(domain.Project)
	p.Name = "test-name"
	p.Description = "test-description"

	ms := new(dTesting.ProjectServiceMock)
	ms.On("Add", p).Return(p, nil)
	domain.GetDefaultProjectService = func(r domain.ProjectRepository) domain.ProjectService {
		return ms
	}
	defer domain.ResetDefaultProjectService()

	mr := new(dTesting.ProjectRepositoryMock)

	uc := usecases.NewProjectUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.Add(p.Name, p.Description)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, p, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseProjectAddErr(t *testing.T) {
	p := new(domain.Project)
	p.Name = "test-name"
	p.Description = "test-description"

	ms := new(dTesting.ProjectServiceMock)
	ms.On("Add", p).Return(new(domain.Project), errors.New("test error"))
	domain.GetDefaultProjectService = func(r domain.ProjectRepository) domain.ProjectService {
		return ms
	}
	defer domain.ResetDefaultProjectService()

	mr := new(dTesting.ProjectRepositoryMock)

	uc := usecases.NewProjectUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.Add(p.Name, p.Description)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseProjectUpdate(t *testing.T) {
	pf := domain.Project{}
	pf.Name = "test-name"
	pf.Description = "test-description"
	pu := pf

	ms := new(dTesting.ProjectServiceMock)
	ms.On("FindByID", pf.ID).Return(pf, nil)
	ms.On("Update", pu).Return(pu, nil)
	domain.GetDefaultProjectService = func(r domain.ProjectRepository) domain.ProjectService {
		return ms
	}
	defer domain.ResetDefaultProjectService()

	mr := new(dTesting.ProjectRepositoryMock)

	uc := usecases.NewProjectUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.Update(pf.ID, pf.Name, pf.Description)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, pu, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseProjectUpdateErr(t *testing.T) {
	pf := domain.Project{}
	pf.Name = "test-name"
	pf.Description = "test-description"
	pu := pf

	ms := new(dTesting.ProjectServiceMock)
	ms.On("FindByID", pf.ID).Return(pf, nil)
	ms.On("Update", pu).Return(pu, errors.New("test error"))
	domain.GetDefaultProjectService = func(r domain.ProjectRepository) domain.ProjectService {
		return ms
	}
	defer domain.ResetDefaultProjectService()

	mr := new(dTesting.ProjectRepositoryMock)

	uc := usecases.NewProjectUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.Update(pf.ID, pf.Name, pf.Description)

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, pu, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseProjectUpdateFindByIDErr(t *testing.T) {
	ms := new(dTesting.ProjectServiceMock)
	ms.On("FindByID", uint(1)).Return(domain.Project{}, errors.New("test error"))
	domain.GetDefaultProjectService = func(r domain.ProjectRepository) domain.ProjectService {
		return ms
	}
	defer domain.ResetDefaultProjectService()

	mr := new(dTesting.ProjectRepositoryMock)

	uc := usecases.NewProjectUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.Update(1, "test-name", "test-description")

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, domain.Project{}, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseProjectFindByID(t *testing.T) {
	p := domain.Project{}
	p.Name = "test-name"
	p.Description = "test-description"

	ms := new(dTesting.ProjectServiceMock)
	ms.On("FindByID", p.ID).Return(p, nil)
	domain.GetDefaultProjectService = func(r domain.ProjectRepository) domain.ProjectService {
		return ms
	}
	defer domain.ResetDefaultProjectService()

	mr := new(dTesting.ProjectRepositoryMock)

	uc := usecases.NewProjectUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.FindByID(p.ID)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, p, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseProjectFindByIDErr(t *testing.T) {
	p := domain.Project{}
	p.Name = "test-name"
	p.Description = "test-description"

	ms := new(dTesting.ProjectServiceMock)
	ms.On("FindByID", p.ID).Return(p, errors.New("test error"))
	domain.GetDefaultProjectService = func(r domain.ProjectRepository) domain.ProjectService {
		return ms
	}
	defer domain.ResetDefaultProjectService()

	mr := new(dTesting.ProjectRepositoryMock)

	uc := usecases.NewProjectUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.FindByID(p.ID)

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, p, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseProjectFind(t *testing.T) {
	projects := []domain.Project{
		domain.Project{
			Name:        "test-name-1",
			Description: "test-description",
		},
		domain.Project{
			Name:        "test-name-2",
			Description: "test-description",
		},
	}

	ms := new(dTesting.ProjectServiceMock)
	ms.On("Find", "test").Return(projects, nil)
	domain.GetDefaultProjectService = func(r domain.ProjectRepository) domain.ProjectService {
		return ms
	}
	defer domain.ResetDefaultProjectService()

	mr := new(dTesting.ProjectRepositoryMock)

	uc := usecases.NewProjectUseCase(mr)

	assert.NotNil(t, uc)

	items, err := uc.Find("test")

	assert.Nil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, projects, items)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseProjectFindErr(t *testing.T) {
	projects := []domain.Project{}

	ms := new(dTesting.ProjectServiceMock)
	ms.On("Find", "test").Return(projects, errors.New("test error"))
	domain.GetDefaultProjectService = func(r domain.ProjectRepository) domain.ProjectService {
		return ms
	}
	defer domain.ResetDefaultProjectService()

	mr := new(dTesting.ProjectRepositoryMock)

	uc := usecases.NewProjectUseCase(mr)

	assert.NotNil(t, uc)

	items, err := uc.Find("test")

	assert.NotNil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, projects, items)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseProjectFindAll(t *testing.T) {
	projects := []domain.Project{
		domain.Project{
			Name:        "test-name-1",
			Description: "test-description",
		},
		domain.Project{
			Name:        "test-name-2",
			Description: "test-description",
		},
	}

	ms := new(dTesting.ProjectServiceMock)
	ms.On("FindAll").Return(projects, nil)
	domain.GetDefaultProjectService = func(r domain.ProjectRepository) domain.ProjectService {
		return ms
	}
	defer domain.ResetDefaultProjectService()

	mr := new(dTesting.ProjectRepositoryMock)

	uc := usecases.NewProjectUseCase(mr)

	assert.NotNil(t, uc)

	items, err := uc.FindAll()

	assert.Nil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, projects, items)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseProjectFindAllErr(t *testing.T) {
	projects := []domain.Project{}

	ms := new(dTesting.ProjectServiceMock)
	ms.On("FindAll").Return(projects, errors.New("test error"))
	domain.GetDefaultProjectService = func(r domain.ProjectRepository) domain.ProjectService {
		return ms
	}
	defer domain.ResetDefaultProjectService()

	mr := new(dTesting.ProjectRepositoryMock)

	uc := usecases.NewProjectUseCase(mr)

	assert.NotNil(t, uc)

	items, err := uc.FindAll()

	assert.NotNil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, projects, items)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseProjectRemove(t *testing.T) {
	ms := new(dTesting.ProjectServiceMock)
	ms.On("Remove", uint(1)).Return(true, nil)
	domain.GetDefaultProjectService = func(r domain.ProjectRepository) domain.ProjectService {
		return ms
	}
	defer domain.ResetDefaultProjectService()

	mr := new(dTesting.ProjectRepositoryMock)

	uc := usecases.NewProjectUseCase(mr)

	assert.NotNil(t, uc)

	status, err := uc.Remove(uint(1))

	assert.Nil(t, err)
	assert.True(t, status)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseProjectRemoveErr(t *testing.T) {
	ms := new(dTesting.ProjectServiceMock)
	ms.On("Remove", uint(1)).Return(false, errors.New("test error"))
	domain.GetDefaultProjectService = func(r domain.ProjectRepository) domain.ProjectService {
		return ms
	}
	defer domain.ResetDefaultProjectService()

	mr := new(dTesting.ProjectRepositoryMock)

	uc := usecases.NewProjectUseCase(mr)

	assert.NotNil(t, uc)

	status, err := uc.Remove(uint(1))

	assert.NotNil(t, err)
	assert.False(t, status)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}
