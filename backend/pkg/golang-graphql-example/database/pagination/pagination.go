package pagination

import (
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/common"
	"gorm.io/gorm"
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

// Paging function in order to have a paginated sorted and filters list of objects.
// Parameters:
// - result: Must be a pointer to a list of objects
// - db: Gorm database
// - p: Pagination input
// - sort: Must be a pointer to an object with *SortOrderEnum objects with tags
// - filter: Must be a pointer to an object with *GenericFilter objects or implementing the GenericFilterBuilder interface and with tags
// - extraFunc: This function is called after filters and before any sorts
// .
func Paging(
	result interface{},
	db *gorm.DB,
	p *PageInput,
	sort interface{},
	filter interface{},
	extraFunc func(db *gorm.DB) (*gorm.DB, error),
) (*PageOutput, error) {
	// Manage default limit
	if p.Limit == 0 {
		p.Limit = 10
	}

	var count int64 = 0

	// Create transaction to avoid situations where count and find are different
	err := db.Transaction(func(db *gorm.DB) error {
		// Apply filter
		db, err := common.ManageFilter(filter, db)
		// Check error
		if err != nil {
			return err
		}

		// Extra function
		if extraFunc != nil {
			db, err = extraFunc(db)
			// Check error
			if err != nil {
				return err
			}
		}

		// Count all objects
		db = db.Model(result).Count(&count)
		// Check error
		if db.Error != nil {
			return db.Error
		}

		// Apply sort
		db, err = common.ManageSortOrder(sort, db)
		// Check error
		if err != nil {
			return err
		}

		// Request to database with limit and offset
		db = db.Limit(p.Limit).Offset(p.Skip).Find(result)
		// Check error
		if db.Error != nil {
			return db.Error
		}

		return nil
	})

	// Check error
	if err != nil {
		return nil, err
	}

	return getPageOutput(p, count), nil
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
