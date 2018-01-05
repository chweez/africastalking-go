# Africa's Talking Golang API Wrapper
The wrapper provides convenient access to the Africa's Talking API from applications written in server-side Golang.

## Installing
You can install the package by running:

```
go get github.com/AndroidStudyOpenSource/africastalking-go
```

## Usage
The package needs to be configured with your app username and API key (which you can get from the dashboard). You can also declare if you are running in production or in sandbox.

```
const (
	apiKey = ""		//Production or Sandbox API Key
	username = ""	    //Your Africa's Talking Username
	option = ""		// Choose either Sandbox or Production
)
```

## Creating the Gateway
We first need to create a Gateway using the constants declared above. We will use this Gateway to invoke Africa's Talking Services - SMS, Voice, Airtime, USSD.

This is how we create the Gateway in our code:

```
gateway, err := sms.NewGateway(username, apiKey, option)
if err != nil {
	log.Fatal(err)
}
```

## SMS 
When sending a message, you need to pass the following data:
* **Recipient(s)** 
* **Message** 

We invoke this function using the following code -  You can declare recipient and message as variables for code neatness:



## Contributing and Issues

Please feel free to contribute or open issues, if any and we will be happy to help out!


