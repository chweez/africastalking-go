package payment

import (
	"africastalking/util"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Service is a service
type Service struct {
	Username string
	APIKey   string
	Env      string
}

// NewService creates a new Service
func NewService() Service {
	return Service{}
}

// RequestB2C sends a B2C request
func (service Service) RequestB2C(body map[string]interface{}) (*B2CResponse, error) {
	host := util.GetAPIHost(service.Env)
	url := host + "/mobile/b2c/request"

	response, err := service.newRequest(url, body)
	if err != nil {
		return nil, err
	}

	var b2cResponse B2CResponse
	json.NewDecoder(response.Body).Decode(&b2cResponse)
	defer response.Body.Close()
	return &b2cResponse, nil
}

// RequestB2B sends a B2B request
func (service Service) RequestB2B(body map[string]interface{}) (*B2BResponse, error) {
	host := util.GetAPIHost(service.Env)
	url := host + "/mobile/b2b/request"

	response, err := service.newRequest(url, body)
	if err != nil {
		return nil, err
	}

	var b2bResponse B2BResponse
	json.NewDecoder(response.Body).Decode(&b2bResponse)
	defer response.Body.Close()
	return &b2bResponse, nil
}

// MobileCheckout requests
func (service Service) MobileCheckout(body map[string]interface{}) (*CheckoutResponse, error) {
	host := util.GetAPIHost(service.Env)
	url := host + "/mobile/checkout/request"

	response, err := service.newRequest(url, body)
	if err != nil {
		return nil, err
	}

	var checkoutResponse CheckoutResponse
	json.NewDecoder(response.Body).Decode(&checkoutResponse)
	defer response.Body.Close()
	return &checkoutResponse, nil
}

// CardCheckoutCharge requests
func (service Service) CardCheckoutCharge(body map[string]interface{}) (*CheckoutResponse, error) {
	host := util.GetAPIHost(service.Env)
	url := host + "/card/checkout/charge"

	response, err := service.newRequest(url, body)
	if err != nil {
		return nil, err
	}

	var checkoutResponse CheckoutResponse
	json.NewDecoder(response.Body).Decode(&checkoutResponse)
	defer response.Body.Close()
	return &checkoutResponse, nil
}

// CardCheckoutValidate requests
func (service Service) CardCheckoutValidate(body map[string]interface{}) (*CheckoutValidateResponse, error) {
	host := util.GetAPIHost(service.Env)
	url := host + "/card/checkout/validate"

	response, err := service.newRequest(url, body)
	if err != nil {
		return nil, err
	}

	var cvr CheckoutValidateResponse
	json.NewDecoder(response.Body).Decode(&cvr)
	defer response.Body.Close()
	return &cvr, nil
}

// BankCheckoutCharge requests
func (service Service) BankCheckoutCharge(body map[string]interface{}) (*CheckoutResponse, error) {
	host := util.GetAPIHost(service.Env)
	url := host + "/bank/checkout/charge"

	response, err := service.newRequest(url, body)
	if err != nil {
		return nil, err
	}

	var checkoutResponse CheckoutResponse
	json.NewDecoder(response.Body).Decode(&checkoutResponse)
	defer response.Body.Close()
	return &checkoutResponse, nil
}

// BankCheckoutValidate requests
func (service Service) BankCheckoutValidate(body map[string]interface{}) (*CheckoutValidateResponse, error) {
	host := util.GetAPIHost(service.Env)
	url := host + "/bank/checkout/validate"

	response, err := service.newRequest(url, body)
	if err != nil {
		return nil, err
	}

	var cvr CheckoutValidateResponse
	json.NewDecoder(response.Body).Decode(&cvr)
	defer response.Body.Close()
	return &cvr, nil
}

// BankTransfer requests
func (service Service) BankTransfer(body map[string]interface{}) (*BankTransferResponse, error) {
	host := util.GetAPIHost(service.Env)
	url := host + "/bank/transfer"

	response, err := service.newRequest(url, body)
	if err != nil {
		return nil, err
	}

	var btr BankTransferResponse
	json.NewDecoder(response.Body).Decode(&btr)
	defer response.Body.Close()
	return &btr, nil
}

func (service Service) newRequest(url string, body map[string]interface{}) (*http.Response, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("could not convert map to body: %v", err)
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	request.Header.Set("apiKey", service.APIKey)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(request)
}
