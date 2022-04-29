package glob

import (
	"context"

	"cloud.google.com/go/firestore"
)

// Firebase context and client used by Firestore functions throughout the program.
var Ctx context.Context
var Client *firestore.Client
