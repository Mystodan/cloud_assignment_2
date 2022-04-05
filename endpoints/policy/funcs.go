package policy

import (
	consts "assignment-2/constants"
	"fmt"
)

func dataExists(list map[string]interface{}, data string) bool {
	return list[data] != nil
}

func formatRequest(countryCode string, date string) string {
	return fmt.Sprintf(consts.POLICY_API+"%s/%s/", countryCode, date)
}

func getValidData(list map[string]interface{}, data string, data_code string) []interface{} {
	alldata := list[data].([]interface{})
	var retVal []interface{}
	for _, policies := range alldata {
		validPolicies := policies.(map[string]interface{})
		if validPolicies[data_code] != "NONE" {
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

	// If there's no data, return an empty struct
	inn := "msg"
	if dataExists(data, inn) && data[inn].(string) == "Data unavailable" {
		return Policy{}
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
		len(policies),
	}
}
