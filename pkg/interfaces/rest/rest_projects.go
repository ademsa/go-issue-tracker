package rest

import (
	"errors"
	"github.com/labstack/echo/v4"
)

// RemoveLabel to add new project
func (m *manager) AddProject(c echo.Context) error {
	name := c.FormValue("name")
	if name == "" {
		return errors.New("name not provided")
	}
	description := c.FormValue("description")

	item, err := m.puc.Add(name, description)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"item": item,
	})
}

// UpdateProject to update project
func (m *manager) UpdateProject(c echo.Context) error {
	id, err := getID(c)
	if err != nil {
		return err
	}

	name := c.FormValue("name")
	if name == "" {
		return errors.New("name not provided")
	}
	description := c.FormValue("description")

	item, err := m.puc.Update(id, name, description)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"item": item,
	})
}

// FindProjectByID to find project by ID
func (m *manager) FindProjectByID(c echo.Context) error {
	id, err := getID(c)
	if err != nil {
		return err
	}

	item, err := m.puc.FindByID(uint(id))
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

// FindProjects to find labels
func (m *manager) FindProjects(c echo.Context) error {
	name := c.QueryParam("name")
	items, err := m.puc.Find(name)
	if err != nil {
		return err
	}
	return c.JSON(200, map[string]interface{}{
		"items": items,
	})
}

// FindAllProjects to find all projects
func (m *manager) FindAllProjects(c echo.Context) error {
	items, err := m.puc.FindAll()
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"items": items,
	})
}

// RemoveProject to remove project
func (m *manager) RemoveProject(c echo.Context) error {
	id, err := getID(c)
	if err != nil {
		return err
	}

	status, err := m.puc.Remove(id)
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
