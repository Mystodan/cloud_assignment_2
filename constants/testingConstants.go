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
const TEST_DEFAULT_MOCK_JSON = "mocked_data.json"

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

// Notifications
const TEST_SERVER_ID_SYMBOLS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const TEST_SERVER_ID_LENGTH = 20
const TEST_NOTIFICATIONS = `{
	"url": "https://webhook.site/7eede626-320e-40c4-81e6-4cb09a3af96d",
	"country": "Norway",
	"calls": 2
}`
const MOCKED_WEBHOOKS = `{"1Mu853KrMUpS9hEXMwGa":{"webhook_id":"HyaZRxKsdQKdu","url":"https://webhook.site/e3a17245-2081-4aeb-a113-223b83960184","country":"Japan","calls":2},"1kMJbsuqy9Ojb3ZXpj0J":{"webhook_id":"fWwrgBOfEbZDh","url":"https://webhook.site/e3a17245-2081-4aeb-a113-223b83960184","country":"Norway","calls":2},"3XDMUGUhb5jv0EaYNIYG":{"webhook_id":"nwCmUMkVLPvQn","url":"https://webhook.site/e3a17245-2081-4aeb-a113-223b83960184","country":"Japan","calls":5}}`
const MOCKED_WEBHOOK_ID = "JjPjzpfRFEgmo"

// Status default
const TEST_STATUS_DEFAULT = `{"cases_api":"200","policy_api":"200","webhooks":0,"version":"v1","uptime":9223372036}`
