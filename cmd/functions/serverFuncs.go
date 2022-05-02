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
func checkNum(inn string) bool {
	for i := range inn {
		if !(uint(inn[i]) >= 48 && uint(inn[i]) <= 57) {
			return false
		}
	}
	return true
}
func validPort(validity bool) string {
	var retVal string
	switch validity {
	case true:
		retVal = constants.PORT_NOTSET + constants.PORT_DEFAULT
	case false:
		retVal = constants.PORT_NOTSET + constants.PORT_INVALID + constants.PORT_DEFAULT
	}
	return retVal
}
func SetPort(inn string) string {
	var port string
	innValidity := checkNum(inn)
	if strings.ToLower(inn) == "default" || inn == constants.DEFAULT_PORT {
		log.Println(constants.PORT_DEFAULT + constants.DEFAULT_PORT)
		port = constants.DEFAULT_PORT
	} else if IsEmpty(inn) || !innValidity {
		log.Println(validPort(innValidity) + constants.DEFAULT_PORT)
		port = constants.DEFAULT_PORT
	} else {
		port = inn
		log.Println(constants.PORT_SET + port)
	}
	return port
}

func CompareLocalA3toCases() {
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

func LoadAllDependancies() {
	globals.AllWebhooks = notifications.LoadWebhooksFromFB() // webhooks from firebase
	globals.MemBuffer = cache.LoadCacheFromFB()              // cache from firebase
	globals.AllCountries = common.LoadCountries()            // Alpha3 from local library
}

func SetListener(inn string) {
	log.Println(constants.PORT_LISTEN + inn)
	log.Fatal(http.ListenAndServe(":"+inn, nil))
}

func SetServiceAcc(ctx context.Context, serviceKey string) *firebase.App {
	servAcc := option.WithCredentialsFile(serviceKey)
	app, err := firebase.NewApp(ctx, nil, servAcc)
	if err != nil {
		log.Fatal(constants.FB_INIT_ERR, err)
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
		log.Fatal(constants.FB_CLOSE_ERR, err)
	}
}
