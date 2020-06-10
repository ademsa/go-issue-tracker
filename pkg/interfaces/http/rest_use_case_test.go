package http_test

import (
	"github.com/stretchr/testify/assert"
	"go-issue-tracker/pkg/interfaces/http"
	ucTesting "go-issue-tracker/pkg/usecases/testing"
	"testing"
)

func TestNewRESTUseCase(t *testing.T) {
	cucm := new(ucTesting.ColorUseCaseMock)
	iucm := new(ucTesting.IssueUseCaseMock)
	lucm := new(ucTesting.LabelUseCaseMock)
	pucm := new(ucTesting.ProjectUseCaseMock)

	ruc := http.NewRESTUseCase(iucm, lucm, pucm, cucm)

	assert.NotNil(t, ruc)
}
