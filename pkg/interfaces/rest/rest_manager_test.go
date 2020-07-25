package rest_test

import (
	"github.com/stretchr/testify/assert"
	"go-issue-tracker/pkg/interfaces/rest"
	ucTesting "go-issue-tracker/pkg/usecases/testing"
	"testing"
)

func TestNewManager(t *testing.T) {
	cucm := new(ucTesting.ColorUseCaseMock)
	iucm := new(ucTesting.IssueUseCaseMock)
	lucm := new(ucTesting.LabelUseCaseMock)
	pucm := new(ucTesting.ProjectUseCaseMock)

	m := rest.NewManager(iucm, lucm, pucm, cucm)

	assert.NotNil(t, m)
}
