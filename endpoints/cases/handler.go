package cases

import (
	consts "assignment-2/constants"
	"assignment-2/endpoints/notifications"
	glob "assignment-2/globals"
	"assignment-2/globals/common"
	funcs "assignment-2/globals/common"
	"encoding/json"
	"net/http"
)

// necessary for routing
var Url string

// necessary for mocking
var (
	GetRequest = common.GetGraphql
)

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
			http.Error(w, consts.NO_COUNTRY_INPUT, http.StatusBadRequest)
			return
		}

		//country = funcs.DesensitizeString(country) - DEPRECATED (succeeded by GetA3 and GetCountry)
		_country, err := funcs.GetCountry(country)
		country = _country.Name
		if common.HandleErr(err, w, http.StatusNotAcceptable) {
			return
		}
		getGraphql, err := GetRequest(country, consts.CASES_API, formatRequest(country))
		if funcs.HandleErr(err, w, http.StatusBadRequest) {
			return
		}

		// wrap response
		formattedResponse := wrapData(getGraphql)

		// invoke webhooks on thread, and send to writer
		if glob.AllowInvocations {
			go notifications.SetInvocation(country)
		}

		err = json.NewEncoder(w).Encode(formattedResponse)
		if funcs.HandleErr(err, w, http.StatusInternalServerError) {
			return
		}

	} else {
		http.Error(w, common.MethodAllowed("GET"), http.StatusMethodNotAllowed)
		return
	}

}
