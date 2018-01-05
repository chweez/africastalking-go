package africastalking

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var responseCode int

const debug = false
const contentJSON = "application/json"

// Gateway is a Gateway
type Gateway struct {
	username    string
	apiKey      string
	environment string
}

// SMS is an sms
type SMS struct {
	username   string
	recipients string
	message    string
}

// SMSResponse is a response
type SMSResponse struct {
	SMSMessageData SMSMessageData
}

// SMSMessageData is a SMSMessageData
type SMSMessageData struct {
	Recipients []Recipient
	Message    string
}

// Recipient is a recipient
type Recipient struct {
	Number string
	Cost   string
	Status string
}

// NewGateway creates a new instance of Gateway and return it or an error
func NewGateway(username, apiKey, environment string) (*Gateway, error) {
	return &Gateway{
		username,
		apiKey,
		environment,
	}, nil
}

// SendSms sends an sms
func (gateway Gateway) SendSms(recipients, message string) ([]Recipient, error) {
	return gateway.sendSms(recipients, message)
}

func (gateway Gateway) sendSms(recipients, messsage string) ([]Recipient, error) {
	data := url.Values{}
	data.Set("username", gateway.username)
	data.Set("to", recipients)
	data.Set("message", messsage)
	body := strings.NewReader(data.Encode())

	client := &http.Client{}
	r, err := http.NewRequest("POST", gateway.getSmsURL(), body)
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Content-Length", strconv.Itoa(body.Len()))
	r.Header.Add("apikey", gateway.apiKey)

	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	var smsResponse SMSResponse
	json.NewDecoder(response.Body).Decode(&smsResponse)
	defer response.Body.Close()

	if len(smsResponse.SMSMessageData.Recipients) > 0 {
		return smsResponse.SMSMessageData.Recipients, nil
	}

	return nil, fmt.Errorf("could not send sms message: %s", smsResponse.SMSMessageData.Message)
}

func (gateway Gateway) getAPIHost() string {
	if gateway.environment == "sandbox" {
		return "https://api.sandbox.africastalking.com"
	}

	return "https://api.africastalking.com"
}

func (gateway Gateway) getSmsURL() string {
	return gateway.getAPIHost() + "/version1/messaging"
}
