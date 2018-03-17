package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response unmarshalls the json response from the lambda event
type Response struct {
	Message string `json:"message"`
}

// Handler implements the lambda handler
func Handler(request events.APIGatewayProxyRequest) (Response, error) {
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)
	return Response{
		Message: "Go Serverless v1.0! Your function executed successfully!",
	}, nil
}

func main() {
	lambda.Start(Handler)
}
