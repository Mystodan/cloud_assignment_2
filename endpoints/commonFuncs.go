package funcs

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"unicode"

	"github.com/machinebox/graphql"
)

func DesensitizeString(inn string) string {
	inn = strings.ToLower(inn)
	r := []rune(inn)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func SplitURL(path string, w http.ResponseWriter, r *http.Request) []string {

	// Handles the Url by splitting its value strating after the CASES_PATH
	urlSplit := HandleURL(r.URL.EscapedPath()[len(path):])
	// Check if the user input enough args
	if len(urlSplit) < 1 {
		http.Error(w, "Not enough arguments, see documentation", http.StatusBadRequest)
		return nil
	}
	return urlSplit
}

/**	checkError logs an error.
 *	@param inn - error value
 */
func checkError(inn error) bool {
	if inn != nil {
		log.Fatal(inn)
		return true
	}
	return false
}

/**	Get issues a GET to the specified URL.
 *	@param inn - URL
 */
func GetURL(inn string) (*http.Response, bool, error) {
	ret, err := http.Get(inn)
	return ret, checkError(err), err
}

/**	HandleErr logs an error.
 *	@param inn - error value
 */
func HandleErr(inn error, w http.ResponseWriter, code int) {
	if inn != nil {
		http.Error(w, inn.Error(), code)
		return
	}
}

func HandleURL(inn string) []string {
	return strings.Split(inn, "/")

}

func GetGraphql(url string, body string) (map[string]interface{}, error) {
	var returnVal map[string]interface{}
	// Send request to graphql api
	urlClientHandler := graphql.NewClient(url)
	urlRequestResponse := graphql.NewRequest(body)

	err := urlClientHandler.Run(context.Background(), urlRequestResponse, &returnVal)
	return returnVal, err
}

func HttpRequest(url string) (map[string]interface{}, error) {
	// Send request to covidtracker API
	resp, state, err := GetURL(url)
	if state {
		return map[string]interface{}{}, err
	}

	// Attempt to decode
	var v map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&v)
	if err != nil {
		return map[string]interface{}{}, err
	}

	// Return
	return v, err
}
