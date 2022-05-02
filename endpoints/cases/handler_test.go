package cases_test

import (
	consts "assignment-2/constants"
	"assignment-2/endpoints/cases"
	glob "assignment-2/globals"
	"assignment-2/globals/common"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerCases(t *testing.T) {
	// Unset Invocations and Caching
	glob.AllowInvocations = false
	glob.AllowCaching = false
	// Set up subtests
	subtests := []struct {
		name         string
		graphql_mock func(name string, url string, body string) (map[string]interface{}, error)
		path         string
		expected     string
		method       string
	}{
		{ // The "normal" path, a successful request.
			name: "normal path",
			graphql_mock: func(name string, url string, body string) (map[string]interface{}, error) {
				var resp map[string]interface{}
				json.Unmarshal([]byte(`{
					"country": {
						"name": "Norway",
						"mostRecent": {
							"date": "2022-04-06",
							"confirmed": 1412969,
							"deaths": 2667,
							"recovered": 0,
							"growthRate": 0.001005277886011831
						}
					}
				}
`), &resp)
				return resp, nil
			},
			path:     consts.CASES_PATH + "Norway",
			expected: `{"country":"Norway","date":"2022-04-06","confirmed":1412969,"recovered":0,"deaths":2667,"growth_rate":0.001005277886011831}`,
			method:   http.MethodGet,
		},
		{ // User typed an invalid country name or code
			name: "invalid country",
			graphql_mock: func(name string, url string, body string) (map[string]interface{}, error) {
				return map[string]interface{}{"country": nil}, errors.New("graphql: Couldn't find data from country Åland Islands")
			},
			path:     consts.CASES_PATH + "ALA",
			expected: `graphql: Couldn't find data from country Åland Islands`,
			method:   http.MethodGet,
		},
		{ // User didn't input a country name or code
			name: "no country",
			graphql_mock: func(name string, url string, body string) (map[string]interface{}, error) {
				return map[string]interface{}{"country": nil}, errors.New("Invalid country name")
			},
			path:     consts.CASES_PATH + "",
			expected: `No country name inputted`,
			method:   http.MethodGet,
		},
		{ // User used wrong method
			name: "wrong method",
			graphql_mock: func(name string, url string, body string) (map[string]interface{}, error) {
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
			cases.Url = consts.CASES_PATH
			origin := common.GetGraphql
			glob.AllCountries = append(glob.AllCountries, []glob.Countries{
				{Name: "Norway", Code: "NOR"},
				{Name: "Åland Islands", Code: "ALA"},
			}...)
			cases.GetRequest = subtest.graphql_mock

			// Send request
			req := httptest.NewRequest(subtest.method, subtest.path, nil)
			// Setup Response
			w := httptest.NewRecorder()
			cases.HandlerCases(w, req)
			// Read result and save as data
			data, err := ioutil.ReadAll(w.Result().Body)
			if err != nil {
				t.Error("Error reading body of HandlerCases result", err)
			}

			strDat := string(data)[:len(string(data))-1] // Remove last character as it's a Line Break
			// compare data to expected value
			if strDat != subtest.expected {
				t.Errorf("Expected '%s' but got '%v'", subtest.expected, strDat)
			}
			// Un-mock
			cases.GetRequest = origin
			glob.AllCountries = []glob.Countries{}
			glob.AllowInvocations = true

		})
	}
}
