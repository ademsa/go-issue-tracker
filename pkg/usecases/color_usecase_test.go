package usecases_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go-issue-tracker/pkg/domain"
	dTesting "go-issue-tracker/pkg/domain/testing"
	"go-issue-tracker/pkg/usecases"
	"testing"
)

func TestUseCaseColorNewColorUseCase(t *testing.T) {
	ms := new(dTesting.ColorServiceMock)
	domain.GetDefaultColorService = func(r domain.ColorRepository) domain.ColorService {
		return ms
	}
	defer domain.ResetDefaultColorService()

	mr := new(dTesting.ColorRepositoryMock)

	uc := usecases.NewColorUseCase(mr)

	assert.NotNil(t, uc)
}

func TestUseCaseColorGetColor(t *testing.T) {
	v := domain.Color{
		HexCode: "FF0000",
	}

	ms := new(dTesting.ColorServiceMock)
	ms.On("GetColor").Return(v, nil)
	domain.GetDefaultColorService = func(r domain.ColorRepository) domain.ColorService {
		return ms
	}
	defer domain.ResetDefaultColorService()

	mr := new(dTesting.ColorRepositoryMock)

	uc := usecases.NewColorUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.GetColor()

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, v, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestUseCaseColorGetColorErr(t *testing.T) {
	ms := new(dTesting.ColorServiceMock)
	ms.On("GetColor").Return(domain.Color{}, errors.New("test error"))
	domain.GetDefaultColorService = func(r domain.ColorRepository) domain.ColorService {
		return ms
	}
	defer domain.ResetDefaultColorService()

	mr := new(dTesting.ColorRepositoryMock)

	uc := usecases.NewColorUseCase(mr)

	assert.NotNil(t, uc)

	item, err := uc.GetColor()

	assert.NotNil(t, err)
	assert.Equal(t, domain.Color{}, item)

	ms.AssertExpectations(t)
	mr.AssertExpectations(t)
}
