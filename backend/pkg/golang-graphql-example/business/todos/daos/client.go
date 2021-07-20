package daos

import (
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/pagination"
)

//go:generate mockgen -destination=./mocks/mock_Doa.go -package=mocks github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/daos Dao
type Dao interface {
	MigrateDB() error
	GetAllPaginated(
		page *pagination.PageInput,
		sort *models.SortOrder,
		filter *models.Filter,
		projection *models.Projection,
	) ([]*models.Todo, *pagination.PageOutput, error)
	CreateOrUpdate(tt *models.Todo) (*models.Todo, error)
	FindByID(id string, projection *models.Projection) (*models.Todo, error)
}

func NewDao(db database.DB) Dao {
	return &dao{db: db}
}
