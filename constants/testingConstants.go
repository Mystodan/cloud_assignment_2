package consts

// default naming for default test
const TEST_DEFAULT_PATH = "default path"

// Cases
// default mock for default test
const TEST_CASES_DEFAULT_MOCK = `{
	"country": {
		"name": "Norway",
		"mostRecent": {
			"date": "2022-04-06",
			"confirmed": 1412969,
			"deaths": 2667,
			"recovered": 0,
			"growthRate": 0.001005277886011831
		}
	}
}
`

// Policy
// default mock for default test
const TEST_POLICY_DEFAULT_MOCK = "mocked_data.json"

// no data response
const TEST_POLICY_NO_DATA = `{
	"policyActions": [
		{
			"policy_type_code": "NONE",
			"policy_type_display": "No data.  Data may be inferred for last 7 days.",
			"flag_value_display_field": "",
			"policy_value_display_field": "No data.  Data may be inferred for last 7 days.",
			"policyvalue": 0,
			"flagged": null,
			"notes": null
		}
	],
	"stringencyData": {
		"msg": "Data unavailable"
	}
}`

// Status default
const TEST_STATUS_DEFAULT = `{"cases_api":"200","policy_api":"200","webhooks":0,"version":"v1","uptime":9223372036}`
