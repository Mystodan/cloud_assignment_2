package status

import (
	consts "assignment-2/constants"
	glob "assignment-2/globals"
	"assignment-2/globals/common"
	"fmt"
	"net/http"
	"time"
)

// Gets current uptime based on timer
func GetUptime(inn time.Time) time.Duration {
	return time.Since(inn)
}

// Handler for checking if api services are down
func checkAPIServices() map[string]string {
	returnAPIs := map[string]string{}
	compare := map[string]string{ // Api list
		"cases_api":  consts.CASES_API,
		"policy_api": consts.POLICY_API + consts.POLICY_TEST,
	}
	for api, url := range compare { // iterate though list of apis
		response, _ := common.GetURL(url) // checks for response
		if response.StatusCode == http.StatusBadRequest || response.StatusCode == http.StatusOK {
			returnAPIs[api] = fmt.Sprint(http.StatusOK) // available
		} else {
			returnAPIs[api] = fmt.Sprint(http.StatusServiceUnavailable)
		} // unavailable
	}
	return returnAPIs //returns status
}

// get amount of webhooks from stored buffer
func getWebHooksAmount(inn map[string]glob.Webhook) int {
	return len(inn)
}

// format status response
func getAPIstatus() statusInterface {
	APIstatuses := checkAPIServices()
	return statusInterface{
		APIstatuses["cases_api"],
		APIstatuses["policy_api"],
		getWebHooksAmount(glob.AllWebhooks), // webhook amount
		consts.APP_VERSION,                  // application version
		int64(GetUptime(Timer).Seconds()),   // getting uptime as seconds
	}
}
