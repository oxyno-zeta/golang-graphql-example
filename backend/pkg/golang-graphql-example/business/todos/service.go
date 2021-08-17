package todos

import (
	"context"
	"fmt"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/daos"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/pagination"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
)

const mainAuthorizationPrefix = "todo"

type service struct {
	dao     daos.Dao
	authSvc AuthorizationService
	dbSvc   database.DB
}

func (s *service) MigrateDB(systemLogger log.Logger) error {
	systemLogger.Debug("Migrate database for Todos")

	return s.dao.MigrateDB()
}

func (s *service) FindByID(ctx context.Context, id string, projection *models.Projection) (*models.Todo, error) {
	// Check authorization
	err := s.authSvc.CheckAuthorized(
		ctx,
		fmt.Sprintf("%s:%s", mainAuthorizationPrefix, "Get"),
		fmt.Sprintf("%s:%s", mainAuthorizationPrefix, id),
	)
	// Check error
	if err != nil {
		return nil, err
	}

	// Find by id
	res, err := s.dao.FindByID(ctx, id, projection)
	// Check error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *service) GetAllPaginated(
	ctx context.Context,
	page *pagination.PageInput,
	sort *models.SortOrder,
	filter *models.Filter,
	projection *models.Projection,
) ([]*models.Todo, *pagination.PageOutput, error) {
	// Check authorization
	err := s.authSvc.CheckAuthorized(
		ctx,
		fmt.Sprintf("%s:%s", mainAuthorizationPrefix, "List"),
		"",
	)
	// Check error
	if err != nil {
		return nil, nil, err
	}

	return s.dao.GetAllPaginated(ctx, page, sort, filter, projection)
}

func (s *service) Create(ctx context.Context, inp *InputCreateTodo) (*models.Todo, error) {
	// Check authorization
	err := s.authSvc.CheckAuthorized(
		ctx,
		fmt.Sprintf("%s:%s", mainAuthorizationPrefix, "Create"),
		"",
	)
	// Check error
	if err != nil {
		return nil, err
	}

	tt := &models.Todo{
		Text: inp.Text,
	}

	return s.dao.CreateOrUpdate(ctx, tt)
}

func (s *service) Update(ctx context.Context, inp *InputUpdateTodo) (*models.Todo, error) {
	// Check authorization
	err := s.authSvc.CheckAuthorized(
		ctx,
		fmt.Sprintf("%s:%s", mainAuthorizationPrefix, "Update"),
		fmt.Sprintf("%s:%s", mainAuthorizationPrefix, inp.ID),
	)
	// Check error
	if err != nil {
		return nil, err
	}

	// Search by id first
	tt, err := s.dao.FindByID(ctx, inp.ID, nil)
	if err != nil {
		return nil, err
	}
	// Update text in existing result
	tt.Text = inp.Text
	// Save
	return s.dao.CreateOrUpdate(ctx, tt)
}

func (s *service) Close(ctx context.Context, id string) (*models.Todo, error) {
	// Check authorization
	err := s.authSvc.CheckAuthorized(
		ctx,
		fmt.Sprintf("%s:%s", mainAuthorizationPrefix, "Close"),
		fmt.Sprintf("%s:%s", mainAuthorizationPrefix, id),
	)
	// Check error
	if err != nil {
		return nil, err
	}

	// Prepare result
	var res *models.Todo

	// Create transaction
	err = s.dbSvc.ExecuteTransaction(ctx, func(c context.Context) error {
		// Search by id first
		tt, err2 := s.dao.FindByID(ctx, id, nil)
		// Check error
		if err2 != nil {
			return err2
		}
		// Update text in existing result
		tt.Done = true
		// Save
		res, err2 = s.dao.CreateOrUpdate(ctx, tt)

		return err2
	})
	// Check error
	if err != nil {
		return nil, err
	}

	return res, nil
}
