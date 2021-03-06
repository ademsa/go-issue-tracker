package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-issue-tracker/pkg/infrastructure/helpers"
	"path/filepath"
)

// RootMessage is displayed in / path
var RootMessage = "Issue Tracker"

// APIRootMessage is displayed in /api path
var APIRootMessage = "Issue Tracker Api"

// PrepareServer function to prepare HTTP Server and REST Use Case
func PrepareServer() *echo.Echo {
	server := echo.New()

	server.HideBanner = true

	server.Use(middleware.Logger())

	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:3001",
			"http://127.0.0.1:3000", "http://127.0.0.1:3001",
			"http://0.0.0.0:3000", "http://0.0.0.0:3001"},
	}))

	return server
}

// PrepareEndpoints function to prepare endpoints
func PrepareEndpoints(e *echo.Echo, m Manager, uiDirPath string) {
	if helpers.PathExists(uiDirPath) {
		e.Static("/static", filepath.Join(uiDirPath, "static"))
		e.File("/favicon.ico", filepath.Join(uiDirPath, "favicon.ico"))
		e.File("/logo192x192.png", filepath.Join(uiDirPath, "logo192x192.png"))
		e.File("/logo512x512.png", filepath.Join(uiDirPath, "logo512x512.png"))
		e.File("/robots.txt", filepath.Join(uiDirPath, "robots.txt"))
		e.File("/manifest.json", filepath.Join(uiDirPath, "manifest.json"))
		e.File("/*", filepath.Join(uiDirPath, "index.html"))
	} else {
		e.GET("/", func(c echo.Context) error {
			return c.String(200, RootMessage)
		})
	}

	e.GET("/api", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"message": APIRootMessage,
		})
	})

	api := e.Group("/api")

	api.POST("/issues/new", m.AddIssue)
	api.POST("/issues/:id", m.UpdateIssue)
	api.GET("/issues/:id", m.FindIssueByID)
	api.GET("/issues/find", m.FindIssues)
	api.GET("/issues", m.FindAllIssues)
	api.DELETE("/issues/:id", m.RemoveIssue)

	api.POST("/labels/new", m.AddLabel)
	api.POST("/labels/:id", m.UpdateLabel)
	api.GET("/labels/:id", m.FindLabelByID)
	api.GET("/labels/find", m.FindLabels)
	api.GET("/labels", m.FindAllLabels)
	api.DELETE("/labels/:id", m.RemoveLabel)

	api.POST("/projects/new", m.AddProject)
	api.POST("/projects/:id", m.UpdateProject)
	api.GET("/projects/:id", m.FindProjectByID)
	api.GET("/projects/find", m.FindProjects)
	api.GET("/projects", m.FindAllProjects)
	api.DELETE("/projects/:id", m.RemoveProject)
}
