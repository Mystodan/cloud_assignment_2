package main

import (
	//constants
	consts "assignment-2/constants"
	glob "assignment-2/globals"
	"time"

	// endpoints
	cases "assignment-2/endpoints/cases"
	notifications "assignment-2/endpoints/notifications"
	policy "assignment-2/endpoints/policy"
	status "assignment-2/endpoints/status"

	// server
	server "assignment-2/server/functions"
	"context"
	"net/http"
	// Firebase Dependancies
)

/**
 *	The main function.
 *	Creates a link/routes to the correct handlers
 */
func main() {
	status.Timer = time.Now()

	port := server.SetPort(consts.DEFAULT_PORT)
	// Firebase initialisation
	glob.Ctx = context.Background()

	// We use a service account, load credentials file that you downloaded from your project's settings menu.
	app := server.SetServiceAcc(glob.Ctx, "serviceKey/serviceAccountKey.json")

	// Instantiate client
	glob.Client = server.InstantiateFBClient(app, glob.Ctx)

	// Load all webhooks
	notifications.SetAllTokens(notifications.LoadWebhooksFromFB())

	// Close down client
	defer server.CloseFB(glob.Client)

	// Routing endpoints
	http.HandleFunc(consts.CASES_PATH, cases.HandlerCases)
	http.HandleFunc(consts.POLICY_PATH, policy.HandlerPolicy)
	http.HandleFunc(consts.STATUS_PATH, status.HandlerStatus)
	http.HandleFunc(consts.NOTIFICATIONS_PATH, notifications.HandlerNotifications)

	// listen for port
	server.SetListener(port)
}
