package main

import(
	
	"log"
	"os"
	"net/smtp"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(){

	// SMTP server Credentials from .env file
	SMTP_USERNAME := os.Getenv("SMTP_USERNAME")
	SMTP_PASSWORD := os.Getenv("SMTP_PASSWORD")
	SMTP_HOST :=os.Getenv("SMTP_HOST")
	FROM_EMAIL :=os.Getenv("FROM_EMAIL")
	SMTP_PORT :=os.Getenv("SMTP_PORT")
	TO_EMAIL :=os.Getenv("TO_EMAIL")
	
	log.Printf("SMTP username:%s  SMTP Password:%s  SMTP HOST:%s  From Address:%s  To address: %s  SMTP PORT:%s",SMTP_USERNAME,SMTP_PASSWORD,SMTP_HOST,FROM_EMAIL,TO_EMAIL,SMTP_PORT)
	
	// Setup authentication variable
	auth:=smtp.PlainAuth("",SMTP_USERNAME,SMTP_PASSWORD,SMTP_HOST)


	// List of emails you want to send the email
	// toList := []string{"email1@gmail.com","email2@gmail.com","email3@gmail.com"}
	toList := []string{TO_EMAIL}
	


	// mail
	subject:="Hello guys"
	body:="This is body"
	reply_to:=""

	var msg []byte
	msg = []byte(
		"Reply-To: "+reply_to+"\r\n"+
		"Subject: "+subject+"\r\n" +
		"\r\n" +
		body+"\r\n")

	
	// send the mail
	err := smtp.SendMail(SMTP_HOST+":"+SMTP_PORT, auth, FROM_EMAIL, toList, msg)

	// handling the errors
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	
	log.Println("Successfully sent mail to all user in toList")
	return
}

func main(){

	lambda.Start(handler)

}