package main

import (
	"log"
	"net/http"
)

func main() {
	SendEmail(EmailRecipient, "testing 123!")
	http.HandleFunc("/sms", smsHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("[main] ListenAndServe error: ", err)
	}
}

func smsHandler(response http.ResponseWriter, request *http.Request) {
	sms := ReceiveSms(request)
	log.Printf("[smsHandler] incoming SMS: %s", sms)
	SendEmail(EmailRecipient, sms.Body)
}
