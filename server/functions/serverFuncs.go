package server

import (
	consts "assignment-2/constants"
	"context"
	"log"
	"net/http"
	"os"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func IsEmpty(inn string) bool {
	return !(len(inn) > 0)
}

/**	checkError logs an error.
 *	@param inn - error value
 */
func checkError(inn error) {
	if inn != nil {
		log.Fatal(inn)
	}
}

func SetPort(inn string) string {
	port := os.Getenv(consts.ENVK_PORT)
	if IsEmpty(port) {
		log.Println("$PORT has not been set. Default: " + inn)
		port = inn
	}
	return port
}

func SetListener(inn string) {
	log.Println("Listening on port " + inn)
	log.Fatal(http.ListenAndServe(":"+inn, nil))
}

func SetServiceAcc(ctx context.Context, serviceKey string) *firebase.App {
	servAcc := option.WithCredentialsFile(serviceKey)
	app, err := firebase.NewApp(ctx, nil, servAcc)
	if err != nil {
		log.Fatal("error initializing app:", err)
	}
	return app
}

func InstantiateFBClient(app *firebase.App, ctx context.Context) *firestore.Client {
	client, err := app.Firestore(ctx)
	checkError(err)
	return client
}

func CloseFB(client *firestore.Client) {
	err := client.Close()
	if err != nil {
		log.Fatal("Closing of the firebase client failed. Error:", err)
	}
}
