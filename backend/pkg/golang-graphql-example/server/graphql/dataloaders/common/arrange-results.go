package common

import (
	"github.com/graph-gophers/dataloader/v7"
	"github.com/thoas/go-funk"
)

// Rearrange results to ensure ids <-> items are at the same place.
func rearrangeResults[T any](input []T, ids []string, idKey string) []*dataloader.Result[T] {
	// Optimization
	if len(input) == 1 {
		return []*dataloader.Result[T]{{Data: input[0]}}
	}

	// Create result
	res := make([]*dataloader.Result[T], len(ids))

	// Create intermediate map
	inMap := funk.ToMap(input, idKey).(map[string]T) //nolint: forcetypeassert // Ignored

	// Rearrange results
	for i, id := range ids {
		res[i] = &dataloader.Result[T]{Data: inMap[id]}
	}

	return res
}
