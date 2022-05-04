package databasehelpers

import (
	"context"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/pkg/errors"
)

func CreateOrUpdate[T any](
	ctx context.Context,
	input T,
	db database.DB,
) (T, error) {
	// Get gorm gdb
	gdb := db.GetTransactionalOrDefaultGormDB(ctx)
	dbres := gdb.Save(input)

	// Check error
	err := dbres.Error
	if err != nil {
		return *new(T), errors.WithStack(err)
	}

	// Return result
	return input, nil
}

func PermanentDelete[T any](
	ctx context.Context,
	input T,
	db database.DB,
) (T, error) {
	// Get gorm gdb
	gdb := db.GetTransactionalOrDefaultGormDB(ctx)
	dbres := gdb.Unscoped().Delete(input)

	// Check error
	err := dbres.Error
	if err != nil {
		return *new(T), errors.WithStack(err)
	}

	// Return result
	return input, nil
}

func Delete[T any](
	ctx context.Context,
	input T,
	db database.DB,
) (T, error) {
	// Get gorm gdb
	gdb := db.GetTransactionalOrDefaultGormDB(ctx)
	dbres := gdb.Delete(input)

	// Check error
	err := dbres.Error
	if err != nil {
		return *new(T), errors.WithStack(err)
	}

	// Return result
	return input, nil
}
