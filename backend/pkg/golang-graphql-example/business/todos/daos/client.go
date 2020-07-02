package daos

import (
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
)

type Dao interface {
	GetAll() ([]*models.Todo, error)
	CreateOrUpdate(tt *models.Todo) (*models.Todo, error)
	FindByID(id string) (*models.Todo, error)
}

func NewDao(db database.DB) Dao {
	return &dao{db: db}
}
