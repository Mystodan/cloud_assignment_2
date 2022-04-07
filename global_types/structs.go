package glob

type Webhook struct {
	ID      string `json:"webhook_id"`
	Url     string `json:"url"`
	Country string `json:"country"`
	Calls   int64  `json:"calls"`
}

/**
 *	The struct for the required covid19-tracker-API 'policies'
 */
type Policy struct {
	Countrycode string `json:"country_code"`
	Scope       string `json:"scope"`
	Stringency  string `json:"stringency"`
	Policies    string `json:"policies"`
}

/**
 *	The struct for the required covid19-API 'cases'
 */
type Case struct {
	Country     string  `json:"country"`
	Date        string  `json:"date"`
	Confirmed   float64 `json:"confirmed"`
	Recovered   float64 `json:"recovered"`
	Deaths      float64 `json:"deaths"`
	Growth_rate float64 `json:"growth_rate"`
}

/**
 *	The struct for the required alpha3 codes
 */
type Countries struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
