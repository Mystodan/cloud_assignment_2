package policy

import (
	consts "assignment-2/constants"
	"assignment-2/globals/common"
)

// Checks if data exists within map [string] interface{}
func dataExists(list map[string]interface{}, data string) bool {
	return list[data] != nil
}

// Formats specific request for current api
func formatRequest(countryCode string, date string) string {
	return common.FormatRequest(countryCode, date, consts.POLICY_API)
}

// Gets valid data from api
func getValidData(list map[string]interface{}, data string, data_code string) []interface{} {
	alldata := list[data].([]interface{}) // set up list with all data
	var retVal []interface{}
	for _, policies := range alldata { // iterates all data
		validPolicies := policies.(map[string]interface{})
		if validPolicies[data_code] != "NONE" { // check for valid data
			retVal = append(retVal, policies)
		}
	}
	return retVal
}

func wrapData(data map[string]interface{}) Policy {
	// Get all valid policies
	policies := getValidData(data, "policyActions", "policy_type_code")
	// Get 'stringencyData'
	data = data["stringencyData"].(map[string]interface{})
	// If there's no data, return a struct with value unavailable
	inn := "msg"
	if dataExists(data, inn) && data[inn].(string) == "Data unavailable" {
		return Policy{
			country,
			scope,
			float64(consts.POLICY_VALUE_UNAVAILABLE),
			consts.POLICY_VALUE_UNAVAILABLE,
		}
	}
	// Setter for stringency
	var actual float64
	actual = data["stringency"].(float64)
	// Check if stringency_acual exists
	inn = "stringency_actual"
	if dataExists(data, inn) {
		actual = data[inn].(float64)
	}

	// Otherwise, fill it with the data form
	return Policy{
		data["country_code"].(string),
		data["date_value"].(string),
		actual,
		int64(len(policies))}
}
