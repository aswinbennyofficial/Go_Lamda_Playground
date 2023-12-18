package main

import(
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"net/smtp"
)

func main(){

	// Loading environment variables
	err := godotenv.Load(".env")
	if err != nil {
        log.Fatalf("Error loading environment variables file")
    }

	// SMTP server Credentials from .env file
	SMTP_USERNAME := os.Getenv("SMTP_USERNAME")
	SMTP_PASSWORD := os.Getenv("SMTP_PASSWORD")
	SMTP_HOST :=os.Getenv("SMTP_HOST")
	FROM_EMAIL :=os.Getenv("FROM_EMAIL")
	SMTP_PORT :=os.Getenv("SMTP_PORT")
	
	log.Println("SMTP CREDS init ",SMTP_USERNAME, " ", SMTP_PASSWORD," ",SMTP_HOST )
	
	// Setup authentication variable
	auth:=smtp.PlainAuth("",SMTP_USERNAME,SMTP_PASSWORD,SMTP_HOST)


	// List of emails you want to send the email
	// toList := []string{"email1@gmail.com","email2@gmail.com","email3@gmail.com"}
	toList := []string{"recipient1@gmail.com"}


	// mail
	subject:="Test Golang Program"
	body:="<html><body><h1>Hello, this is an HTML-rich email template!</h1></body></html>"
	reply_to:=""

	if reply_to==""{
		reply_to=FROM_EMAIL
	}

	var msg []byte
	//For basic text
	// msg = []byte(
	// 	"Reply-To: "+reply_to+"\r\n"+
	// 	"Subject: "+subject+"\r\n" +
	// 	"\r\n" +
	// 	body+"\r\n")

	//For rich html support
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
		os.Exit(1)
	}
	
	fmt.Println("Successfully sent mail to all user in toList")

}