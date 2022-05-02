package consts

// Environmental variable key
const ENVK_PORT = "PORT"

// Port values
const DEFAULT_PORT = "8080"
const CURRENT_PORT = "DEFAULT" // Set to "DEFAULT" for default port

// Port Messages
const PORT_LISTEN = "Listening on port "
const PORT_SET = "$PORT has been set: "
const PORT_DEFAULT = "Default port has been set: "

// default path for endpoint
const CASES_PATH = "/corona/" + APP_VERSION + "/cases/"
const POLICY_PATH = "/corona/" + APP_VERSION + "/policy/"
const STATUS_PATH = "/corona/" + APP_VERSION + "/status/"
const NOTIFICATIONS_PATH = "/corona/" + APP_VERSION + "/notifications/"

// stub path
const STUBBING = "/corona/" + APP_VERSION + "/stubbing/"

// Alpha 3 compare msgs
const COMPARE_A3_CASES_NOT_FOUND = "No Inconsistancies found between Alpha 3 library and Cases dependancy"
const COMPARE_A3_CASES_FOUND = "Inconsistancies found between Alpha 3 library and Cases dependancy:"
const COMPARE_A3_CASES_HEADER = ">> FOUND \t\t( %d ):\n"
const COMPARE_A3_CASES_IDENTIFIER = ">" + COMPARE_SHIFT + "Country(missing):\t[%s]\n"
const COMPARE_A3_CASES_BODY = ">" + COMPARE_SHIFT + "Incorrect amount:\t[%d/%d]\n>" + COMPARE_SHIFT + "Correct amount:\t[%d/%d]\nNB! Check (globals/alpha3.json) Alpha3 Local dependancy for inconsistencies\n"
const COMPARE_SHIFT = "  "

// Cache timer
const DAYS = 24
const CACHE_TIMER = 2 * DAYS
