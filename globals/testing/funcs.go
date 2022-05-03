package testing_funcs

import (
	consts "assignment-2/constants"
	"assignment-2/endpoints/notifications"
	glob "assignment-2/globals"
	"encoding/json"
	"errors"
	"os"
)

/**
 *	Sets current application bound rules to be inactive/active
 *
 *	@param status - set status for the current rules.
 *
 */
func HandleAllRules(status bool) {
	glob.AllowFBCaching = status
	glob.AllowFBWebhooks = status
	glob.AllowInvocations = status
}

func Mocking_Policy(inn string, hasFile bool) func(name string, url string) (map[string]interface{}, error) {
	return func(name string, url string) (map[string]interface{}, error) {
		if hasFile {
			var resp map[string]interface{}
			jsonData, _ := os.ReadFile(inn)
			json.Unmarshal(jsonData, &resp)
			return resp, nil
		}
		var resp map[string]interface{}
		json.Unmarshal([]byte(inn), &resp)
		return resp, nil
	}
}
func Mocking_Case(inn string, hasBody bool) func(name string, url string, body string) (map[string]interface{}, error) {
	if hasBody {
		return func(name string, url string, body string) (map[string]interface{}, error) {
			var resp map[string]interface{}
			json.Unmarshal([]byte(inn), &resp)
			return resp, nil
		}
	}
	return func(name string, url string, body string) (map[string]interface{}, error) {
		return map[string]interface{}{"country": nil}, errors.New(inn)
	}
}

func resetRandomToken() {
	notifications.TOKEN_LENGTH = consts.APP_TOKEN_LENGTH
	notifications.TOKEN_SYMBOLS = consts.APP_TOKEN_SYMBOLS
}

func Mocking_Server_Webhook_ID() string {
	defer resetRandomToken()
	notifications.TOKEN_LENGTH = consts.TEST_SERVER_ID_LENGTH
	notifications.TOKEN_SYMBOLS = consts.TEST_SERVER_ID_SYMBOLS
	return notifications.GenerateRandomToken()
}
