package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"what_is_your_name"`
	Age  int    `json:"how_old_are_you"`
}

type MyResponse struct {
	Message string `json:"answer"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse{
		Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age),
	}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
