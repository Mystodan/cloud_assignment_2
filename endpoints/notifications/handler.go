package notifications

import (
	consts "assignment-2/constants"
	glob "assignment-2/globals"
	"assignment-2/globals/common"
	"encoding/json"
	"net/http"
)

var (
	// necessary for routing
	Url string
	// necessary for mocking
	Webhook_server_id string
	Gen_Token         = GenerateRandomToken
	TOKEN_SYMBOLS     = consts.APP_TOKEN_SYMBOLS
	TOKEN_LENGTH      = consts.APP_TOKEN_LENGTH
)

/**
 *	Handler for 'notifications' endpoint
 */
func HandlerNotifications(w http.ResponseWriter, r *http.Request) {
	// Handles the Url by splitting its value strating after the NOTIFICATION_PATH
	urlSplit := common.SplitURL(Url, w, r)
	switch r.Method {
	case http.MethodPost:
		{
			w.Header().Add("content-type", "application/json")
			webHook := glob.Webhook{}
			err := json.NewDecoder(r.Body).Decode(&webHook)
			if common.HandleErr(err, w, http.StatusBadRequest) {
				return
			}
			// create a new token
			webHook.ID = handleNewToken()
			country, err := common.GetCountry(webHook.Country)
			webHook.Country = country.Name
			if common.HandleErr(err, w, http.StatusInternalServerError) {
				return
			}
			// sends new webhook to firestore
			if glob.AllowFBWebhooks {
				id, err := SendWebhookToFB(webHook)
				Webhook_server_id = id
				if common.HandleErr(err, w, http.StatusInternalServerError) {
					return
				}
			}
			// saves webhook to local storage
			glob.AllWebhooks[Webhook_server_id] = webHook
			// outputs
			w.WriteHeader(http.StatusCreated)
			err = json.NewEncoder(w).Encode(map[string]string{"webhook_id": webHook.ID})
			if common.HandleErr(err, w, http.StatusInternalServerError) {
				return
			}
		}

	case http.MethodDelete:
		{
			// Check that the args are valid
			if len(urlSplit) < 1 || urlSplit[0] == "" {
				http.Error(w, consts.INPUT_NOT_FOUND, http.StatusNotAcceptable)
				return
			}
			// Delete webhook from database

			deleted, delErr := DeleteWebhooks(urlSplit[0], &glob.AllWebhooks)
			if common.HandleErr(delErr, w, http.StatusInternalServerError) {
				return
			} // Handles status messages on deleted or not
			if !deleted {
				w.WriteHeader(http.StatusNotAcceptable)
			} else {
				w.WriteHeader(http.StatusOK)
			}
			w.Write([]byte(urlSplit[0] + handleDeleted(deleted)))
			//err := json.NewEncoder(w).Encode(urlSplit[0] + handleDeleted(deleted)) -- DEPRECATED
		}

	case http.MethodGet:
		{
			w.Header().Add("content-type", "application/json")
			if len(glob.AllWebhooks) > 0 {
				// Getting ALL registered webhooks
				if len(urlSplit) < 1 || len(urlSplit[0]) < 1 {
					err := json.NewEncoder(w).Encode(glob.AllWebhooks)
					if common.HandleErr(err, w, http.StatusInternalServerError) {
						return
					}
					// Getting a specific webhook
				} else {
					ID := urlSplit[0]
					webhook, err := GetWebhook(ID, glob.AllWebhooks)
					if err != nil {
						if common.HandleErr(err, w, http.StatusInternalServerError) {
							return
						}
					}
					// outputs
					err = json.NewEncoder(w).Encode(webhook)
					if err != nil {
						if common.HandleErr(err, w, http.StatusInternalServerError) {
							return
						}
					}

				}
			} else {
				w.WriteHeader(http.StatusNoContent)
				w.Write([]byte("No Webhooks registered yet"))
			}
		}
	default:
		{
			http.Error(w, common.MethodAllowed("GET, POST or DELETE"), http.StatusMethodNotAllowed)
			return
		}
	}
}
