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

// GetVoiceUrl returns the voice endpoint
func GetVoiceUrl(env string) string {
	return GetVoiceHost(env)
}

// GetSubscriptionUrl returns the Subscription endpoint
func GetSubscriptionUrl(env string) string {
	return GetAPIHost(env) + "/version1/subscription"
}

// GetUserDataUrl returns the  user data endpoint
func GetUserDataUrl(env string) string {
	return GetAPIHost(env) + "/version1/user"
}

// GetAirtimeUrl returns the airtime endpoint
func GetAirtimeUrl(env string) string {
	return GetAPIHost(env) + "/version1/airtime"
}

// GetMobilePaymentCheckoutUrl returns the mobile payments checkout endpoint
func GetMobilePaymentCheckoutUrl(env string) string {
	return GetPaymentHost(env) + "/mobile/checkout/request"
}

// GetMobilePaymentB2BUrl returns the Mobile Payments B2B endpoint
func GetMobilePaymentB2BUrl(env string) string {
	return GetPaymentHost(env) + "/mobile/b2b/request"
}

// GetMobilePaymentB2CUrl returns  the Mobile Payment B2C endpoint
func GetMobilePaymentB2CUrl(env string) string {
	return GetPaymentHost(env) + "/mobile/b2c/request"
}

func getHost(env, service string) string {
	if env != "sandbox" {
		return fmt.Sprintf("https://%s.africastalking.com", service)
	}
	return fmt.Sprintf("https://%s.sandbox.africastalking.com", service)

}
