package http_test

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go-issue-tracker/pkg/domain"
	"strings"
	"testing"
)

func TestAddLabel(t *testing.T) {
	l := &domain.Label{
		Name:         "test-name",
		ColorHexCode: "FFFFFF",
	}

	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	lucm.On("Add", l.Name, l.ColorHexCode).Return(l, nil)

	body := strings.NewReader("name=test-name&color_hex_code=FFFFFF")
	c, rec := prepareHTTP(echo.POST, "/api/labels/new", body)

	err := ruc.AddLabel(c)

	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestAddLabelValueErrs(t *testing.T) {
	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	var tests = []struct {
		body *strings.Reader
		err  error
	}{
		{
			strings.NewReader("name=&color_hex_code=FFFFFF"),
			errors.New("name not provided"),
		},
	}

	for _, ts := range tests {
		c, _ := prepareHTTP(echo.POST, "/api/labels/new", ts.body)

		err := ruc.AddLabel(c)

		assert.NotNil(t, err)
		assert.Equal(t, ts.err.Error(), err.Error())

		checkAssertions(t, cucm, iucm, lucm, pucm)
	}
}

func TestAddLabelColorFromExternalApi(t *testing.T) {
	l := &domain.Label{
		Name:         "test-name",
		ColorHexCode: "FF0000",
	}

	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	cucm.On("GetColor").Return(domain.Color{HexCode: "FF0000"}, nil)
	lucm.On("Add", l.Name, l.ColorHexCode).Return(l, nil)

	body := strings.NewReader("name=test-name&color_hex_code=")
	c, rec := prepareHTTP(echo.POST, "/api/labels/new", body)

	err := ruc.AddLabel(c)

	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestAddLabelColorErr(t *testing.T) {
	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	cucm.On("GetColor").Return(domain.Color{}, errors.New("test error"))

	body := strings.NewReader("name=test-name&color_hex_code=")
	c, _ := prepareHTTP(echo.POST, "/api/labels/new", body)

	err := ruc.AddLabel(c)

	assert.NotNil(t, err)
	assert.Equal(t, "test error", err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestAddLabelErr(t *testing.T) {
	p := &domain.Label{
		Name:         "test-name",
		ColorHexCode: "FFFFFF",
	}

	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	lucm.On("Add", p.Name, p.ColorHexCode).Return(p, errors.New("test error"))

	body := strings.NewReader("name=test-name&color_hex_code=FFFFFF")
	c, _ := prepareHTTP(echo.POST, "/api/labels/new", body)

	err := ruc.AddLabel(c)

	assert.NotNil(t, err)
	assert.Equal(t, "test error", err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestUpdateLabel(t *testing.T) {
	p := domain.Label{
		ID:           1,
		Name:         "test-name",
		ColorHexCode: "FFFFFF",
	}

	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	lucm.On("Update", p.ID, p.Name, p.ColorHexCode).Return(p, nil)

	body := strings.NewReader("name=test-name&color_hex_code=FFFFFF")
	c, rec := prepareHTTP(echo.POST, "/api/labels/:id", body)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := ruc.UpdateLabel(c)

	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestUpdateLabelIDErr(t *testing.T) {
	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	body := strings.NewReader("name=test-name&color_hex_code=FFFFFF")
	c, _ := prepareHTTP(echo.POST, "/api/labels/:id", body)
	c.SetParamNames("id")
	c.SetParamValues("test")

	err := ruc.UpdateLabel(c)

	assert.NotNil(t, err)
	assert.Equal(t, fmt.Errorf("strconv.Atoi: parsing \"%s\": invalid syntax", "test").Error(), err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestUpdateLabelValueErrs(t *testing.T) {
	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	var tests = []struct {
		body *strings.Reader
		err  error
	}{
		{
			strings.NewReader("name=&color_hex_code=FFFFFF"),
			errors.New("name not provided"),
		},
	}

	for _, ts := range tests {
		c, _ := prepareHTTP(echo.POST, "/api/labels/:id", ts.body)
		c.SetParamNames("id")
		c.SetParamValues("1")

		err := ruc.UpdateLabel(c)

		assert.NotNil(t, err)
		assert.Equal(t, ts.err.Error(), err.Error())

		checkAssertions(t, cucm, iucm, lucm, pucm)
	}
}

func TestUpdateLabelColorFromExternalApi(t *testing.T) {
	l := domain.Label{
		ID:           1,
		Name:         "test-name",
		ColorHexCode: "FF0000",
	}

	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	cucm.On("GetColor").Return(domain.Color{HexCode: "FF0000"}, nil)
	lucm.On("Update", l.ID, l.Name, l.ColorHexCode).Return(l, nil)

	body := strings.NewReader("name=test-name&color_hex_code=")
	c, rec := prepareHTTP(echo.POST, "/api/labels/:id", body)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := ruc.UpdateLabel(c)

	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestUpdateLabelColorErr(t *testing.T) {
	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	cucm.On("GetColor").Return(domain.Color{}, errors.New("test error"))

	body := strings.NewReader("name=test-name&color_hex_code=")
	c, _ := prepareHTTP(echo.POST, "/api/labels/:id", body)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := ruc.UpdateLabel(c)

	assert.NotNil(t, err)
	assert.Equal(t, "test error", err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestUpdateLabelErr(t *testing.T) {
	p := domain.Label{
		ID:           1,
		Name:         "test-name",
		ColorHexCode: "FFFFFF",
	}

	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	lucm.On("Update", p.ID, p.Name, p.ColorHexCode).Return(p, errors.New("test error"))

	body := strings.NewReader("name=test-name&color_hex_code=FFFFFF")
	c, _ := prepareHTTP(echo.POST, "/api/labels/:id", body)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := ruc.UpdateLabel(c)

	assert.NotNil(t, err)
	assert.Equal(t, "test error", err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestFindLabelByID(t *testing.T) {
	l := domain.Label{
		ID:           1,
		Name:         "test-name",
		ColorHexCode: "FFFFFF",
	}

	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	lucm.On("FindByID", l.ID).Return(l, nil)

	c, rec := prepareHTTP(echo.GET, "/api/labels/:id", nil)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := ruc.FindLabelByID(c)

	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestFindLabelByIDIDErr(t *testing.T) {
	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	c, _ := prepareHTTP(echo.GET, "/api/labels/:id", nil)
	c.SetParamNames("id")
	c.SetParamValues("test")

	err := ruc.FindLabelByID(c)

	assert.NotNil(t, err)
	assert.Equal(t, fmt.Errorf("strconv.Atoi: parsing \"%s\": invalid syntax", "test").Error(), err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestFindLabelByIDNotFoundNoErr(t *testing.T) {
	l := domain.Label{
		ID:           1,
		Name:         "test-name",
		ColorHexCode: "FFFFFF",
	}

	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	lucm.On("FindByID", l.ID).Return(l, errors.New("record not found"))

	c, _ := prepareHTTP(echo.GET, "/api/labels/:id", nil)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := ruc.FindLabelByID(c)

	assert.Nil(t, err)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestFindLabelByIDOtherErr(t *testing.T) {
	l := domain.Label{
		ID:           1,
		Name:         "test-name",
		ColorHexCode: "FFFFFF",
	}

	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	lucm.On("FindByID", l.ID).Return(l, errors.New("test error"))

	c, _ := prepareHTTP(echo.GET, "/api/labels/:id", nil)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := ruc.FindLabelByID(c)

	assert.NotNil(t, err)
	assert.Equal(t, "test error", err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestFindLabelByName(t *testing.T) {
	l := domain.Label{
		ID:           1,
		Name:         "test-name",
		ColorHexCode: "FFFFFF",
	}

	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	lucm.On("FindByName", l.Name).Return(l, nil)

	c, rec := prepareHTTP(echo.GET, "/api/labels/findbyname/:name", nil)
	c.SetParamNames("name")
	c.SetParamValues(l.Name)

	err := ruc.FindLabelByName(c)

	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestFindLabelByNameNameErr(t *testing.T) {
	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	c, _ := prepareHTTP(echo.GET, "/api/labels/findbyname/:name", nil)
	c.SetParamNames("name")
	c.SetParamValues("")

	err := ruc.FindLabelByName(c)

	assert.NotNil(t, err)
	assert.Equal(t, "label name not provided", err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestFindLabelByNameNotFoundNoErr(t *testing.T) {
	l := domain.Label{
		ID:           1,
		Name:         "test-name",
		ColorHexCode: "FFFFFF",
	}

	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	lucm.On("FindByName", l.Name).Return(l, errors.New("record not found"))

	c, _ := prepareHTTP(echo.GET, "/api/labels/findbyname/:name", nil)
	c.SetParamNames("name")
	c.SetParamValues(l.Name)

	err := ruc.FindLabelByName(c)

	assert.Nil(t, err)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestFindLabelByNameOtherErr(t *testing.T) {
	l := domain.Label{
		ID:           1,
		Name:         "test-name",
		ColorHexCode: "FFFFFF",
	}

	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	lucm.On("FindByName", l.Name).Return(l, errors.New("test error"))

	c, _ := prepareHTTP(echo.GET, "/api/labels/findbyname/:name", nil)
	c.SetParamNames("name")
	c.SetParamValues(l.Name)

	err := ruc.FindLabelByName(c)

	assert.NotNil(t, err)
	assert.Equal(t, "test error", err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestFindAllLabels(t *testing.T) {
	l := []domain.Label{}

	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	lucm.On("FindAll").Return(l, nil)

	c, rec := prepareHTTP(echo.GET, "/api/labels", nil)

	err := ruc.FindAllLabels(c)

	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestFindAllLabelsErr(t *testing.T) {
	l := []domain.Label{}

	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	lucm.On("FindAll").Return(l, errors.New("test error"))

	c, _ := prepareHTTP(echo.GET, "/api/labels", nil)

	err := ruc.FindAllLabels(c)

	assert.NotNil(t, err)
	assert.Equal(t, "test error", err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestRemoveLabel(t *testing.T) {
	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	lucm.On("Remove", uint(1)).Return(true, nil)

	c, rec := prepareHTTP(echo.DELETE, "/api/labels/:id", nil)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := ruc.RemoveLabel(c)

	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestRemoveLabelNotFoundNoErr(t *testing.T) {
	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	lucm.On("Remove", uint(1)).Return(false, errors.New("record not found"))

	c, _ := prepareHTTP(echo.DELETE, "/api/labels/:id", nil)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := ruc.RemoveLabel(c)

	assert.Nil(t, err)

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestRemoveLabelErr(t *testing.T) {
	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	lucm.On("Remove", uint(1)).Return(false, errors.New("test error"))

	c, _ := prepareHTTP(echo.DELETE, "/api/labels/:id", nil)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := ruc.RemoveLabel(c)

	assert.NotNil(t, err)
	assert.Equal(t, "test error", err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}

func TestRemoveLabelIDErr(t *testing.T) {
	cucm, iucm, lucm, pucm, ruc := prepareMocksAndRUC()

	c, _ := prepareHTTP(echo.DELETE, "/api/labels/:id", nil)
	c.SetParamNames("id")
	c.SetParamValues("test")

	err := ruc.RemoveLabel(c)

	assert.NotNil(t, err)
	assert.Equal(t, fmt.Errorf("strconv.Atoi: parsing \"%s\": invalid syntax", "test").Error(), err.Error())

	checkAssertions(t, cucm, iucm, lucm, pucm)
}
