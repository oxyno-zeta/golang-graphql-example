package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.70

import (
	"context"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/dataloaders"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/dataloaders/common"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/generated"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/mappers"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/model"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/utils"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*models.Todo, error) {
	inp := &todos.InputCreateTodo{Text: input.Text}
	tt, err := r.BusiServices.TodoSvc.Create(ctx, inp)
	// Check error
	if err != nil {
		return nil, err
	}

	return tt, nil
}

// CloseTodo is the resolver for the closeTodo field.
func (r *mutationResolver) CloseTodo(ctx context.Context, todoID string) (*models.Todo, error) {
	// Manage relay id
	bid, err := utils.FromIDRelay(todoID, mappers.TodoIDPrefix)
	// Check error
	if err != nil {
		return nil, err
	}

	// Get projection
	proj := &models.Projection{}
	err = utils.ManageSimpleProjection(ctx, proj)
	// Check error
	if err != nil {
		return nil, err
	}

	res, err := r.BusiServices.TodoSvc.Close(ctx, bid, proj)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, input *model.UpdateTodo) (*models.Todo, error) {
	// Manage relay id
	bid, err := utils.FromIDRelay(input.ID, mappers.TodoIDPrefix)
	// Check error
	if err != nil {
		return nil, err
	}

	inp := &todos.InputUpdateTodo{ID: bid, Text: input.Text}
	tt, err := r.BusiServices.TodoSvc.Update(ctx, inp)
	// Check error
	if err != nil {
		return nil, err
	}

	return tt, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context, after *string, before *string, first *int, last *int, sort *models.SortOrder, sorts []*models.SortOrder, filter *models.Filter) (*model.TodoConnection, error) {
	// Create pagination input
	pageInput, err := utils.GetPageInput(after, before, first, last)
	// Check error
	if err != nil {
		return nil, err
	}

	// Build projection from graphql fields
	projection := &models.Projection{}
	err = utils.ManageConnectionNodeProjection(ctx, projection)
	// Check error
	if err != nil {
		return nil, err
	}

	// Manage deprecated sort
	if len(sorts) == 0 && sort != nil {
		sorts = []*models.SortOrder{sort}
	}

	// Call business
	allTodos, pageOut, err := r.BusiServices.TodoSvc.GetAllPaginated(ctx, pageInput, sorts, filter, projection)
	// Check error
	if err != nil {
		return nil, err
	}

	var res model.TodoConnection
	err = utils.MapConnection(&res, allTodos, pageOut)
	// Check error
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Todo is the resolver for the todo field.
func (r *queryResolver) Todo(ctx context.Context, id string) (*models.Todo, error) {
	// Get projection
	proj := &models.Projection{}
	err := utils.ManageSimpleProjection(ctx, proj)
	// Check error
	if err != nil {
		return nil, err
	}

	uuid, err := utils.FromIDRelay(id, mappers.TodoIDPrefix)
	// Check error
	if err != nil {
		return nil, err
	}

	// Example of dataloader query
	dl := dataloaders.GetDataloadersFromContext(ctx)
	res, err := dl.Todos.GenericLoader.Load(ctx, &common.IDProjectionKey{ID: uuid, Projection: proj})()

	// Call business (example)
	// res, err := r.BusiServices.TodoSvc.FindByID(ctx, uuid, proj)

	return res, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
