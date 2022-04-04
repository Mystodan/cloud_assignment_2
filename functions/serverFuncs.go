package funcs

import (
	consts "assignment-2/constants"
	"log"
	"net/http"
	"os"
)

func IsEmpty(inn string) bool {
	return !(len(inn) > 0)
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
