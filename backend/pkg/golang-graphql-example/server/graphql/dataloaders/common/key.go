package common

import (
	"fmt"
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

type idsProjectionGroup struct {
	IDs        []string
	Projection interface{}
}
