package databasehelpers

import (
	"context"

	"emperror.dev/errors"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/common"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/pagination"
	"gorm.io/gorm"
)

func Find[T any](
	ctx context.Context,
	res []T,
	db database.DB,
	sort interface{},
	filter interface{},
	projection interface{},
) ([]T, error) {
	// Get gorm gdb
	gdb := db.GetTransactionalOrDefaultGormDB(ctx)
	// Apply filter
	gdb, err := common.ManageFilter(filter, gdb)
	// Check error
	if err != nil {
		return nil, err
	}

	// Apply sort
	gdb, err = common.ManageSortOrder(sort, gdb)
	// Check error
	if err != nil {
		return nil, err
	}

	// Apply projection
	gdb, err = common.ManageProjection(projection, gdb)
	// Check error
	if err != nil {
		return nil, err
	}

	// Request to database with limit and offset
	gdb = gdb.Find(&res)
	// Check error
	if gdb.Error != nil {
		return nil, errors.WithStack(gdb.Error)
	}

	return res, nil
}

func FindWithPagination[T any](
	ctx context.Context,
	res []T,
	db database.DB,
	page *pagination.PageInput,
	sort interface{},
	filter interface{},
	projection interface{},
) ([]T, error) {
	// Manage default limit
	if page.Limit == 0 {
		page.Limit = 10
	}

	// Get gorm gdb
	gdb := db.GetTransactionalOrDefaultGormDB(ctx)
	// Apply filter
	gdb, err := common.ManageFilter(filter, gdb)
	// Check error
	if err != nil {
		return nil, err
	}

	// Apply sort
	gdb, err = common.ManageSortOrder(sort, gdb)
	// Check error
	if err != nil {
		return nil, err
	}

	// Apply projection
	gdb, err = common.ManageProjection(projection, gdb)
	// Check error
	if err != nil {
		return nil, err
	}

	// Request to database with limit and offset
	gdb = gdb.Limit(page.Limit).Offset(page.Skip).Find(&res)
	// Check error
	if gdb.Error != nil {
		return nil, errors.WithStack(gdb.Error)
	}

	return res, nil
}

func CountPaginated[T any](
	ctx context.Context,
	db database.DB,
	input T,
	page *pagination.PageInput,
	sort interface{},
	filter interface{},
) (int64, error) {
	// Initialize count
	var res int64

	// Get gorm gdb
	gdb := db.GetTransactionalOrDefaultGormDB(ctx)
	// Apply filter
	gdb, err := common.ManageFilter(filter, gdb)
	// Check error
	if err != nil {
		return 0, err
	}

	// Apply sort
	gdb, err = common.ManageSortOrder(sort, gdb)
	// Check error
	if err != nil {
		return 0, err
	}

	// Request to database with limit and offset
	gdb = gdb.Model(input).Limit(page.Limit).Offset(page.Skip).Count(&res)
	// Check error
	if gdb.Error != nil {
		return 0, errors.WithStack(gdb.Error)
	}

	return res, nil
}

func Count[T any](
	ctx context.Context,
	db database.DB,
	input T,
	sort interface{},
	filter interface{},
) (int64, error) {
	// Initialize count
	var res int64

	// Get gorm gdb
	gdb := db.GetTransactionalOrDefaultGormDB(ctx)
	// Apply filter
	gdb, err := common.ManageFilter(filter, gdb)
	// Check error
	if err != nil {
		return 0, err
	}

	// Apply sort
	gdb, err = common.ManageSortOrder(sort, gdb)
	// Check error
	if err != nil {
		return 0, err
	}

	// Request to database with limit and offset
	gdb = gdb.Model(input).Count(&res)
	// Check error
	if gdb.Error != nil {
		return 0, errors.WithStack(gdb.Error)
	}

	return res, nil
}

func GetAllPaginated[T any](
	ctx context.Context,
	res []T,
	db database.DB,
	page *pagination.PageInput,
	sort interface{},
	filter interface{},
	projection interface{},
	tOpts ...database.TransactionOption,
) ([]T, *pagination.PageOutput, error) {
	// Find
	pageOut, err := pagination.Paging(ctx, &res, &pagination.PagingOptions{
		DBSvc:      db,
		PageInput:  page,
		Filter:     filter,
		Sort:       sort,
		Projection: projection,
		TOpts:      tOpts,
	})
	// Check error
	if err != nil {
		return nil, nil, err
	}

	return res, pageOut, nil
}

func FindByID[T any](
	ctx context.Context,
	res T,
	db database.DB,
	id string,
	projection interface{},
) (T, error) {
	// Get gorm db
	gdb := db.GetTransactionalOrDefaultGormDB(ctx)
	// Apply projection
	gdb, err := common.ManageProjection(projection, gdb)
	// Check error
	if err != nil {
		return *new(T), err
	}

	// Find in db
	dbres := gdb.Where("id = ?", id).First(res)

	// Check error
	err = dbres.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return *new(T), nil
		}

		return *new(T), errors.WithStack(err)
	}

	return res, nil
}

func FindOne[T any](
	ctx context.Context,
	res T,
	db database.DB,
	filter interface{},
	projection interface{},
) (T, error) {
	// Get gorm db
	gdb := db.GetTransactionalOrDefaultGormDB(ctx)
	// Apply filter
	gdb, err := common.ManageFilter(filter, gdb)
	// Check error
	if err != nil {
		return *new(T), err
	}

	// Apply projection
	gdb, err = common.ManageProjection(projection, gdb)
	// Check error
	if err != nil {
		return *new(T), err
	}

	// Find in db
	dbres := gdb.First(res)

	// Check error
	err = dbres.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return *new(T), nil
		}

		return *new(T), errors.WithStack(err)
	}

	return res, nil
}
