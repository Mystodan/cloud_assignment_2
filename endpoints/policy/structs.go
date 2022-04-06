package policy

type Policy struct {
	Countrycode string `json:"country_code"`
	Scope       string `json:"scope"`
	Stringency  string `json:"stringency"`
	Policies    string `json:"policies"`
}
