package main

import (
	//constants
	consts "assignment-2/constants"

	glob "assignment-2/globals"
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
	server.RouteCases()
	server.RoutePolicy()
	server.RouteNotifications()
	server.RouteStatus()

	// listen for port
	server.SetListener(port)
}
