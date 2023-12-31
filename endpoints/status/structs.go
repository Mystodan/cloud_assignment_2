package status

type statusInterface struct {
	Cases_api  string `json:"cases_api" ` // http status code for *Covid 19 Cases API
	Policy_api string `json:"policy_api"` // http status code for *Corona Policy Stringency API
	Webhooks   int    `json:"webhooks"`   // <number of registered webhooks>,
	Version    string `json:"version"`
	Uptime     int64  `json:"uptime"` // <time in seconds from the last service restart>
}
