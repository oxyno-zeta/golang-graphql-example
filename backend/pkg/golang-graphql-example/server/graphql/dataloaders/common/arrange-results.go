package common

import (
	"github.com/graph-gophers/dataloader/v6"
	"github.com/thoas/go-funk"
)

// Rearrange results to ensure ids <-> items are at the same place.
func RearrangeResults[T any](input []T, ids []string) []*dataloader.Result {
	// Optimization
	if len(input) == 1 {
		return []*dataloader.Result{{Data: input[0]}}
	}

	// Create result
	res := make([]*dataloader.Result, len(ids))

	// Create intermediate map
	inMap := funk.ToMap(input, "ID").(map[string]T) // nolint: forcetypeassert // Ignored

	// Rearrange results
	for i, id := range ids {
		res[i] = &dataloader.Result{Data: inMap[id]}
	}

	return res
}
