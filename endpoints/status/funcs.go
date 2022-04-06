package status

import (
	consts "assignment-2/constants"
	funcs "assignment-2/endpoints"
	"assignment-2/endpoints/notifications"
	server "assignment-2/server/functions"
	"fmt"
	"net/http"
)

func checkAPIServices() map[string]string {
	returnAPIs := map[string]string{}
	compare := map[string]string{
		"cases_api":  consts.CASES_API,
		"policy_api": consts.POLICY_API + consts.POLICY_TEST,
	}
	for api, url := range compare {
		response, _, _ := funcs.GetURL(url)
		if response.StatusCode == http.StatusBadRequest || response.StatusCode == http.StatusOK {
			returnAPIs[api] = http.StatusText(http.StatusOK)
		} else {
			returnAPIs[api] = http.StatusText(http.StatusServiceUnavailable)
		}
	}
	return returnAPIs
}

func getWebHooksAmount(inn map[string]notifications.Webhook) string {
	return fmt.Sprintf("%d registered webhooks", len(inn))
}

func getAPIstatus() statusInterface {
	APIstatuses := checkAPIServices()
	return statusInterface{
		APIstatuses["cases_api"],
		APIstatuses["policy_api"],
		getWebHooksAmount(notifications.GetAllWebhooks()),
		consts.APP_VERSION,
		fmt.Sprintf("%f", server.GetUptime(Timer).Seconds()) + "s",
	}
}
