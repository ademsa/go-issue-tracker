package http

import (
	"errors"
	"github.com/labstack/echo/v4"
)

// AddLabel to add new label
func (ruc *restUseCase) AddLabel(c echo.Context) error {
	name := c.FormValue("name")
	if name == "" {
		return errors.New("name not provided")
	}

	colorHexCode := c.FormValue("color_hex_code")
	if colorHexCode == "" {
		cl, err := ruc.cuc.GetColor()
		if err != nil {
			return err
		}
		colorHexCode = cl.HexCode
	}

	item, err := ruc.luc.Add(name, colorHexCode)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"item": item,
	})
}

// UpdateLabel to update label
func (ruc *restUseCase) UpdateLabel(c echo.Context) error {
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
		cl, err := ruc.cuc.GetColor()
		if err != nil {
			return err
		}
		colorHexCode = cl.HexCode
	}

	item, err := ruc.luc.Update(id, name, colorHexCode)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"item": item,
	})
}

// FindLabelByID to find label by ID
func (ruc *restUseCase) FindLabelByID(c echo.Context) error {
	id, err := getID(c)
	if err != nil {
		return err
	}

	item, err := ruc.luc.FindByID(uint(id))
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

// FindLabelByName to find label by name
func (ruc *restUseCase) FindLabelByName(c echo.Context) error {
	name := c.Param("name")
	if name == "" {
		return errors.New("label name not provided")
	}

	item, err := ruc.luc.FindByName(name)
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

// FindAllLabels to find all labels
func (ruc *restUseCase) FindAllLabels(c echo.Context) error {
	items, err := ruc.luc.FindAll()
	if err != nil {
		return err
	}
	return c.JSON(200, map[string]interface{}{
		"items": items,
	})
}

// RemoveLabel to remove label
func (ruc *restUseCase) RemoveLabel(c echo.Context) error {
	id, err := getID(c)
	if err != nil {
		return err
	}

	status, err := ruc.luc.Remove(id)
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
