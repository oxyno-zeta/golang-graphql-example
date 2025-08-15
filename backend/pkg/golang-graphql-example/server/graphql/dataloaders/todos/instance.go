package todosdataloaders

import (
	"context"

	"github.com/graph-gophers/dataloader/v7"

	dataloadertracing "github.com/graph-gophers/dataloader/v7/trace/otel"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/common"
	dataloaderscommon "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/dataloaders/common"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/utils"
)

type TodosDataloaders struct {
	EntitiesLoader dataloader.Interface[string, *models.Todo]
	GenericLoader  dataloader.Interface[*dataloaderscommon.IDProjectionKey, *models.Todo]
}

func New(busiServices *business.Services) *TodosDataloaders {
	return &TodosDataloaders{
		GenericLoader: dataloader.NewBatchedLoader(
			dataloaderscommon.GenericLoader(
				func(ctx context.Context, ids []string, projection any) ([]*models.Todo, error) {
					return busiServices.TodoSvc.Find(
						ctx,
						nil,
						&models.Filter{ID: &common.GenericFilter{In: ids}},
						projection.(*models.Projection), //nolint:forcetypeassert // We know that it is that
					)
				},
			),
			dataloader.WithTracer[*dataloaderscommon.IDProjectionKey, *models.Todo](
				dataloadertracing.NewTracer[*dataloaderscommon.IDProjectionKey, *models.Todo](nil),
			),
			dataloader.WithWait[*dataloaderscommon.IDProjectionKey, *models.Todo](
				dataloaderscommon.DefaultWait,
			),
			dataloader.WithBatchCapacity[*dataloaderscommon.IDProjectionKey, *models.Todo](
				dataloaderscommon.DefaultBatchCapacity,
			),
			dataloader.WithCache[*dataloaderscommon.IDProjectionKey, *models.Todo](
				dataloaderscommon.NewCache[*dataloaderscommon.IDProjectionKey, *models.Todo](),
			),
		),
		EntitiesLoader: dataloader.NewBatchedLoader(
			dataloaderscommon.GenericEntitiesLoader(
				func(ctx context.Context) (any, error) {
					// Create projection
					var projection models.Projection
					// Get projection from context
					// This is context is one of the requesting context
					// They are all coming from the entities request so projection will be the same as
					// dataloaders are created per request and each key is unique.
					err := utils.ManageSimpleProjection(ctx, &projection)
					// Check error
					if err != nil {
						return nil, err
					}

					// Return default
					return &projection, nil
				},
				func(ctx context.Context, ids []string, projection any) ([]*models.Todo, error) {
					return busiServices.TodoSvc.Find(
						ctx,
						nil,
						&models.Filter{ID: &common.GenericFilter{In: ids}},
						projection.(*models.Projection), //nolint:forcetypeassert // We know that it is that
					)
				},
			),
			dataloader.WithTracer[string, *models.Todo](
				&dataloadertracing.Tracer[string, *models.Todo]{},
			),
			dataloader.WithWait[string, *models.Todo](dataloaderscommon.DefaultWait),
			dataloader.WithBatchCapacity[string, *models.Todo](
				dataloaderscommon.DefaultBatchCapacity,
			),
			dataloader.WithCache[string, *models.Todo](
				dataloaderscommon.NewCache[string, *models.Todo](),
			),
		),
	}
}
