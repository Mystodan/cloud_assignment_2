package policy

import (
	consts "assignment-2/constants"
	funcs "assignment-2/endpoints"
	"assignment-2/endpoints/notifications"
	"encoding/json"
	"net/http"
	"time"
)

/**
 *	Handler for 'policy' endpoint
 */
func HandlerPolicy(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	// Handles the Url by splitting its value strating after the CASES_PATH
	urlSplit := funcs.SplitURL(consts.POLICY_PATH, w, r)
	urlQuery := r.URL.Query()

	if r.Method == http.MethodGet {

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

			if _, err := time.Parse(consts.POLICY_DATE, optParam); funcs.HandleErr(err, w, http.StatusBadRequest) {
				return
			}
		} else {
			optParam = time.Now().AddDate(0, 0, -3).Format(consts.POLICY_DATE)
		}

		// convert to A3 code
		country = funcs.GetA3(country)

		// Send request to api
		getRequest, err := funcs.RequestURL(formatRequest(country, optParam))
		if err != nil {
			http.Error(w, "Error sending request to covidtracker API", http.StatusFailedDependency)
			return
		}
		// wrap response
		formattedResponse := wrapData(getRequest)

		// invoke webhooks, annd send to writer
		notifications.SetInvocation(funcs.GetCountry(country))
		// send to writer
		err = json.NewEncoder(w).Encode(formattedResponse)
		if funcs.HandleErr(err, w, http.StatusInternalServerError) {
			return
		}
	} else {
		http.Error(w, "Method not allowed, use GET", http.StatusMethodNotAllowed)
		return
	}

}
