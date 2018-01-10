package payment

// B2BResponse is a service
type B2BResponse struct {
	Status          string `json:"status"`
	TransactionID   string `json:"transactionId"`
	TransactionFee  string `json:"transactionFee"`
	ProviderChannel string `json:"providerChannel"`
}

// B2CResponse is a service
type B2CResponse struct {
	NumQueued           int        `json:"numQueued"`
	TotalValue          string     `json:"totalValue"`
	TotalTransactionFee string     `json:"totalTransactionFee"`
	Entries             []B2CEntry `json:"entries"`
}

//B2CEntry is a model
type B2CEntry struct {
	PhoneNumber     string `json:"phoneNumber"`
	Status          string `json:"status"`
	Provider        string `json:"provider"`
	ProviderChannel string `json:"providerChannel"`
	Value           string `json:"value"`
	TransactionID   string `json:"transactionId"`
	TransactionFee  string `json:"transactionFee"`
	ErrorMessage    string `json:"errorMessage"`
}

// CheckoutResponse is a model
type CheckoutResponse struct {
	Status        string `json:"status"`
	TransactioID  string `json:"transactionId"`
	Description   string `json:"description"`
	CheckoutToken string `json:"checkoutToken"`
}

// CheckoutValidateResponse is a model
type CheckoutValidateResponse struct {
	Status        string `json:"status"`
	Description   string `json:"description"`
	CheckoutToken string `json:"checkoutToken"`
}

// BankTransferResponse is a model
type BankTransferResponse struct {
	ErrorMessage string      `json:"errorMessage"`
	Entries      []BankEntry `json:"entries"`
}

// BankEntry is a model
type BankEntry struct {
	AccountNumber  string `json:"accountNumber"`
	Status         string `json:"status"`
	TransactioID   string `json:"transactionId"`
	TransactionFee string `json:"transactionFee"`
	ErrorMessage   string `json:"errorMessage"`
}
