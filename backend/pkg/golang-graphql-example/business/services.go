package business

import (
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos"
	todoModels "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/lockdistributor"
)

type Services struct {
	db      database.DB
	TodoSvc todos.Service
}

func (s *Services) MigrateDB() error {
	gdb := s.db.GetGormDB()

	// Run automigrate on all structures
	res := gdb.AutoMigrate(&todoModels.Todo{})

	return res.Error
}

func NewServices(db database.DB, ld lockdistributor.Service) *Services {
	// Create todos service
	todoSvc := todos.NewService(db)

	return &Services{
		db:      db,
		TodoSvc: todoSvc,
	}
}
