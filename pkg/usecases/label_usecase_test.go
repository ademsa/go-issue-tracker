package usecases_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go-issue-tracker/pkg/domain"
	dTesting "go-issue-tracker/pkg/domain/testing"
	"go-issue-tracker/pkg/usecases"
	"testing"
)

func TestUseCaseLabelNewLabelUseCase(t *testing.T) {
	ms := new(dTesting.LabelServiceMock)
	domain.GetDefaultLabelService = func(r domain.LabelRepository) domain.LabelService {
		return ms
	}
	defer domain.ResetDefaultLabelService()

	mr := new(dTesting.LabelRepositoryMock)

	uc := usecases.NewLabelUseCase(mr)

	assert.NotNil(t, uc)
}

func TestUseCaseLabelAdd(t *testing.T) {
	l := new(domain.Label)
	l.Name = "test-name"
	l.ColorHexCode = "FFFFFF"

	ms := new(dTesting.LabelServiceMock)
	ms.On("Add", l).Return(l, nil)
	domain.GetDefaultLabelService = func(r domain.LabelRepository) domain.LabelService {
		return ms
	}
	defer domain.ResetDefaultLabelService()

	mr := new(dTesting.LabelRepositoryMock)

	uc := usecases.NewLabelUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.Add(l.Name, l.ColorHexCode)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, l, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseLabelAddErr(t *testing.T) {
	l := new(domain.Label)
	l.Name = "test-name"
	l.ColorHexCode = "FFFFFF"

	ms := new(dTesting.LabelServiceMock)
	ms.On("Add", l).Return(new(domain.Label), errors.New("test error"))
	domain.GetDefaultLabelService = func(r domain.LabelRepository) domain.LabelService {
		return ms
	}
	defer domain.ResetDefaultLabelService()

	mr := new(dTesting.LabelRepositoryMock)

	uc := usecases.NewLabelUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.Add(l.Name, l.ColorHexCode)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseLabelUpdate(t *testing.T) {
	lf := domain.Label{}
	lf.Name = "test-name"
	lf.ColorHexCode = "FFFFFF"
	lu := lf

	ms := new(dTesting.LabelServiceMock)
	ms.On("FindByID", lf.ID).Return(lf, nil)
	ms.On("Update", lu).Return(lu, nil)
	domain.GetDefaultLabelService = func(r domain.LabelRepository) domain.LabelService {
		return ms
	}
	defer domain.ResetDefaultLabelService()

	mr := new(dTesting.LabelRepositoryMock)

	uc := usecases.NewLabelUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.Update(lf.ID, lf.Name, lf.ColorHexCode)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, lu, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseLabelUpdateErr(t *testing.T) {
	lf := domain.Label{}
	lf.Name = "test-name"
	lf.ColorHexCode = "FFFFFF"
	lu := lf

	ms := new(dTesting.LabelServiceMock)
	ms.On("FindByID", lf.ID).Return(lf, nil)
	ms.On("Update", lu).Return(lu, errors.New("test error"))
	domain.GetDefaultLabelService = func(r domain.LabelRepository) domain.LabelService {
		return ms
	}
	defer domain.ResetDefaultLabelService()

	mr := new(dTesting.LabelRepositoryMock)

	uc := usecases.NewLabelUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.Update(lf.ID, lf.Name, lf.ColorHexCode)

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, lu, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseLabelUpdateFindByIDErr(t *testing.T) {
	ms := new(dTesting.LabelServiceMock)
	ms.On("FindByID", uint(1)).Return(domain.Label{}, errors.New("test error"))
	domain.GetDefaultLabelService = func(r domain.LabelRepository) domain.LabelService {
		return ms
	}
	defer domain.ResetDefaultLabelService()

	mr := new(dTesting.LabelRepositoryMock)

	uc := usecases.NewLabelUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.Update(1, "test-name", "FFFFFF")

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, domain.Label{}, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseLabelFindByID(t *testing.T) {
	l := domain.Label{}
	l.Name = "test-name"
	l.ColorHexCode = "FFFFFF"

	ms := new(dTesting.LabelServiceMock)
	ms.On("FindByID", l.ID).Return(l, nil)
	domain.GetDefaultLabelService = func(r domain.LabelRepository) domain.LabelService {
		return ms
	}
	defer domain.ResetDefaultLabelService()

	mr := new(dTesting.LabelRepositoryMock)

	uc := usecases.NewLabelUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.FindByID(l.ID)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, l, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseLabelFindByIDErr(t *testing.T) {
	l := domain.Label{}
	l.Name = "test-name"
	l.ColorHexCode = "FFFFFF"

	ms := new(dTesting.LabelServiceMock)
	ms.On("FindByID", l.ID).Return(l, errors.New("test error"))
	domain.GetDefaultLabelService = func(r domain.LabelRepository) domain.LabelService {
		return ms
	}
	defer domain.ResetDefaultLabelService()

	mr := new(dTesting.LabelRepositoryMock)

	uc := usecases.NewLabelUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.FindByID(l.ID)

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, l, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseLabelFindByName(t *testing.T) {
	l := domain.Label{}
	l.Name = "test-name"
	l.ColorHexCode = "FFFFFF"

	ms := new(dTesting.LabelServiceMock)
	ms.On("FindByName", l.Name).Return(l, nil)
	domain.GetDefaultLabelService = func(r domain.LabelRepository) domain.LabelService {
		return ms
	}
	defer domain.ResetDefaultLabelService()

	mr := new(dTesting.LabelRepositoryMock)

	uc := usecases.NewLabelUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.FindByName(l.Name)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, l, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseLabelFindByNameErr(t *testing.T) {
	l := domain.Label{}
	l.Name = "test-name"
	l.ColorHexCode = "FFFFFF"

	ms := new(dTesting.LabelServiceMock)
	ms.On("FindByName", l.Name).Return(l, errors.New("test error"))
	domain.GetDefaultLabelService = func(r domain.LabelRepository) domain.LabelService {
		return ms
	}
	defer domain.ResetDefaultLabelService()

	mr := new(dTesting.LabelRepositoryMock)

	uc := usecases.NewLabelUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.FindByName(l.Name)

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, l, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseLabelFind(t *testing.T) {
	labels := []domain.Label{
		domain.Label{
			Name:         "test-name-1",
			ColorHexCode: "FFFFFF",
		},
		domain.Label{
			Name:         "test-name-2",
			ColorHexCode: "FFFFFF",
		},
	}

	ms := new(dTesting.LabelServiceMock)
	ms.On("Find", "test").Return(labels, nil)
	domain.GetDefaultLabelService = func(r domain.LabelRepository) domain.LabelService {
		return ms
	}
	defer domain.ResetDefaultLabelService()

	mr := new(dTesting.LabelRepositoryMock)

	uc := usecases.NewLabelUseCase(mr)

	assert.NotNil(t, uc)

	items, err := uc.Find("test")

	assert.Nil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, labels, items)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseLabelFindErr(t *testing.T) {
	labels := []domain.Label{}

	ms := new(dTesting.LabelServiceMock)
	ms.On("Find", "test").Return(labels, errors.New("test error"))
	domain.GetDefaultLabelService = func(r domain.LabelRepository) domain.LabelService {
		return ms
	}
	defer domain.ResetDefaultLabelService()

	mr := new(dTesting.LabelRepositoryMock)

	uc := usecases.NewLabelUseCase(mr)

	assert.NotNil(t, uc)

	items, err := uc.Find("test")

	assert.NotNil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, labels, items)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseLabelFindAll(t *testing.T) {
	labels := []domain.Label{
		domain.Label{
			Name:         "test-name-1",
			ColorHexCode: "FFFFFF",
		},
		domain.Label{
			Name:         "test-name-2",
			ColorHexCode: "FFFFFF",
		},
	}

	ms := new(dTesting.LabelServiceMock)
	ms.On("FindAll").Return(labels, nil)
	domain.GetDefaultLabelService = func(r domain.LabelRepository) domain.LabelService {
		return ms
	}
	defer domain.ResetDefaultLabelService()

	mr := new(dTesting.LabelRepositoryMock)

	uc := usecases.NewLabelUseCase(mr)

	assert.NotNil(t, uc)

	items, err := uc.FindAll()

	assert.Nil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, labels, items)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseLabelFindAllErr(t *testing.T) {
	labels := []domain.Label{}

	ms := new(dTesting.LabelServiceMock)
	ms.On("FindAll").Return(labels, errors.New("test error"))
	domain.GetDefaultLabelService = func(r domain.LabelRepository) domain.LabelService {
		return ms
	}
	defer domain.ResetDefaultLabelService()

	mr := new(dTesting.LabelRepositoryMock)

	uc := usecases.NewLabelUseCase(mr)

	assert.NotNil(t, uc)

	items, err := uc.FindAll()

	assert.NotNil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, labels, items)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseLabelRemove(t *testing.T) {
	ms := new(dTesting.LabelServiceMock)
	ms.On("Remove", uint(1)).Return(true, nil)
	domain.GetDefaultLabelService = func(r domain.LabelRepository) domain.LabelService {
		return ms
	}
	defer domain.ResetDefaultLabelService()

	mr := new(dTesting.LabelRepositoryMock)

	uc := usecases.NewLabelUseCase(mr)

	assert.NotNil(t, uc)

	status, err := uc.Remove(uint(1))

	assert.Nil(t, err)
	assert.True(t, status)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseLabelRemoveErr(t *testing.T) {
	ms := new(dTesting.LabelServiceMock)
	ms.On("Remove", uint(1)).Return(false, errors.New("test error"))
	domain.GetDefaultLabelService = func(r domain.LabelRepository) domain.LabelService {
		return ms
	}
	defer domain.ResetDefaultLabelService()

	mr := new(dTesting.LabelRepositoryMock)

	uc := usecases.NewLabelUseCase(mr)

	assert.NotNil(t, uc)

	status, err := uc.Remove(uint(1))

	assert.NotNil(t, err)
	assert.False(t, status)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}
