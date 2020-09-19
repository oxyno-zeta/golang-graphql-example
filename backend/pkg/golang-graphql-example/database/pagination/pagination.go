package pagination

import (
	"gorm.io/gorm"
)

const dbColTagName = "dbcolumn"

type PageInput struct {
	Skip  int
	Limit int
}

// PageOutput.
type PageOutput struct {
	TotalRecord int
	Offset      int
	Limit       int
	Skip        int
	HasPrevious bool
	HasNext     bool
}

// Paging.
func Paging(db *gorm.DB, filterFunc func(db *gorm.DB) *gorm.DB, orderBy []string, p *PageInput, result interface{}) (*PageOutput, error) {
	if p.Limit == 0 {
		p.Limit = 10
	}

	// Filter function
	if filterFunc != nil {
		db = filterFunc(db)
	}

	var count int64 = 0
	db = db.Model(result).Count(&count)
	// Check error
	if db.Error != nil {
		return nil, db.Error
	}

	// Check if order by exists
	if len(orderBy) == 0 {
		// Set default
		orderBy = []string{"created_at DESC"}
	}

	// Apply order by
	for _, o := range orderBy {
		db = db.Order(o)
	}

	var paginator PageOutput

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
