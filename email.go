package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendEmail(to, body string) {
	// Set up authentication information.
	auth := smtp.PlainAuth("", EmailAddress, EmailPassword, EmailHost)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		fmt.Sprintf("%s:587", EmailHost),
		auth,
		EmailAddress,
		[]string{to},
		[]byte(body),
	)
	if err != nil {
		log.Printf("[email.SendEmail] error sending email: ", err)
	}
}
