package notifications_test

import (
	consts "assignment-2/constants"
	"assignment-2/endpoints/notifications"
	"assignment-2/endpoints/policy"
	glob "assignment-2/globals"
	"assignment-2/globals/common"
	testfuncs "assignment-2/globals/testing"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestHandlerNotifications(t *testing.T) {
	// Unset Invocations and Caching
	testfuncs.HandleAllRules(false)

	// Set up subtests
	subtests := []struct {
		name     string
		path     string
		expected string
		method   string
	}{
		{ // The "default" path.
			name:     consts.TEST_DEFAULT_PATH,
			path:     consts.NOTIFICATIONS_PATH,
			expected: consts.MOCKED_WEBHOOKS,
			method:   http.MethodGet,
		},
		{ // The "default+" path.
			name:     consts.TEST_DEFAULT_PATH,
			path:     consts.NOTIFICATIONS_PATH + "nwCmUMkVLPvQn",
			expected: `{"webhook_id":"nwCmUMkVLPvQn","url":"https://webhook.site/e3a17245-2081-4aeb-a113-223b83960184","country":"Japan","calls":5}`,
			method:   http.MethodGet,
		},
		{ // The "default-" path.
			name:     consts.TEST_DEFAULT_PATH + "-",
			path:     consts.NOTIFICATIONS_PATH + "nwCmUMkVLPvQn",
			expected: `No Webhooks registered yet`,
			method:   http.MethodGet,
		},
		{ // The "wrong method" path.
			name:     "wrong method",
			path:     consts.NOTIFICATIONS_PATH,
			expected: "Method not allowed, use GET, POST or DELETE",
			method:   http.MethodPatch,
		},
		{ // The "post" path.
			name:     "post method",
			path:     consts.NOTIFICATIONS_PATH,
			expected: `{"webhook_id":"` + consts.MOCKED_WEBHOOK_ID + `"}`,
			method:   http.MethodPost,
		},
		{ // The wrong "post" path.
			name:     "wrong post method",
			path:     consts.NOTIFICATIONS_PATH,
			expected: `{"webhook_id":"` + consts.MOCKED_WEBHOOK_ID + `"}`,
			method:   http.MethodPost,
		},
		{ // The "delete existing" path.
			name:     "post method",
			path:     consts.NOTIFICATIONS_PATH + "nwCmUMkVLPvQn",
			expected: `nwCmUMkVLPvQn - Webhook was deleted`,
			method:   http.MethodDelete,
		},
		{ // The "delete, not exist" path.
			name:     "post method",
			path:     consts.NOTIFICATIONS_PATH + "nwCUMMkVLPvQn",
			expected: `nwCUMMkVLPvQn - Webhook does not exist`,
			method:   http.MethodDelete,
		},
	}

	// Test subtests
	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			// Mock values
			notifications.Url = consts.NOTIFICATIONS_PATH
			origin := common.RequestURL
			notifications.Gen_Token = func() string { return consts.MOCKED_WEBHOOK_ID }

			glob.AllCountries = append(glob.AllCountries, []glob.Countries{
				{Name: "Norway", Code: "NOR"},
				{Name: "Japan", Code: "JPN"},
			}...)
			notifications.Webhook_server_id = testfuncs.Mocking_Server_Webhook_ID()
			if subtest.name != consts.TEST_DEFAULT_PATH+"-" {
				resp, _ := os.ReadFile(consts.TEST_DEFAULT_MOCK_JSON)
				json.Unmarshal(resp, &glob.AllWebhooks)
			}
			// Send request
			var req = httptest.NewRequest(subtest.method, subtest.path, nil)

			if subtest.method == http.MethodPost {
				bodyToSend := consts.TEST_NOTIFICATIONS
				if subtest.name == "wrong post method" {
					bodyToSend = `{"OWO":"UWU","country":"JPN"}`
				}
				req = httptest.NewRequest(subtest.method, subtest.path, strings.NewReader(bodyToSend))
			}
			// Setup Response
			w := httptest.NewRecorder()

			notifications.HandlerNotifications(w, req)
			// Read result and save as data
			res := w.Result()
			defer res.Body.Close()

			data, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Error("Error reading body of HandlerPolicy result", err)
			}
			var sens = 1
			if subtest.method == http.MethodDelete || subtest.name == consts.TEST_DEFAULT_PATH+"-" {
				sens = 0
			}
			strDat := string(data)[:len(string(data))-sens]
			// compare data to expected value
			if strDat != subtest.expected {
				t.Errorf("Expected '%s'\n but got '%v'", subtest.expected, strDat)
			}

			// Un-mock
			glob.AllWebhooks = map[string]glob.Webhook{}
			policy.GetRequest = origin
			glob.AllCountries = []glob.Countries{}

		})
	}
	testfuncs.HandleAllRules(true)
}
