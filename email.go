package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendEmail(to, subject, body string) {
	// set up the email message to contain the subject header and the body
	message := fmt.Sprintf("From:%s\nTo:%s\nSubject:%s\r\n%s",
		config.EmailAddress, to, subject, body)

	// Set up authentication information.
	auth := smtp.PlainAuth("", config.EmailAddress, config.EmailPassword, config.EmailHost)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		fmt.Sprintf("%s:%s", config.EmailHost, config.EmailPort),
		auth,
		config.EmailAddress,
		[]string{to},
		[]byte(message),
	)
	if err != nil {
		log.Printf("[email.SendEmail] error sending email: ", err)
	}
}

func FormatEmail(sms string) (subject, body string) {
	if config.SmsNotebook != "" {
		subject = fmt.Sprintf("@%s", config.SmsNotebook)
	} else {
		subject = sms
	}
	body = sms
	return
}
