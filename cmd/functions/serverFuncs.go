package server

import (
	"assignment-2/cmd/cache"
	consts "assignment-2/constants"
	"assignment-2/endpoints/notifications"
	glob "assignment-2/globals"
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
		retVal = consts.PORT_NOTSET + consts.PORT_DEFAULT
	case false:
		retVal = consts.PORT_NOTSET + consts.PORT_INVALID + consts.PORT_DEFAULT
	}
	return retVal
}
func SetPort(inn string) string {
	var port string
	innValidity := checkNum(inn)
	if strings.ToLower(inn) == "default" || inn == "8080" {
		log.Println(consts.PORT_DEFAULT + consts.DEFAULT_PORT)
		port = consts.DEFAULT_PORT
	} else if IsEmpty(inn) || !innValidity {
		log.Println(validPort(innValidity) + consts.DEFAULT_PORT)
		port = consts.DEFAULT_PORT
	} else {
		port = inn
		log.Println(consts.PORT_SET + port)
	}
	return port
}

func CompareLocalA3toCases() {
	Inconsistancies, err, amount, cases_size := common.CompareGraphCountryNames()
	if err {
		log.Println("No Inconsistancies found between Alpha 3 library and Cases dependancy")
	} else {
		log.Println("Inconsistancies found between Alpha 3 library and Cases dependancy:")
		fmt.Println(">>\t	FOUND (", amount, "):")
		for _, val := range Inconsistancies {
			fmt.Println("> Country:\t[" + val + "]")
		}
		fmt.Println("> Incorrect amount:\t[" + fmt.Sprint(amount) + "/" + fmt.Sprint(cases_size) + "]\n> Correct amount:\t[" + fmt.Sprint(cases_size-amount) + "/" + fmt.Sprint(cases_size) + "]")
	}
}

func LoadAllDependancies() {
	glob.AllWebhooks = notifications.LoadWebhooksFromFB() // webhooks from firebase
	glob.MemBuffer = cache.LoadCacheFromFB()              // cache from firebase
	glob.AllCountries = common.LoadCountries()            // Alpha3 from local library
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
