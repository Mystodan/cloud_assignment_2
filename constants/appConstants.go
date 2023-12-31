package consts

// app details
const APP_VERSION = "v1"

// Collections for FireBase
const COLLECTION_WEBHOOKS = "webhooks"
const COLLECTION_CACHE = "cache"

// webhook symbol range
const APP_TOKEN_SYMBOLS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const APP_TOKEN_LENGTH = 13

// Cases
const CASES_REQUEST = `
query {
	country(name: "%s") {
		name
		mostRecent {
			date(format: "yyyy-MM-dd")
			confirmed
			recovered
			deaths
			growthRate
		}
	}
}`
const CASES_GET_ALL = `
query {
	countries(names:[]){
		name
		}
}`

// Policy
const POLICY_TEST = "NOR/2022-04-04"
const POLICY_DATE = "2006-01-02"
const POLICY_TIME_DAY = -2 // 2 days before current date, in order to get correct data
const POLICY_TIME_MONTH = 0
const POLICY_TIME_YEAR = 0

// Cache
const CACHE_NEXT_PARAM = "<nextParam>"
const CACHE_NO_DATA = "currently no data exists within cache"
