package status

import (
	consts "assignment-2/constants"
	glob "assignment-2/globals"
	"assignment-2/globals/common"
	"fmt"
	"net/http"
	"time"
)

func GetUptime(inn time.Time) time.Duration {
	return time.Since(inn)
}

func checkAPIServices() map[string]string {
	returnAPIs := map[string]string{}
	compare := map[string]string{
		"cases_api":  consts.CASES_API,
		"policy_api": consts.POLICY_API + consts.POLICY_TEST,
	}
	for api, url := range compare {
		response, _ := common.GetURL(url)
		if response.StatusCode == http.StatusBadRequest || response.StatusCode == http.StatusOK {
			returnAPIs[api] = fmt.Sprint(http.StatusOK)
		} else {
			returnAPIs[api] = fmt.Sprint(http.StatusServiceUnavailable)
		}
	}
	return returnAPIs
}

func getWebHooksAmount(inn map[string]glob.Webhook) int {
	return len(inn)
}

func getAPIstatus() statusInterface {
	APIstatuses := checkAPIServices()
	return statusInterface{
		APIstatuses["cases_api"],
		APIstatuses["policy_api"],
		getWebHooksAmount(glob.AllWebhooks),
		consts.APP_VERSION,
		int64(GetUptime(Timer).Seconds()),
	}
}
