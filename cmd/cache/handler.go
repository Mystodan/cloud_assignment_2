package cache

import (
	consts "assignment-2/constants"
	glob "assignment-2/globals"
	"log"
)

/**
 *	Loads webhooks from CACHE (firestore database).
 *
 *	@return A Struct containing all cached data.
 */
func LoadCacheFromFB() map[string]map[string]interface{} {
	retVal := make(map[string]map[string]interface{}) // Prepare return list

	log.Println("Loading cache...")
	// Loops through firestore database...
	loopThroughFireBase := glob.Client.Collection(consts.COLLECTION_CACHE).Documents(glob.Ctx)
	all, _ := loopThroughFireBase.GetAll()
	for i := range all {
		// reads data
		doc := all[i]

		if checkTime(doc) {
			doc.Ref.Delete(glob.Ctx)

			continue
		}

		// Thereafter appends to return list.
		data := doc.Data()
		for index, value := range data {
			retVal[index] = value.(map[string]interface{})
		}
	}
	if len(all) < 1 {
		log.Println("Nothing in cache to load!")
	} else {
		log.Println("Done!")
	}
	return retVal
}
