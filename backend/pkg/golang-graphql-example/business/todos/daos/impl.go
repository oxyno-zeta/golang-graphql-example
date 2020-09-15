package daos

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/pagination"
)

type dao struct {
	db database.DB
}

func (d *dao) MigrateDB() error {
	// Get gorm database
	gdb := d.db.GetGormDB()
	// Migrate
	res := gdb.AutoMigrate(&models.Todo{})

	return res.Error
}

func (d *dao) FindByID(id string) (*models.Todo, error) {
	// Get gorm db
	db := d.db.GetGormDB()
	// result
	res := &models.Todo{}
	// Find in db
	dbres := db.Where("id = ?", id).First(res)

	// Check error
	err := dbres.Error
	if err != nil {
		// Check if it is a not found error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return res, nil
}

func (d *dao) CreateOrUpdate(tt *models.Todo) (*models.Todo, error) {
	// Get gorm db
	db := d.db.GetGormDB()
	dbres := db.Save(tt)

	// Check error
	err := dbres.Error
	if err != nil {
		return nil, err
	}

	// Return result
	return tt, nil
}

func (d *dao) GetAllPaginated(page *pagination.PageInput, sort *models.SortOrder) ([]*models.Todo, *pagination.PageOutput, error) {
	// Get gorm db
	db := d.db.GetGormDB()
	// result
	res := make([]*models.Todo, 0)
	// Find todos
	pageOut, err := pagination.Paging(db, nil, nil, page, &res)
	// Check error
	if err != nil {
		return nil, nil, err
	}

	return res, pageOut, nil
}
