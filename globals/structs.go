package glob

type Webhook struct {
	ID      string `json:"webhook_id"`
	Url     string `json:"url"`
	Country string `json:"country"`
	Calls   int64  `json:"calls"`
}

/**
 *	The struct for the required alpha3 codes
 */
type Countries struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
