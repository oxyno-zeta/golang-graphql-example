package daos

import (
	"context"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/pagination"
)

//go:generate mockgen -destination=./mocks/mock_Doa.go -package=mocks github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/daos Dao
type Dao interface {
	Find(
		ctx context.Context,
		sort []*models.SortOrder,
		filter *models.Filter,
		projection *models.Projection,
	) ([]*models.Todo, error)
	GetAllPaginated(
		ctx context.Context,
		page *pagination.PageInput,
		sort []*models.SortOrder,
		filter *models.Filter,
		projection *models.Projection,
	) ([]*models.Todo, *pagination.PageOutput, error)
	CreateOrUpdate(ctx context.Context, tt *models.Todo) (*models.Todo, error)
	PatchUpdate(
		ctx context.Context,
		tt *models.Todo,
		input map[string]interface{},
	) (*models.Todo, error)
	FindByID(ctx context.Context, id string, projection *models.Projection) (*models.Todo, error)
}

func NewDao(db database.DB) Dao {
	return &dao{db: db}
}
