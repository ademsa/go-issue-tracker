package rest

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

func getID(c echo.Context) (uint, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return uint(id), err
	}
	return uint(id), nil
}
