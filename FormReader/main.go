package main

import (
	//"encoding/json"
	"log"
	"net/http"
	//"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type ContactForm struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	

	// Parse form values
    name := request.QueryStringParameters["name"]
    email := request.QueryStringParameters["email"]
    message := request.QueryStringParameters["message"]

	// Process the form data
	log.Printf("Received form data - Name: %s, Email: %s, Message: %s", name, email, message)

	// Perform any additional processing or send emails, etc.

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "Form data received successfully",
	}, nil
}

func main() {
	lambda.Start(handler)
}
