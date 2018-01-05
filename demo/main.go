package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	africastkng "github.com/AndroidStudyOpenSource/africastalking-go"
)

const (
	apiKey   = "6108b25d63a8fa8db3c5fd40a57df0a4f1b139b718ffd06fa4ebdf416a159bf0"
	username = "Me4u"
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

	//Call the Gateway, and pass the constants here!
	gateway, err := africastkng.NewGateway(username, apiKey, *env)
	if err != nil {
		log.Fatal(err)
	}

	// Entered at the commandline
	recipients, err := gateway.SendSms(*recipient, *message)
	if err != nil {
		fmt.Println(err)
	}

	//For loop to log all the recipients
	for _, recipient := range recipients {
		fmt.Println(recipient)
	}
}
