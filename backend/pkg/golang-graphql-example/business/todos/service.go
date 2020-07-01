package todos

import (
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
)

type service struct {
}

func (s *service) GetAll() ([]*models.Todo, error)                   { return nil, nil }
func (s *service) Create(inp *InputCreateTodo) (*models.Todo, error) { return nil, nil }
func (s *service) Close(id string) (*models.Todo, error)             { return nil, nil }
