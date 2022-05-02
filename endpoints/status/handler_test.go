package status_test

import (
	consts "assignment-2/constants"
	"assignment-2/endpoints/status"
	glob "assignment-2/globals"
	testfuncs "assignment-2/globals/testing"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerStatus(t *testing.T) {
	// Unset Invocations and Caching
	testfuncs.HandleAllRules(false)

	// Set up subtests
	subtests := []struct {
		name     string
		expected string
		method   string
	}{
		{ // The "default" path, a successful request.
			name:     consts.TEST_DEFAULT_PATH,
			expected: consts.TEST_STATUS_DEFAULT,
			method:   http.MethodGet,
		},
		{ // INVALID METHOD
			name:     "invalid method",
			expected: `Method not allowed, use GET`,
			method:   http.MethodDelete,
		},
	}
	// Set Path
	path := consts.STATUS_PATH

	// Test subtests
	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			// Mock values

			// Send request
			req := httptest.NewRequest(subtest.method, path, nil)
			// Setup Response
			w := httptest.NewRecorder()
			status.HandlerStatus(w, req)
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
			glob.AllCountries = []glob.Countries{}
			testfuncs.HandleAllRules(true)

		})
	}
}
