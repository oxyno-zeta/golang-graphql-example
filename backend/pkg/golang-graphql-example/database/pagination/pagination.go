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

	// Apply filter
	db, err := common.ManageFilter(filter, db)
	// Check error
	if err != nil {
		return nil, err
	}

	// Extra function
	if extraFunc != nil {
		db, err = extraFunc(db)
		// Check error
		if err != nil {
			return nil, err
		}
	}

	// Count all objects
	var count int64 = 0
	db = db.Model(result).Count(&count)
	// Check error
	if db.Error != nil {
		return nil, db.Error
	}

	// Apply sort
	db, err = common.ManageSortOrder(sort, db)
	// Check error
	if err != nil {
		return nil, err
	}

	// Create paginator output
	var paginator PageOutput

	// Request to database with limit and offset
	db = db.Limit(p.Limit).Offset(p.Skip).Find(result)
	// Check error
	if db.Error != nil {
		return nil, db.Error
	}

	paginator.TotalRecord = int(count)
	paginator.Skip = p.Skip
	paginator.Limit = p.Limit

	paginator.HasNext = (p.Limit+p.Skip < paginator.TotalRecord)
	paginator.HasPrevious = (p.Skip != 0)

	return &paginator, nil
}
