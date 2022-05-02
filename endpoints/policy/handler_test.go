package policy_test

import (
	consts "assignment-2/constants"
	"assignment-2/endpoints/policy"
	glob "assignment-2/globals"
	"assignment-2/globals/common"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandlerPolicy(t *testing.T) {
	// Unset Invocations and Caching
	glob.AllowInvocations = false
	glob.AllowCaching = false
	// Set up subtests
	subtests := []struct {
		name      string
		http_mock func(name string, url string) (map[string]interface{}, error)
		path      string
		expected  string
		method    string
	}{
		{ // The "normal" path.
			name: "normal path",
			http_mock: func(name string, url string) (map[string]interface{}, error) {
				var resp map[string]interface{}
				jsonData, _ := os.ReadFile("mocked_data.json")
				json.Unmarshal(jsonData, &resp)
				return resp, nil
			},
			path:     consts.POLICY_PATH + "Norway?scope=2022-03-13",
			expected: `{"country_code":"NOR","scope":"2022-03-13","stringency":11.11,"policies":21}`,
			method:   http.MethodGet,
		},
		{ // No data for the given date+country.
			name: "no data available path",
			http_mock: func(name string, url string) (map[string]interface{}, error) {
				var resp map[string]interface{}
				json.Unmarshal([]byte(`{
					"policyActions": [
						{
							"policy_type_code": "NONE",
							"policy_type_display": "No data.  Data may be inferred for last 7 days.",
							"flag_value_display_field": "",
							"policy_value_display_field": "No data.  Data may be inferred for last 7 days.",
							"policyvalue": 0,
							"flagged": null,
							"notes": null
						}
					],
					"stringencyData": {
						"msg": "Data unavailable"
					}
				}`), &resp)
				return resp, nil
			},
			path:     consts.POLICY_PATH + "NOR?scope=2022-04-31", // At the time this is written, no data exists for this entry
			expected: `parsing time "2022-04-31": day out of range`,
			method:   http.MethodGet,
		},
		{ // User used wrong method
			name: "wrong method",
			http_mock: func(name string, url string) (map[string]interface{}, error) {
				return map[string]interface{}{"country": nil}, errors.New("Invalid country name")
			},
			path:     consts.CASES_PATH + "",
			expected: `Method not allowed, use GET`,
			method:   http.MethodDelete,
		},
	}

	// Test subtests
	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			// Mock values
			policy.Url = consts.POLICY_PATH
			origin := common.RequestURL
			glob.AllCountries = append(glob.AllCountries, []glob.Countries{
				{Name: "Norway", Code: "NOR"},
				{Name: "Åland Islands", Code: "ALA"},
			}...)
			policy.GetRequest = subtest.http_mock
			// Send request
			req := httptest.NewRequest(subtest.method, subtest.path, nil)
			// Setup Response
			w := httptest.NewRecorder()
			policy.HandlerPolicy(w, req)

			// Read result and save as data
			res := w.Result()
			defer res.Body.Close()

			data, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Error("Error reading body of HandlerPolicy result", err)
			}
			strDat := string(data)[:len(string(data))-1] // Remove last character as it's a Line Break
			// compare data to expected value
			if strDat != subtest.expected {
				t.Errorf("Expected '%s' but got '%v'", subtest.expected, strDat)
			}

			// Un-mock
			policy.GetRequest = origin
			glob.AllowInvocations = true
		})
	}
}
