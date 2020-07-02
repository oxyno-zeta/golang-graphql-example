package daos

import (
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
)

type dao struct {
	db database.DB
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

func (d *dao) GetAll() ([]*models.Todo, error) {
	// Get gorm db
	db := d.db.GetGormDB()
	// result
	res := make([]*models.Todo, 0)
	// Find todos
	dbres := db.Find(&res)

	err := dbres.Error
	// Check error
	if err != nil {
		return nil, err
	}

	return res, nil
}
