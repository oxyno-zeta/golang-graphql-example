package databasehelpers

import (
	"context"

	"emperror.dev/errors"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/common"
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

func SoftDelete[T any](
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

/**
 * PatchUpdate will update specific columns and return the updated object/model.
 * Params:
 * - ctx context
 * - originalObject Original object
 * - input is a map with gorm key with values that should be updated.
 */
func PatchUpdate[T any](
	ctx context.Context,
	originalObject T,
	input map[string]interface{},
	db database.DB,
) (T, error) {
	// Get gorm gdb
	gdb := db.GetTransactionalOrDefaultGormDB(ctx)
	dbres := gdb.Model(originalObject).Updates(input)

	// Check error
	err := dbres.Error
	if err != nil {
		return *new(T), errors.WithStack(err)
	}

	// Return result
	return originalObject, nil
}

/**
 * PatchUpdateAllFiltered will update specific columns filtered on where.
 * Params:
 * - ctx context
 * - model Empty object to have object column (!! This is different from PatchUpdate function !!)
 * - input is a map with gorm key with values that should be updated.
 * - filter is a filter object that will be used to filter lines where to apply patch.
 */
func PatchUpdateAllFiltered[T any](
	ctx context.Context,
	model T,
	input map[string]interface{},
	filter interface{},
	db database.DB,
) error {
	// Get gorm gdb
	gdb := db.GetTransactionalOrDefaultGormDB(ctx)

	// Apply filter
	gdb, err := common.ManageFilter(filter, gdb)
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	// Updates
	dbres := gdb.Model(model).Updates(input)

	// Check error
	err = dbres.Error
	if err != nil {
		return errors.WithStack(err)
	}

	// Return result
	return nil
}
