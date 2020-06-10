package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-issue-tracker/pkg/usecases"
)

// RootMessage is displayed in / path
var RootMessage = "Issue Tracker"

// APIRootMessage is displayed in /api path
var APIRootMessage = "Issue Tracker Api"

// PrepareHTTPServer function to prepare HTTP Server and Use Case Manager
func PrepareHTTPServer(iuc usecases.IssueUseCase, luc usecases.LabelUseCase, puc usecases.ProjectUseCase, cuc usecases.ColorUseCase) (*echo.Echo, RESTUseCase) {
	ruc := NewRESTUseCase(iuc, luc, puc, cuc)

	httpServer := echo.New()

	httpServer.HideBanner = true

	httpServer.Use(middleware.Logger())

	return httpServer, ruc
}

// PrepareEndpoints function to prepare endpoints
func PrepareEndpoints(e *echo.Echo, ruc RESTUseCase) {
	e.GET("/", func(c echo.Context) error {
		return c.String(200, RootMessage)
	})
	e.GET("/api", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"message": APIRootMessage,
		})
	})

	api := e.Group("/api")

	api.POST("/issues/new", ruc.AddIssue)
	api.POST("/issues/:id", ruc.UpdateIssue)
	api.GET("/issues/:id", ruc.FindIssueByID)
	api.GET("/issues", ruc.FindAllIssues)
	api.DELETE("/issues/:id", ruc.RemoveIssue)

	api.POST("/labels/new", ruc.AddLabel)
	api.POST("/labels/:id", ruc.UpdateLabel)
	api.GET("/labels/:id", ruc.FindLabelByID)
	api.GET("/labels/findbyname/:name", ruc.FindLabelByName)
	api.GET("/labels", ruc.FindAllLabels)
	api.DELETE("/labels/:id", ruc.RemoveLabel)

	api.POST("/projects/new", ruc.AddProject)
	api.POST("/projects/:id", ruc.UpdateProject)
	api.GET("/projects/:id", ruc.FindProjectByID)
	api.GET("/projects", ruc.FindAllProjects)
	api.DELETE("/projects/:id", ruc.RemoveProject)
}
