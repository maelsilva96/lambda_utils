package apigateway

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"net/http"
)

func CreateResponse(statusCode int, headers map[string]string, body interface{}) events.APIGatewayProxyResponse {
	data, err := json.Marshal(body)
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers:    headers,
			Body:       "{\"success\":false,\"message\":\"Internal error!\"}",
		}
	}
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    headers,
		Body:       string(data),
	}
}
