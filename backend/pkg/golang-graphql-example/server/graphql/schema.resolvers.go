package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/generated"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/mappers"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	inp := &todos.InputCreateTodo{Text: input.Text}
	tt, err := r.BusiServices.TodoSvc.Create(inp)
	// Check error
	if err != nil {
		return nil, err
	}

	return mappers.MapTodo(tt), nil
}

func (r *mutationResolver) CloseTodo(ctx context.Context, todoID string) (*model.Todo, error) {
	res, err := r.BusiServices.TodoSvc.Close(todoID)
	if err != nil {
		return nil, err
	}

	return mappers.MapTodo(res), nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input *model.UpdateTodo) (*model.Todo, error) {
	inp := &todos.InputUpdateTodo{ID: input.ID, Text: input.Text}
	tt, err := r.BusiServices.TodoSvc.Update(inp)
	// Check error
	if err != nil {
		return nil, err
	}

	return mappers.MapTodo(tt), nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	res, err := r.BusiServices.TodoSvc.GetAll()
	if err != nil {
		return nil, err
	}

	return mappers.MapTodos(res), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
