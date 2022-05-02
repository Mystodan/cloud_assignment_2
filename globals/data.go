package glob

// current port
var PORT string

//invocations
var CountryInvocations map[string]int = make(map[string]int)

//local webhooks
var AllWebhooks map[string]Webhook = make(map[string]Webhook)

//local ALPHA3 library
var AllCountries []Countries = []Countries{}

// memory buffer for caching
var MemBuffer map[string]map[string]interface{}
