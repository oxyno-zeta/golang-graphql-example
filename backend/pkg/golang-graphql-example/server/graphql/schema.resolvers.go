package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/generated"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/mappers"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/model"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/utils"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*models.Todo, error) {
	inp := &todos.InputCreateTodo{Text: input.Text}
	tt, err := r.BusiServices.TodoSvc.Create(ctx, inp)
	// Check error
	if err != nil {
		return nil, err
	}

	return tt, nil
}

func (r *mutationResolver) CloseTodo(ctx context.Context, todoID string) (*models.Todo, error) {
	res, err := r.BusiServices.TodoSvc.Close(ctx, todoID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input *model.UpdateTodo) (*models.Todo, error) {
	inp := &todos.InputUpdateTodo{ID: input.ID, Text: input.Text}
	tt, err := r.BusiServices.TodoSvc.Update(ctx, inp)
	// Check error
	if err != nil {
		return nil, err
	}

	return tt, nil
}

func (r *queryResolver) Todos(ctx context.Context, after *string, before *string, first *int, last *int) (*model.TodoConnection, error) {
	// Create pagination input
	pageInput, err := utils.GetPageInput(after, before, first, last)
	// Check error
	if err != nil {
		return nil, err
	}

	// Call business
	allTodos, pageOut, err := r.BusiServices.TodoSvc.GetAllPaginated(ctx, pageInput)
	// Check error
	if err != nil {
		return nil, err
	}

	return mappers.MapTodoConnection(allTodos, pageOut), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
