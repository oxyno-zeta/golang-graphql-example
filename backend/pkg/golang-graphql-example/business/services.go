package business

import (
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/lockdistributor"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
)

type Services struct {
	db           database.DB
	systemLogger log.Logger
	TodoSvc      todos.Service
}

func (s *Services) MigrateDB() error {
	return s.TodoSvc.MigrateDB(s.systemLogger)
}

func NewServices(systemLogger log.Logger, db database.DB, ld lockdistributor.Service) *Services {
	// Create todos service
	todoSvc := todos.NewService(db)

	return &Services{
		db:           db,
		systemLogger: systemLogger,
		TodoSvc:      todoSvc,
	}
}
