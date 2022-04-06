package notifications

type Webhook struct {
	ID      string `json:"webhook_id"`
	Url     string `json:"url"`
	Country string `json:"country"`
	Calls   int64  `json:"calls"`
}
