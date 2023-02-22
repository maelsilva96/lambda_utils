package apigateway

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/maelsilva96/lambda_utils/entities"
	"log"
	"net/http"
)

func CreateResponse(statusCode int, headers map[string]string, body entities.ResponseBody) events.APIGatewayProxyResponse {
	data, err := json.Marshal(body)
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers:    headers,
			Body:       "{\"success\":false,\"message\":\"Internal error!\",\"error_code\":1}",
		}
	}
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    headers,
		Body:       string(data),
	}
}
