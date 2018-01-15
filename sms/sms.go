package sms

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"africastalking/util"
)

const (
	Sandbox = "sandbox"
	Prod    = "production"
)

// Service is a model
type Service struct {
	Username string
	APIKey   string
	Env      string
}

// NewService returns a new service
func NewService(username, apiKey, env string) Service {
	return Service{username, apiKey, env}
}

// SendToMany is a utility method to send to many recipients at the same time
func (service Service) SendToMany(from, message string, to []string) (*SendMessageResponse, error) {
	recipients := strings.Join(to, ",")
	return service.Send(from, recipients, message)
}

// Send - POST
func (service Service) Send(from, to, message string) (*SendMessageResponse, error) {
	values := url.Values{}
	values.Set("username", service.Username)
	values.Set("to", to)
	values.Set("from", from)
	values.Set("message", message)

	url := util.GetSmsURL(service.Env)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"

	res, err := service.newPostRequest(url, values, headers)
	if err != nil {
		return nil, err
	}

	var smsMessageResponse SendMessageResponse
	json.NewDecoder(res.Body).Decode(&smsMessageResponse)
	defer res.Body.Close()

	return &smsMessageResponse, nil
}

// SendBulkToMany is a utility method to send to many recipients at the same time
func (service Service) SendBulkToMany(from, message string, to []string, bulkMode int, enqueue string) (*SendMessageResponse, error) {
	recipients := strings.Join(to, ",")
	return service.SendBulk(from, recipients, message, bulkMode, enqueue)
}

// SendBulk - POST
func (service Service) SendBulk(from, to, message string, bulkMode int, enqueue string) (*SendMessageResponse, error) {
	values := url.Values{}
	values.Set("username", service.Username)
	values.Set("to", to)
	values.Set("from", from)
	values.Set("message", message)
	values.Set("bulkMode", string(bulkMode))
	values.Set("enqueue", enqueue)

	url := util.GetSmsURL(service.Env)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"

	res, err := service.newPostRequest(url, values, headers)
	if err != nil {
		return nil, err
	}

	var smsMessageResponse SendMessageResponse
	json.NewDecoder(res.Body).Decode(&smsMessageResponse)
	defer res.Body.Close()

	return &smsMessageResponse, nil
}

// SendPremium - POST
func (service Service) SendPremium(username, to, from, message, keyword,
	linkID, retryDurationInHours string, bulkMode int) (*SendMessageResponse, error) {
	values := url.Values{}
	values.Set("username", username)
	values.Set("to", to)
	values.Set("from", from)
	values.Set("message", message)
	values.Set("keyword", keyword)
	values.Set("bulkMode", string(bulkMode))
	values.Set("linkId", linkID)
	values.Set("retryDurationInHours", retryDurationInHours)

	url := util.GetSmsURL(service.Env)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"

	res, err := service.newPostRequest(url, values, headers)
	if err != nil {
		return nil, err
	}

	var smsMessageResponse SendMessageResponse
	json.NewDecoder(res.Body).Decode(&smsMessageResponse)
	defer res.Body.Close()

	return &smsMessageResponse, nil
}

// FetchMessage - username = query
func (service Service) FetchMessage(username, lastReceivedID string) (*FetchMessageResponse, error) {
	url := util.GetAPIHost(service.Env)
	queries := make(map[string]string)
	queries["username"] = username
	queries["lastReceivedID"] = lastReceivedID

	res, err := service.newGetRequest(url, queries)
	if err != nil {
		return nil, fmt.Errorf("could not get response: %v", err)
	}

	var fmr FetchMessageResponse
	json.NewDecoder(res.Body).Decode(&fmr)
	defer res.Body.Close()

	return &fmr, nil
}

// FetchSubscription - query
func (service Service) FetchSubscription(username, shortCode, keyword, lastReceivedID string) (*FetchSubscriptionResponse, error) {
	url := util.GetAPIHost(service.Env) + "/version1/subscription"
	queries := make(map[string]string)
	queries["username"] = username
	queries["shortCode"] = shortCode
	queries["keyword"] = keyword
	queries["lastReceivedID"] = lastReceivedID

	res, err := service.newGetRequest(url, queries)
	if err != nil {
		return nil, fmt.Errorf("could not get response: %v", err)
	}

	var fsr FetchSubscriptionResponse
	json.NewDecoder(res.Body).Decode(&fsr)
	defer res.Body.Close()

	return &fsr, nil
}

// CreateSubscription - POST
func (service Service) CreateSubscription(username, shortCode, keyword, phoneNumber, checkoutToken string) (*SubscriptionResponse, error) {
	values := url.Values{}
	values.Set("username", service.Username)
	values.Set("shortCode", shortCode)
	values.Set("keyword", keyword)
	values.Set("phoneNumber", phoneNumber)
	values.Set("checkoutToken", checkoutToken)

	headers := make(map[string]string)

	url := util.GetAPIHost(service.Env) + "/version1/subscription/create"
	res, err := service.newPostRequest(url, values, headers)
	if err != nil {
		return nil, fmt.Errorf("could not get response: %v", err)
	}

	var subsciprionResponse SubscriptionResponse
	json.NewDecoder(res.Body).Decode(&subsciprionResponse)
	defer res.Body.Close()

	return &subsciprionResponse, nil
}

func (service Service) newPostRequest(url string, values url.Values, headers map[string]string) (*http.Response, error) {
	reader := strings.NewReader(values.Encode())

	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Length", strconv.Itoa(reader.Len()))

	client := &http.Client{Timeout: 10 * time.Second}
	return client.Do(req)
}

func (service Service) newGetRequest(url string, queries map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	values := req.URL.Query()
	for key, value := range queries {
		values.Add(key, value)
	}
	req.URL.RawQuery = values.Encode()

	client := &http.Client{Timeout: 10 * time.Second}
	return client.Do(req)
}
