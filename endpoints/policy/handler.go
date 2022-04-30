package policy

import (
	consts "assignment-2/constants"
	"assignment-2/globals/common"

	"assignment-2/endpoints/notifications"
	"encoding/json"
	"net/http"
	"time"
)

var Url string

/**
 *	Handler for 'policy' endpoint
 */
func HandlerPolicy(w http.ResponseWriter, r *http.Request) {
	// Handles the Url by splitting its value strating after the CASES_PATH
	urlSplit := common.SplitURL(Url, w, r)
	urlQuery := r.URL.Query()

	if r.Method == http.MethodGet {
		w.Header().Add("content-type", "application/json")
		country := urlSplit[0]
		if len(country) < 1 {
			http.Error(w, "No country name inputted", http.StatusBadRequest)
			return
		}
		if len(urlSplit) < 1 {
			http.Error(w, "Not enough parameters, try again", http.StatusBadRequest)
			return
		}

		// get optional parameter
		var optParam string
		if _, isScope := urlQuery["scope"]; isScope {
			optParam = urlQuery["scope"][0]

			if _, err := time.Parse(consts.POLICY_DATE, optParam); common.HandleErr(err, w, http.StatusRequestedRangeNotSatisfiable) {
				return
			}
		} else {
			optParam = time.Now().AddDate(consts.POLICY_TIME_YEAR, consts.POLICY_TIME_MONTH, consts.POLICY_TIME_DAY).Format(consts.POLICY_DATE)
		}

		// convert to A3 code
		country, err := common.GetA3(country)
		if common.HandleErr(err, w, http.StatusNotAcceptable) {
			return
		}
		// Send request to api
		getRequest, err := common.RequestURL(country, formatRequest(country, optParam))
		if err != nil {
			http.Error(w, "Error sending request to covidtracker API", http.StatusFailedDependency)
			return
		}
		// wrap response
		formattedResponse := wrapData(getRequest)

		// invoke webhooks on thread, and send to writer
		DesensitizeString, _ := common.GetCountry(country)
		go notifications.SetInvocation(DesensitizeString)
		// send to writer
		err = json.NewEncoder(w).Encode(formattedResponse)
		if common.HandleErr(err, w, http.StatusInternalServerError) {
			return
		}
	} else {
		http.Error(w, common.MethodAllowed("GET"), http.StatusMethodNotAllowed)
		return
	}

}
