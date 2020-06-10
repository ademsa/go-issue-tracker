package http

import (
	"github.com/labstack/echo/v4"
	"go-issue-tracker/pkg/usecases"
)

// RESTUseCase interface
type RESTUseCase interface {
	AddIssue(c echo.Context) error
	UpdateIssue(c echo.Context) error
	FindIssueByID(c echo.Context) error
	FindAllIssues(c echo.Context) error
	RemoveIssue(c echo.Context) error
	AddLabel(c echo.Context) error
	UpdateLabel(c echo.Context) error
	FindLabelByID(c echo.Context) error
	FindLabelByName(c echo.Context) error
	FindAllLabels(c echo.Context) error
	RemoveLabel(c echo.Context) error
	AddProject(c echo.Context) error
	UpdateProject(c echo.Context) error
	FindProjectByID(c echo.Context) error
	FindAllProjects(c echo.Context) error
	RemoveProject(c echo.Context) error
}

// restUseCase contains use cases
type restUseCase struct {
	iuc usecases.IssueUseCase
	luc usecases.LabelUseCase
	puc usecases.ProjectUseCase
	cuc usecases.ColorUseCase
}

// NewRESTUseCase to init RESTUseCase
func NewRESTUseCase(iuc usecases.IssueUseCase, luc usecases.LabelUseCase, puc usecases.ProjectUseCase, cuc usecases.ColorUseCase) RESTUseCase {
	return &restUseCase{
		iuc: iuc,
		luc: luc,
		puc: puc,
		cuc: cuc,
	}
}
