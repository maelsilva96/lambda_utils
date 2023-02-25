package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/maelsilva96/lambda_utils/apigateway"
	"github.com/maelsilva96/lambda_utils/entities"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type fakerHandler struct {
}

func (hand *fakerHandler) Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response := apigateway.NewResponse("POST")
	return response.CreateResponse(200, &entities.ResponseBody{
		Data: map[string]string{
			"key": "value",
		},
	}), nil
}

func TestHandlerRequest(t *testing.T) {
	values := map[string]string{"body": "{\"key_in\":\"value_in\"}"}
	jsonData, err := json.Marshal(values)

	req, err := http.NewRequest("POST", "/handler", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := &fakerHandler{}
	localHandler := apigateway.NewLocalHandler(handler)
	handlerTest := http.HandlerFunc(localHandler.HandlerProcess)
	handlerTest.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	respBody, _ := io.ReadAll(rr.Body)
	assert.Equal(t, "{\"success\":false,\"error_code\":0,\"message\":\"\",\"data\":{\"key\":\"value\"},\"field_error\":null}", string(respBody))

	log.Println(rr.Header())
	assert.Equal(t, 6, len(rr.Header()))
}

type fakerHandlerHeader struct {
}

func (hand *fakerHandlerHeader) Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response := apigateway.NewResponse("POST")
	response.AddHeader("Test", "ok")
	return response.CreateResponse(200, &entities.ResponseBody{
		Data: map[string]string{
			"key": "value",
		},
	}), nil
}

func TestHandlerAddHeaderResponse(t *testing.T) {
	values := map[string]string{"body": "{\"key_in\":\"value_in\"}"}
	jsonData, err := json.Marshal(values)

	req, err := http.NewRequest("POST", "/handler", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := &fakerHandlerHeader{}
	localHandler := apigateway.NewLocalHandler(handler)
	handlerTest := http.HandlerFunc(localHandler.HandlerProcess)
	handlerTest.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
	headers := rr.Header()
	assert.Equal(t, 7, len(headers))
	if item, ok := headers["Test"]; ok {
		assert.Equal(t, "ok", item[0])
	}
}
