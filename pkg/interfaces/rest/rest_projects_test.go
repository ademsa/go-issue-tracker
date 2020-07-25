package rest_test

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go-issue-tracker/pkg/domain"
	"strings"
	"testing"
)

func TestAddProject(t *testing.T) {
	p := &domain.Project{
		Name:        "test-name",
		Description: "test-description",
	}

	cucm, iucm, lucm, pucm, m := prepareMocksAndRUC()

	pucm.On("Add", p.Name, p.Description).Return(p, nil)

	body := strings.NewReader("name=test-name&description=test-description")
	c, rec := prepareHTTP(echo.POST, "/api/projects/new", body)

	err := m.AddProject(c)

	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestAddProjectValueErrs(t *testing.T) {
	cucm, iucm, lucm, pucm, m := prepareMocksAndRUC()

	tests := []struct {
		body *strings.Reader
		err  error
	}{
		{
			strings.NewReader("name=&description=test-description"),
			errors.New("name not provided"),
		},
	}

	for _, ts := range tests {
		c, _ := prepareHTTP(echo.POST, "/api/projects/new", ts.body)

		err := m.AddProject(c)

		assert.NotNil(t, err)
		assert.Equal(t, ts.err.Error(), err.Error())

		checkAssertions(t, cucm, iucm, lucm, pucm)
	}
}

func TestAddProjectErr(t *testing.T) {
	p := &domain.Project{
		Name:        "test-name",
		Description: "test-description",
	}

	cucm, iucm, lucm, pucm, m := prepareMocksAndRUC()

	pucm.On("Add", p.Name, p.Description).Return(p, errors.New("test error"))

	body := strings.NewReader("name=test-name&description=test-description")
	c, _ := prepareHTTP(echo.POST, "/api/projects/new", body)

	err := m.AddProject(c)

	assert.NotNil(t, err)
	assert.Equal(t, "test error", err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestUpdateProject(t *testing.T) {
	p := domain.Project{
		ID:          1,
		Name:        "test-name",
		Description: "test-description",
	}

	cucm, iucm, lucm, pucm, m := prepareMocksAndRUC()

	pucm.On("Update", p.ID, p.Name, p.Description).Return(p, nil)

	body := strings.NewReader("name=test-name&description=test-description")
	c, rec := prepareHTTP(echo.POST, "/api/projects/:id", body)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := m.UpdateProject(c)

	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestUpdateProjectIDErr(t *testing.T) {
	cucm, iucm, lucm, pucm, m := prepareMocksAndRUC()

	body := strings.NewReader("name=test-name&description=test-description")
	c, _ := prepareHTTP(echo.POST, "/api/projects/:id", body)
	c.SetParamNames("id")
	c.SetParamValues("test")

	err := m.UpdateProject(c)

	assert.NotNil(t, err)
	assert.Equal(t, fmt.Errorf("strconv.Atoi: parsing \"%s\": invalid syntax", "test").Error(), err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestUpdateProjectValueErrs(t *testing.T) {
	cucm, iucm, lucm, pucm, m := prepareMocksAndRUC()

	tests := []struct {
		body *strings.Reader
		err  error
	}{
		{
			strings.NewReader("name=&description=test-description"),
			errors.New("name not provided"),
		},
	}

	for _, ts := range tests {
		c, _ := prepareHTTP(echo.POST, "/api/projects/:id", ts.body)
		c.SetParamNames("id")
		c.SetParamValues("1")

		err := m.UpdateProject(c)

		assert.NotNil(t, err)
		assert.Equal(t, ts.err.Error(), err.Error())

		checkAssertions(t, cucm, iucm, lucm, pucm)
	}
}

func TestUpdateProjectErr(t *testing.T) {
	p := domain.Project{
		ID:          1,
		Name:        "test-name",
		Description: "test-description",
	}

	cucm, iucm, lucm, pucm, m := prepareMocksAndRUC()

	pucm.On("Update", p.ID, p.Name, p.Description).Return(p, errors.New("test error"))

	body := strings.NewReader("name=test-name&description=test-description")
	c, _ := prepareHTTP(echo.POST, "/api/projects/:id", body)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := m.UpdateProject(c)

	assert.NotNil(t, err)
	assert.Equal(t, "test error", err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestFindProjectByID(t *testing.T) {
	p := domain.Project{
		ID:          1,
		Name:        "test-name",
		Description: "test-description",
	}

	cucm, iucm, lucm, pucm, m := prepareMocksAndRUC()

	pucm.On("FindByID", p.ID).Return(p, nil)

	c, rec := prepareHTTP(echo.GET, "/api/projects/:id", nil)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := m.FindProjectByID(c)

	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestFindProjectByIDIDErr(t *testing.T) {
	cucm, iucm, lucm, pucm, m := prepareMocksAndRUC()

	c, _ := prepareHTTP(echo.GET, "/api/projects/:id", nil)
	c.SetParamNames("id")
	c.SetParamValues("test")

	err := m.FindProjectByID(c)

	assert.NotNil(t, err)
	assert.Equal(t, fmt.Errorf("strconv.Atoi: parsing \"%s\": invalid syntax", "test").Error(), err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestFindProjectByIDNotFoundNoErr(t *testing.T) {
	p := domain.Project{
		ID:          1,
		Name:        "test-name",
		Description: "test-description",
	}

	cucm, iucm, lucm, pucm, m := prepareMocksAndRUC()

	pucm.On("FindByID", p.ID).Return(p, errors.New("record not found"))

	c, _ := prepareHTTP(echo.GET, "/api/projects/:id", nil)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := m.FindProjectByID(c)

	assert.Nil(t, err)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestFindProjectByIDOtherErr(t *testing.T) {
	p := domain.Project{
		ID:          1,
		Name:        "test-name",
		Description: "test-description",
	}

	cucm, iucm, lucm, pucm, m := prepareMocksAndRUC()

	pucm.On("FindByID", p.ID).Return(p, errors.New("test error"))

	c, _ := prepareHTTP(echo.GET, "/api/projects/:id", nil)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := m.FindProjectByID(c)

	assert.NotNil(t, err)
	assert.Equal(t, "test error", err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestFindProjects(t *testing.T) {
	p := []domain.Project{}

	cucm, iucm, lucm, pucm, m := prepareMocksAndRUC()

	pucm.On("Find", "test").Return(p, nil)

	c, rec := prepareHTTP(echo.GET, "/api/projects/find?name=test", nil)

	err := m.FindProjects(c)

	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestFindProjectsErr(t *testing.T) {
	p := []domain.Project{}

	cucm, iucm, lucm, pucm, m := prepareMocksAndRUC()

	pucm.On("Find", "test").Return(p, errors.New("test error"))

	c, _ := prepareHTTP(echo.GET, "/api/projects/find?name=test", nil)

	err := m.FindProjects(c)

	assert.NotNil(t, err)
	assert.Equal(t, "test error", err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestFindAllProjects(t *testing.T) {
	p := []domain.Project{}

	cucm, iucm, lucm, pucm, m := prepareMocksAndRUC()

	pucm.On("FindAll").Return(p, nil)

	c, rec := prepareHTTP(echo.GET, "/api/projects", nil)

	err := m.FindAllProjects(c)

	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestFindAllProjectsErr(t *testing.T) {
	p := []domain.Project{}

	cucm, iucm, lucm, pucm, m := prepareMocksAndRUC()

	pucm.On("FindAll").Return(p, errors.New("test error"))

	c, _ := prepareHTTP(echo.GET, "/api/projects", nil)

	err := m.FindAllProjects(c)

	assert.NotNil(t, err)
	assert.Equal(t, "test error", err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestRemoveProject(t *testing.T) {
	cucm, iucm, lucm, pucm, m := prepareMocksAndRUC()

	pucm.On("Remove", uint(1)).Return(true, nil)

	c, rec := prepareHTTP(echo.DELETE, "/api/projects/:id", nil)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := m.RemoveProject(c)

	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestRemoveProjectNotFoundNoErr(t *testing.T) {
	cucm, iucm, lucm, pucm, m := prepareMocksAndRUC()

	pucm.On("Remove", uint(1)).Return(false, errors.New("record not found"))

	c, _ := prepareHTTP(echo.DELETE, "/api/projects/:id", nil)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := m.RemoveProject(c)

	assert.Nil(t, err)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestRemoveProjectErr(t *testing.T) {
	cucm, iucm, lucm, pucm, m := prepareMocksAndRUC()

	pucm.On("Remove", uint(1)).Return(false, errors.New("test error"))

	c, _ := prepareHTTP(echo.DELETE, "/api/projects/:id", nil)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := m.RemoveProject(c)

	assert.NotNil(t, err)
	assert.Equal(t, "test error", err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestRemoveProjectIDErr(t *testing.T) {
	cucm, iucm, lucm, pucm, m := prepareMocksAndRUC()

	c, _ := prepareHTTP(echo.DELETE, "/api/projects/:id", nil)
	c.SetParamNames("id")
	c.SetParamValues("test")

	err := m.RemoveProject(c)

	assert.NotNil(t, err)
	assert.Equal(t, fmt.Errorf("strconv.Atoi: parsing \"%s\": invalid syntax", "test").Error(), err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}
