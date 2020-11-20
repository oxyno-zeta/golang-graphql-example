package todos

import (
	"context"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/daos"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/pagination"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
)

type authorizationService interface {
	CheckAuthorized(ctx context.Context, action, resource string) error
}

type Service interface {
	MigrateDB(systemLogger log.Logger) error
	GetAllPaginated(
		ctx context.Context,
		page *pagination.PageInput,
		sort *models.SortOrder,
		filter *models.Filter,
		projection *models.Projection,
	) ([]*models.Todo, *pagination.PageOutput, error)
	FindByID(ctx context.Context, id string, projection *models.Projection) (*models.Todo, error)
	Create(ctx context.Context, inp *InputCreateTodo) (*models.Todo, error)
	Update(ctx context.Context, inp *InputUpdateTodo) (*models.Todo, error)
	Close(ctx context.Context, id string) (*models.Todo, error)
}

type InputCreateTodo struct {
	Text string
}

type InputUpdateTodo struct {
	ID   string
	Text string
}

func NewService(db database.DB, authSvc authorizationService) Service {
	// Create dao
	dao := daos.NewDao(db)

	return &service{dao: dao, authSvc: authSvc}
}
