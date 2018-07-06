package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// Response contains the repsonse
type Response struct {
	Message string `json:"message"`
}

// Request contains the request
type Request struct {
	ID int `json:"id"`
}

func handler(r Request) Response {
	return Response{
		Message: fmt.Sprint("Responding to request no.", r.ID),
	}
}

func main() {
	lambda.Start(handler)
}
