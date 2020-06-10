package domain_test

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-issue-tracker/pkg/domain"
	dTesting "go-issue-tracker/pkg/domain/testing"
	"testing"
)

func TestDomainLabelResetDefaultLabelService(t *testing.T) {
	assert.NotNil(t, domain.GetDefaultLabelService)

	domain.GetDefaultLabelService = nil
	defer domain.ResetDefaultLabelService()

	assert.Nil(t, domain.GetDefaultLabelService)

	domain.ResetDefaultLabelService()

	assert.NotNil(t, domain.GetDefaultLabelService)
}

func TestDomainLabelGetDefaultLabelService(t *testing.T) {
	m := new(dTesting.LabelRepositoryMock)

	s := domain.GetDefaultLabelService(m)

	assert.NotNil(t, s)
}

func TestDomainLabelAdd(t *testing.T) {
	v := new(domain.Label)
	v.Name = "test-name"
	v.ColorHexCode = "FFFFFF"

	m := new(dTesting.LabelRepositoryMock)
	m.On("Add", v).Return(v, nil)
	m.On("FindByName", v.Name).Return(*v, nil)

	s := domain.GetDefaultLabelService(m)

	item, err := s.Add(v)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, v, item)

	m.AssertExpectations(t)
}

func TestDomainLabelAddAlreadyExists(t *testing.T) {
	var tests = []struct {
		v   *domain.Label
		err error
	}{
		{
			&domain.Label{
				ID:           1,
				Name:         "test-name",
				ColorHexCode: "FFFFFF",
			},
			fmt.Errorf("%s label already exists", "test"),
		},
		{
			&domain.Label{
				Name:         "test-name",
				ColorHexCode: "FFFFFF",
			},
			fmt.Errorf("different error"),
		},
	}

	for _, ts := range tests {
		m := new(dTesting.LabelRepositoryMock)
		m.On("FindByName", ts.v.Name).Return(*ts.v, ts.err)

		s := domain.GetDefaultLabelService(m)

		item, err := s.Add(ts.v)

		assert.NotNil(t, err)
		assert.Nil(t, item)

		m.AssertExpectations(t)
	}
}

func TestDomainLabelAddErr(t *testing.T) {
	v := new(domain.Label)

	m := new(dTesting.LabelRepositoryMock)
	m.On("Add", v).Return(v, errors.New("test error"))
	m.On("FindByName", v.Name).Return(*v, nil)

	s := domain.GetDefaultLabelService(m)

	item, err := s.Add(v)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	m.AssertExpectations(t)
}

func TestDomainLabelUpdate(t *testing.T) {
	v := domain.Label{
		ID:           1,
		Name:         "test-name",
		ColorHexCode: "FFFFFF",
	}

	m := new(dTesting.LabelRepositoryMock)
	m.On("Update", v).Return(v, nil)

	s := domain.GetDefaultLabelService(m)

	item, err := s.Update(v)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, v, item)

	m.AssertExpectations(t)
}

func TestDomainLabelUpdateErr(t *testing.T) {
	v := domain.Label{}

	m := new(dTesting.LabelRepositoryMock)
	m.On("Update", v).Return(v, errors.New("test error"))

	s := domain.GetDefaultLabelService(m)

	item, err := s.Update(v)

	assert.NotNil(t, err)
	assert.Equal(t, v, item)

	m.AssertExpectations(t)
}

func TestDomainLabelFindByID(t *testing.T) {
	v := domain.Label{
		ID:           1,
		Name:         "test-name",
		ColorHexCode: "FFFFFF",
	}

	m := new(dTesting.LabelRepositoryMock)
	m.On("FindByID", v.ID).Return(v, nil)

	s := domain.GetDefaultLabelService(m)

	item, err := s.FindByID(v.ID)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, v, item)

	m.AssertExpectations(t)
}

func TestDomainLabelFindByIDErr(t *testing.T) {
	v := domain.Label{}

	m := new(dTesting.LabelRepositoryMock)
	m.On("FindByID", uint(1)).Return(v, errors.New("test error"))

	s := domain.GetDefaultLabelService(m)

	item, err := s.FindByID(uint(1))

	assert.NotNil(t, err)
	assert.Equal(t, v, item)

	m.AssertExpectations(t)
}

func TestDomainLabelFindByName(t *testing.T) {
	v := domain.Label{
		ID:           1,
		Name:         "test-name",
		ColorHexCode: "FFFFFF",
	}

	m := new(dTesting.LabelRepositoryMock)
	m.On("FindByName", v.Name).Return(v, nil)

	s := domain.GetDefaultLabelService(m)

	item, err := s.FindByName(v.Name)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, v, item)

	m.AssertExpectations(t)
}

func TestDomainLabelFindByNameErr(t *testing.T) {
	v := domain.Label{}

	m := new(dTesting.LabelRepositoryMock)
	m.On("FindByName", "test-name").Return(v, errors.New("test error"))

	s := domain.GetDefaultLabelService(m)

	item, err := s.FindByName("test-name")

	assert.NotNil(t, err)
	assert.Equal(t, v, item)

	m.AssertExpectations(t)
}

func TestDomainLabelFindAll(t *testing.T) {
	v := []domain.Label{}

	m := new(dTesting.LabelRepositoryMock)
	m.On("FindAll").Return(v, nil)

	s := domain.GetDefaultLabelService(m)

	items, err := s.FindAll()

	assert.Nil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, v, items)

	m.AssertExpectations(t)
}

func TestDomainLabelFindAllErr(t *testing.T) {
	v := []domain.Label{}

	m := new(dTesting.LabelRepositoryMock)
	m.On("FindAll").Return(v, errors.New("test error"))

	s := domain.GetDefaultLabelService(m)

	items, err := s.FindAll()

	assert.NotNil(t, err)
	assert.Equal(t, v, items)

	m.AssertExpectations(t)
}

func TestDomainLabelRemove(t *testing.T) {
	m := new(dTesting.LabelRepositoryMock)
	m.On("Remove", uint(1)).Return(true, nil)

	s := domain.GetDefaultLabelService(m)

	status, err := s.Remove(uint(1))

	assert.Nil(t, err)
	assert.True(t, status)

	m.AssertExpectations(t)
}

func TestDomainLabelRemoveErr(t *testing.T) {
	m := new(dTesting.LabelRepositoryMock)
	m.On("Remove", uint(1)).Return(false, errors.New("test error"))

	s := domain.GetDefaultLabelService(m)

	status, err := s.Remove(uint(1))

	assert.NotNil(t, err)
	assert.False(t, status)

	m.AssertExpectations(t)
}
