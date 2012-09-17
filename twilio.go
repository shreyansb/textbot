package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var (
	sendSMSUrl = fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/SMS/Messages.json",
		config.TwilioAccountSID)
	client *http.Client
)

type TwilioResponse struct {
	Sid          string `json:"sid"`
	Date_created string `json:"date_created"`
	Date_updated string `json:"date_updated"`
	Account_sid  string `json:"account_sid"`
	To           string `json:"to"`
	From         string `json:"from"`
	Body         string `json:"body"`
	Status       string `json:"status"`
	Direction    string `json:"direction"`
	Api_version  string `json:"api_version"`
	Uri          string `json:"uri"`
}

type TwilioSms struct {
	SmsSid, AccountSid, From, To, Body string
}

func init() {
	client = &http.Client{}
}

func SendSms(from, to, message string) (response TwilioResponse, err error) {
	// prepare request values
	values := make(url.Values)
	values.Set("From", from)
	values.Set("To", to)
	values.Set("Body", message)
	reqBody := strings.NewReader(values.Encode())

	// create request object
	req, err := http.NewRequest("POST", sendSMSUrl, reqBody)
	req.SetBasicAuth(config.TwilioAccountSID, config.TwilioAuthToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// send request
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[twilio.sendSms] error sending request: ", err)
		return
	}

	// parse response
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	log.Printf("[twilio.sendSms] response body: %s", body)

	// convert to object
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("[twilio.sendSms] error unmarshaling response: ", err)
		return
	}

	return
}

func ReceiveSms(request *http.Request) TwilioSms {
	request.ParseForm()
	sms := TwilioSms{
		request.FormValue("SmsSid"),
		request.FormValue("AccountSid"),
		request.FormValue("From"),
		request.FormValue("To"),
		request.FormValue("Body"),
	}
	return sms
}
