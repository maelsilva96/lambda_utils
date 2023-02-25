package tests

import (
	"github.com/maelsilva96/lambda_utils/apigateway"
	"github.com/maelsilva96/lambda_utils/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateBodyIsEmpty(t *testing.T) {
	response := apigateway.NewResponse("POST")
	result := response.CreateResponse(201, nil)
	assert.Equal(t, "", result.Body)
	assert.Equal(t, 201, result.StatusCode)
}

func TestCreateBodyNotEmpty(t *testing.T) {
	response := apigateway.NewResponse("POST")
	result := response.CreateResponse(200, &entities.ResponseBody{
		Success: true,
	})
	assert.NotEqual(t, "", result.Body)
	assert.Equal(t, "{\"success\":true,\"error_code\":0,\"message\":\"\",\"data\":null,\"field_error\":null}", result.Body)
	assert.Equal(t, 200, result.StatusCode)
}
