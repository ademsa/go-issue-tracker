package gql

import (
	"encoding/json"
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
	"go-issue-tracker/pkg/usecases"
	"log"
)

// PrepareGraphQL function to prepare GraphQL
func PrepareGraphQL(iuc usecases.IssueUseCase, luc usecases.LabelUseCase, puc usecases.ProjectUseCase, cuc usecases.ColorUseCase) graphql.Schema {
	resolver := GetResolver(iuc, luc, puc, cuc)

	SetTypesAndNodeDefinitions(resolver)

	query := GetQuery(resolver)
	mutation := GetMutation(resolver)
	schema, err := GetSchema(query, mutation)
	if err != nil {
		log.Fatal(err)
	}
	return schema
}

// PrepareEndpoints function to prepare GraphQL endpoints
func PrepareEndpoints(e *echo.Echo, rm RequestManager) {
	e.POST("/graphql", rm.Handler)
}

// Handler handler
func (rm requestManager) Handler(c echo.Context) error {
	var actual map[string]interface{}
	json.NewDecoder(c.Request().Body).Decode(&actual)
	requestString, requestStringOK := actual["query"].(string)
	if !requestStringOK {
		return c.JSON(500, map[string]interface{}{
			"error": errors.New("request not complete"),
		})
	}
	variables, _ := actual["variables"].(map[string]interface{})
	result := rm.Execute(rm.schema, requestString, variables)
	if result.HasErrors() {
		return c.JSON(200, map[string]interface{}{
			"error": result.Errors,
		})
	}
	return c.JSON(200, result)
}

// Execute func
func (rm requestManager) Execute(schema graphql.Schema, requestString string, variableValues map[string]interface{}) *graphql.Result {
	return graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  requestString,
		VariableValues: variableValues,
	})
}
