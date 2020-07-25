package domain_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go-issue-tracker/pkg/domain"
	dTesting "go-issue-tracker/pkg/domain/testing"
	"testing"
)

func TestDomainColorResetDefaultColorService(t *testing.T) {
	assert.NotNil(t, domain.GetDefaultColorService)

	domain.GetDefaultColorService = nil
	defer domain.ResetDefaultColorService()

	assert.Nil(t, domain.GetDefaultColorService)

	domain.ResetDefaultColorService()

	assert.NotNil(t, domain.GetDefaultColorService)
}

func TestDomainColorGetDefaultColorService(t *testing.T) {
	m := new(dTesting.ColorRepositoryMock)

	s := domain.GetDefaultColorService(m)

	assert.NotNil(t, s)
}

func TestDomainColorGetColor(t *testing.T) {
	c := domain.Color{
		HexCode: "FFFFFF",
	}

	m := new(dTesting.ColorRepositoryMock)
	m.On("GetColor").Return(c, nil)

	s := domain.GetDefaultColorService(m)

	item, err := s.GetColor()

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, c, item)

	m.AssertExpectations(t)
}

func TestDomainColorGetColorErr(t *testing.T) {
	c := domain.Color{}

	m := new(dTesting.ColorRepositoryMock)
	m.On("GetColor").Return(c, errors.New("test error"))

	s := domain.GetDefaultColorService(m)

	item, err := s.GetColor()

	assert.NotNil(t, err)
	assert.Equal(t, c, item)

	m.AssertExpectations(t)
}
