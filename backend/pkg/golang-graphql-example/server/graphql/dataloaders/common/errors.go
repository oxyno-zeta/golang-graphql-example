package common

import (
	"github.com/graph-gophers/dataloader/v7"
)

// FillWithError will fill a dataloader result array with the same error for all.
func FillWithError[V any](length int, err error) []*dataloader.Result[V] {
	res := make([]*dataloader.Result[V], 0, length)

	for range length {
		res = append(res, &dataloader.Result[V]{Error: err})
	}

	return res
}
