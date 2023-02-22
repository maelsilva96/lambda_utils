package apigateway

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"io"
	"log"
	"net/http"
)

type LocalHandler struct {
	host    string
	port    int
	handler Handler
}

func NewLocalHandler(handler Handler) *LocalHandler {
	return &LocalHandler{
		host:    "",
		port:    8080,
		handler: handler,
	}
}

func (hand LocalHandler) SetHost(host string) {
	hand.host = host
}

func (hand LocalHandler) SetPort(port int) {
	hand.port = port
}

func (hand LocalHandler) handlerProcess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var requestBody events.APIGatewayProxyRequest
	reqBody, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &requestBody)
	if err != nil {
		log.Println(err)
		return
	}
	result, err := hand.handler.Handler(context.TODO(), requestBody)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(result.StatusCode)
	if len(result.Headers) > 0 {
		for key, val := range result.Headers {
			if w.Header().Get(key) != "" {
				w.Header().Del(key)
			}
			w.Header().Set(key, val)
		}
	}
	if result.Body != "" {
		_, err = w.Write([]byte(result.Body))
		if err != nil {
			log.Println(err)
		}
	}
}

func (hand LocalHandler) Start() {
	http.HandleFunc("/handler", hand.handlerProcess)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", hand.host, hand.port), nil))
}
