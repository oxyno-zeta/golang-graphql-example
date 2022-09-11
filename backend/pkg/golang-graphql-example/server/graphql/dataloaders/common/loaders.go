package common

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader/v7"
)

func GenericEntitiesLoader[V any](
	buildProjection func(ctx context.Context) (interface{}, error),
	findAll func(ctx context.Context, ids []string, projection interface{}) ([]V, error),
) dataloader.BatchFunc[string, V] {
	return func(ctx context.Context, ids []string) []*dataloader.Result[V] {
		// Create result
		res := make([]*dataloader.Result[V], len(ids))

		// Build projection
		projection, err := buildProjection(ctx)
		// Check error
		if err != nil {
			// Fill array with errors
			FillWithError(res, err)
			// Return result
			return res
		}

		// Call business
		data, err := findAll(ctx, ids, projection)
		// Check error
		if err != nil {
			// Fill array with errors
			FillWithError(res, err)
			// Return result
			return res
		}

		// Default
		return RearrangeResults(data, ids)
	}
}

func GenericLoader[V any](
	findAll func(ctx context.Context, ids []string, projection interface{}) ([]V, error),
) dataloader.BatchFunc[*IDProjectionKey, V] {
	return func(ctx context.Context, keys []*IDProjectionKey) []*dataloader.Result[V] {
		// Create result
		res := make([]*dataloader.Result[V], len(keys))

		// Create ids
		ids := []string{}

		// Create temporary map
		tmp := map[string]*idsProjectionGroup{}

		// Loop over keys
		for _, key := range keys {
			// Add id in list
			ids = append(ids, key.ID)

			// Compute tmp key
			tmpKey := fmt.Sprintf("%+v", key.Projection)

			// Check if it isn't already populated
			if tmp[tmpKey] == nil {
				tmp[tmpKey] = &idsProjectionGroup{
					IDs:        []string{},
					Projection: key.Projection,
				}
			}

			// Add id in group
			tmp[tmpKey].IDs = append(tmp[tmpKey].IDs, key.ID)
		}

		// Make data list results
		data := make([]V, 0)
		// Loop over the tmp map
		for _, group := range tmp {
			// Call business
			dataTmp, err := findAll(ctx, group.IDs, group.Projection)
			// Check error
			if err != nil {
				// Fill array with errors
				FillWithError(res, err)
				// Return result
				return res
			}
			// Save results
			data = append(data, dataTmp...)
		}

		// Default
		return RearrangeResults(data, ids)
	}
}
