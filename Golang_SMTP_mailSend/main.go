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

	// SMTP server Credentials
	SMTP_USERNAME := os.Getenv("SMTP_USERNAME")
	SMTP_PASSWORD := os.Getenv("SMTP_PASSWORD")
	SMTP_HOST :=os.Getenv("SMTP_HOST")
	FROM_EMAIL :=os.Getenv("FROM_EMAIL")
	
	log.Println("SMTP CREDS init ",SMTP_USERNAME, " ", SMTP_PASSWORD," ",SMTP_HOST )
	
	// Setup authentication variable
	auth:=smtp.PlainAuth("",SMTP_USERNAME,SMTP_PASSWORD,SMTP_HOST)


	// Setup other variables

	// List of emails you want to send the email
	toList := []string{"ses@aswinbenny.in"}

	

	// Port of SMTP server
	port := "587"

	// msg
	// msg := []byte("To: recipient@example.net\r\n" +
	// 	"Subject: discount Gophers!\r\n" +
	// 	"\r\n" +
	// 	"This is the email body.\r\n")

	msg := []byte(
		"Subject: Hello Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")

		err = smtp.SendMail(SMTP_HOST+":"+port, auth, FROM_EMAIL, toList, msg)
 
		// handling the errors
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	 
		fmt.Println("Successfully sent mail to all user in toList")




}