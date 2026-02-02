package common

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader/v7"
	"github.com/thoas/go-funk"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/common/errors"
)

func GenericEntitiesLoader[V any](
	buildProjection func(ctx context.Context) (any, error),
	findAll func(ctx context.Context, ids []string, projection any) ([]V, error),
	options ...func(*LoaderOption),
) dataloader.BatchFunc[string, V] {
	return func(ctx context.Context, ids []string) []*dataloader.Result[V] {
		length := len(ids)

		// Get options
		opts := getOptions(options)

		// Build projection
		projection, err := buildProjection(ctx)
		// Check error
		if err != nil {
			// Fill array with errors
			return FillWithError[V](length, err)
		}

		// Check if projection exists
		if projection != nil {
			// Force set id in projection
			err = funk.Set(projection, true, opts.IDKey)
			// Check error
			if err != nil {
				// Fill array with errors
				return FillWithError[V](length, errors.NewInternalServerErrorWithError(err))
			}
		}

		// Call business
		data, err := findAll(ctx, ids, projection)
		// Check error
		if err != nil {
			// Fill array with errors
			return FillWithError[V](length, err)
		}

		// Default
		return rearrangeResults(data, ids, opts.IDKey)
	}
}

func GenericLoader[V any](
	findAll func(ctx context.Context, ids []string, projection any) ([]V, error),
	options ...func(*LoaderOption),
) dataloader.BatchFunc[*IDProjectionKey, V] {
	return func(ctx context.Context, keys []*IDProjectionKey) []*dataloader.Result[V] {
		length := len(keys)
		// Get options
		opts := getOptions(options)

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
			// Get projection
			projection := group.Projection
			// Check if projection exists and id key exists
			if projection != nil && opts.IDKey != "" {
				// Force set id in projection
				err := funk.Set(projection, true, opts.IDKey)
				// Check error
				if err != nil {
					// Fill array with errors
					return FillWithError[V](length, errors.NewInternalServerErrorWithError(err))
				}
			}

			// Call business
			dataTmp, err := findAll(ctx, group.IDs, projection)
			// Check error
			if err != nil {
				// Fill array with errors
				return FillWithError[V](length, err)
			}
			// Save results
			data = append(data, dataTmp...)
		}

		// Default
		return rearrangeResults(data, ids, opts.IDKey)
	}
}
