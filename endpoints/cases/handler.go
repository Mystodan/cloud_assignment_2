package cases

import (
	consts "assignment-2/constants"
	funcs "assignment-2/endpoints"
	"encoding/json"
	"net/http"
	"unicode"
)

/**
 *	Handler for 'cases' endpoint
 */
func HandlerCases(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	// Handles the Url by splitting its value strating after the CASES_PATH
	urlSplit := funcs.HandleURL(r.URL.EscapedPath()[len(consts.CASES_PATH):])

	// Check if the user input enough args
	if len(urlSplit) < 1 {
		http.Error(w, "Not enough arguments, see documentation", http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodGet {
		country := urlSplit[0]
		if len(country) < 1 {
			http.Error(w, "No country name inputted", http.StatusBadRequest)
			return
		}
		r := []rune(country)
		r[0] = unicode.ToUpper(r[0])
		country = string(r)

		getGraphql, err := funcs.GetGraphql("https://covid19-graphql.vercel.app/", reqBody(country))
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
