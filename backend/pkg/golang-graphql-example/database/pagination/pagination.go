package pagination

import (
	"gorm.io/gorm"
)

const dbColTagName = "dbfield"

// PageInput represents an input pagination configuration.
type PageInput struct {
	Skip  int
	Limit int
}

// PageOutput represents an output pagination structure.
type PageOutput struct {
	TotalRecord int
	Offset      int
	Limit       int
	Skip        int
	HasPrevious bool
	HasNext     bool
}

// Paging.
func Paging(result interface{}, db *gorm.DB, p *PageInput, sort interface{}, filter interface{}, extraFunc func(db *gorm.DB) *gorm.DB) (*PageOutput, error) {
	// Manage default limit
	if p.Limit == 0 {
		p.Limit = 10
	}

	// Count all objects
	var count int64 = 0
	db = db.Model(result).Count(&count)
	// Check error
	if db.Error != nil {
		return nil, db.Error
	}

	// Apply sort
	db, err := manageSortOrder(sort, db)
	// Check error
	if err != nil {
		return nil, err
	}

	// Apply filter
	db, err = manageFilter(filter, db, db, false)
	// Check error
	if err != nil {
		return nil, err
	}

	// Extra function
	if extraFunc != nil {
		db = extraFunc(db)
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
