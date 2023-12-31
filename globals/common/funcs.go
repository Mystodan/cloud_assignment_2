package common

import (
	"assignment-2/cmd/cache"
	consts "assignment-2/constants"
	glob "assignment-2/globals"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"unicode"

	"github.com/machinebox/graphql"
)

// Function which splits url into readable slices
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

// Handler for formatting request based on api
func FormatRequest(countryVal string, date string, requestType string) string {
	switch requestType {
	case consts.CASES_API: // cases api request
		return fmt.Sprintf(consts.CASES_REQUEST, countryVal)
	case consts.POLICY_API: // policy api request
		return fmt.Sprintf(consts.POLICY_API+"%s/%s/", countryVal, date)
	default:
		return ""
	}
}

// Reads from Cases api for country names and compares with local Alpha3 library
func CompareGraphCountryNames() ([]string, bool, int, int) {
	// reads from cases api
	graphqlClient := graphql.NewClient(consts.CASES_API)
	graphqlRequest := graphql.NewRequest(consts.CASES_GET_ALL)
	var graphqlResponse map[string]interface{} // gets data from response
	graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse)
	var serverCountryNames []string                      // stores all country names from api
	data := graphqlResponse["countries"].([]interface{}) // unwraps response into readable data
	for _, server_Name := range data {                   // iterates and reads data
		var shouldAppend = true                                                     // handles append
		server_CountryName := server_Name.(map[string]interface{})["name"].(string) // gets country name from api
		for _, localCountry := range glob.AllCountries {                            // iterates local country names from alpha 3
			if localCountry.Name == server_CountryName { // checks for consistencies
				shouldAppend = false // appends only inconsistencies
				break                // breaks out of loop since the data already exists in local database
			}
		}
		if shouldAppend { // Appends inconstency
			serverCountryNames = append(serverCountryNames, server_Name.(map[string]interface{})["name"].(string))
		}
	}
	if len(serverCountryNames) > 0 { // check for name
		return serverCountryNames, false, len(serverCountryNames), len(data)
	} else { // check for no name
		return serverCountryNames, true, len(serverCountryNames), len(data)
	}
}

// Loads all data from local alpha 3 database/library/dependancy
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
		var value glob.Countries
		value.Code = val["code"].(string)
		value.Name = val["name"].(string)
		setAllCountries = append(setAllCountries, value)
	}
	log.Println("Done!")
	return setAllCountries
}

// Function which reads bad symbols and converts them into readable data for handler
func fixSymbols(inn string) string {
	inn = strings.Replace(inn, "%20", " ", -1)
	inn = strings.Replace(inn, "%C3%85", "Å", -1)
	inn = strings.Replace(inn, "%C3%A5", "å", -1)
	inn = strings.Replace(inn, "%C3%86", "Æ", -1)
	inn = strings.Replace(inn, "%C3%A6", "æ", -1)
	inn = strings.Replace(inn, "%C3%98", "Æ", -1)
	inn = strings.Replace(inn, "%C3%B8", "ø", -1)
	return inn
}

// function which gets a certain country(code and name)
func GetCountry(inn string) (glob.Countries, error) {
	if inn != "_None" { // checks for non-Alpha3 variables
		inn = fixSymbols(inn)                   // fixes bad unreadable symbols
		for _, val := range glob.AllCountries { //iterates current buffer
			if strings.EqualFold(val.Name, inn) || strings.EqualFold(val.Code, inn) {
				if val.Code == "_None" && strings.EqualFold(val.Code, inn) {
					return val, errors.New(consts.COUNTRY_NOT_REGISTERED) //not registered
				}
				return val, nil // If no problems return correct country
			}
		}
	}
	fmt.Println(len(glob.AllCountries))
	var retVal glob.Countries
	retVal.Code = inn
	retVal.Name = inn // returns invalid country and error
	return retVal, errors.New(consts.COUNTRY_NOT_VALID)
}

//Deprecated
func GetA3(inn string) (string, error) {
	if inn != "_None" {
		inn = strings.Replace(inn, "%20", " ", -1)
		for _, val := range glob.AllCountries {
			if strings.EqualFold(val.Name, inn) || strings.EqualFold(val.Code, inn) {
				if val.Code == "_None" {
					return val.Code, errors.New(consts.COUNTRY_NOT_REGISTERED)
				}
				return val.Code, nil
			}
		}
	}
	return inn, errors.New(consts.COUNTRY_NOT_VALID)
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

// function which makes writing errors easier
func MethodAllowed(method string) string {
	return consts.METHOD_NOT_ALLOWED + method
}

/**	Get issues a GET to the specified URL.(for caching :))
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

// splits url on "/"
func HandleURL(inn string) []string {
	return strings.Split(inn, "/")
}

// Graphql request for cases api
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

// http request handling for policy api
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
