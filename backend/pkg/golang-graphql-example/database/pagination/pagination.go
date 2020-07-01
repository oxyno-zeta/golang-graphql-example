package pagination

import (
	"math"

	"github.com/jinzhu/gorm"
)

type PaginationInput struct {
	Page    int
	Limit   int
	OrderBy []string
}

// Paginator
type Paginator struct {
	TotalRecord int
	TotalPage   int
	Offset      int
	Limit       int
	Page        int
	PrevPage    int
	NextPage    int
}

// Paging
func Paging(db *gorm.DB, p *PaginationInput, result interface{}) (*Paginator, error) {
	if p.Page < 1 {
		p.Page = 1
	}

	if p.Limit == 0 {
		p.Limit = 10
	}

	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			db = db.Order(o)
		}
	}

	var paginator Paginator

	var offset int

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.Limit
	}

	// TODO Try to do all request in //
	db = db.Limit(p.Limit).Offset(offset).Find(result)
	if db.Error != nil {
		return nil, db.Error
	}

	count, err := countRecords(db, result)
	if err != nil {
		return nil, err
	}

	paginator.TotalRecord = count
	paginator.Page = p.Page

	paginator.Offset = offset
	paginator.Limit = p.Limit
	paginator.TotalPage = int(math.Ceil(float64(count) / float64(p.Limit)))

	if p.Page > 1 {
		paginator.PrevPage = p.Page - 1
	} else {
		paginator.PrevPage = p.Page
	}

	if p.Page == paginator.TotalPage {
		paginator.NextPage = p.Page
	} else {
		paginator.NextPage = p.Page + 1
	}

	return &paginator, nil
}

func countRecords(db *gorm.DB, anyType interface{}) (int, error) {
	var count int
	qres := db.Model(anyType).Count(count)
	return count, qres.Error
}
