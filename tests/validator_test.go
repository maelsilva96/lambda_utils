package tests

import (
	"github.com/maelsilva96/lambda_utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type modelTest struct {
	Name   string   `json:"name" name:"Nome" validate:"required,gte=8,lte=120"`
	Fields []string `json:"fields" name:"Campos" validate:"max=1"`
}

func TestValidField(t *testing.T) {
	model := modelTest{
		Name: "test",
	}
	fieldErrors := lambda_utils.Valid(model)
	assert.Len(t, fieldErrors, 1)
	assert.Equal(t, "name", fieldErrors[0].Field)
}
