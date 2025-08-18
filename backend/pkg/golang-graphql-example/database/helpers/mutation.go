package databasehelpers

import (
	"context"

	"emperror.dev/errors"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
)

func CreateOrUpdate[T any](
	ctx context.Context,
	input T,
	db database.DB,
	opts ...GormOpt,
) (T, error) {
	// Get gorm gdb
	gdb := db.GetTransactionalOrDefaultGormDB(ctx)

	// Define error
	var err error
	// Apply options
	for _, o := range opts {
		gdb, err = o(ctx, gdb)
		// Check error
		if err != nil {
			return *new(T), errors.WithStack(err)
		}
	}

	// Save
	dbres := gdb.Save(input)

	// Check error
	err = dbres.Error
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
	opts ...GormOpt,
) (T, error) {
	// Get gorm gdb
	gdb := db.GetTransactionalOrDefaultGormDB(ctx)

	// Define error
	var err error
	// Apply options
	for _, o := range opts {
		gdb, err = o(ctx, gdb)
		// Check error
		if err != nil {
			return *new(T), errors.WithStack(err)
		}
	}

	dbres := gdb.Unscoped().Delete(input)

	// Check error
	err = dbres.Error
	if err != nil {
		return *new(T), errors.WithStack(err)
	}

	// Return result
	return input, nil
}

func PermanentDeleteFiltered[T any](
	ctx context.Context,
	input T,
	filter any,
	db database.DB,
	opts ...GormOpt,
) error {
	// Create new options
	lOpts := make([]GormOpt, 0)
	// Save input options
	lOpts = append(lOpts, opts...)
	// Append filter
	lOpts = append(lOpts, WithFilterGormOpt(filter))

	// Patch
	_, err := PermanentDelete(ctx, input, db, lOpts...)
	// Return result
	return err
}

func SoftDelete[T any](
	ctx context.Context,
	input T,
	db database.DB,
	opts ...GormOpt,
) (T, error) {
	// Get gorm gdb
	gdb := db.GetTransactionalOrDefaultGormDB(ctx)

	// Define error
	var err error
	// Apply options
	for _, o := range opts {
		gdb, err = o(ctx, gdb)
		// Check error
		if err != nil {
			return *new(T), errors.WithStack(err)
		}
	}

	dbres := gdb.Delete(input)

	// Check error
	err = dbres.Error
	if err != nil {
		return *new(T), errors.WithStack(err)
	}

	// Return result
	return input, nil
}

func SoftDeleteFiltered[T any](
	ctx context.Context,
	input T,
	filter any,
	db database.DB,
	opts ...GormOpt,
) error {
	// Create new options
	lOpts := make([]GormOpt, 0)
	// Save input options
	lOpts = append(lOpts, opts...)
	// Append filter
	lOpts = append(lOpts, WithFilterGormOpt(filter))

	// Patch
	_, err := SoftDelete(ctx, input, db, lOpts...)
	// Return result
	return err
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
	input map[string]any,
	db database.DB,
	opts ...GormOpt,
) (T, error) {
	// Get gorm gdb
	gdb := db.GetTransactionalOrDefaultGormDB(ctx)

	// Define error
	var err error
	// Apply options
	for _, o := range opts {
		gdb, err = o(ctx, gdb)
		// Check error
		if err != nil {
			return *new(T), errors.WithStack(err)
		}
	}

	dbres := gdb.Model(originalObject).Updates(input)

	// Check error
	err = dbres.Error
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
	input map[string]any,
	filter any,
	db database.DB,
	opts ...GormOpt,
) error {
	// Create new options
	lOpts := make([]GormOpt, 0)
	// Save input options
	lOpts = append(lOpts, opts...)
	// Append filter
	lOpts = append(lOpts, WithFilterGormOpt(filter))

	// Patch
	_, err := PatchUpdate(ctx, model, input, db, lOpts...)
	// Return result
	return err
}
