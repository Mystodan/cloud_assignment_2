package stub

import (
	consts "assignment-2/constants"
	"encoding/json"
	"io/ioutil"
)

func LoadAllStubs() map[string]interface{} {
	//return values
	var getStubs map[string]interface{}
	// load from file
	readStubs, _ := ioutil.ReadFile(consts.STUB_FILEPATH)
	// read data as map[string]interface{}
	json.Unmarshal(readStubs, &getStubs)
	return getStubs
}
func getValue(inn string, want string, stub string, list map[string]interface{}) interface{} {
	getVal := list[stub].([]interface{})
	var retVal interface{}
	for _, val := range getVal {
		check := val.(map[string]interface{})
		if check[want] == inn {
			retVal = val
		}
	}
	return retVal
}
