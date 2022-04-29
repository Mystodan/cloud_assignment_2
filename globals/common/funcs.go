package common

import (
	"assignment-2/cmd/cache"
	consts "assignment-2/constants"
	glob "assignment-2/globals"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"unicode"

	"github.com/machinebox/graphql"
)

func SplitURL(path string, w http.ResponseWriter, r *http.Request) []string {

	// Handles when path incorrecly adds "/" to the end of url
	Sensitivity := len(path)
	if Sensitivity > len(r.URL.EscapedPath()) {
		Sensitivity--
	}
	// Handles the Url by splitting its value strating after the CASES_PATH
	urlSplit := HandleURL(r.URL.EscapedPath()[Sensitivity:])
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
	log.Println("Loading Alpha3 library...")
	for i := range getCountries {
		val := getCountries[i].(map[string]interface{})
		setAllCountries = append(setAllCountries, glob.Countries{
			val["code"].(string),
			val["name"].(string),
		})
	}
	log.Println("Done!")
	return setAllCountries
}

func GetCountry(inn string) (string, error) {
	inn = strings.Replace(inn, "%20", " ", -1)
	for _, val := range glob.AllCountries {
		if strings.EqualFold(val.Name, inn) || strings.EqualFold(val.Code, inn) {
			return val.Name, nil
		}
	}
	return inn, errors.New(consts.CODE_NOT_REGISTERED)
}

func GetA3(inn string) string {
	for _, val := range glob.AllCountries {
		if strings.EqualFold(val.Name, inn) || strings.EqualFold(val.Code, inn) {
			return val.Code
		}
	}
	return inn
}

//Deprecated
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
		return false
	}
	log.Fatal(inn)
	return true
}

func MethodAllowed(method string) string {
	return consts.METHOD_NOT_ALLOWED + method
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

func GetGraphql(name string, url string, body string) (map[string]interface{}, error) {
	returnVal, err := cache.GetCache(url, body)
	if err != nil {
		// Send request to graphql api
		urlClientHandler := graphql.NewClient(url)
		urlRequestResponse := graphql.NewRequest(body)

		err = urlClientHandler.Run(context.Background(), urlRequestResponse, &returnVal)
		if err == nil { // If no errors
			cache.AddToCache(name, returnVal, url, body) // Adds to cache
		}
	}

	return returnVal, err
}

func RequestURL(name string, url string) (map[string]interface{}, error) {
	returnVal, err := cache.GetCache(url)

	if !checkError(err) {
		// Send request to API
		resp, _ := GetURL(url)

		// Attempt to decode
		err = json.NewDecoder(resp.Body).Decode(&returnVal)
		if err == nil {
			cache.AddToCache(name, returnVal, url)
		}
	}
	// Return
	return returnVal, err
}
