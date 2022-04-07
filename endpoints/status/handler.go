package status

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
func HandlerStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	// Handles the Url by splitting its value strating after the CASES_PATH
	urlSplit := funcs.SplitURL(consts.STATUS_PATH, w, r)
	if r.Method == http.MethodGet {
		if urlSplit[0] == "" {
			err := json.NewEncoder(w).Encode(getAPIstatus())
			if funcs.HandleErr(err, w, http.StatusInternalServerError) {
				return
			}
		}
	} else {
		http.Error(w, "Method not allowed, use GET", http.StatusMethodNotAllowed)
		return
	}

}
