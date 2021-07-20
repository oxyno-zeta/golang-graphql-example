package daos

import (
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/common"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/pagination"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type dao struct {
	db database.DB
}

func (d *dao) MigrateDB() error {
	// Get gorm database
	gdb := d.db.GetGormDB()
	// Migrate
	err := gdb.AutoMigrate(&models.Todo{})
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (d *dao) FindByID(id string, projection *models.Projection) (*models.Todo, error) {
	// Get gorm db
	db := d.db.GetGormDB()
	// result
	res := &models.Todo{}
	// Apply projection
	db, err := common.ManageProjection(projection, db)
	// Check error
	if err != nil {
		return nil, err
	}
	// Find in db
	dbres := db.Where("id = ?", id).First(res)

	// Check error
	err = dbres.Error
	if err != nil {
		// Check if it is a not found error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.WithStack(err)
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
		return nil, errors.WithStack(err)
	}

	// Return result
	return tt, nil
}

func (d *dao) GetAllPaginated(
	page *pagination.PageInput,
	sort *models.SortOrder,
	filter *models.Filter,
	projection *models.Projection,
) ([]*models.Todo, *pagination.PageOutput, error) {
	// Get gorm db
	db := d.db.GetGormDB()
	// result
	res := make([]*models.Todo, 0)
	// Find todos
	pageOut, err := pagination.Paging(&res, &pagination.PagingOptions{
		DB:         db,
		PageInput:  page,
		Sort:       sort,
		Filter:     filter,
		Projection: projection,
		ExtraFunc:  nil,
	})
	// Check error
	if err != nil {
		return nil, nil, err
	}

	return res, pageOut, nil
}
