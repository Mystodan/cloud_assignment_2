package glob

/**
 *	Update globals/testing/funcs.go:5 (HandleAllRules())
 *	when adding more rules
 *	usecase: For use of mocking when testing
 */

// Rule 1: Allow Application to invoke webhooks *default: true
var AllowInvocations = true

// Rule 2: Allow Application to store cache *default: true
var AllowFBCaching = true

// Rule 3: Allow Application to store webhooks *default: true
var AllowFBWebhooks = true
