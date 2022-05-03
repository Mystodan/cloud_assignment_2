package server

import (
	"assignment-2/cmd/cache"
	constants "assignment-2/constants"
	"assignment-2/endpoints/notifications"
	globals "assignment-2/globals"
	"assignment-2/globals/common"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

//checks if string is empty
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

//Checks if a string contains a number
func checkNum(inn string) bool {
	for i := range inn {
		if !(uint(inn[i]) >= 48 && uint(inn[i]) <= 57) {
			return false
		}
	}
	return true
}

// handler for messages for valid ports
func validPort(validity bool) string {
	var retVal string
	switch validity {
	case true: // if port is valid
		retVal = constants.PORT_NOTSET + constants.PORT_DEFAULT
	case false: // if port is invalid
		retVal = constants.PORT_NOTSET + constants.PORT_INVALID + constants.PORT_DEFAULT
	}
	return retVal
}

// function which handles port msgs
func SetPort(inn string) string {
	var port string
	innValidity := checkNum(inn)                                            // checks for numbers
	if strings.ToLower(inn) == "default" || inn == constants.DEFAULT_PORT { // checks for default port
		log.Println(constants.PORT_DEFAULT + constants.DEFAULT_PORT)
		port = constants.DEFAULT_PORT
	} else if IsEmpty(inn) || !innValidity { // checks for empty inn if so set to default
		log.Println(validPort(innValidity) + constants.DEFAULT_PORT)
		port = constants.DEFAULT_PORT
	} else {
		port = inn // set different than default port
		log.Println(constants.PORT_SET + port)
	}
	return port
}

// Handles logging of comparing local alpha3 library/dependancy to Cases API countries
func CompareLocalA3toCases() { // gets necessary data from compare function
	Inconsistancies, err, amount, cases_size := common.CompareGraphCountryNames()
	if err {
		log.Println(constants.COMPARE_A3_CASES_NOT_FOUND)
	} else {
		log.Println(constants.COMPARE_A3_CASES_FOUND)
		fmt.Printf(constants.COMPARE_A3_CASES_HEADER, amount)
		for _, val := range Inconsistancies {
			fmt.Printf(constants.COMPARE_A3_CASES_IDENTIFIER, val)
		}
		fmt.Printf(constants.COMPARE_A3_CASES_BODY, amount, cases_size, cases_size-amount, cases_size)
	}
}

// Loads all data from dependencies, firebase and local
func LoadAllDependancies() {
	globals.AllWebhooks = notifications.LoadWebhooksFromFB() // webhooks from firebase
	globals.MemBuffer = cache.LoadCacheFromFB()              // cache from firebase
	globals.AllCountries = common.LoadCountries()            // Alpha3 from local library
}

// sets port listener
func SetListener(inn string) {
	log.Println(constants.PORT_LISTEN + inn)
	log.Fatal(http.ListenAndServe(":"+inn, nil))
}

// Access to Firebase established
func SetServiceAcc(ctx context.Context, serviceKey string) *firebase.App {
	servAcc := option.WithCredentialsFile(serviceKey) // using local servicekey
	app, err := firebase.NewApp(ctx, nil, servAcc)
	if err != nil {
		log.Fatal(constants.FB_INIT_ERR, err)
	}
	return app
}

// Instantiates Firebase client and returns it as *firestore.Client
func InstantiateFBClient(app *firebase.App, ctx context.Context) *firestore.Client {
	client, err := app.Firestore(ctx)
	checkError(err)
	return client
}

// Closes firebases client
func CloseFB(client *firestore.Client) {
	err := client.Close()
	if err != nil {
		log.Fatal(constants.FB_CLOSE_ERR, err)
	}
}
