package stub

import (
	consts "assignment-2/constants"
	funcs "assignment-2/endpoints"
	"encoding/json"
	"net/http"
	"time"
)

var Timer time.Time

/**
 *	Handler for 'status' endpoint
 */
func HandlerStub(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	// Handles the Url by splitting its value strating after the CASES_PATH
	urlSplit := funcs.SplitURL(consts.STUBBING, w, r)
	allStubs := LoadAllStubs()
	if r.Method == http.MethodGet {
		if urlSplit[0] == "" {
			http.Error(w, "For stubbing use either stubbing/cases, or stubbing/policy ", http.StatusBadGateway)
			return
		} else if urlSplit[0] == "cases" {
			if len(urlSplit) > 1 && urlSplit[1] != "" {
				getCase := getValue(funcs.GetCountry(urlSplit[1]), "country", "cases_stub", allStubs)
				err := json.NewEncoder(w).Encode(getCase)
				if funcs.HandleErr(err, w, http.StatusInternalServerError) {
					return
				}
			} else {
				err := json.NewEncoder(w).Encode(allStubs["cases_stub"])
				if funcs.HandleErr(err, w, http.StatusInternalServerError) {
					return
				}
			}
		} else if urlSplit[0] == "policy" {
			if len(urlSplit) > 1 && urlSplit[1] != "" {
				getPolicy := getValue(funcs.GetA3(urlSplit[1]), "country_code", "policy_stub", allStubs)
				err := json.NewEncoder(w).Encode(getPolicy)
				if funcs.HandleErr(err, w, http.StatusInternalServerError) {
					return
				}
			} else {
				err := json.NewEncoder(w).Encode(allStubs["policy_stub"])
				if funcs.HandleErr(err, w, http.StatusInternalServerError) {
					return
				}
			}
		} else {
			http.Error(w, "Method not allowed, use GET", http.StatusMethodNotAllowed)
			return
		}
	}
}
