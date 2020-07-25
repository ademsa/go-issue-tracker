package rest_test

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-issue-tracker/pkg/infrastructure/helpers"
	"go-issue-tracker/pkg/interfaces/rest"
	restTesting "go-issue-tracker/pkg/interfaces/rest/testing"
	ucTesting "go-issue-tracker/pkg/usecases/testing"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPrepareServer(t *testing.T) {
	httpServer := rest.PrepareServer()

	assert.NotNil(t, httpServer)
	assert.True(t, httpServer.HideBanner)
}

func TestPrepareEndpoints(t *testing.T) {
	e := echo.New()

	rm := new(restTesting.ManagerMock)

	rootDirPath, _ := helpers.GetProjectDirPath()
	uiDirPath := filepath.Join(rootDirPath, "ui")
	rest.PrepareEndpoints(e, rm, uiDirPath)

	// /api GET
	request := httptest.NewRequest(echo.GET, "/api", nil)
	recorder := httptest.NewRecorder()
	e.ServeHTTP(recorder, request)
	assert.Equal(t, 200, recorder.Code)
	var actual map[string]interface{}
	json.NewDecoder(recorder.Body).Decode(&actual)
	assert.Equal(t, map[string]interface{}{
		"message": rest.APIRootMessage,
	}, actual)

	// /api/issues/new POST
	checkPath(t, rm, e, echo.POST, "/api/issues/new", "AddIssue")

	// /api/issues/:id POST
	checkPath(t, rm, e, echo.POST, "/api/issues/:id", "UpdateIssue")

	// /api/issues/:id GET
	checkPath(t, rm, e, echo.GET, "/api/issues/:id", "FindIssueByID")

	// /api/issues/find GET
	checkPath(t, rm, e, echo.GET, "/api/issues/find", "FindIssues")

	// /api/issues GET
	checkPath(t, rm, e, echo.GET, "/api/issues", "FindAllIssues")

	// /api/issues/:id DELETE
	checkPath(t, rm, e, echo.DELETE, "/api/issues/:id", "RemoveIssue")

	// /api/labels/new POST
	checkPath(t, rm, e, echo.POST, "/api/labels/new", "AddLabel")

	// /api/labels/:id POST
	checkPath(t, rm, e, echo.POST, "/api/labels/:id", "UpdateLabel")

	// /api/labels/:id GET
	checkPath(t, rm, e, echo.GET, "/api/labels/:id", "FindLabelByID")

	// /api/labels/find GET
	checkPath(t, rm, e, echo.GET, "/api/labels/find", "FindLabels")

	// /api/labels GET
	checkPath(t, rm, e, echo.GET, "/api/labels", "FindAllLabels")

	// /api/labels/:id DELETE
	checkPath(t, rm, e, echo.DELETE, "/api/labels/:id", "RemoveLabel")

	// /api/projects/new POST
	checkPath(t, rm, e, echo.POST, "/api/projects/new", "AddProject")

	// /api/projects/:id POST
	checkPath(t, rm, e, echo.POST, "/api/projects/:id", "UpdateProject")

	// /api/projects/:id GET
	checkPath(t, rm, e, echo.GET, "/api/projects/:id", "FindProjectByID")

	// /api/projects/find GET
	checkPath(t, rm, e, echo.GET, "/api/projects/find", "FindProjects")

	// /api/projects GET
	checkPath(t, rm, e, echo.GET, "/api/projects", "FindAllProjects")

	// /api/projects/:id DELETE
	checkPath(t, rm, e, echo.DELETE, "/api/projects/:id", "RemoveProject")

	rm.AssertExpectations(t)
}

func TestPrepareEndpointsRootStatic(t *testing.T) {
	e := echo.New()

	rm := new(restTesting.ManagerMock)

	rootDirPath, _ := helpers.GetProjectDirPath()
	uiDirPath := filepath.Join(rootDirPath, "ui")
	helpers.Stat = func(path string) (os.FileInfo, error) {
		return nil, nil
	}
	defer func() {
		helpers.Stat = os.Stat
	}()
	rest.PrepareEndpoints(e, rm, uiDirPath)

	// / GET
	request := httptest.NewRequest(echo.GET, "/", nil)
	recorder := httptest.NewRecorder()
	e.ServeHTTP(recorder, request)
	assert.Equal(t, 404, recorder.Code)

	rm.AssertExpectations(t)
}

func TestPrepareEndpointsRootNotStatic(t *testing.T) {
	e := echo.New()

	rm := new(restTesting.ManagerMock)

	rootDirPath, _ := helpers.GetProjectDirPath()
	uiDirPath := filepath.Join(rootDirPath, "ui")
	rest.PrepareEndpoints(e, rm, uiDirPath)

	// / GET
	request := httptest.NewRequest(echo.GET, "/", nil)
	recorder := httptest.NewRecorder()
	e.ServeHTTP(recorder, request)
	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, rest.RootMessage, recorder.Body.String())

	rm.AssertExpectations(t)
}

func checkPath(t *testing.T, rm *restTesting.ManagerMock, e *echo.Echo, httpMethod string, path string, mockMethod string) {
	request := httptest.NewRequest(httpMethod, path, nil)
	recorder := httptest.NewRecorder()

	rm.On(mockMethod, mock.AnythingOfType("*echo.context")).Return(nil)

	e.ServeHTTP(recorder, request)

	assert.Equal(t, 200, recorder.Code)
}

func prepareHTTP(method string, path string, body *strings.Reader) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	var req *http.Request
	if body != nil {
		req = httptest.NewRequest(method, path, body)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req = httptest.NewRequest(method, path, nil)
	}
	rec := httptest.NewRecorder()
	return e.NewContext(req, rec), rec
}

func prepareMocksAndRUC() (*ucTesting.ColorUseCaseMock, *ucTesting.IssueUseCaseMock, *ucTesting.LabelUseCaseMock, *ucTesting.ProjectUseCaseMock, rest.Manager) {
	cucm := new(ucTesting.ColorUseCaseMock)
	iucm := new(ucTesting.IssueUseCaseMock)
	lucm := new(ucTesting.LabelUseCaseMock)
	pucm := new(ucTesting.ProjectUseCaseMock)
	return cucm, iucm, lucm, pucm, rest.NewManager(iucm, lucm, pucm, cucm)
}

func checkAssertions(t *testing.T, cucm *ucTesting.ColorUseCaseMock, iucm *ucTesting.IssueUseCaseMock, lucm *ucTesting.LabelUseCaseMock, pucm *ucTesting.ProjectUseCaseMock) {
	cucm.AssertExpectations(t)
	iucm.AssertExpectations(t)
	lucm.AssertExpectations(t)
	pucm.AssertExpectations(t)
}
