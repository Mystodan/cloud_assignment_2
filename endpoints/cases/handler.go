package cases

import (
	consts "assignment-2/constants"
	"assignment-2/endpoints/notifications"
	"assignment-2/globals/common"
	funcs "assignment-2/globals/common"
	"encoding/json"
	"net/http"
)

var Url string

/**
 *	Handler for 'cases' endpoint
 */
func HandlerCases(w http.ResponseWriter, r *http.Request) {
	// Handles the Url by splitting its value strating after the CASES_PATH
	urlSplit := funcs.SplitURL(Url, w, r)

	if r.Method == http.MethodGet {
		w.Header().Add("content-type", "application/json")
		country := urlSplit[0]
		if len(country) < 1 {
			http.Error(w, "No country name inputted", http.StatusBadRequest)
			return
		}

		//country = funcs.DesensitizeString(country) - DEPRECATED (succeeded by GetA3 and GetCountry)
		country, err := funcs.GetCountry(country)
		if common.HandleErr(err, w, http.StatusNotAcceptable) {
			return
		}
		getGraphql, err := funcs.GetGraphql(country, consts.CASES_API, formatRequest(country))
		if funcs.HandleErr(err, w, http.StatusBadRequest) {
			return
		}

		// wrap response
		formattedResponse := wrapData(getGraphql)

		// invoke webhooks on thread, and send to writer
		go notifications.SetInvocation(country)

		err = json.NewEncoder(w).Encode(formattedResponse)
		if funcs.HandleErr(err, w, http.StatusInternalServerError) {
			return
		}

	} else {
		http.Error(w, common.MethodAllowed("GET"), http.StatusMethodNotAllowed)
		return
	}

}
