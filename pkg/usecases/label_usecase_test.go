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
	v := new(domain.Label)
	v.Name = "test-name"
	v.ColorHexCode = "FFFFFF"

	ms := new(dTesting.LabelServiceMock)
	ms.On("Add", v).Return(v, nil)
	domain.GetDefaultLabelService = func(r domain.LabelRepository) domain.LabelService {
		return ms
	}
	defer domain.ResetDefaultLabelService()

	mr := new(dTesting.LabelRepositoryMock)

	uc := usecases.NewLabelUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.Add(v.Name, v.ColorHexCode)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, v, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseLabelAddErr(t *testing.T) {
	v := new(domain.Label)
	v.Name = "test-name"
	v.ColorHexCode = "FFFFFF"

	ms := new(dTesting.LabelServiceMock)
	ms.On("Add", v).Return(new(domain.Label), errors.New("test error"))
	domain.GetDefaultLabelService = func(r domain.LabelRepository) domain.LabelService {
		return ms
	}
	defer domain.ResetDefaultLabelService()

	mr := new(dTesting.LabelRepositoryMock)

	uc := usecases.NewLabelUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.Add(v.Name, v.ColorHexCode)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseLabelUpdate(t *testing.T) {
	vf := domain.Label{}
	vf.Name = "test-name"
	vf.ColorHexCode = "FFFFFF"
	vu := vf

	ms := new(dTesting.LabelServiceMock)
	ms.On("FindByID", vf.ID).Return(vf, nil)
	ms.On("Update", vu).Return(vu, nil)
	domain.GetDefaultLabelService = func(r domain.LabelRepository) domain.LabelService {
		return ms
	}
	defer domain.ResetDefaultLabelService()

	mr := new(dTesting.LabelRepositoryMock)

	uc := usecases.NewLabelUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.Update(vf.ID, vf.Name, vf.ColorHexCode)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, vu, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseLabelUpdateErr(t *testing.T) {
	vf := domain.Label{}
	vf.Name = "test-name"
	vf.ColorHexCode = "FFFFFF"
	vu := vf

	ms := new(dTesting.LabelServiceMock)
	ms.On("FindByID", vf.ID).Return(vf, nil)
	ms.On("Update", vu).Return(vu, errors.New("test error"))
	domain.GetDefaultLabelService = func(r domain.LabelRepository) domain.LabelService {
		return ms
	}
	defer domain.ResetDefaultLabelService()

	mr := new(dTesting.LabelRepositoryMock)

	uc := usecases.NewLabelUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.Update(vf.ID, vf.Name, vf.ColorHexCode)

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, vu, item)

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
	v := domain.Label{}
	v.Name = "test-name"
	v.ColorHexCode = "FFFFFF"

	ms := new(dTesting.LabelServiceMock)
	ms.On("FindByID", v.ID).Return(v, nil)
	domain.GetDefaultLabelService = func(r domain.LabelRepository) domain.LabelService {
		return ms
	}
	defer domain.ResetDefaultLabelService()

	mr := new(dTesting.LabelRepositoryMock)

	uc := usecases.NewLabelUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.FindByID(v.ID)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, v, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseLabelFindByIDErr(t *testing.T) {
	v := domain.Label{}
	v.Name = "test-name"
	v.ColorHexCode = "FFFFFF"

	ms := new(dTesting.LabelServiceMock)
	ms.On("FindByID", v.ID).Return(v, errors.New("test error"))
	domain.GetDefaultLabelService = func(r domain.LabelRepository) domain.LabelService {
		return ms
	}
	defer domain.ResetDefaultLabelService()

	mr := new(dTesting.LabelRepositoryMock)

	uc := usecases.NewLabelUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.FindByID(v.ID)

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, v, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseLabelFindByName(t *testing.T) {
	v := domain.Label{}
	v.Name = "test-name"
	v.ColorHexCode = "FFFFFF"

	ms := new(dTesting.LabelServiceMock)
	ms.On("FindByName", v.Name).Return(v, nil)
	domain.GetDefaultLabelService = func(r domain.LabelRepository) domain.LabelService {
		return ms
	}
	defer domain.ResetDefaultLabelService()

	mr := new(dTesting.LabelRepositoryMock)

	uc := usecases.NewLabelUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.FindByName(v.Name)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, v, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseLabelFindByNameErr(t *testing.T) {
	v := domain.Label{}
	v.Name = "test-name"
	v.ColorHexCode = "FFFFFF"

	ms := new(dTesting.LabelServiceMock)
	ms.On("FindByName", v.Name).Return(v, errors.New("test error"))
	domain.GetDefaultLabelService = func(r domain.LabelRepository) domain.LabelService {
		return ms
	}
	defer domain.ResetDefaultLabelService()

	mr := new(dTesting.LabelRepositoryMock)

	uc := usecases.NewLabelUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.FindByName(v.Name)

	assert.NotNil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, v, item)

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
