package payment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AndroidStudyOpenSource/africastalking-go/util"
)

// Service is a service
type Service struct {
	Username string
	APIKey   string
	Env      string
}

// NewService creates a new Service
func NewService(username, apiKey, env string) Service {
	return Service{username, apiKey, env}
}

// RequestB2C sends a B2C request
func (service Service) RequestB2C(body B2CRequest) (*B2CResponse, error) {
	url := util.GetMobilePaymentB2CURL(service.Env)

	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("could not marshal b2c req body %v: ", err)
	}

	response, err := service.newRequest(url, reqBody)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var b2cResponse B2CResponse
	json.NewDecoder(response.Body).Decode(&b2cResponse)
	return &b2cResponse, nil
}

// RequestB2B sends a B2B request
func (service Service) RequestB2B(body B2BRequest) (*B2BResponse, error) {
	url := util.GetMobilePaymentB2BURL(service.Env)

	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("could not marshal b2b req body %v: ", err)
	}

	response, err := service.newRequest(url, reqBody)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var b2bResponse B2BResponse
	json.NewDecoder(response.Body).Decode(&b2bResponse)
	return &b2bResponse, nil
}

// MobileCheckout requests
func (service Service) MobileCheckout(body MobileCheckoutRequest) (*CheckoutResponse, error) {
	url := util.GetMobilePaymentCheckoutURL(service.Env)

	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("could not marshal mobile checkout req body %v: ", err)
	}

	response, err := service.newRequest(url, reqBody)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var checkoutResponse CheckoutResponse
	json.NewDecoder(response.Body).Decode(&checkoutResponse)
	return &checkoutResponse, nil
}

// CardCheckoutCharge requests
func (service Service) CardCheckoutCharge(body CardCheckoutRequest) (*CheckoutResponse, error) {
	if !body.Card.IsValid() {
		return nil, fmt.Errorf("card invalid")
	}

	host := util.GetPaymentHost(service.Env)
	url := host + "/card/checkout/charge"

	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("could not marshal card checkout req body %v: ", err)
	}

	response, err := service.newRequest(url, reqBody)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var checkoutResponse CheckoutResponse
	json.NewDecoder(response.Body).Decode(&checkoutResponse)
	return &checkoutResponse, nil
}

// CardCheckoutValidate requests
func (service Service) CardCheckoutValidate(body CardValidateCheckoutRequest) (*CheckoutValidateResponse, error) {
	host := util.GetPaymentHost(service.Env)
	url := host + "/card/checkout/validate"

	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("could not marshal card validate checkout req body %v: ", err)
	}

	response, err := service.newRequest(url, reqBody)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var cvr CheckoutValidateResponse
	json.NewDecoder(response.Body).Decode(&cvr)
	return &cvr, nil
}

// BankCheckoutCharge requests
func (service Service) BankCheckoutCharge(body BankCheckoutRequest) (*CheckoutResponse, error) {
	host := util.GetPaymentHost(service.Env)
	url := host + "/bank/checkout/charge"

	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("could not marshal bank checkout charge req body %v: ", err)
	}

	response, err := service.newRequest(url, reqBody)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var checkoutResponse CheckoutResponse
	json.NewDecoder(response.Body).Decode(&checkoutResponse)
	return &checkoutResponse, nil
}

// BankCheckoutValidate requests
func (service Service) BankCheckoutValidate(body BankValidateCheckoutRequest) (*CheckoutValidateResponse, error) {
	host := util.GetPaymentHost(service.Env)
	url := host + "/bank/checkout/validate"

	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("could not marshal bank validate checkout req body %v: ", err)
	}

	response, err := service.newRequest(url, reqBody)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var cvr CheckoutValidateResponse
	json.NewDecoder(response.Body).Decode(&cvr)
	return &cvr, nil
}

// BankTransfer requests
func (service Service) BankTransfer(body BankTransferRequest) (*BankTransferResponse, error) {
	host := util.GetPaymentHost(service.Env)
	url := host + "/bank/transfer"

	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("could not marshal bank transfer req body %v: ", err)
	}

	response, err := service.newRequest(url, reqBody)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var btr BankTransferResponse
	json.NewDecoder(response.Body).Decode(&btr)
	return &btr, nil
}

func (service Service) newRequest(url string, body []byte) (*http.Response, error) {
	buffer := bytes.NewBuffer(body)
	request, err := http.NewRequest(http.MethodPost, url, buffer)
	if err != nil {
		return nil, err
	}

	request.Header.Set("apikey", service.APIKey)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Content-Length", strconv.Itoa(buffer.Len()))

	client := &http.Client{}
	return client.Do(request)
}
