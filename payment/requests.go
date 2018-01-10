package payment

// CheckoutRequest is a model
type CheckoutRequest struct {
	Username     string            `json:"username"`
	ProductName  string            `json:"productName"`
	Narration    string            `json:"narration"`
	Metadata     map[string]string `json:"metadata"`
	CurrencyCode string            `json:"currencyCode"`
	Amount       float64           `json:"amount"`
}

// CheckoutValidateRequest is a model
type CheckoutValidateRequest struct {
	Username      string `json:"username"`
	TransactionID string `json:"transactionId"`
	OTP           string `json:"otp"`
}

// B2CRequest is a model
type B2CRequest struct {
	Username    string     `json:"username"`
	ProductName string     `json:"productName"`
	Recipients  []Consumer `json:"recipients"`
}

// B2BRequest is a model
type B2BRequest struct {
	Username    string `json:"username"`
	ProductName string `json:"productName"`
	Business
}

// MobileCheckoutRequest is a model
type MobileCheckoutRequest struct {
	CheckoutRequest
	PhoneNumber string `json:"phoneNumber"`
}

// CardCheckoutRequest is a checkout type for card payment
type CardCheckoutRequest struct {
	CheckoutRequest
	Card  Card   `json:"paymentCard"`
	Token string `json:"checkoutToken"`
}

// CardValidateCheckoutRequest is a model
type CardValidateCheckoutRequest struct {
	CheckoutValidateRequest
}
