package databasehelpers

import (
	"context"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/common"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/pagination"
	"github.com/pkg/errors"
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

func GetAllPaginated[T any](
	ctx context.Context,
	res []T,
	db database.DB,
	page *pagination.PageInput,
	sort interface{},
	filter interface{},
	projection interface{},
) ([]T, *pagination.PageOutput, error) {
	// Get gorm db
	gdb := db.GetTransactionalOrDefaultGormDB(ctx)
	// Find
	pageOut, err := pagination.Paging(&res, &pagination.PagingOptions{
		DB:         gdb,
		PageInput:  page,
		Filter:     filter,
		Sort:       sort,
		Projection: projection,
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
			return *new(T), nil // nolint: nilnil // not needed here
		}

		return *new(T), errors.WithStack(err)
	}

	return res, nil
}

func FindByOne[T any](
	ctx context.Context,
	res T,
	db database.DB,
	filter interface{},
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

	// Apply filter
	gdb, err = common.ManageFilter(filter, gdb)
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
			return *new(T), nil // nolint: nilnil // not needed here
		}

		return *new(T), errors.WithStack(err)
	}

	return res, nil
}
