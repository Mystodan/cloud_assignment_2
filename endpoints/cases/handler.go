package cases

import (
	consts "assignment-2/constants"
	funcs "assignment-2/endpoints"
	"assignment-2/endpoints/notifications"
	glob "assignment-2/global_types"
	"encoding/json"
	"net/http"
)

/**
 *	Handler for 'cases' endpoint
 */
func HandlerCases(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	// Handles the Url by splitting its value strating after the CASES_PATH
	urlSplit := funcs.SplitURL(consts.CASES_PATH, w, r)

	if r.Method == http.MethodGet {
		country := urlSplit[0]
		if len(country) < 1 {
			http.Error(w, "No country name inputted", http.StatusBadRequest)
			return
		}

		//country = funcs.DesensitizeString(country) - DEPRECATED

		country = funcs.GetCountry(country)

		getGraphql, err := funcs.GetGraphql(consts.CASES_API, formatRequest(country))
		if funcs.HandleErr(err, w, http.StatusBadRequest) {
			return
		}

		// wrap response
		formattedResponse := wrapData(getGraphql)

		// invoke webhooks, annd send to writer
		if _, invoked := glob.CountryInvocations[country]; !invoked {
			glob.CountryInvocations[country] = 0
		}

		glob.CountryInvocations[country]++
		notifications.InvokeWebhooks(country, glob.CountryInvocations, glob.AllWebhooks)

		err = json.NewEncoder(w).Encode(formattedResponse)
		if funcs.HandleErr(err, w, http.StatusInternalServerError) {
			return
		}

	} else {
		http.Error(w, "Method not allowed, use GET", http.StatusMethodNotAllowed)
		return
	}

}
