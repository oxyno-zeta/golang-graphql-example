package common

import (
	"fmt"

	"github.com/graph-gophers/dataloader/v6"
)

type IDProjectionKey struct {
	ID         string
	Projection interface{}
}

func (k *IDProjectionKey) String() string {
	return fmt.Sprintf("%s%v", k.ID, k.Projection)
}

func (k *IDProjectionKey) Raw() interface{} {
	return k
}

func GetIDsFromKeys(k dataloader.Keys) []string {
	// Get key ids
	keys := []string{}
	// Loop overs keys objects
	for _, keyObj := range k {
		// Get ID Projection Key
		raw, ok := keyObj.Raw().(*IDProjectionKey)
		// Check if cast is possible to save id
		if ok {
			keys = append(keys, raw.ID)
		} else {
			// Consider key object
			keys = append(keys, keyObj.String())
		}
	}

	return keys
}
