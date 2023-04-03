package common

import (
	"fmt"
)

type IDProjectionKey struct {
	Projection interface{}
	ID         string
}

func (k *IDProjectionKey) String() string {
	return fmt.Sprintf("%s%v", k.ID, k.Projection)
}

func (k *IDProjectionKey) Raw() interface{} {
	return k
}

type idsProjectionGroup struct {
	Projection interface{}
	IDs        []string
}
