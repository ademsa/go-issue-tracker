package externalapi_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-issue-tracker/pkg/domain"
	"go-issue-tracker/pkg/interfaces/externalapi"
	externalapiTesting "go-issue-tracker/pkg/interfaces/externalapi/testing"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testUrl = "http://0.0.0.0:3001/externalapimock/color"

func TestNewColorRepository(t *testing.T) {
	r := externalapi.NewColorRepository(testUrl, http.DefaultClient)

	assert.Equal(t, testUrl, r.Endpoint)
	assert.NotNil(t, r.HTTPClient)
}

func TestGetColor(t *testing.T) {
	ec := domain.Color{
		HexCode: "2979FF",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"hexCode": "2979FF",
		})
	}))
	defer ts.Close()

	r := externalapi.NewColorRepository(ts.URL, http.DefaultClient)

	ac, err := r.GetColor()

	assert.Nil(t, err)
	assert.NotNil(t, ac)
	assert.Equal(t, ec, ac)
}

func TestGetColorErr(t *testing.T) {
	m := new(externalapiTesting.HTTPClientMock)

	v := new(http.Response)
	m.On("Get", testUrl).Return(v, errors.New("test error"))

	r := externalapi.NewColorRepository(testUrl, m)

	actualColor, err := r.GetColor()

	assert.NotNil(t, err)
	assert.Equal(t, domain.Color{}, actualColor)
}

func TestGetColorBodyErr(t *testing.T) {
	cm := new(externalapiTesting.HTTPClientMock)

	rm := new(externalapiTesting.ReadCloserMock)
	rm.On("Close").Return(nil)
	rm.On("Read", mock.AnythingOfType("[]uint8")).Return(0, errors.New("test error"))

	v := new(http.Response)
	v.Body = rm
	v.StatusCode = 200
	cm.On("Get", testUrl).Return(v, nil)

	r := externalapi.NewColorRepository(testUrl, cm)

	actualColor, err := r.GetColor()

	assert.NotNil(t, err)
	assert.Equal(t, domain.Color{}, actualColor)
}

func TestGetColor404(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	}))
	defer ts.Close()

	r := externalapi.NewColorRepository(ts.URL, http.DefaultClient)

	actualColor, err := r.GetColor()

	assert.NotNil(t, err)
	assert.Equal(t, domain.Color{}, actualColor)
}

func TestGetColorJSONErr(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "\"hexCode\": 2979FF\"")
	}))
	defer ts.Close()

	r := externalapi.NewColorRepository(ts.URL, http.DefaultClient)

	actualColor, err := r.GetColor()

	assert.NotNil(t, err)
	assert.Equal(t, domain.Color{}, actualColor)
}
