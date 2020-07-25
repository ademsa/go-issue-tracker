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
	l := new(domain.Label)
	l.Name = "test-name"
	l.ColorHexCode = "FFFFFF"

	m := new(dTesting.LabelRepositoryMock)
	m.On("Add", l).Return(l, nil)
	m.On("FindByName", l.Name).Return(*l, nil)

	s := domain.GetDefaultLabelService(m)

	item, err := s.Add(l)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, l, item)

	m.AssertExpectations(t)
}

func TestDomainLabelAddAlreadyExists(t *testing.T) {
	tests := []struct {
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
	l := new(domain.Label)

	m := new(dTesting.LabelRepositoryMock)
	m.On("Add", l).Return(l, errors.New("test error"))
	m.On("FindByName", l.Name).Return(*l, nil)

	s := domain.GetDefaultLabelService(m)

	item, err := s.Add(l)

	assert.NotNil(t, err)
	assert.Nil(t, item)

	m.AssertExpectations(t)
}

func TestDomainLabelUpdate(t *testing.T) {
	l := domain.Label{
		ID:           1,
		Name:         "test-name",
		ColorHexCode: "FFFFFF",
	}

	m := new(dTesting.LabelRepositoryMock)
	m.On("Update", l).Return(l, nil)

	s := domain.GetDefaultLabelService(m)

	item, err := s.Update(l)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, l, item)

	m.AssertExpectations(t)
}

func TestDomainLabelUpdateErr(t *testing.T) {
	l := domain.Label{}

	m := new(dTesting.LabelRepositoryMock)
	m.On("Update", l).Return(l, errors.New("test error"))

	s := domain.GetDefaultLabelService(m)

	item, err := s.Update(l)

	assert.NotNil(t, err)
	assert.Equal(t, l, item)

	m.AssertExpectations(t)
}

func TestDomainLabelFindByID(t *testing.T) {
	l := domain.Label{
		ID:           1,
		Name:         "test-name",
		ColorHexCode: "FFFFFF",
	}

	m := new(dTesting.LabelRepositoryMock)
	m.On("FindByID", l.ID).Return(l, nil)

	s := domain.GetDefaultLabelService(m)

	item, err := s.FindByID(l.ID)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, l, item)

	m.AssertExpectations(t)
}

func TestDomainLabelFindByIDErr(t *testing.T) {
	l := domain.Label{}

	m := new(dTesting.LabelRepositoryMock)
	m.On("FindByID", uint(1)).Return(l, errors.New("test error"))

	s := domain.GetDefaultLabelService(m)

	item, err := s.FindByID(uint(1))

	assert.NotNil(t, err)
	assert.Equal(t, l, item)

	m.AssertExpectations(t)
}

func TestDomainLabelFindByName(t *testing.T) {
	l := domain.Label{
		ID:           1,
		Name:         "test-name",
		ColorHexCode: "FFFFFF",
	}

	m := new(dTesting.LabelRepositoryMock)
	m.On("FindByName", l.Name).Return(l, nil)

	s := domain.GetDefaultLabelService(m)

	item, err := s.FindByName(l.Name)

	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, l, item)

	m.AssertExpectations(t)
}

func TestDomainLabelFindByNameErr(t *testing.T) {
	l := domain.Label{}

	m := new(dTesting.LabelRepositoryMock)
	m.On("FindByName", "test-name").Return(l, errors.New("test error"))

	s := domain.GetDefaultLabelService(m)

	item, err := s.FindByName("test-name")

	assert.NotNil(t, err)
	assert.Equal(t, l, item)

	m.AssertExpectations(t)
}

func TestDomainLabelFind(t *testing.T) {
	v := []domain.Label{}

	m := new(dTesting.LabelRepositoryMock)
	m.On("Find", "test").Return(v, nil)

	s := domain.GetDefaultLabelService(m)

	items, err := s.Find("test")

	assert.Nil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, v, items)

	m.AssertExpectations(t)
}

func TestDomainLabelFindErr(t *testing.T) {
	v := []domain.Label{}

	m := new(dTesting.LabelRepositoryMock)
	m.On("Find", "test").Return(v, errors.New("test error"))

	s := domain.GetDefaultLabelService(m)

	items, err := s.Find("test")

	assert.NotNil(t, err)
	assert.Equal(t, v, items)

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
