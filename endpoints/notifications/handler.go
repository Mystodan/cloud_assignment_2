package notifications

import (
	consts "assignment-2/constants"
	funcs "assignment-2/endpoints"
	glob "assignment-2/global_types"
	"encoding/json"
	"net/http"
)

/**
 *	Handler for 'notifications' endpoint
 */
func HandlerNotifications(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	// Handles the Url by splitting its value strating after the NOTIFICATION_PATH
	urlSplit := funcs.SplitURL(consts.NOTIFICATIONS_PATH, w, r)

	switch r.Method {

	case http.MethodPost:
		{
			webHook := glob.Webhook{}
			err := json.NewDecoder(r.Body).Decode(&webHook)
			if funcs.HandleErr(err, w, http.StatusBadRequest) {
				return
			}
			// create a new token
			webHook.ID = handleNewToken()
			// sends new webhook to firestore
			id, err := sendWebhookToFB(webHook)
			if funcs.HandleErr(err, w, http.StatusInternalServerError) {
				return
			}
			// saves webhook to local storage
			glob.AllWebhooks[id] = webHook
			// outputs
			err = json.NewEncoder(w).Encode(map[string]string{"webhook_id": webHook.ID})
			if funcs.HandleErr(err, w, http.StatusInternalServerError) {
				return
			}
		}

	case http.MethodDelete:
		{
			// Check that the args are valid
			if len(urlSplit) < 1 || urlSplit[0] == "" {
				http.Error(w, "missing webhook(id) for deletion", http.StatusBadRequest)
				return
			}
			// Delete webhook from database
			deleted, delErr := DeleteWebhook(urlSplit[0], &glob.AllWebhooks)
			if funcs.HandleErr(delErr, w, http.StatusInternalServerError) {
				return
			}
			err := json.NewEncoder(w).Encode(urlSplit[0] + handleDeleted(deleted))
			if funcs.HandleErr(err, w, http.StatusInternalServerError) {
				return
			}

		}

	case http.MethodGet:
		{
			// Getting ALL registered webhooks
			if len(urlSplit) < 1 || len(urlSplit[0]) < 1 {
				err := json.NewEncoder(w).Encode(glob.AllWebhooks)
				if funcs.HandleErr(err, w, http.StatusInternalServerError) {
					return
				}
				// Getting a specific webhook
			} else {
				ID := urlSplit[0]
				webhook, err := GetWebhook(ID, glob.AllWebhooks)
				if err != nil {
					if funcs.HandleErr(err, w, http.StatusInternalServerError) {
						return
					}
				}
				// outputs
				err = json.NewEncoder(w).Encode(webhook)
				if err != nil {
					if funcs.HandleErr(err, w, http.StatusInternalServerError) {
						return
					}
				}

			}
		}

	}
}
