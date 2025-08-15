package pagination

import (
	"context"

	"emperror.dev/errors"
	"gorm.io/gorm"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/common"
)

// PageInput represents an input pagination configuration.
type PageInput struct {
	Skip  int
	Limit int
}

// PageOutput represents an output pagination structure.
type PageOutput struct {
	TotalRecord int
	Limit       int
	Skip        int
	HasPrevious bool
	HasNext     bool
}

// PagingOptions represents pagination options.
type PagingOptions struct {
	// DB Service
	DBSvc database.DB
	// Pagination input
	PageInput *PageInput
	// Must be a pointer to an object with *SortOrderEnum objects with tags
	Sort any
	// Must be a pointer to an object with *GenericFilter objects or implementing the GenericFilterBuilder interface and with tags
	Filter any
	// Must be a pointer to an object with booleans with tags
	Projection any
	// This function is called after filters and before any sorts
	ExtraFunc func(db *gorm.DB) (*gorm.DB, error)
	// Transaction options
	TOpts []database.TransactionOption
}

// Paging function in order to have a paginated sorted and filters list of objects.
// Parameters:
// - result: Must be a pointer to a list of objects
// - options: Pagination options
// .
func Paging(
	ctx context.Context,
	result any,
	options *PagingOptions,
) (*PageOutput, error) {
	// Manage default limit
	if options.PageInput.Limit == 0 {
		options.PageInput.Limit = 10
	}

	// Initialize
	var count int64

	// Create local transaction options
	localTOpts := options.TOpts
	// Check if nil
	if localTOpts == nil {
		// Init it
		localTOpts = make([]database.TransactionOption, 0)
	}
	// Add local options
	localTOpts = append(localTOpts, database.WithReadTransactionOpt)

	// Create transaction to avoid situations where count and find are different
	err := options.DBSvc.ExecuteTransaction(ctx, func(ctx context.Context) error {
		// Get gorm db
		db := options.DBSvc.GetTransactionalOrDefaultGormDB(ctx)
		// Apply filter
		db, err := common.ManageFilter(options.Filter, db)
		// Check error
		if err != nil {
			return err
		}

		// Extra function
		if options.ExtraFunc != nil {
			db, err = options.ExtraFunc(db)
			// Check error
			if err != nil {
				return err
			}
		}

		// Count all objects
		db = db.Model(result).Count(&count)
		// Check error
		if db.Error != nil {
			return errors.WithStack(db.Error)
		}

		// Apply sort
		db, err = common.ManageSortOrder(options.Sort, db)
		// Check error
		if err != nil {
			return err
		}

		// Apply projection
		db, err = common.ManageProjection(options.Projection, db)
		// Check error
		if err != nil {
			return err
		}

		// Request to database with limit and offset
		db = db.Limit(options.PageInput.Limit).Offset(options.PageInput.Skip).Find(result)
		// Check error
		if db.Error != nil {
			return errors.WithStack(db.Error)
		}

		return nil
	}, localTOpts...)
	// Check error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return getPageOutput(options.PageInput, count), nil
}

func getPageOutput(p *PageInput, count int64) *PageOutput {
	var paginator PageOutput
	// Create total record
	paginator.TotalRecord = int(count)
	// Store skip
	paginator.Skip = p.Skip
	// Store limit
	paginator.Limit = p.Limit
	// Calculate has next page
	paginator.HasNext = (p.Limit+p.Skip < paginator.TotalRecord)
	// Calculate has previous page
	paginator.HasPrevious = (p.Skip != 0)

	return &paginator
}
