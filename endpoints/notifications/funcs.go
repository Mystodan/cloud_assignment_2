package notifications

import (
	consts "assignment-2/constants"
	glob "assignment-2/global_types"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

// Functions for webhook tokens
func handleNewToken() string {
	token := GenerateRandomToken()
	for checkIfTokenExists(token) {
		token = GenerateRandomToken()
	}
	return token
}

func checkIfTokenExists(inn string) bool {
	for i := range glob.AllWebhooks {
		if glob.AllWebhooks[i].ID == inn {
			return true
		}
	}
	return false
}

func GenerateRandomToken() string {
	Token := make([]byte, consts.APP_TOKEN_LENGTH)
	for i := range Token {
		Token[i] = consts.APP_TOKEN_SYMBOLS[rand.Intn(len(consts.APP_TOKEN_SYMBOLS))]
	}
	return string(Token)
}

func sendWebhookToFB(webhook glob.Webhook) (string, error) {
	id, _, err := glob.Client.Collection(consts.COLLECTION_WEBHOOKS).Add(glob.Ctx, webhook)
	return id.ID, err
}

func GetWebhook(id string, webhooks map[string]glob.Webhook) (glob.Webhook, error) {
	for _, webhook := range webhooks {
		if webhook.ID == id {
			return webhook, nil
		}
	}
	return glob.Webhook{}, errors.New("webhook(id): not found")
}

/**
 *	Loads webhooks from the firestore database.
 *
 *	@return A map linking each webhook's doc ref/ID to their struct.
 */
func LoadWebhooksFromFB() map[string]glob.Webhook {
	retVal := make(map[string]glob.Webhook) // Prepare return list

	// Iterate through firestore database...
	loopThroughFireBase := glob.Client.Collection(consts.COLLECTION_WEBHOOKS).Documents(glob.Ctx)
	all, _ := loopThroughFireBase.GetAll()
	log.Println("Loading webhooks...")

	for i := 0; i < len(all); i++ {
		// reads data
		doc := all[i]
		// Thereafter fits it into the Webhook struct.
		data := doc.Data()
		retVal[doc.Ref.ID] = glob.Webhook{
			ID:      data["ID"].(string),
			Url:     data["Url"].(string),
			Country: data["Country"].(string),
			Calls:   data["Calls"].(int64),
		}
	}
	if len(all) < 1 {
		log.Println("No webhooks to load!")
	} else {
		log.Println("Done!")
	}
	return retVal
}
func DeleteWebhook(id string, webhooks *map[string]glob.Webhook) (bool, error) {
	deleted := false
	// Delete from local webhooks, temporarily storing deleted entries
	deletedWebhooks := []string{}
	for webhook_string, value := range *webhooks {
		if value.ID == id {
			deleted = true
			deletedWebhooks = append(deletedWebhooks, webhook_string)
			delete(*webhooks, webhook_string)
		}
	}
	// Loop through temporarily stored webhooks, matching, thereafter deleting them from FireStore database
	var retVal error = nil
	for _, webhook_string := range deletedWebhooks {
		_, delErr := glob.Client.Collection(consts.COLLECTION_WEBHOOKS).Doc(webhook_string).Delete(glob.Ctx)
		if delErr != nil {
			retVal = delErr
		}
	}
	return deleted, retVal
}

func handleDeleted(inn bool) string {
	var retVal string
	switch inn {
	case true:
		retVal = consts.TOKEN_DELETED_FOUND
	case false:
		retVal = consts.TOKEN_DELETED_NOT_FOUND
	}
	return retVal
}

// functions for invocations

/**
 *	Invokes all webhooks (if applicable).
 *
 *	countryCalls - A map of each country and how many times they've been called.
 */
func InvokeWebhooks(country string, countryCalls map[string]int, webhooks map[string]glob.Webhook) {
	invocations := countryCalls[country]
	for _, webhook := range webhooks {
		if webhook.Country == country && invocations%int(webhook.Calls) == 0 {
			InvokeWebhook(webhook)
		}
	}
}

/**
 *	Invokes a given webhook.
 *
 *	@param webhook - The webhook.
 */
func InvokeWebhook(webhook glob.Webhook) (interface{}, error) {
	// prepare/wrap data for invocation
	data := bytes.NewReader([]byte(fmt.Sprintf(`{
		"webhook_id": "%s",
		"country": "%s",
		"calls": "%d"
	}`, webhook.ID, webhook.Country, webhook.Calls)))

	// create a post request to URL for invocation
	POSTrequest, err := http.NewRequest(http.MethodPost, webhook.Url, data)
	if err != nil {
		return "", err
	}

	var client http.Client = http.Client{}
	resp, err := client.Do(POSTrequest)
	if err != nil {
		return "", err
	}

	// Read the response
	var response interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}
func SetInvocation(country string) {
	if _, invoked := glob.CountryInvocations[country]; !invoked {
		glob.CountryInvocations[country] = 0
	}

	glob.CountryInvocations[country]++
	InvokeWebhooks(country, glob.CountryInvocations, glob.AllWebhooks)
}
