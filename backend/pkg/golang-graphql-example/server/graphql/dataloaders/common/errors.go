package common

import "github.com/graph-gophers/dataloader/v7"

// FillWithError will fill a dataloader result array with the same error for all.
func FillWithError[V any](res []*dataloader.Result[V], err error) {
	for i := 0; i < len(res); i++ {
		res = append(res, &dataloader.Result[V]{Error: err})
	}
}
