package cases

import (
	"fmt"
)

func formatRequest(inn string) string {
	return fmt.Sprintf(`
	query {
		country(name: "%s") {
			name
			mostRecent {
				date(format: "yyyy-MM-dd")
				confirmed
				recovered
				deaths
				growthRate
			}
		}
	}
	`, inn)
}

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
