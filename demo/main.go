package main

import (
	"fmt"
	"log"

	africastkng "github.com/AndroidStudyOpenSource/africastalking-go"
)

const (
	username = "" //Your Africa's Talking Username
	apiKey   = "" //Production or Sandbox API Key
	env      = "" // Choose either Sandbox or Production
)

func main() {
	//Call the Gateway, and pass the constants here!
	gateway, err := africastkng.NewGateway(username, apiKey, env)
	if err != nil {
		log.Fatal(err)
	}

	//Send SMS - REPLACE Recipient and Message with REAL Values
	recipients, err := gateway.SendSms("Recipient", "Message To Send")
	if err != nil {
		fmt.Println(err)
	}

	//For loop to log all the recipients
	for _, recipient := range recipients {
		fmt.Println(recipient)
	}
}
