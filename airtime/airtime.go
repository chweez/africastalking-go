package airtime

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/AndroidStudyOpenSource/africastalking-go/util"
)

// Response is the reponse from the api
type Response struct {
	NumSent       int
	TotalAmount   string
	TotalDiscount string
	ErrorMessage  string
	Responses     []Entry
}

// Entry is the entry for each airtime response
type Entry struct {
	ErrorMessage string
	PhoneNumber  string
	Amount       string
	Discount     string
	Status       string
	RequestID    string
}

// Service is the airtime service
type Service struct {
	Username string
	APIKey   string
	Env      string
}

// NewService returns a new service
func NewService(username, apiKey, env string) Service {
	return Service{username, apiKey, env}
}

// Send sends a new airtime request
func (service Service) Send() (*Response, error) {
	host := util.GetAPIHost(service.Env)
	url := host + "/version1/airtime"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request %v", err)
	}

	values := request.URL.Query()
	values.Add("username", service.Username)
	request.URL.RawQuery = values.Encode()

	request.Header.Set("apikey", service.APIKey)
	request.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("could not get rsponse %v", err)
	}
	defer response.Body.Close()

	var airtimeResponse Response
	json.NewDecoder(response.Body).Decode(&airtimeResponse)
	return &airtimeResponse, nil
}
