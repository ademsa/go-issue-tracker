package http_test

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-issue-tracker/pkg/interfaces/http"
	httpTesting "go-issue-tracker/pkg/interfaces/http/testing"
	ucTesting "go-issue-tracker/pkg/usecases/testing"
	netHttp "net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPrepareHTTPServer(t *testing.T) {
	cucm := new(ucTesting.ColorUseCaseMock)
	iucm := new(ucTesting.IssueUseCaseMock)
	lucm := new(ucTesting.LabelUseCaseMock)
	pucm := new(ucTesting.ProjectUseCaseMock)

	httpServer, ruc := http.PrepareHTTPServer(iucm, lucm, pucm, cucm)

	assert.NotNil(t, httpServer)
	assert.True(t, httpServer.HideBanner)
	assert.NotNil(t, ruc)
}

func TestPrepareEndpoints(t *testing.T) {
	e := echo.New()

	rucm := new(httpTesting.RESTUseCaseMock)

	http.PrepareEndpoints(e, rucm)

	// / GET
	request := httptest.NewRequest(echo.GET, "/", nil)
	recorder := httptest.NewRecorder()
	e.ServeHTTP(recorder, request)
	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, http.RootMessage, recorder.Body.String())

	// /api GET
	request = httptest.NewRequest(echo.GET, "/api", nil)
	recorder = httptest.NewRecorder()
	e.ServeHTTP(recorder, request)
	assert.Equal(t, 200, recorder.Code)
	var actual map[string]interface{}
	json.NewDecoder(recorder.Body).Decode(&actual)
	assert.Equal(t, map[string]interface{}{
		"message": http.APIRootMessage,
	}, actual)

	// /api/issues/new POST
	checkPath(t, rucm, e, echo.POST, "/api/issues/new", "AddIssue")

	// /api/issues/:id POST
	checkPath(t, rucm, e, echo.POST, "/api/issues/:id", "UpdateIssue")

	// /api/issues/:id GET
	checkPath(t, rucm, e, echo.GET, "/api/issues/:id", "FindIssueByID")

	// /api/issues GET
	checkPath(t, rucm, e, echo.GET, "/api/issues", "FindAllIssues")

	// /api/issues/:id DELETE
	checkPath(t, rucm, e, echo.DELETE, "/api/issues/:id", "RemoveIssue")

	// /api/labels/new POST
	checkPath(t, rucm, e, echo.POST, "/api/labels/new", "AddLabel")

	// /api/labels/:id POST
	checkPath(t, rucm, e, echo.POST, "/api/labels/:id", "UpdateLabel")

	// /api/labels/:id GET
	checkPath(t, rucm, e, echo.GET, "/api/labels/:id", "FindLabelByID")

	// /api/labels/findbyname/:name GET
	checkPath(t, rucm, e, echo.GET, "/api/labels/findbyname/:name", "FindLabelByName")

	// /api/labels GET
	checkPath(t, rucm, e, echo.GET, "/api/labels", "FindAllLabels")

	// /api/labels/:id DELETE
	checkPath(t, rucm, e, echo.DELETE, "/api/labels/:id", "RemoveLabel")

	// /api/projects/new POST
	checkPath(t, rucm, e, echo.POST, "/api/projects/new", "AddProject")

	// /api/projects/:id POST
	checkPath(t, rucm, e, echo.POST, "/api/projects/:id", "UpdateProject")

	// /api/projects/:id GET
	checkPath(t, rucm, e, echo.GET, "/api/projects/:id", "FindProjectByID")

	// /api/projects GET
	checkPath(t, rucm, e, echo.GET, "/api/projects", "FindAllProjects")

	// /api/projects/:id DELETE
	checkPath(t, rucm, e, echo.DELETE, "/api/projects/:id", "RemoveProject")

	rucm.AssertExpectations(t)
}

func checkPath(t *testing.T, rucm *httpTesting.RESTUseCaseMock, e *echo.Echo, httpMethod string, path string, mockMethod string) {
	request := httptest.NewRequest(httpMethod, path, nil)
	recorder := httptest.NewRecorder()

	rucm.On(mockMethod, mock.AnythingOfType("*echo.context")).Return(nil)

	e.ServeHTTP(recorder, request)

	assert.Equal(t, 200, recorder.Code)
}

func prepareHTTP(method string, path string, body *strings.Reader) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	var req *netHttp.Request
	if body != nil {
		req = httptest.NewRequest(method, path, body)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req = httptest.NewRequest(method, path, nil)
	}
	rec := httptest.NewRecorder()
	return e.NewContext(req, rec), rec
}

func prepareMocksAndRUC() (*ucTesting.ColorUseCaseMock, *ucTesting.IssueUseCaseMock, *ucTesting.LabelUseCaseMock, *ucTesting.ProjectUseCaseMock, http.RESTUseCase) {
	cucm := new(ucTesting.ColorUseCaseMock)
	iucm := new(ucTesting.IssueUseCaseMock)
	lucm := new(ucTesting.LabelUseCaseMock)
	pucm := new(ucTesting.ProjectUseCaseMock)
	return cucm, iucm, lucm, pucm, http.NewRESTUseCase(iucm, lucm, pucm, cucm)
}

func checkAssertions(t *testing.T, cucm *ucTesting.ColorUseCaseMock, iucm *ucTesting.IssueUseCaseMock, lucm *ucTesting.LabelUseCaseMock, pucm *ucTesting.ProjectUseCaseMock) {
	cucm.AssertExpectations(t)
	iucm.AssertExpectations(t)
	lucm.AssertExpectations(t)
	pucm.AssertExpectations(t)
}
