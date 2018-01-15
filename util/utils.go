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
	return "https://africastalking.com"
}

// GetSmsURL is the sms endpoint
func GetSmsURL(env string) string {
	return GetAPIHost(env) + "/version1/messaging"
}
