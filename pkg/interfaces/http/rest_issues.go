package http

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"go-issue-tracker/pkg/domain"
	"strconv"
	"strings"
)

// getLabels to get/validate labels from echo.Context
func (ruc *restUseCase) getLabels(c echo.Context) (map[string]domain.Label, error) {
	labelsRaw := strings.Split(strings.Trim(c.FormValue("labels"), " "), ",")
	labels := make(map[string]domain.Label)
	for _, lR := range labelsRaw {
		if labels[lR].ID == 0 && lR != "" {
			label, err := ruc.luc.FindByName(lR)
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
func (ruc *restUseCase) AddIssue(c echo.Context) error {
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
	projectID, err := strconv.Atoi(c.FormValue("project_id"))
	if err != nil {
		return err
	}
	project, err := ruc.puc.FindByID(uint(projectID))
	if err != nil {
		return errors.New("project not found")
	}
	labels, err := ruc.getLabels(c)
	if err != nil {
		return err
	}

	item, err := ruc.iuc.Add(title, description, status, project, labels)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"item": item,
	})
}

// UpdateIssue to update issue
func (ruc *restUseCase) UpdateIssue(c echo.Context) error {
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
	labels, err := ruc.getLabels(c)
	if err != nil {
		return err
	}

	item, err := ruc.iuc.Update(id, title, description, status, labels)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"item": item,
	})
}

// FindIssueByID to find issue by ID
func (ruc *restUseCase) FindIssueByID(c echo.Context) error {
	id, err := getID(c)
	if err != nil {
		return err
	}

	item, err := ruc.iuc.FindByID(id)
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

// FindAllIssues to find all issues
func (ruc *restUseCase) FindAllIssues(c echo.Context) error {
	items, err := ruc.iuc.FindAll()
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"items": items,
	})
}

// RemoveIssue to remove issue
func (ruc *restUseCase) RemoveIssue(c echo.Context) error {
	id, err := getID(c)
	if err != nil {
		return err
	}

	status, err := ruc.iuc.Remove(id)
	if err != nil {
		if err.Error() == "record not found" {
			return c.JSON(200, map[string]interface{}{
				"status": false,
			})
		}
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"item": status,
	})
}
