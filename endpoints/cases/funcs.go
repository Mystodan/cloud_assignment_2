package cases

import (
	consts "assignment-2/constants"
	"assignment-2/globals/common"
)

// Formats specific request for current api
func formatRequest(countryCode string) string {
	return common.FormatRequest(countryCode, "", consts.CASES_API)
}

// Wraps data into structs from map[string]interface{} into a Case struct
func wrapData(data map[string]interface{}) Case {
	data = (data["country"].(map[string]interface{}))
	mostRecentData := data["mostRecent"].(map[string]interface{})

	return Case{
		data["name"].(string),
		mostRecentData["date"].(string),
		mostRecentData["confirmed"].(float64),
		mostRecentData["recovered"].(float64),
		mostRecentData["deaths"].(float64),
		mostRecentData["growthRate"].(float64),
	}
}
