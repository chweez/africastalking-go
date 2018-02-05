package payment

import (
	"regexp"
	"time"
	"strconv"
)

const (
	ProviderMpesa  = "Mpesa"
	ProviderAthena = "Athena"

	TransferTypeBuyGoods = "BusinessBuyGoods"
	TransferTypePayBill  = "BusinessPayBill"
	TransferTypeDisburse = "DisburseFundsToBusiness"
	TransferTypeB2B      = "BusinessToBusinessTransfer"

	ReasonSalary             = "SalaryPayment"
	ReasonSalaryWithCharge   = "SalaryPaymentWithWithdrawalChargePaid"
	ReasonBusiness           = "BusinessPayment"
	ReasonBusinessWithCharge = "BusinessPaymentWithWithdrawalChargePaid"
	ReasonPromotion          = "PromotionPayment"
)

// BankCode really?
type BankCode int

const (
	FCMB_NG       BankCode = 234001
	Zenith_NG     BankCode = 234002
	Access_NG     BankCode = 234003
	GTBank_NG     BankCode = 234004
	Ecobank_NG    BankCode = 234005
	Diamond_NG    BankCode = 234006
	Providus_NG   BankCode = 234007
	Unity_NG      BankCode = 234008
	Stanbic_NG    BankCode = 234009
	Sterling_NG   BankCode = 234010
	Parkway_NG    BankCode = 234011
	Afribank_NG   BankCode = 234012
	Enterprise_NG BankCode = 234013
	Fidelity_NG   BankCode = 234014
	Heritage_NG   BankCode = 234015
	Keystone_NG   BankCode = 234016
	Skye_NG       BankCode = 234017
	Stanchart_NG  BankCode = 234018
	Union_NG      BankCode = 234019
	Uba_NG        BankCode = 234020
	Wema_NG       BankCode = 234021
	First_NG      BankCode = 234022
	CBA_KE        BankCode = 254001
	UNKNOWN       BankCode = -1
)

// Bank is a.. Bank
type Bank struct {
	CurrencyCode string            `json:"currencyCode"`
	Amount       float64           `json:"amount"`
	BankAccount  BankAccount       `json:"bankAccount"`
	Narration    string            `json:"narration"`
	Metadata     map[string]string `json:"metadata"`
}

// BankAccount is a model
type BankAccount struct {
	AccountName   string   `json:"accountName"`
	AccountNumber string   `json:"accountNumber"`
	BankCode      BankCode `json:"bankCode"`
	DateOfBirth   string   `json:"dateOfBirth"`
}

// Business is a business
type Business struct {
	CurrencyCode       string            `json:"currencyCode"`
	Amount             float64           `json:"amount"`
	Provider           string            `json:"provider"`
	TransferType       string            `json:"transferType"`
	DestinationChannel string            `json:"destinationChannel"`
	DestinationAccount string            `json:"destinationAccount"`
	Metadata           map[string]string `json:"metadata"`
}

// Consumer is a model
type Consumer struct {
	Name            string            `json:"name"`
	PhoneNumber     string            `json:"phoneNumber"`
	CurrencyCode    string            `json:"currencyCode"`
	Amount          float64           `json:"amount"`
	ProviderChannel string            `json:"providerChannel"`
	Reason          string            `json:"reason"`
	Metadata        map[string]string `json:"metadata"`
}

// Card is a model
type Card struct {
	Number      string `json:"number"`
	CVVNumber   int    `json:"cvvNumber"`
	ExpiryMonth int    `json:"expiryMonth"`
	ExpiryYear  int    `json:"expiryYear"`
	CountryCode string `json:"countryCode"`
	AuthToken   string `json:"authToken"`
}

var numberPattern, cvvPattern, countryCodePattern *regexp.Regexp

func init() {
	// ignore errors here, hope it won't bite us in the back
	numberPattern, _ = regexp.Compile("^\\d{12,19}$")
	cvvPattern, _ = regexp.Compile("^\\d{3,4}$")
	countryCodePattern, _ = regexp.Compile("^[A-Z]{2}$")
}

// IsValid checks whether card details are valid
func (card Card) IsValid() bool {
	if !numberPattern.MatchString(card.Number) {
		return false
	}
	if !cvvPattern.MatchString(strconv.Itoa(card.CVVNumber)) {
		return false
	}

	if !countryCodePattern.MatchString(card.CountryCode) {
		return false
	}

	if card.ExpiryMonth < 1 || card.ExpiryMonth > 12 {
		return false
	}

	if card.ExpiryYear < time.Now().Year() {
		return false
	}

	return card.AuthToken != ""
}
