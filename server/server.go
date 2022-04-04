package main

import (
	consts "assignment-2/constants"
	funcs "assignment-2/functions"
	"context"

	"cloud.google.com/go/firestore"
	// Firebase Dependancies
)

// Firebase context and client used by Firestore functions throughout the program.
var ctx context.Context
var client *firestore.Client

// Collection name in Firestore
const collection = "messages"

/**
 *	The main function.
 *	Creates a link/routes to the correct handlers
 */
func main() {
	funcs.TimerStart(funcs.Timer)
	port := funcs.SetPort(consts.DEFAULT_PORT)
	// Firebase initialisation
	ctx = context.Background()

	// We use a service account, load credentials file that you downloaded from your project's settings menu.
	app := funcs.SetServiceAcc(ctx, "serviceKey/serviceAccountKey.json")

	// Instantiate client
	client := funcs.InstantiateFBClient(app, ctx)

	// Close down client
	defer funcs.CloseFB(client)

	// listen for port
	funcs.SetListener(port)
}
