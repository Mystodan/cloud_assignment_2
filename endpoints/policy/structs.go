package policy

/**
 *	The struct for the required covid19-tracker-API 'policies'
 */
type Policy struct {
	Countrycode string  `json:"country_code"`
	Scope       string  `json:"scope"`
	Stringency  float64 `json:"stringency"`
	Policies    int64   `json:"policies"`
}
