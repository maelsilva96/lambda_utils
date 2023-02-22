package apigateway

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
)

type Handler interface {
	Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}
