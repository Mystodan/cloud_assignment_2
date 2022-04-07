package policy

import (
	consts "assignment-2/constants"
	funcs "assignment-2/endpoints"
	glob "assignment-2/global_types"
	"fmt"
	"strconv"
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

func wrapData(data map[string]interface{}) glob.Policy {

	// Get all valid policies
	policies := getValidData(data, "policyActions", "policy_type_code")
	// Get 'stringencyData'
	data = data["stringencyData"].(map[string]interface{})
	// If there's no data, return a struct with value unavailable
	inn := "msg"
	if dataExists(data, inn) && data[inn].(string) == "Data unavailable" {
		return funcs.POLICY_UNAVAILABLE()
	}
	// Setter for stringency
	var actual float64
	actual = data["stringency"].(float64)
	// Check if stringency_acual exists
	inn = "stringency_actual"
	if dataExists(data, inn) {
		actual = data[inn].(float64)
	}
	// convert rest data to string
	stringency := fmt.Sprintf("%f", actual)
	policiesAmount := strconv.Itoa((len(policies)))

	// Otherwise, fill it with the data form
	return funcs.POLICY_AVAILABLE(
		data["country_code"].(string),
		data["date_value"].(string),
		stringency,
		policiesAmount)
}
