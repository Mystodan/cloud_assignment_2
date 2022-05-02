package main

import (
	//constants
	constants "assignment-2/constants"
	globals "assignment-2/globals"
	"math/rand"
	"time"

	// endpoints
	status "assignment-2/endpoints/status"

	// server
	server "assignment-2/cmd/functions"
	"context"
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
	globals.PORT = server.SetPort(constants.CURRENT_PORT)
	// Firebase initialisation
	globals.Ctx = context.Background()

	// We use a service account, load credentials file that you downloaded from your project's settings menu.
	app := server.SetServiceAcc(globals.Ctx, constants.SERVICEKEY_PATH)

	// Instantiate client
	globals.Client = server.InstantiateFBClient(app, globals.Ctx)

	// Load all dependancies
	server.LoadAllDependancies()

	// Compare local Alpha 3 lib to Cases
	server.CompareLocalA3toCases()

	// Close down client
	defer server.CloseFB(globals.Client)

	// Routing endpoints
	server.RouteCases()
	server.RoutePolicy()
	server.RouteNotifications()
	server.RouteStatus()
	server.RouteStubbing()

	// listen for port
	server.SetListener(globals.PORT)
}
