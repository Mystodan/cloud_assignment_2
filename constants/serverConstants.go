package consts

// Environmental variable key
const ENVK_PORT = "PORT"

// Port values
const DEFAULT_PORT = "8080"
const CURRENT_PORT = "DEFAULT" // Set to "DEFAULT" for default port

// Port Messages
const PORT_SET = "$PORT has been set: "

const PORT_NOTSET = "$PORT has not been set. "
const PORT_INVALID = "Input contained illegal symbols. "
const PORT_DEFAULT = "Default port has been set: "

// default path for endpoint
const CASES_PATH = "/corona/" + APP_VERSION + "/cases/"
const POLICY_PATH = "/corona/" + APP_VERSION + "/policy/"
const STATUS_PATH = "/corona/" + APP_VERSION + "/status/"
const NOTIFICATIONS_PATH = "/corona/" + APP_VERSION + "/notifications/"

// stub path
const STUBBING = "/corona/" + APP_VERSION + "/stubbing/"

// Cache timer
const DAYS = 24
const CACHE_TIMER = 2 * DAYS
