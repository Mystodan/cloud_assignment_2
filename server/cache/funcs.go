package cache

import (
	consts "assignment-2/constants"
	glob "assignment-2/global_types"
	"errors"
	"time"

	"cloud.google.com/go/firestore"
)

/**
 *	takes multiple strings inn, encloses them end returns them as a string
 *
 *	@param inn - The strings.
 *
 *	@return a string containing the params, separated by a seperator.
 */
func encloseParam(inn ...string) string {
	var retVal string = ""
	for _, param := range inn {
		retVal += param + consts.NEXT_PARAM
	}
	return retVal
}

/**
 *	Gets existing data from cache, unless...
 */
func GetCache(inn ...string) (map[string]interface{}, error) {
	var encloseParams string
	if len(inn) > 1 {
		encloseParams = encloseParam(inn...)
	} else {
		encloseParams = inn[0]
	}
	if val, Get := glob.MemBuffer[encloseParams]; Get {
		return val, nil
	}
	return map[string]interface{}{}, errors.New("currently no data exists within cache")
}

func CheckIfCached(inn string) bool {
	for i := range glob.MemBuffer {
		if i == inn {
			return true
		}
	}
	return false
}

func checkTime(doc *firestore.DocumentSnapshot) bool {
	// Check if the data is too old, if so, delete it
	return time.Since(doc.CreateTime).Hours() > consts.CACHE_TIMER
}

/**
 *	Adds something to the cache
 */
func AddToCache(name string, inn map[string]interface{}, params ...string) error {
	var err error = nil
	if !CheckIfCached(name) {
		glob.MemBuffer[name] = inn
		_, _, err = glob.Client.Collection(consts.COLLECTION_CACHE).Add(glob.Ctx, map[string]map[string]interface{}{name: inn})
	}
	return err
}
