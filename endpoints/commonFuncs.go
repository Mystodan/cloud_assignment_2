package funcs

import (
	consts "assignment-2/constants"
	glob "assignment-2/global_types"
	"assignment-2/server/cache"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"unicode"

	"github.com/machinebox/graphql"
)

func POLICY_AVAILABLE(CountryCode string, Scope string, Stringency string, Policies string) glob.Policy {
	return glob.Policy{CountryCode, Scope, Stringency, Policies}
}
func CASE_AVAILABLE(Country string, Date string, Confirmed float64, Recovered float64, Deaths float64, Growth_rate float64) glob.Case {
	return glob.Case{Country, Date, Confirmed, Recovered, Deaths, Growth_rate}
}
func POLICY_UNAVAILABLE() glob.Policy {
	return POLICY_AVAILABLE(
		consts.APP_VALUE_UNAVAILABLE,
		consts.APP_VALUE_UNAVAILABLE,
		consts.APP_VALUE_UNAVAILABLE,
		consts.APP_VALUE_UNAVAILABLE,
	)
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

func LoadCountries() []glob.Countries {
	//return values
	var getAllCountries map[string]interface{}
	var setAllCountries []glob.Countries
	// load from file
	readCountries, _ := ioutil.ReadFile(consts.ALPHA3_PATH)
	// read data as map[string]interface{}
	json.Unmarshal(readCountries, &getAllCountries)
	getCountries := getAllCountries["countries"].([]interface{})
	for i := range getCountries {
		val := getCountries[i].(map[string]interface{})
		setAllCountries = append(setAllCountries, glob.Countries{
			val["code"].(string),
			val["name"].(string),
		})
	}
	return setAllCountries
}

func GetCountry(inn string) string {
	for _, val := range glob.AllCountries {
		if strings.EqualFold(val.Name, inn) || strings.EqualFold(val.Code, inn) {
			return val.Name
		}
	}
	return inn
}

func GetA3(inn string) string {
	for _, val := range glob.AllCountries {
		if strings.EqualFold(val.Name, inn) || strings.EqualFold(val.Code, inn) {
			return val.Code
		}
	}
	return inn
}

func DesensitizeString(inn string) string {
	inn = strings.ToLower(inn)
	r := []rune(inn)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

/**	checkError logs an error.
 *	@param inn - error value
 */
func checkError(inn error) bool {
	if inn != nil {
		log.Fatal(inn)
		return false
	}
	return true
}

/**	Get issues a GET to the specified URL.
 *	@param inn - URL
 */
func GetURL(inn string) (*http.Response, error) {
	return http.Get(inn)
}

/**	HandleErr logs an error.
 *	@param inn - error value
 */
func HandleErr(err error, w http.ResponseWriter, code int) bool {
	if err != nil {
		http.Error(w, err.Error(), code)
		return true
	}
	return false
}

func HandleURL(inn string) []string {
	return strings.Split(inn, "/")

}

func GetGraphql(url string, body string) (map[string]interface{}, error) {
	returnVal, err := cache.GetCache(url, body)
	if err != nil {
		// Send request to graphql api
		urlClientHandler := graphql.NewClient(url)
		urlRequestResponse := graphql.NewRequest(body)

		err = urlClientHandler.Run(context.Background(), urlRequestResponse, &returnVal)
		if err == nil { // If no errors
			cache.AddToCache(returnVal, url, body) // Adds to cache
		}
	}
	return returnVal, err
}

func RequestURL(url string) (map[string]interface{}, error) {
	returnVal, err := cache.GetCache(url)
	if !checkError(err) {
		// Send request to API
		resp, err := GetURL(url)
		if !checkError(err) {
			return map[string]interface{}{}, err
		} // Attempt to decode
		err = json.NewDecoder(resp.Body).Decode(&returnVal)
		if err != nil {
			return map[string]interface{}{}, err
		}
		cache.AddToCache(returnVal, url)
	}

	// Return
	return returnVal, err
}
