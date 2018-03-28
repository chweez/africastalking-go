package util

import "fmt"

// Constants for responses
const (
	PendingConfirmation = "PendingConfirmation"
	PendingValidation   = "PendingValidation"
	InvalidRequest      = "InvalidRequest"
	NotSupported        = "NotSupported"
	SUCCESS             = "Success"
	FAILED              = "Failed"
	QUEUED              = "Queued"
	SENT                = "Sent"
)

// GetAPIHost returns either sandbox or prod
func GetAPIHost(env string) string {
	return getHost(env, "api")
}

// GetSmsURL is the sms endpoint
func GetSmsURL(env string) string {
	return GetAPIHost(env) + "/version1/messaging"
}

// GetPaymentHost returns the payments domain
func GetPaymentHost(env string) string {
	return getHost(env, "payments")
}

// GetVoiceHost returns the voice host domain
func GetVoiceHost(env string) string {
	return getHost(env, "voice")
}

// GetVoiceURL returns the voice endpoint
func GetVoiceURL(env string) string {
	return GetVoiceHost(env)
}

// GetSubURL returns the Subscription endpoint
func GetSubURL(env string) string {
	return GetAPIHost(env) + "/version1/subscription"
}

// GetCreateSubURL returns the Subscription create endpoint
func GetCreateSubURL(env string) string {
	return GetAPIHost(env) + "/version1/subscription/create"
}

// GetUserDataURL returns the  user data endpoint
func GetUserDataURL(env string) string {
	return GetAPIHost(env) + "/version1/user"
}

// GetAirtimeURL returns the airtime endpoint
func GetAirtimeURL(env string) string {
	return GetAPIHost(env) + "/version1/airtime"
}

// GetMobilePaymentCheckoutURL returns the mobile payments checkout endpoint
func GetMobilePaymentCheckoutURL(env string) string {
	return GetPaymentHost(env) + "/mobile/checkout/request"
}

// GetMobilePaymentB2BURL returns the Mobile Payments B2B endpoint
func GetMobilePaymentB2BURL(env string) string {
	return GetPaymentHost(env) + "/mobile/b2b/request"
}

// GetMobilePaymentB2CURL returns  the Mobile Payment B2C endpoint
func GetMobilePaymentB2CURL(env string) string {
	return GetPaymentHost(env) + "/mobile/b2c/request"
}

func GetCreateCheckoutTokenURL(env string) string {
	return GetAPIHost(env) + "/checkout/token/create"
}

func GetGenerateAuthTokenURL(env string) string {
	return GetAPIHost(env) + "/auth-token/generate"
}

func getHost(env, service string) string {
	if env != "sandbox" {
		return fmt.Sprintf("https://%s.africastalking.com", service)
	}
	return fmt.Sprintf("https://%s.sandbox.africastalking.com", service)

}
