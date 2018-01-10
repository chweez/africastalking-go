package payment

import (
	"regexp"
	"time"
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
	CurrencyCode string
	Amount       float64
	BankAccount  BankAccount
	Narration    string
	Metadata     map[string]string
}

// BankAccount is a model
type BankAccount struct {
	AccountName   string
	AccountNumber string
	BankCode      BankCode
	DateOfBirth   string
}

// Business is a business
type Business struct {
	CurrencyCode       string
	Amount             float64
	Provider           string
	TransferType       string
	DestinationChannel string
	DestinationAccount string
	Metadata           map[string]string
}

// Consumer is a model
type Consumer struct {
	Name            string
	PhoneNumber     string
	CurrencyCode    string
	Amount          float64
	ProviderChannel string
	Reason          string
	Metadata        map[string]string
}

// Card is a model
type Card struct {
	Number      string
	CVVNumber   int
	ExpiryMonth int
	ExpiryYear  int
	CountryCOde string
	AuthToken   string
}

var numberPattern, cvvPattern, countryCodePattern *regexp.Regexp

func init() {
	// ignore errors here, hope it won't back us in the back hehe
	numberPattern, _ = regexp.Compile("^\\d{12,19}$")
	cvvPattern, _ = regexp.Compile("^\\d{3,4}$")
	countryCodePattern, _ = regexp.Compile("^[A-Z]{2}$")
}

// IsValid checks whether card details are valid
func (card Card) IsValid() bool {
	if !numberPattern.MatchString(card.Number) {
		return false
	}

	if !cvvPattern.MatchString(string(card.CVVNumber)) {
		return false
	}

	if !countryCodePattern.MatchString(string(card.CountryCOde)) {
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
