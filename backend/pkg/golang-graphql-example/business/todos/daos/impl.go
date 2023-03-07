package daos

import (
	"context"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	databasehelpers "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/helpers"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/pagination"
)

type dao struct {
	db database.DB
}

func (d *dao) FindByID(ctx context.Context, id string, projection *models.Projection) (*models.Todo, error) {
	return databasehelpers.FindByID(
		ctx,
		&models.Todo{},
		d.db,
		id,
		projection,
	)
}

func (d *dao) CreateOrUpdate(ctx context.Context, tt *models.Todo) (*models.Todo, error) {
	return databasehelpers.CreateOrUpdate(
		ctx,
		tt,
		d.db,
	)
}

func (d *dao) GetAllPaginated(
	ctx context.Context,
	page *pagination.PageInput,
	sort []*models.SortOrder,
	filter *models.Filter,
	projection *models.Projection,
) ([]*models.Todo, *pagination.PageOutput, error) {
	return databasehelpers.GetAllPaginated(
		ctx,
		make([]*models.Todo, 0),
		d.db,
		page,
		sort,
		filter,
		projection,
	)
}

func (d *dao) Find(
	ctx context.Context,
	sort []*models.SortOrder,
	filter *models.Filter,
	projection *models.Projection,
) ([]*models.Todo, error) {
	return databasehelpers.Find(
		ctx,
		make([]*models.Todo, 0),
		d.db,
		sort,
		filter,
		projection,
	)
}
