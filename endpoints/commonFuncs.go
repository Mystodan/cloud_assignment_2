package funcs

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/machinebox/graphql"
)

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
	resp, err := http.Get(url)
	if err != nil {
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
