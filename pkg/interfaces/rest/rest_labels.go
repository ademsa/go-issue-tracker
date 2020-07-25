package rest

import (
	"errors"
	"github.com/labstack/echo/v4"
)

// AddLabel to add new label
func (m *manager) AddLabel(c echo.Context) error {
	name := c.FormValue("name")
	if name == "" {
		return errors.New("name not provided")
	}

	colorHexCode := c.FormValue("color_hex_code")
	if colorHexCode == "" {
		cl, err := m.cuc.GetColor()
		if err != nil {
			return err
		}
		colorHexCode = cl.HexCode
	}

	item, err := m.luc.Add(name, colorHexCode)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"item": item,
	})
}

// UpdateLabel to update label
func (m *manager) UpdateLabel(c echo.Context) error {
	id, err := getID(c)
	if err != nil {
		return err
	}

	name := c.FormValue("name")
	if name == "" {
		return errors.New("name not provided")
	}

	colorHexCode := c.FormValue("color_hex_code")
	if colorHexCode == "" {
		cl, err := m.cuc.GetColor()
		if err != nil {
			return err
		}
		colorHexCode = cl.HexCode
	}

	item, err := m.luc.Update(id, name, colorHexCode)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"item": item,
	})
}

// FindLabelByID to find label by ID
func (m *manager) FindLabelByID(c echo.Context) error {
	id, err := getID(c)
	if err != nil {
		return err
	}

	item, err := m.luc.FindByID(uint(id))
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

// FindLabels to find labels
func (m *manager) FindLabels(c echo.Context) error {
	name := c.QueryParam("name")
	items, err := m.luc.Find(name)
	if err != nil {
		return err
	}
	return c.JSON(200, map[string]interface{}{
		"items": items,
	})
}

// FindAllLabels to find all labels
func (m *manager) FindAllLabels(c echo.Context) error {
	items, err := m.luc.FindAll()
	if err != nil {
		return err
	}
	return c.JSON(200, map[string]interface{}{
		"items": items,
	})
}

// RemoveLabel to remove label
func (m *manager) RemoveLabel(c echo.Context) error {
	id, err := getID(c)
	if err != nil {
		return err
	}

	status, err := m.luc.Remove(id)
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
