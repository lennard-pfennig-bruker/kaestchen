// main.go
package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Funcy(r events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	fmt.Println(r.Body)
	response := events.LambdaFunctionURLResponse{
		StatusCode:      0,
		Headers:         nil,
		Body:            "is it?",
		IsBase64Encoded: false,
		Cookies:         nil,
	}
	return response, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(Funcy)
}
