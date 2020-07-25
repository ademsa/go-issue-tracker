package rest

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"go-issue-tracker/pkg/domain"
	"strconv"
	"strings"
)

// getLabels to get/validate labels from echo.Context
func (m *manager) getLabels(c echo.Context) (map[string]domain.Label, error) {
	labelsRaw := strings.Split(strings.Trim(c.FormValue("labels"), " "), ",")
	labels := make(map[string]domain.Label)
	for _, lR := range labelsRaw {
		if labels[lR].ID == 0 && lR != "" {
			label, err := m.luc.FindByName(lR)
			if err != nil {
				return labels, fmt.Errorf("label %s is not valid", lR)
			}
			labels[lR] = label
		}
	}
	if len(labels) == 0 {
		return labels, errors.New("no labels assigned")
	}

	return labels, nil
}

// AddIssue to add new issue
func (m *manager) AddIssue(c echo.Context) error {
	title := c.FormValue("title")
	if title == "" {
		return errors.New("title not provided")
	}
	description := c.FormValue("description")
	if description == "" {
		return errors.New("description not provided")
	}
	status, err := strconv.Atoi(c.FormValue("status"))
	if err != nil {
		return err
	}
	projectID, err := strconv.Atoi(c.FormValue("projectId"))
	if err != nil {
		return err
	}
	project, err := m.puc.FindByID(uint(projectID))
	if err != nil {
		return errors.New("project not found")
	}
	labels, err := m.getLabels(c)
	if err != nil {
		return err
	}

	item, err := m.iuc.Add(title, description, status, project, labels)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"item": item,
	})
}

// UpdateIssue to update issue
func (m *manager) UpdateIssue(c echo.Context) error {
	id, err := getID(c)
	if err != nil {
		return err
	}

	title := c.FormValue("title")
	if title == "" {
		return errors.New("title not provided")
	}
	description := c.FormValue("description")
	if description == "" {
		return errors.New("description not provided")
	}
	status, err := strconv.Atoi(c.FormValue("status"))
	if err != nil {
		return err
	}
	labels, err := m.getLabels(c)
	if err != nil {
		return err
	}

	item, err := m.iuc.Update(id, title, description, status, labels)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"item": item,
	})
}

// FindIssueByID to find issue by ID
func (m *manager) FindIssueByID(c echo.Context) error {
	id, err := getID(c)
	if err != nil {
		return err
	}

	item, err := m.iuc.FindByID(id)
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

// FindIssues to find issues
func (m *manager) FindIssues(c echo.Context) error {
	title := c.QueryParam("title")
	projectID, err := strconv.Atoi(c.QueryParam("projectId"))
	if err != nil {
		return err
	}
	labelsRaw := strings.Split(strings.Trim(c.QueryParam("labels"), " "), ",")
	labels := []string{}
	for _, lR := range labelsRaw {
		if lR != "" {
			labels = append(labels, lR)
		}
	}

	items, err := m.iuc.Find(title, uint(projectID), labels)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"items": items,
	})
}

// FindAllIssues to find all issues
func (m *manager) FindAllIssues(c echo.Context) error {
	items, err := m.iuc.FindAll()
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"items": items,
	})
}

// RemoveIssue to remove issue
func (m *manager) RemoveIssue(c echo.Context) error {
	id, err := getID(c)
	if err != nil {
		return err
	}

	status, err := m.iuc.Remove(id)
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
