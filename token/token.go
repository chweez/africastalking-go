package token

import (
	"net/url"
	"strings"
	"net/http"
	"time"
	"fmt"
	"encoding/json"
	"strconv"
	"bytes"
	"github.com/AndroidStudyOpenSource/africastalking-go/util"
)

// CheckoutTokenResponse is the response from a create checkout token request
type CheckoutTokenResponse struct {
	Token       string `json:"token"`
	Description string `json:"description"`
}

// AuthTokenResponse is the response from a generate auth token request
type AuthTokenResponse struct {
	Token           string `json:"token"`
	LifetimeSeconds int64  `json:"lifetimeInSeconds"`
}

// Service is a service
type Service struct {
	Username string
	APIKey   string
	Env      string
}

func NewService(username, apiKey, env string) Service {
	return Service{username, apiKey, env}
}

func (s Service) CreateCheckoutToken(phoneNumber string) (*CheckoutTokenResponse, error) {
	checkoutURL := util.GetCreateCheckoutTokenURL(s.Env)

	values := url.Values{}
	values.Set("phoneNumber", phoneNumber)

	reader := strings.NewReader(values.Encode())

	req, err := http.NewRequest(http.MethodPost, checkoutURL, reader)
	if err != nil {
		return nil, fmt.Errorf("could not create checkout token create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", strconv.Itoa(reader.Len()))
	req.Header.Set("apikey", s.APIKey)

	client := http.Client{Timeout: 10 * time.Second}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not send checkout request: %v", err)
	}
	defer res.Body.Close()

	var ctr CheckoutTokenResponse
	if err := json.NewDecoder(res.Body).Decode(&ctr); err != nil {
		return nil, fmt.Errorf("could not decode create checkout token response reader: %v", err)
	}

	return &ctr, nil
}

func (s Service) GenerateAuthToken() (*AuthTokenResponse, error) {
	generateTokenURL := util.GetGenerateAuthTokenURL(s.Env)

	authRequest := struct {
		Username string `json:"username"`
	}{Username: s.Username}

	atrBytes, err := json.Marshal(authRequest)
	if err != nil {
		return nil, fmt.Errorf("could not marshal generate token struct: %v", err)
	}
	reqBody := bytes.NewReader(atrBytes)

	req, err := http.NewRequest(http.MethodPost, generateTokenURL, reqBody)
	if err != nil {
		return nil, fmt.Errorf("could not create generate auth token request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Length", strconv.Itoa(reqBody.Len()))
	req.Header.Set("apikey", s.APIKey)

	client := http.Client{Timeout: 10 * time.Second}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not send generate auth token request: %v", err)
	}
	defer res.Body.Close()

	var authResponse AuthTokenResponse
	if err = json.NewDecoder(res.Body).Decode(&authResponse); err != nil {
		return nil, fmt.Errorf("could not decode generate auth token response: %v", err)
	}

	return &authResponse, nil
}
