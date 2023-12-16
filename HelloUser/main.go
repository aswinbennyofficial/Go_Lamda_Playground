package main

import (
	"encoding/json"
	"fmt"
	//"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

)

type Person struct{
	FirstName *string `json:"firstName"`
	LastName *string `json:"lastName"`
}

type ResponseBody struct{
	Msg *string `json:"message"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error){
	var person1 Person
	
	err:=json.Unmarshal([]byte(request.Body),&person1)

	if err!=nil{
		return events.APIGatewayProxyResponse{},err
	}

	var msg string
	if person1.FirstName != nil && person1.LastName != nil {
		msg = fmt.Sprintf("Hello %v %v", *person1.FirstName, *person1.LastName)
	} else {
		msg = "Hello, unnamed person"
	}

	Responsebody1:= ResponseBody{
		Msg: &msg,
	}

	responsebytes,err:=json.Marshal(Responsebody1)
	if err!=nil{
		return events.APIGatewayProxyResponse{},err
	}


	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body: string(responsebytes),
	}

	return response,nil
}

func main(){
	lambda.Start(handler)

}