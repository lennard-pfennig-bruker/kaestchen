package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

func handler(r events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	fmt.Println(r.Body)
	respHeader := map[string]string{"Content-Type": "application/json"}
	response := events.LambdaFunctionURLResponse{
		StatusCode: http.StatusOK,
		Headers:    respHeader,
		Body:       r.Body,
	}
	return response, nil
}

func main() {
	lambda.Start(handler)
}
