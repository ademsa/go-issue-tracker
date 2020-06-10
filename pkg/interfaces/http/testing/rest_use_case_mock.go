package testing

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

// RESTUseCaseMock is a mock of RESTUseCase
type RESTUseCaseMock struct {
	mock.Mock
}

// AddIssue mock
func (m *RESTUseCaseMock) AddIssue(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// UpdateIssue mock
func (m *RESTUseCaseMock) UpdateIssue(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// FindIssueByID mock
func (m *RESTUseCaseMock) FindIssueByID(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// FindAllIssues mock
func (m *RESTUseCaseMock) FindAllIssues(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// RemoveIssue mock
func (m *RESTUseCaseMock) RemoveIssue(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// AddLabel mock
func (m *RESTUseCaseMock) AddLabel(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// UpdateLabel mock
func (m *RESTUseCaseMock) UpdateLabel(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// FindLabelByID mock
func (m *RESTUseCaseMock) FindLabelByID(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// FindLabelByName mock
func (m *RESTUseCaseMock) FindLabelByName(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// FindAllLabels mock
func (m *RESTUseCaseMock) FindAllLabels(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// RemoveLabel mock
func (m *RESTUseCaseMock) RemoveLabel(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// AddProject mock
func (m *RESTUseCaseMock) AddProject(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// UpdateProject mock
func (m *RESTUseCaseMock) UpdateProject(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// FindProjectByID mock
func (m *RESTUseCaseMock) FindProjectByID(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// FindAllProjects mock
func (m *RESTUseCaseMock) FindAllProjects(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// RemoveProject mock
func (m *RESTUseCaseMock) RemoveProject(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}
