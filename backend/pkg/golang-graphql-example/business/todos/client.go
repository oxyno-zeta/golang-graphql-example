package todos

import (
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/daos"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
)

type Service interface {
	MigrateDB(systemLogger log.Logger) error
	GetAll() ([]*models.Todo, error)
	Create(inp *InputCreateTodo) (*models.Todo, error)
	Update(inp *InputUpdateTodo) (*models.Todo, error)
	Close(id string) (*models.Todo, error)
}

type InputCreateTodo struct {
	Text string
}

type InputUpdateTodo struct {
	ID   string
	Text string
}

func NewService(db database.DB) Service {
	// Create dao
	dao := daos.NewDao(db)

	return &service{dao: dao}
}
