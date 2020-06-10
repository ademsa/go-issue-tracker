package http_test

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-issue-tracker/pkg/domain"
	"strings"
	"testing"
)

func TestAddIssue(t *testing.T) {
	p := domain.Project{}
	i := &domain.Issue{
		Title:       "test-title",
		Description: "test-description",
		Status:      1,
		ProjectID:   1,
	}
	labels := map[string]domain.Label{
		"test1": domain.Label{},
		"test2": domain.Label{},
		"test3": domain.Label{},
	}

	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	iucm.On("Add", i.Title, i.Description, i.Status, p, labels).Return(i, nil)
	lucm.On("FindByName", mock.AnythingOfType("string")).Return(domain.Label{}, nil)
	pucm.On("FindByID", mock.AnythingOfType("uint")).Return(p, nil)

	body := strings.NewReader("project_id=1&title=test-title&description=test-description&status=1&labels=test1,test2,test3")
	c, rec := prepareHTTP(echo.POST, "/api/issues/new", body)

	err := ruc.AddIssue(c)

	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestAddIssueValueErrs(t *testing.T) {
	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	var tests = []struct {
		body *strings.Reader
		err  error
	}{
		{
			strings.NewReader("project_id=1&title=&description=test-description&status=1&labels=test1,test2,test3"),
			errors.New("title not provided"),
		},
		{
			strings.NewReader("project_id=1&title=test-title&description=&status=1&labels=test1,test2,test3"),
			errors.New("description not provided"),
		},
		{
			strings.NewReader("project_id=1&title=test-title&description=test-description&status=test&labels=test1,test2,test3"),
			fmt.Errorf("strconv.Atoi: parsing \"%s\": invalid syntax", "test"),
		},
		{
			strings.NewReader("project_id=test&title=test-title&description=test-description&status=1&labels=test1,test2,test3"),
			fmt.Errorf("strconv.Atoi: parsing \"%s\": invalid syntax", "test"),
		},
	}

	for _, ts := range tests {
		c, _ := prepareHTTP(echo.POST, "/api/issues/new", ts.body)

		err := ruc.AddIssue(c)

		assert.NotNil(t, err)
		assert.Equal(t, ts.err.Error(), err.Error())

		checkAssertions(t, cucm, iucm, lucm, pucm)
	}
}

func TestAddIssueValueProjectNotFoundErr(t *testing.T) {
	p := domain.Project{}

	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	pucm.On("FindByID", mock.AnythingOfType("uint")).Return(p, errors.New("project not found"))

	body := strings.NewReader("project_id=1&title=test-title&description=test-description&status=1&labels=test1,test2,test3")
	c, _ := prepareHTTP(echo.POST, "/api/issues/new", body)

	err := ruc.AddIssue(c)

	assert.NotNil(t, err)
	assert.Equal(t, "project not found", err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestAddIssueValueLabelErrs(t *testing.T) {
	p := domain.Project{}

	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	pucm.On("FindByID", mock.AnythingOfType("uint")).Return(p, nil)

	var tests = []struct {
		body   *strings.Reader
		err    error
		mockOn bool
	}{
		{
			strings.NewReader("project_id=1&title=test-title&description=test-description&status=1&labels="),
			errors.New("no labels assigned"),
			false,
		},
		{
			strings.NewReader("project_id=1&title=test-title&description=test-description&status=1&labels=test1,test2,test3"),
			errors.New("label test1 is not valid"),
			true,
		},
	}

	for _, ts := range tests {
		if ts.mockOn {
			lucm.On("FindByName", mock.AnythingOfType("string")).Return(domain.Label{}, ts.err)
		}
		c, _ := prepareHTTP(echo.POST, "/api/issues/new", ts.body)

		err := ruc.AddIssue(c)

		assert.NotNil(t, err)
		assert.Equal(t, ts.err.Error(), err.Error())

		checkAssertions(t, cucm, iucm, lucm, pucm)
	}
}

func TestAddIssueErr(t *testing.T) {
	p := domain.Project{
		ID: 1,
	}
	i := &domain.Issue{
		Title:       "test-title",
		Description: "test-description",
		Status:      1,
		ProjectID:   1,
	}
	labels := map[string]domain.Label{
		"test1": domain.Label{},
		"test2": domain.Label{},
		"test3": domain.Label{},
	}

	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	pucm.On("FindByID", mock.AnythingOfType("uint")).Return(p, nil)
	lucm.On("FindByName", mock.AnythingOfType("string")).Return(domain.Label{}, nil)
	iucm.On("Add", i.Title, i.Description, i.Status, p, labels).Return(i, errors.New("test error"))

	body := strings.NewReader("project_id=1&title=test-title&description=test-description&status=1&labels=test1,test2,test3")
	c, _ := prepareHTTP(echo.POST, "/api/issues/new", body)

	err := ruc.AddIssue(c)

	assert.NotNil(t, err)
	assert.Equal(t, "test error", err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestUpdateIssue(t *testing.T) {
	i := domain.Issue{
		ID:          1,
		Title:       "test-title",
		Description: "test-description",
		Status:      1,
		ProjectID:   1,
	}
	labels := map[string]domain.Label{
		"test1": domain.Label{},
		"test2": domain.Label{},
		"test3": domain.Label{},
	}

	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	lucm.On("FindByName", mock.AnythingOfType("string")).Return(domain.Label{}, nil)
	iucm.On("Update", i.ID, i.Title, i.Description, i.Status, labels).Return(i, nil)

	body := strings.NewReader("title=test-title&description=test-description&status=1&labels=test1,test2,test3")
	c, rec := prepareHTTP(echo.POST, "/api/issues/:id", body)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := ruc.UpdateIssue(c)

	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestUpdateIssueIDErr(t *testing.T) {
	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	body := strings.NewReader("title=test-title&description=test-description&status=1&labels=test1,test2,test3")
	c, _ := prepareHTTP(echo.POST, "/api/issues/:id", body)
	c.SetParamNames("id")
	c.SetParamValues("test")

	err := ruc.UpdateIssue(c)

	assert.NotNil(t, err)
	assert.Equal(t, fmt.Errorf("strconv.Atoi: parsing \"%s\": invalid syntax", "test").Error(), err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestUpdateIssueValueErrs(t *testing.T) {
	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	var tests = []struct {
		body *strings.Reader
		err  error
	}{
		{
			strings.NewReader("title=&description=test-description&status=1&labels=test1,test2,test3"),
			errors.New("title not provided"),
		},
		{
			strings.NewReader("title=test-title&description=&status=1&labels=test1,test2,test3"),
			errors.New("description not provided"),
		},
		{
			strings.NewReader("title=test-title&description=test-description&status=test&labels=test1,test2,test3"),
			fmt.Errorf("strconv.Atoi: parsing \"%s\": invalid syntax", "test"),
		},
	}

	for _, ts := range tests {
		c, _ := prepareHTTP(echo.POST, "/api/issues/:id", ts.body)
		c.SetParamNames("id")
		c.SetParamValues("1")

		err := ruc.UpdateIssue(c)

		assert.NotNil(t, err)
		assert.Equal(t, ts.err.Error(), err.Error())

		checkAssertions(t, cucm, iucm, lucm, pucm)
	}
}

func TestUpdateIssueValueLabelErrs(t *testing.T) {
	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	var tests = []struct {
		body   *strings.Reader
		err    error
		mockOn bool
	}{
		{
			strings.NewReader("title=test-title&description=test-description&status=1&labels="),
			errors.New("no labels assigned"),
			false,
		},
		{
			strings.NewReader("title=test-title&description=test-description&status=1&labels=test1,test2,test3"),
			errors.New("label test1 is not valid"),
			true,
		},
	}

	for _, ts := range tests {
		if ts.mockOn {
			lucm.On("FindByName", mock.AnythingOfType("string")).Return(domain.Label{}, ts.err)
		}
		c, _ := prepareHTTP(echo.POST, "/api/issues/:id", ts.body)
		c.SetParamNames("id")
		c.SetParamValues("1")

		err := ruc.UpdateIssue(c)

		assert.NotNil(t, err)
		assert.Equal(t, ts.err.Error(), err.Error())

		checkAssertions(t, cucm, iucm, lucm, pucm)
	}
}

func TestUpdateIssueErr(t *testing.T) {
	i := domain.Issue{
		ID:          1,
		Title:       "test-title",
		Description: "test-description",
		Status:      1,
		ProjectID:   1,
	}
	labels := map[string]domain.Label{
		"test1": domain.Label{},
		"test2": domain.Label{},
		"test3": domain.Label{},
	}

	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	lucm.On("FindByName", mock.AnythingOfType("string")).Return(domain.Label{}, nil)
	iucm.On("Update", i.ID, i.Title, i.Description, i.Status, labels).Return(i, errors.New("test error"))

	body := strings.NewReader("title=test-title&description=test-description&status=1&labels=test1,test2,test3")
	c, _ := prepareHTTP(echo.POST, "/api/issues/:id", body)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := ruc.UpdateIssue(c)

	assert.NotNil(t, err)
	assert.Equal(t, "test error", err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestFindIssueByID(t *testing.T) {
	i := domain.Issue{
		ID:          1,
		Title:       "test-title",
		Description: "test-description",
		Status:      1,
		ProjectID:   1,
	}

	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	iucm.On("FindByID", i.ID).Return(i, nil)

	c, rec := prepareHTTP(echo.GET, "/api/issues/:id", nil)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := ruc.FindIssueByID(c)

	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestFindIssueByIDIDErr(t *testing.T) {
	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	c, _ := prepareHTTP(echo.GET, "/api/issues/:id", nil)
	c.SetParamNames("id")
	c.SetParamValues("test")

	err := ruc.FindIssueByID(c)

	assert.NotNil(t, err)
	assert.Equal(t, fmt.Errorf("strconv.Atoi: parsing \"%s\": invalid syntax", "test").Error(), err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestFindIssueByIDNotFoundNoErr(t *testing.T) {
	i := domain.Issue{
		ID:          1,
		Title:       "test-title",
		Description: "test-description",
		Status:      1,
		ProjectID:   1,
	}

	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	iucm.On("FindByID", i.ID).Return(i, errors.New("record not found"))

	c, _ := prepareHTTP(echo.GET, "/api/issues/:id", nil)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := ruc.FindIssueByID(c)

	assert.Nil(t, err)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestFindIssueByIDOtherErr(t *testing.T) {
	i := domain.Issue{
		ID:          1,
		Title:       "test-title",
		Description: "test-description",
		Status:      1,
		ProjectID:   1,
	}

	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	iucm.On("FindByID", i.ID).Return(i, errors.New("test error"))

	c, _ := prepareHTTP(echo.GET, "/api/issues/:id", nil)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := ruc.FindIssueByID(c)

	assert.NotNil(t, err)
	assert.Equal(t, "test error", err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestFindAllIssues(t *testing.T) {
	i := []domain.Issue{}

	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	iucm.On("FindAll").Return(i, nil)

	c, rec := prepareHTTP(echo.GET, "/api/issues", nil)

	err := ruc.FindAllIssues(c)

	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestFindAllIssuesErr(t *testing.T) {
	i := []domain.Issue{}

	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	iucm.On("FindAll").Return(i, errors.New("test error"))

	c, _ := prepareHTTP(echo.GET, "/api/issues", nil)

	err := ruc.FindAllIssues(c)

	assert.NotNil(t, err)
	assert.Equal(t, "test error", err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestRemoveIssue(t *testing.T) {
	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	iucm.On("Remove", uint(1)).Return(true, nil)

	c, rec := prepareHTTP(echo.DELETE, "/api/issues/:id", nil)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := ruc.RemoveIssue(c)

	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestRemoveIssueNotFoundNoErr(t *testing.T) {
	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	iucm.On("Remove", uint(1)).Return(false, errors.New("record not found"))

	c, _ := prepareHTTP(echo.DELETE, "/api/issues/:id", nil)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := ruc.RemoveIssue(c)

	assert.Nil(t, err)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestRemoveIssueErr(t *testing.T) {
	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	iucm.On("Remove", uint(1)).Return(false, errors.New("test error"))

	c, _ := prepareHTTP(echo.DELETE, "/api/issues/:id", nil)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := ruc.RemoveIssue(c)

	assert.NotNil(t, err)
	assert.Equal(t, "test error", err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestRemoveIssueIDErr(t *testing.T) {
	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	c, _ := prepareHTTP(echo.DELETE, "/api/issues/:id", nil)
	c.SetParamNames("id")
	c.SetParamValues("test")

	err := ruc.RemoveIssue(c)

	assert.NotNil(t, err)
	assert.Equal(t, fmt.Errorf("strconv.Atoi: parsing \"%s\": invalid syntax", "test").Error(), err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}
