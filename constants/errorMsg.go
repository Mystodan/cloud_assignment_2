package consts

// Port Error messages
const PORT_NOTSET = "$PORT has not been set. "
const PORT_INVALID = "Input contained illegal symbols. "

// Method Error messages
const METHOD_NOT_ALLOWED = "Method not allowed, use "

// Stubbing method err
const STUBBING_ERR = "For stubbing use either stubbing/cases, or stubbing/policy "

// Policy fetch error
const POLICY_VALUE_UNAVAILABLE = int64(-1)

// Alpha3 Error messages
const COUNTRY_NOT_VALID = "invalid countryname, or countrycode"
const COUNTRY_NOT_REGISTERED = "countryname has no assosiated countrycode"

// webhook Error messages
const INPUT_NOT_FOUND = "webhook (id): not inputted"
const TOKEN_DELETED_FOUND = " - Webhook was deleted"
const TOKEN_DELETED_NOT_FOUND = " - Webhook does not exist"