package common

import "github.com/graph-gophers/dataloader/v6"

// FillWithError will fill a dataloader result array with the same error for all.
func FillWithError(res []*dataloader.Result, err error) {
	for i := 0; i < len(res); i++ {
		res = append(res, &dataloader.Result{Error: err})
	}
}
