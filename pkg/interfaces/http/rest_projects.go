package http

import (
	"errors"
	"github.com/labstack/echo/v4"
)

// RemoveLabel to add new project
func (ruc *restUseCase) AddProject(c echo.Context) error {
	name := c.FormValue("name")
	if name == "" {
		return errors.New("name not provided")
	}
	description := c.FormValue("description")

	item, err := ruc.puc.Add(name, description)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"item": item,
	})
}

// UpdateProject to update project
func (ruc *restUseCase) UpdateProject(c echo.Context) error {
	id, err := getID(c)
	if err != nil {
		return err
	}

	name := c.FormValue("name")
	if name == "" {
		return errors.New("name not provided")
	}
	description := c.FormValue("description")

	item, err := ruc.puc.Update(id, name, description)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"item": item,
	})
}

// FindProjectByID to find project by ID
func (ruc *restUseCase) FindProjectByID(c echo.Context) error {
	id, err := getID(c)
	if err != nil {
		return err
	}

	item, err := ruc.puc.FindByID(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			return c.JSON(200, map[string]interface{}{
				"item": nil,
			})
		}
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"item": item,
	})
}

// FindAllProjects to find all projects
func (ruc *restUseCase) FindAllProjects(c echo.Context) error {
	items, err := ruc.puc.FindAll()
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"items": items,
	})
}

// RemoveProject to remove project
func (ruc *restUseCase) RemoveProject(c echo.Context) error {
	id, err := getID(c)
	if err != nil {
		return err
	}

	status, err := ruc.puc.Remove(id)
	if err != nil {
		if err.Error() == "record not found" {
			return c.JSON(200, map[string]interface{}{
				"status": false,
			})
		}
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"status": status,
	})
}
