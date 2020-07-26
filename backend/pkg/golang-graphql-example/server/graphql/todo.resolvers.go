package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/generated"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/utils"
)

const TodoIDPrefix = "todos"

func (r *todoResolver) ID(ctx context.Context, obj *models.Todo) (string, error) {
	return utils.ToIDRelay(TodoIDPrefix, obj.ID), nil
}

func (r *todoResolver) CreationDate(ctx context.Context, obj *models.Todo) (string, error) {
	return utils.FormatTime(obj.CreatedAt), nil
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
