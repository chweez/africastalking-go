package sms

// SendMessageResponse is a model
type SendMessageResponse struct {
	SMS SMS2 `json:"SMSMessageData"`
}

// SMS2 is a model
type SMS2 struct {
	Recipients []Recipient `json:"recipients"`
}

// SubscriptionResponse is a model
type SubscriptionResponse struct {
	Success     string `json:"success"`
	Description string `json:"description"`
}

// FetchMessageResponse is a model
type FetchMessageResponse struct {
	SMS SMS `json:"SMSMessageData"`
}

// SMS is a model
type SMS struct {
	Messages []Message `json:"Recipients"`
}

// Message is a model
type Message struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Text   string `json:"text"`
	LinkID string `json:"linkId"`
	Date   string `json:"date"`
	ID     int64  `json:"id"`
}

// FetchSubscriptionResponse is a model
type FetchSubscriptionResponse struct {
	Subscriptions []Subscription
}

// Subscription is a model
type Subscription struct {
	ID          int64  `json:"id"`
	PhoneNumber string `json:"phoneNumber"`
	Date        string `json:"date"`
}

// Recipient is a model
type Recipient struct {
	Number    string `json:"number"`
	Cost      string `json:"cost"`
	Status    string `json:"status"`
	MessageID string `json:"messageId"`
}