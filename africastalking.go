package africastalking

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var responseCode int

const contentJSON = "application/json"

// Gateway is a Gateway
type Gateway struct {
	Username    string
	APIKey      string
	Environment string
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
	data := url.Values{}
	data.Set("username", gateway.Username)
	data.Set("to", recipients)
	data.Set("message", message)
	body := strings.NewReader(data.Encode())

	client := &http.Client{}
	r, err := http.NewRequest("POST", gateway.getSmsURL(), body)
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Content-Length", strconv.Itoa(body.Len()))
	r.Header.Add("apikey", gateway.APIKey)

	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	// make sure we have a response body
	if response == nil || response.Body == nil {
		return nil, fmt.Errorf("received empty response")
	}

	log.Println(response.Body)

	var smsResponse SMSResponse
	json.NewDecoder(response.Body).Decode(&smsResponse)
	defer response.Body.Close()

	if len(smsResponse.SMSMessageData.Recipients) > 0 {
		return smsResponse.SMSMessageData.Recipients, nil
	}

	return nil, fmt.Errorf("could not send sms message: %s to: %s", message, recipients)
}

func (gateway Gateway) getAPIHost() string {
	if gateway.Environment == "sandbox" {
		return "https://api.sandbox.africastalking.com"
	}

	return "https://api.africastalking.com"
}

func (gateway Gateway) getSmsURL() string {
	return gateway.getAPIHost() + "/version1/messaging"
}
