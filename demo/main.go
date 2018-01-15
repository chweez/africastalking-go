package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"africastalking/sms"
)

const (
	apiKey   = ""
	username = ""
)

func main() {
	recipient := flag.String("r", "", "The phone number of the recipient of the message")
	message := flag.String("m", "", "The message to be sent")
	env := flag.String("e", "production", "The environment of the api")

	flag.Parse()
	if *recipient == "" || *message == "" {
		log.Println("please enter recipient and message")
		os.Exit(1)
	}

	smsService := sms.NewService(username, apiKey, *env)
	// Entered at the commandline
	sendResponse, err := smsService.Send("Me4u", *recipient, *message)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(sendResponse)
}
