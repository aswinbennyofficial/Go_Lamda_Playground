package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error){
	log.Println("Hello world")
	
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body: "This is body",
	}

	return response,nil
}

func main(){
	lambda.Start(handler)
}