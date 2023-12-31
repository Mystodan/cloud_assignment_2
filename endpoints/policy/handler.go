package policy

import (
	consts "assignment-2/constants"
	glob "assignment-2/globals"
	"assignment-2/globals/common"

	"assignment-2/endpoints/notifications"
	"encoding/json"
	"net/http"
	"time"
)

var (
	// necessary for routing
	Url string
	// necessary for mocking
	GetRequest = common.RequestURL
	// necessary for returning when scope returns unavailable data
	country string
	scope   string
)

/**
 *	Handler for 'policy' endpoint
 */
func HandlerPolicy(w http.ResponseWriter, r *http.Request) {
	// Handles the Url by splitting its value strating after the CASES_PATH
	urlSplit := common.SplitURL(Url, w, r)
	urlQuery := r.URL.Query()

	if r.Method == http.MethodGet {
		w.Header().Add("content-type", "application/json")
		country = urlSplit[0]
		if len(country) < 1 {
			http.Error(w, consts.NO_COUNTRY_INPUT, http.StatusBadRequest)
			return
		}
		if len(urlSplit) < 1 {
			http.Error(w, "Not enough parameters, try again", http.StatusLengthRequired)
			return
		}

		// get optional parameter
		var optParam string
		if _, isScope := urlQuery["scope"]; isScope {
			optParam = urlQuery["scope"][0]
			if _, err := time.Parse(consts.POLICY_DATE, optParam); common.HandleErr(err, w, http.StatusRequestedRangeNotSatisfiable) {
				return
			}
			scope = optParam
		} else {
			optParam = time.Now().AddDate(consts.POLICY_TIME_YEAR, consts.POLICY_TIME_MONTH, consts.POLICY_TIME_DAY).Format(consts.POLICY_DATE)
			scope = optParam
		}

		// convert to A3
		country, err := common.GetCountry(country)

		if common.HandleErr(err, w, http.StatusNotAcceptable) {
			return
		}
		// Send request to api
		getRequest, err := common.RequestURL(country.Code, formatRequest(country.Code, optParam))
		if err != nil {
			http.Error(w, consts.POLICY_API_REQUEST_ERR, http.StatusFailedDependency)
			return
		}
		// wrap response
		formattedResponse := wrapData(getRequest)

		// invoke webhooks on thread, and send to writer
		if glob.AllowInvocations {
			go notifications.SetInvocation(country.Name)
		}
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
