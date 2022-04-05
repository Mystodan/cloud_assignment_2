package policy

type Policy struct {
	Countrycode string  `json:"country_code"`
	Scope       string  `json:"scope"`
	Stringency  float64 `json:"stringency"`
	Policies    int     `json:"policies"`
}
