# Unofficial Africa's Talking Golang API Wrapper
The wrapper provides convenient access to the Africa's Talking API from applications written in server-side Golang. 

**This is Work in Progress!**

## Installing
You can install the package by running:

```
go get github.com/AndroidStudyOpenSource/africastalking-go
```

## Usage
The package needs to be configured with your app **username** and **API key** (which you can get from the dashboard). You can also declare if you are running in production or in sandbox.

```
const (
	apiKey = "YOUR_API_KEY"		    //Production or Sandbox API Key
	username = "YOUR_USERNAME"	    //Your Africa's Talking Username
	env = ""		                // Choose either Sandbox or Production
)
```

## Creating the Gateway
We first need to create a Gateway using the constants declared above. We will use this Gateway to invoke Africa's Talking Services - SMS, Voice, Airtime, USSD.

This is how we create the Gateway in our code:

```
gateway, err := africastkng.NewGateway(username, apiKey, env)
if err != nil {
	log.Fatal(err)
}
```

## SMS 
When sending a message, you need to pass the following data:
* **Recipient(s)** 
* **Message** 

We invoke this function using the following code -  You can declare recipient and message as variables for code neatness:

```
//Send SMS - REPLACE Recipient and Message with REAL Values
recipients, err := gateway.SendSms("Recipient", "Message To Send")
if err != nil {
	fmt.Println(err)
}

//For loop to log all the recipients
for _, recipient := range recipients {
	fmt.Println(recipient)
}
```

This is the complete sample code. Try to understand how this works first!!!
```
package main

import (
	"fmt"
	"log"

	africastkng "github.com/AndroidStudyOpenSource/africastalking-go"
)

const (
	username = "" //Your Africa's Talking Username
	apiKey   = "" //Production or Sandbox API Key
	env   = "" // Choose either Sandbox or Production
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

```
You can easily test this using [Postman](https://www.getpostman.com) or [Insomnia](https://insomnia.rest) Clients!

## Contributing and Issues

Please feel free to contribute or open issues, if any and we will be happy to help out!


