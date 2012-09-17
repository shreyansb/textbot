package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var config Config

type Config struct {
	TwilioAccountSID, TwilioAuthToken, TwilioNumber                   string
	EmailRecipient, EmailAddress, EmailPassword, EmailHost, EmailPort string
	SmsNotebook                                                       string
}

func main() {
	parseConfig()

	http.HandleFunc("/sms", smsHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("[main] ListenAndServe error: ", err)
	}
}

func smsHandler(response http.ResponseWriter, request *http.Request) {
	sms := ReceiveSms(request)
	log.Printf("[smsHandler] incoming SMS: %s", sms)
	subject, body := FormatEmail(sms.Body)
	SendEmail(config.EmailRecipient, subject, body)
}

func parseConfig() {
	path := "./textbot.conf"
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("[parseConfig] error reading config file: ", err)
	}

	err = json.Unmarshal(b, &config)
	if err != nil {
		log.Fatal("[parseConfig] error parsing config json: ", err)
	}
}
