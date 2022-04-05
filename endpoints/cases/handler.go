package cases

import (
	consts "assignment-2/constants"
	funcs "assignment-2/endpoints"
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
		country = funcs.DesensitizeString(country)

		getGraphql, err := funcs.GetGraphql("https://covid19-graphql.vercel.app/", formatRequest(country))
		funcs.HandleErr(err, w, http.StatusBadRequest)

		// wrap response
		formattedResponse := wrapData(getGraphql)

		// send to writer
		err = json.NewEncoder(w).Encode(formattedResponse)
		funcs.HandleErr(err, w, http.StatusInternalServerError)

	} else {
		http.Error(w, "Method not allowed, use GET", http.StatusMethodNotAllowed)
		return
	}

}
