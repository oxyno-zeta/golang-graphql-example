package todos

import (
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
)

type Service interface {
	GetAll() ([]*models.Todo, error)
	Create(inp *InputCreateTodo) (*models.Todo, error)
	Close(id string) (*models.Todo, error)
}

type InputCreateTodo struct {
	Text string
}

func NewService(db database.DB) Service {
	return &service{}
}
