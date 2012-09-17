package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendEmail(to, body string) {
	// Set up authentication information.
	auth := smtp.PlainAuth("", config.EmailAddress, config.EmailPassword, config.EmailHost)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		fmt.Sprintf("%s:%s", config.EmailHost, config.EmailPort),
		auth,
		config.EmailAddress,
		[]string{to},
		[]byte(body),
	)
	if err != nil {
		log.Printf("[email.SendEmail] error sending email: ", err)
	}
}

func FormatEmail(body string) {
	return fmt.Sprintf("%s @%s", body, config.SmsNotebook)
}
