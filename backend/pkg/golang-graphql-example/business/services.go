package business

import (
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos"
	todoModels "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
)

type Services struct {
	db      database.DB
	TodoSvc todos.Service
}

func (s *Services) Migrate() error {
	gdb := s.db.GetGormDB()

	// Run automigrate on all structures
	res := gdb.AutoMigrate(&todoModels.Todo{})

	return res.Error
}

func NewServices(db database.DB) *Services {
	// Create todos service
	todoSvc := todos.NewService(db)

	return &Services{
		db:      db,
		TodoSvc: todoSvc,
	}
}
