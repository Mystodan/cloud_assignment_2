package main

import (
	//constants
	consts "assignment-2/constants"

	glob "assignment-2/global_types"
	"math/rand"
	"time"

	// endpoints
	cases "assignment-2/endpoints/cases"
	notifications "assignment-2/endpoints/notifications"
	policy "assignment-2/endpoints/policy"
	status "assignment-2/endpoints/status"
	stub "assignment-2/endpoints/testing"

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
	rand.Seed(time.Now().UTC().UnixNano())
	status.Timer = time.Now()
	// sets port to defalt
	port := server.SetPort(consts.CURRENT_PORT)
	// Firebase initialisation
	glob.Ctx = context.Background()

	// We use a service account, load credentials file that you downloaded from your project's settings menu.
	app := server.SetServiceAcc(glob.Ctx, consts.SERVICEKEY_PATH)

	// Instantiate client
	glob.Client = server.InstantiateFBClient(app, glob.Ctx)

	// Load all dependancies
	server.LoadAllDependancies()

	// Close down client
	defer server.CloseFB(glob.Client)

	// Routing endpoints
	http.HandleFunc(consts.CASES_PATH, cases.HandlerCases)
	http.HandleFunc(consts.POLICY_PATH, policy.HandlerPolicy)
	http.HandleFunc(consts.STATUS_PATH, status.HandlerStatus)
	http.HandleFunc(consts.NOTIFICATIONS_PATH, notifications.HandlerNotifications)
	http.HandleFunc(consts.STUBBING, stub.HandlerStub)

	// listen for port
	server.SetListener(port)
}
