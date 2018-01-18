package util

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
	if env == "sandbox" {
		return "https://api.sandbox.africastalking.com"
	}

	// env == "production"
	return "https://api.africastalking.com"
}

// GetSmsURL is the sms endpoint
func GetSmsURL(env string) string {
	return GetAPIHost(env) + "/version1/messaging"
}

func GetPaymentHost(env string) string {
	if env == "sandbox" {
		return "https://payments.sandbox.africastalking.com"
	}
	return "https://payments.africastalking.com"
}

func GetVoiceHost(env string) string {
	if env == "sandbox" {
		return "https://voice.sandbox.africastalking.com"
	}
	return "https://voice.africastalking.com"
}

func GetVoiceUrl(env string) string {
	return GetVoiceHost(env)
}

func GetSubscriptionUrl(env string) string {
	return GetAPIHost(env) + "/version1/subscription"
}

func GetUserDataUrl(env string) string {
	return GetAPIHost(env) + "/version1/user"
}

func GetAirtimeUrl(env string) string {
	return GetAPIHost(env) + "/version1/airtime"
}

func GetMobilePaymentCheckoutUrl(env string) string {
	return GetPaymentHost(env) + "/mobile/checkout/request"
}

func GetMobilePaymentB2BUrl(env string) string {
	return GetPaymentHost(env) + "/mobile/b2b/request"
}

func GetMobilePaymentB2CUrl(env string) string {
	return GetPaymentHost(env) + "/mobile/b2c/request"
}
