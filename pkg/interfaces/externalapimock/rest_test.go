package externalapimock_test

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go-issue-tracker/pkg/domain"
	"go-issue-tracker/pkg/interfaces/externalapimock"
	"net/http/httptest"
	"testing"
)

func TestPrepareEndpoints(t *testing.T) {
	expectedColor := domain.Color{
		HexCode: "2979FF",
	}

	e := echo.New()

	externalapimock.PrepareEndpoints(e)

	request := httptest.NewRequest(echo.GET, externalapimock.ExternalAPIMockPath, nil)
	recorder := httptest.NewRecorder()
	e.ServeHTTP(recorder, request)
	assert.Equal(t, 200, recorder.Code)
	var actualColor domain.Color
	json.Unmarshal(recorder.Body.Bytes(), &actualColor)
	assert.Equal(t, expectedColor, actualColor)
}
