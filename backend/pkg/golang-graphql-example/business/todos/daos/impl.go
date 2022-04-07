package daos

import (
	"context"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/common"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/pagination"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type dao struct {
	db database.DB
}

func (d *dao) FindByID(ctx context.Context, id string, projection *models.Projection) (*models.Todo, error) {
	// Get gorm db
	db := d.db.GetTransactionalOrDefaultGormDB(ctx)
	// result
	res := &models.Todo{}
	// Apply projection
	db, err := common.ManageProjection(projection, db)
	// Check error
	if err != nil {
		return nil, err
	}
	// Find in db
	dbres := db.Where("id = ?", id).First(res)

	// Check error
	err = dbres.Error
	if err != nil {
		// Check if it is a not found error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // nolint: nilnil // not needed here
		}

		return nil, errors.WithStack(err)
	}

	return res, nil
}

func (d *dao) CreateOrUpdate(ctx context.Context, tt *models.Todo) (*models.Todo, error) {
	// Get gorm db
	db := d.db.GetTransactionalOrDefaultGormDB(ctx)
	dbres := db.Save(tt)

	// Check error
	err := dbres.Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Return result
	return tt, nil
}

func (d *dao) GetAllPaginated(
	ctx context.Context,
	page *pagination.PageInput,
	sort *models.SortOrder,
	filter *models.Filter,
	projection *models.Projection,
) ([]*models.Todo, *pagination.PageOutput, error) {
	// Get gorm db
	db := d.db.GetTransactionalOrDefaultGormDB(ctx)
	// result
	res := make([]*models.Todo, 0)
	// Find todos
	pageOut, err := pagination.Paging(&res, &pagination.PagingOptions{
		DB:         db,
		PageInput:  page,
		Sort:       sort,
		Filter:     filter,
		Projection: projection,
		ExtraFunc:  nil,
	})
	// Check error
	if err != nil {
		return nil, nil, err
	}

	return res, pageOut, nil
}

func (d *dao) Find(
	ctx context.Context,
	sort *models.SortOrder,
	filter *models.Filter,
	projection *models.Projection,
) ([]*models.Todo, error) {
	// Get gorm db
	db := d.db.GetTransactionalOrDefaultGormDB(ctx)
	// result
	res := make([]*models.Todo, 0)
	// Apply filter
	db, err := common.ManageFilter(filter, db)
	// Check error
	if err != nil {
		return nil, err
	}

	// Apply sort
	db, err = common.ManageSortOrder(sort, db)
	// Check error
	if err != nil {
		return nil, err
	}

	// Apply projection
	db, err = common.ManageProjection(projection, db)
	// Check error
	if err != nil {
		return nil, err
	}

	// Request to database with limit and offset
	db = db.Find(res)
	// Check error
	if db.Error != nil {
		return nil, errors.WithStack(db.Error)
	}

	return res, nil
}
