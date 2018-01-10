package account

import (
	"africastalking/util"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Response is a model
type Response struct {
	User User `json:"UserData"`
}

// User is the model for a user
type User struct {
	Balance string `json:"balance"`
}

// Service is a service
type Service struct {
	Username string
	APIKey   string
	Env      string
}

// NewService returns a new service
func NewService(username, apiKey, env string) Service {
	return Service{username, apiKey, env}
}

// GetUser gets the user for the account
func (service Service) GetUser() (*User, error) {
	host := util.GetAPIHost(service.Env)
	url := host + "/version1/user"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request %v", err)
	}

	q := request.URL.Query()
	q.Add("username", service.Username)
	request.URL.RawQuery = q.Encode()
	request.Header.Set("apiKey", service.APIKey)
	request.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("could not get rsponse %v", err)
	}

	var accountResponse Response
	json.NewDecoder(response.Body).Decode(&accountResponse)
	return &accountResponse.User, nil
}
