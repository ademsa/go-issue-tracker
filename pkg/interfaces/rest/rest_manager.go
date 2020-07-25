package rest

import (
	"github.com/labstack/echo/v4"
	"go-issue-tracker/pkg/usecases"
)

// Manager interface
type Manager interface {
	AddIssue(c echo.Context) error
	UpdateIssue(c echo.Context) error
	FindIssueByID(c echo.Context) error
	FindIssues(c echo.Context) error
	FindAllIssues(c echo.Context) error
	RemoveIssue(c echo.Context) error
	AddLabel(c echo.Context) error
	UpdateLabel(c echo.Context) error
	FindLabelByID(c echo.Context) error
	FindLabels(c echo.Context) error
	FindAllLabels(c echo.Context) error
	RemoveLabel(c echo.Context) error
	AddProject(c echo.Context) error
	UpdateProject(c echo.Context) error
	FindProjectByID(c echo.Context) error
	FindProjects(c echo.Context) error
	FindAllProjects(c echo.Context) error
	RemoveProject(c echo.Context) error
}

// manager contains use cases
type manager struct {
	iuc usecases.IssueUseCase
	luc usecases.LabelUseCase
	puc usecases.ProjectUseCase
	cuc usecases.ColorUseCase
}

// NewManager to init Manager
func NewManager(iuc usecases.IssueUseCase, luc usecases.LabelUseCase, puc usecases.ProjectUseCase, cuc usecases.ColorUseCase) Manager {
	return &manager{
		iuc: iuc,
		luc: luc,
		puc: puc,
		cuc: cuc,
	}
}
