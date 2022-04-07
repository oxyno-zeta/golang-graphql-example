package todosdataloaders

import (
	"context"

	"github.com/graph-gophers/dataloader/v6"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/common"
	dataloaderscommon "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/dataloaders/common"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/utils"
)

type TodosDataloaders struct {
	EntitiesLoader dataloader.Interface
}

func New(busiServices *business.Services) *TodosDataloaders {
	return &TodosDataloaders{
		EntitiesLoader: dataloader.NewBatchedLoader(
			todosDataloaderGen(busiServices),
			dataloader.WithOpenTracingTracer(),
			dataloader.WithWait(dataloaderscommon.DefaultWait),
			dataloader.WithBatchCapacity(dataloaderscommon.DefaultBatchCapacity),
		),
	}
}

func todosDataloaderGen(busiServices *business.Services) func(context.Context, dataloader.Keys) []*dataloader.Result {
	return func(ctx context.Context, k dataloader.Keys) []*dataloader.Result {
		// Get keys
		keys := k.Keys()
		// Create result
		res := make([]*dataloader.Result, len(keys))

		// Create projection
		var projection models.Projection
		// Get projection from context
		// This is context is one of the requesting context
		// They are all coming from the entities request so projection will be the same as
		// dataloaders are created per request.
		err := utils.ManageSimpleProjection(ctx, &projection)
		// Check error
		if err != nil {
			// Fill array with errors
			dataloaderscommon.FillWithError(res, err)
			// Return result
			return res
		}
		// Force id fetch in projection
		// ID is used to rearrange items after
		projection.ID = true

		// Call business
		data, err := busiServices.TodoSvc.Find(
			ctx,
			nil,
			&models.Filter{
				ID: &common.GenericFilter{In: keys},
			},
			&projection,
		)
		// Check error
		if err != nil {
			// Fill array with errors
			dataloaderscommon.FillWithError(res, err)
			// Return result
			return res
		}

		// Default
		return dataloaderscommon.RearrangeResults(data, keys)
	}
}
