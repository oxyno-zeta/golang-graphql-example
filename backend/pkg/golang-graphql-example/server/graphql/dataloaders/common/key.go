package common

import (
	"fmt"
)

type IDProjectionKey struct {
	Projection any
	ID         string
}

func (k *IDProjectionKey) String() string {
	return fmt.Sprintf("%s%v", k.ID, k.Projection)
}

func (k *IDProjectionKey) Raw() any {
	return k
}

type idsProjectionGroup struct {
	Projection any
	IDs        []string
}
