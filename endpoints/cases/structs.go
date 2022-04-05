package cases

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
