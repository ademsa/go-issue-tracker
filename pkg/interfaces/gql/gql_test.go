package gql_test

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-issue-tracker/pkg/interfaces/gql"
	gqlTesting "go-issue-tracker/pkg/interfaces/gql/testing"
	ucTesting "go-issue-tracker/pkg/usecases/testing"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func prepareHTTP(method string, path string, body *strings.Reader) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	var req *http.Request
	req = httptest.NewRequest(method, path, body)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	return e.NewContext(req, rec), rec
}

func TestPrepareGraphQL(t *testing.T) {
	cucm := new(ucTesting.ColorUseCaseMock)
	iucm := new(ucTesting.IssueUseCaseMock)
	lucm := new(ucTesting.LabelUseCaseMock)
	pucm := new(ucTesting.ProjectUseCaseMock)

	schema := gql.PrepareGraphQL(iucm, lucm, pucm, cucm)

	assert.NotNil(t, schema)
}

func TestPrepareEndpoints(t *testing.T) {
	e := echo.New()

	gqlm := new(gqlTesting.RequestManagerMock)

	gql.PrepareEndpoints(e, gqlm)

	// /graphql POST
	request := httptest.NewRequest(echo.POST, "/graphql", nil)
	recorder := httptest.NewRecorder()

	gqlm.On("Handler", mock.AnythingOfType("*echo.context")).Return(nil)

	e.ServeHTTP(recorder, request)

	assert.Equal(t, 200, recorder.Code)

	gqlm.AssertExpectations(t)
}

func TestHandler(t *testing.T) {
	gqlSchema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    nil,
		Mutation: nil,
	})
	gqlm := gql.NewRequestManager(gqlSchema)

	json := fmt.Sprintf(`
	{"query":"query {\n    label (id: 1){\n        id\n        name\n        createdAt\n    }\n}"}
	`)
	body := strings.NewReader(json)
	c, rec := prepareHTTP(echo.POST, "/graphql", body)

	err := gqlm.Handler(c)

	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)
}

func TestHandlerMutation(t *testing.T) {
	cucm := new(ucTesting.ColorUseCaseMock)
	iucm := new(ucTesting.IssueUseCaseMock)
	lucm := new(ucTesting.LabelUseCaseMock)
	pucm := new(ucTesting.ProjectUseCaseMock)

	lucm.On("Remove", uint(1)).Return(true, nil)

	schema := gql.PrepareGraphQL(iucm, lucm, pucm, cucm)

	gqlm := gql.NewRequestManager(schema)

	json := fmt.Sprintf(`
	{"query":"mutation {\n    removeLabel (id: \"%s\"){\n        id\n    }\n}"}
	`, relay.ToGlobalID("Label", "1"))
	body := strings.NewReader(json)
	c, rec := prepareHTTP(echo.POST, "/graphql", body)

	err := gqlm.Handler(c)

	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)
}

func TestHandlerMissingQuery(t *testing.T) {
	cucm := new(ucTesting.ColorUseCaseMock)
	iucm := new(ucTesting.IssueUseCaseMock)
	lucm := new(ucTesting.LabelUseCaseMock)
	pucm := new(ucTesting.ProjectUseCaseMock)

	lucm.On("Remove", uint(1)).Return(true, nil)

	schema := gql.PrepareGraphQL(iucm, lucm, pucm, cucm)

	gqlm := gql.NewRequestManager(schema)

	json := fmt.Sprintf(`
	{"query2":"query {\n    label (id: 1){\n        id\n        name\n        createdAt\n    }\n}"}
	`)
	body := strings.NewReader(json)
	c, rec := prepareHTTP(echo.POST, "/graphql", body)

	err := gqlm.Handler(c)

	assert.Nil(t, err)
	assert.Equal(t, 500, rec.Code)
}

func TestHandlerExecutionErr(t *testing.T) {
	gqlSchema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    nil,
		Mutation: nil,
	})
	gqlm := gql.NewRequestManager(gqlSchema)

	json := fmt.Sprintf(`
	{"query":"queryWRONG {\n    label (id: 1){\n        id\n        name\n        createdAt\n    }\n}"}
	`)
	body := strings.NewReader(json)
	c, rec := prepareHTTP(echo.POST, "/graphql", body)

	err := gqlm.Handler(c)

	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)
}
