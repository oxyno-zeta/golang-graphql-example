package mappers

import (
	"time"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/model"
	"github.com/thoas/go-funk"
)

func MapTodo(tt *models.Todo) *model.Todo {
	return &model.Todo{
		CreationDate: tt.CreatedAt.Format(time.RFC3339),
		Done:         tt.Done,
		ID:           tt.ID,
		Text:         tt.Text,
	}
}

func MapTodos(tt []*models.Todo) []*model.Todo {
	// Create result
	res := make([]*model.Todo, 0)

	if tt == nil {
		return res
	}

	return funk.Map(tt, MapTodo).([]*model.Todo)
}
