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
	ec := domain.Color{
		HexCode: "2979FF",
	}

	e := echo.New()

	externalapimock.PrepareEndpoints(e)

	request := httptest.NewRequest(echo.GET, externalapimock.ExternalAPIMockPath, nil)
	recorder := httptest.NewRecorder()
	e.ServeHTTP(recorder, request)
	assert.Equal(t, 200, recorder.Code)
	var ac domain.Color
	json.Unmarshal(recorder.Body.Bytes(), &ac)
	assert.Equal(t, ec, ac)
}
