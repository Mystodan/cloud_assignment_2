package status

import (
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
func HandlerStatus(w http.ResponseWriter, r *http.Request) {
	// Handles the Url by splitting its value strating after the CASES_PATH
	urlSplit := common.SplitURL(Url, w, r)
	if r.Method == http.MethodGet {
		w.Header().Add("content-type", "application/json")
		if urlSplit[0] == "" {
			err := json.NewEncoder(w).Encode(getAPIstatus())
			if common.HandleErr(err, w, http.StatusInternalServerError) {
				return
			}
		}
	} else {
		http.Error(w, common.MethodAllowed("GET"), http.StatusMethodNotAllowed)
		return
	}

}
