// main.go
package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

func Funcy(r events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	fmt.Println(r.Body)

	// responseBody, _ := json.Marshal("is all good")
	responseBody := "is it"
	// respHeader := map[string]string{"Content-Type": "application/json"}
	response := events.LambdaFunctionURLResponse{
		StatusCode: http.StatusOK,
		// Headers:    respHeader,
		Body: string(responseBody),
	}
	return response, nil
}

func main() {
	lambda.Start(Funcy)
}
