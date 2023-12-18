package main

import (
	"encoding/json"
	"log"
	"net/smtp"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type ContactForm struct {
	Replyto    *string `json:"from"`
	Subject *string `json:"subject"`
	MsgBody  *string `json:"msgbody"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Check if the request body is empty
	if request.Body == "" {
		log.Println("Error: Empty request body")
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error: Empty request body",
		}, nil
	}

	var formObj ContactForm

	err := json.Unmarshal([]byte(request.Body), &formObj)

	log.Printf("Replyto address:%s  Subject:%s  Msgbody:%s",*formObj.Replyto, *formObj.Subject, *formObj.MsgBody)

	if err != nil {
		log.Println("Error decoding JSON:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error decoding JSON",
		}, nil
	}

	// SMTP server Credentials from .env file
	SMTP_USERNAME := os.Getenv("SMTP_USERNAME")
	SMTP_PASSWORD := os.Getenv("SMTP_PASSWORD")
	SMTP_HOST := os.Getenv("SMTP_HOST")
	FROM_EMAIL := os.Getenv("FROM_EMAIL")
	SMTP_PORT := os.Getenv("SMTP_PORT")
	TO_EMAIL := os.Getenv("TO_EMAIL")

	log.Printf("SMTP username:%s  SMTP Password:%s  SMTP HOST:%s  From Address:%s  To address: %s  SMTP PORT:%s",
		SMTP_USERNAME, SMTP_PASSWORD, SMTP_HOST, FROM_EMAIL, TO_EMAIL, SMTP_PORT)

	// Setup authentication variable
	auth := smtp.PlainAuth("", SMTP_USERNAME, SMTP_PASSWORD, SMTP_HOST)

	// List of emails you want to send the email
	toList := []string{TO_EMAIL}

	// mail
	subject := *formObj.Subject
	body := *formObj.MsgBody
	reply_to := *formObj.Replyto
	

	var msg []byte
	msg = []byte(
		"From: "+FROM_EMAIL+"\r\n"+
		"Reply-To: " + reply_to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n" +
			"\r\n" +
			body + "\r\n")

	// send the mail
	err = smtp.SendMail(SMTP_HOST+":"+SMTP_PORT, auth, FROM_EMAIL, toList, msg)

	// handling the errors
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error sending email",
		}, nil
	}

	log.Println("Successfully sent mail")

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Msg sent successfully",
	}, nil
}

func main() {
	lambda.Start(handler)
}
