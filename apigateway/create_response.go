package apigateway

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/maelsilva96/lambda_utils/entities"
	"log"
	"net/http"
)

type Response struct {
	headers map[string]string
}

func NewResponse(httpMethod string) *Response {
	return &Response{
		headers: map[string]string{
			"Strict-Transport-Security":      "max-age=31536000; includeSubDomains; preload",
			"X-Content-Type-Options":         "nosniff",
			"X-Frame-Options":                "SAMEORIGIN",
			"Access-Control-Allow-Origin":    "*",
			"Access-Control-Request-Method:": httpMethod,
		},
	}
}

func (resp *Response) AddHeader(key string, val string) {
	resp.headers[key] = val
}

func (resp *Response) RemoveHeader(key string) {
	delete(resp.headers, key)
}

func (resp *Response) CreateResponse(statusCode int, body entities.ResponseBody) events.APIGatewayProxyResponse {
	data, err := json.Marshal(body)
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers:    resp.headers,
			Body:       "{\"success\":false,\"message\":\"Internal error!\",\"error_code\":1}",
		}
	}
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    resp.headers,
		Body:       string(data),
	}
}
