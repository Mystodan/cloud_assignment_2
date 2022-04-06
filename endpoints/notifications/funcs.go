package notifications

import (
	consts "assignment-2/constants"
	glob "assignment-2/globals"
	"errors"
	"log"
	"math/rand"
)

func handleNewToken() string {
	token := GenerateRandomToken()
	for checkIfTokenExists(token) {
		token = GenerateRandomToken()
	}
	return token
}

func checkIfTokenExists(inn string) bool {
	for i := range allWebhooks {
		if allWebhooks[i].ID == inn {
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

func sendWebhookToFB(webhook Webhook) (string, error) {
	id, _, err := glob.Client.Collection(consts.Collection).Add(glob.Ctx, webhook)
	return id.ID, err
}

func GetWebhook(id string, webhooks map[string]Webhook) (Webhook, error) {
	for _, webhook := range webhooks {
		if webhook.ID == id {
			return webhook, nil
		}
	}
	return Webhook{}, errors.New("webhook(id): not found")
}

func GetAllWebhooks() map[string]Webhook {
	return allWebhooks
}
func SetAllTokens(inn map[string]Webhook) {
	allWebhooks = inn
}

/**
 *	Loads webhooks from the firestore database.
 *
 *	@return A map linking each webhook's doc ref/ID to their struct.
 */
func LoadWebhooksFromFB() map[string]Webhook {
	ret := make(map[string]Webhook) // Prepare return list

	// Iterate through firestore database...
	loopThroughFireBase := glob.Client.Collection(consts.Collection).Documents(glob.Ctx)
	all, _ := loopThroughFireBase.GetAll()
	log.Println("Loading webhooks...")

	for i := 0; i < len(all); i++ {
		doc := all[i]
		// And fitting it into the Webhook struct.
		data := doc.Data()
		ret[doc.Ref.ID] = Webhook{
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
	return ret
}
func DeleteWebhook(id string, webhooks *map[string]Webhook) (bool, error) {
	var deleted = false
	// Delete from local webhooks, temporarily storing deleted entries
	deletedWebhook := []string{}
	for webhook_string, value := range *webhooks {
		if value.ID == id {
			deleted = true
			deletedWebhook = append(deletedWebhook, webhook_string)
			delete(*webhooks, webhook_string)
		}
	}
	// Loop through temporarily stored webhooks, matching, thereafter deleting them from FireStore database
	var retVal error = nil
	for _, webhook_string := range deletedWebhook {
		_, delErr := glob.Client.Collection(consts.Collection).Doc(webhook_string).Delete(glob.Ctx)
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
