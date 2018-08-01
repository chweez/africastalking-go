package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/AndroidStudyOpenSource/africastalking-go/account"
	"github.com/AndroidStudyOpenSource/africastalking-go/sms"
)

var (
	apiKey    = os.Getenv("AT_APIKEY")
	username  = os.Getenv("AT_USERNAME")
	shortcode = os.Getenv("AT_SHORTCODE")
)

func main() {
	recipient := flag.String("r", "", "The phone number of the recipient of the message")
	message := flag.String("m", "", "The message to be sent")
	env := flag.String("e", "production", "The environment of the api")

	flag.Parse()
	if *recipient == "" || *message == "" {
		log.Fatal("please enter all required arguments. see --help")
	}

	if apiKey == "" || username == "" {
		log.Fatal("missing required environment variables: AT_APIKEY, AT_USERNAME")
	}

	smsService := sms.NewService(username, apiKey, *env)

	sendResponse, err := smsService.Send(shortcode, *recipient, *message)
	if err != nil {
		fmt.Printf("Failed to send sms: %v", err)
	}
	fmt.Printf("SMS Send reponse: %v\n", sendResponse)

	accountService := account.NewService(username, apiKey, *env)
	user, err := accountService.GetUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("User: %v\n", user)

}
