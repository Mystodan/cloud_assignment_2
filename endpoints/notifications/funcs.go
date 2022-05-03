package notifications

import (
	consts "assignment-2/constants"
	glob "assignment-2/globals"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// Functions for webhook tokens
func handleNewToken() string {
	token := Gen_Token()
	for checkIfTokenExists(token) {
		token = Gen_Token()
	}
	return token
}

// Function for checking if token already exists within local buffer
func checkIfTokenExists(inn string) bool {
	for i := range glob.AllWebhooks {
		if glob.AllWebhooks[i].ID == inn {
			return true
		}
	}
	return false
}

// Generates a random token based on constants
func GenerateRandomToken() string {
	Token := make([]byte, TOKEN_LENGTH)
	for i := range Token {
		Token[i] = TOKEN_SYMBOLS[rand.Intn(len(TOKEN_SYMBOLS))]
	}
	return string(Token)
}

// Deletes a webhook from firebase database
func RemoveWebhookFromFB(webhook_id string) error {
	_, err := glob.Client.Collection(consts.COLLECTION_WEBHOOKS).Doc(webhook_id).Delete(glob.Ctx)
	return err
}

// Sends a webhook to firebase database
func SendWebhookToFB(webhook glob.Webhook) (string, error) {
	id, _, err := glob.Client.Collection(consts.COLLECTION_WEBHOOKS).Add(glob.Ctx, webhook)
	return id.ID, err
}

// returns a document iterator to iterate through firebase database
func iterateWebhooksFromFB() *firestore.DocumentIterator {
	return glob.Client.Collection(consts.COLLECTION_WEBHOOKS).Documents(glob.Ctx)
}

// gets a specific webhook from local buffer
func GetWebhook(id string, webhooks map[string]glob.Webhook) (glob.Webhook, error) {
	for _, webhook := range webhooks {
		if webhook.ID == id {
			return webhook, nil
		}
	}
	return glob.Webhook{}, errors.New(consts.INPUT_NOT_FOUND)
}

/**
 *	Loads webhooks from the firestore database.
 *
 *	@return A map linking each webhook's doc ref/ID to their struct.
 */
func LoadWebhooksFromFB() map[string]glob.Webhook {
	retVal := make(map[string]glob.Webhook) // Prepare return list

	// Iterate through firestore database...
	IterateWebhooks := iterateWebhooksFromFB()
	log.Println("Loading webhooks...")

	for {
		// reads data
		doc, err := IterateWebhooks.Next()
		if err == iterator.Done {
			log.Println("Done!")
			IterateWebhooks.Stop()
			break
		}
		if err != nil {
			panic("Failed to load webhooks from firestore!")
		}

		// Thereafter fits it into the Webhook struct.
		data := doc.Data()
		retVal[doc.Ref.ID] = glob.Webhook{
			ID:      data["ID"].(string),
			Url:     data["Url"].(string),
			Country: data["Country"].(string),
			Calls:   data["Calls"].(int64),
		}
	}
	return retVal
}

//Deletes webhooks from local and firebase database/buffer
func DeleteWebhooks(id string, webhooks *map[string]glob.Webhook) (bool, error) {
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
	if glob.AllowFBWebhooks {
		for _, webhook_id := range deletedWebhooks {
			delErr := RemoveWebhookFromFB(webhook_id)
			if delErr != nil {
				retVal = delErr
			}
		}
	}
	return deleted, retVal
}

// Handles msgs on deleting webhooks
func handleDeleted(inn bool) string {
	var retVal string
	switch inn {
	case true: // if specific webhook is found
		retVal = consts.TOKEN_DELETED_FOUND
	case false: // if specific webhook does not exist
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

// sets invocation based on country name
func SetInvocation(country string) {
	if _, invoked := glob.CountryInvocations[country]; !invoked {
		glob.CountryInvocations[country] = 0
	}
	glob.CountryInvocations[country]++
	InvokeWebhooks(country, glob.CountryInvocations, glob.AllWebhooks)
}
