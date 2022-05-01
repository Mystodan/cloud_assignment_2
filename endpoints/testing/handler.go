package stub

import (
	consts "assignment-2/constants"
	"assignment-2/globals/common"

	"encoding/json"
	"net/http"
	"time"
)

var Timer time.Time
var Url string

/**
 *	Handler for 'status' endpoint
 */
func HandlerStub(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	// Handles the Url by splitting its value strating after the CASES_PATH
	urlSplit := common.SplitURL(Url, w, r)
	allStubs := LoadAllStubs()
	if r.Method == http.MethodGet {
		if urlSplit[0] == "" {
			http.Error(w, consts.STUBBING_ERR, http.StatusBadGateway)
			return
		} else if urlSplit[0] == "cases" {
			if len(urlSplit) > 1 && urlSplit[1] != "" {
				DesensitizeString, err := common.GetCountry(urlSplit[1])
				if common.HandleErr(err, w, http.StatusNotAcceptable) {
					return
				}
				getCase := getValue(DesensitizeString, "country", "cases_stub", allStubs)
				err = json.NewEncoder(w).Encode(getCase)
				if common.HandleErr(err, w, http.StatusInternalServerError) {
					return
				}
			} else {
				err := json.NewEncoder(w).Encode(allStubs["cases_stub"])
				if common.HandleErr(err, w, http.StatusInternalServerError) {
					return
				}
			}
		} else if urlSplit[0] == "policy" {
			if len(urlSplit) > 1 && urlSplit[1] != "" {
				DesensitizeString, err := common.GetA3(urlSplit[1])
				if common.HandleErr(err, w, http.StatusNotAcceptable) {
					return
				}
				getPolicy := getValue(DesensitizeString, "country_code", "policy_stub", allStubs)
				err = json.NewEncoder(w).Encode(getPolicy)
				if common.HandleErr(err, w, http.StatusInternalServerError) {
					return
				}
			} else {
				err := json.NewEncoder(w).Encode(allStubs["policy_stub"])
				if common.HandleErr(err, w, http.StatusInternalServerError) {
					return
				}
			}
		} else {
			http.Error(w, common.MethodAllowed("GET"), http.StatusMethodNotAllowed)
			return
		}
	}
}
