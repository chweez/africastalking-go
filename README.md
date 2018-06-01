# Unofficial Africa's Talking Golang API Wrapper  [![CircleCI](https://circleci.com/gh/AndroidStudyOpenSource/africastalking-go.svg?style=shield)](https://circleci.com/gh/AndroidStudyOpenSource/africastalking-go)
The wrapper provides convenient access to the Africa's Talking API from applications written in server-side Golang. This is Work in Progress

## Installing
You can install any of the packages as shown below:

### Sms
```sh
go get github.com/AndroidStudyOpenSource/africastalking-go/sms
```

### Airtime
```sh
go get github.com/AndroidStudyOpenSource/africastalking-go/airtime
```

### Account
```sh
go get github.com/AndroidStudyOpenSource/africastalking-go/account
```

### Payment
```sh
go get github.com/AndroidStudyOpenSource/africastalking-go/payment
```

### Tokens
```sh
go get github.com/AndroidStudyOpenSource/africastalking-go/token
```

### Demo
In order to run the demo, export the the following values to your environment. They can be found/generated at the Africa's Talking Dashboard.

``` sh
export AT_APIKEY=Your-AfricasTalking-API-KEY
export AT_USERNAME=Your-AfricasTalking-APP-USERNAME
export AT_SHORTCODE=Your-AfricasTalking-APP-Shortcode
```

That's all's that required. Assuming its a sandbox app, here's to run the demo:

``` sh
cd demo
go run main.go -e sandbox -m "Hello gopher!" -r "+254700000000"
```

### Usage
The package needs to be configured with your app username and API key (which you can get from the dashboard). You can also declare if you are running in production or in sandbox.

```golang
const (
	apiKey = "YOUR_API_KEY"		    //Production or Sandbox API Key
	username = "YOUR_USERNAME"	    //Your Africa's Talking Username
	env = ""		                // Choose either Sandbox or Production
)
```

### Creating the Gateway
We first need to create a Gateway using the constants declared above. We will use this Gateway to invoke Africa's Talking Services - SMS, Voice, Airtime, USSD.

This is how we create the Gateway in our code:

```golang
smsService := sms.NewService(username, apiKey, env)
```

### SMS 
When sending a message, you need to pass the following data:
* **Recipient(s)** 
* **Message** 

We invoke this function using the following code -  You can declare recipient and message as variables for code neatness:

```golang
//Send SMS - REPLACE Recipient and Message with REAL Values
smsResponse, err := smsService.Send("Recipient", "Message To Send", "")
if err != nil {
	fmt.Println(err)
}

fmt.Println(smsResponse)
```

This is the complete sample code. Try to understand how this works first!!!
```golang
package main

import (
	"fmt"
	"log"

	"github.com/AndroidStudyOpenSource/africastalking-go/sms"
)

const (
	username = "" //Your Africa's Talking Username
	apiKey   = "" //Production or Sandbox API Key
	env   = "" // Choose either Sandbox or Production
)

func main() {
	//Call the Gateway, and pass the constants here!
	smsService := sms.NewService(username, apiKey, env)

	//Send SMS - REPLACE Recipient and Message with REAL Values
	recipients, err := smsService.Send("Recipient", "Message To Send", "ShortCode" //Leave blank, "", if you don't have one)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(smsResponse)
}

```
You can easily test this using [Postman](https://www.getpostman.com) or [Insomnia](https://insomnia.rest) Clients!

### Contributing and Issues

We’re glad you’re interested in Africas Talking Golang SDK, and we’d love to see where you take it. If you would like to contribute code to this project you can do so through GitHub by Forking the Repository and creating a Pull Request.

When submitting code, please make every effort to follow existing conventions and style in order to keep the code as readable as possible. We look forward to you submitting a Pull Request.

Use [gitflow](https://www.atlassian.com/git/tutorials/comparing-workflows#gitflow-workflow).
Always tag releases to `develop` and `master`.

Thanks, and please do take it for a joyride!

### License

```text
MIT License

Copyright (c) 2018 Android Study Open Source

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```


