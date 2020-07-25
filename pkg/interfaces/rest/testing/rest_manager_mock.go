package testing

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

// ManagerMock is a mock of Manager
type ManagerMock struct {
	mock.Mock
}

// AddIssue mock
func (m *ManagerMock) AddIssue(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// UpdateIssue mock
func (m *ManagerMock) UpdateIssue(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// FindIssueByID mock
func (m *ManagerMock) FindIssueByID(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// FindIssues mock
func (m *ManagerMock) FindIssues(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// FindAllIssues mock
func (m *ManagerMock) FindAllIssues(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// RemoveIssue mock
func (m *ManagerMock) RemoveIssue(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// AddLabel mock
func (m *ManagerMock) AddLabel(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// UpdateLabel mock
func (m *ManagerMock) UpdateLabel(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// FindLabelByID mock
func (m *ManagerMock) FindLabelByID(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// FindLabels mock
func (m *ManagerMock) FindLabels(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// FindAllLabels mock
func (m *ManagerMock) FindAllLabels(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// RemoveLabel mock
func (m *ManagerMock) RemoveLabel(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// AddProject mock
func (m *ManagerMock) AddProject(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// UpdateProject mock
func (m *ManagerMock) UpdateProject(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// FindProjectByID mock
func (m *ManagerMock) FindProjectByID(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// FindProjects mock
func (m *ManagerMock) FindProjects(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// FindAllProjects mock
func (m *ManagerMock) FindAllProjects(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// RemoveProject mock
func (m *ManagerMock) RemoveProject(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}
