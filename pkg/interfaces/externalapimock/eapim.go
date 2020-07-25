package externalapimock

import (
	"github.com/labstack/echo/v4"
)

// ExternalAPIMockPath is a path for getting default color
var ExternalAPIMockPath = "/externalapimock/color"

// getColor to get default color
func getColor(c echo.Context) error {
	return c.JSON(200, map[string]interface{}{
		"hexCode": "2979FF",
	})
}

// PrepareEndpoints function to prepare mock endpoints
func PrepareEndpoints(e *echo.Echo) {
	e.GET(ExternalAPIMockPath, getColor)
}
