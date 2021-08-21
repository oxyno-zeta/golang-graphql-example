package business

import (
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/authorization"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/migration"
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
	// Create migration service
	migrationSvc := migration.New(s.db)

	return migrationSvc.Migrate()
}

func NewServices(systemLogger log.Logger, db database.DB, authSvc authorization.Service, ld lockdistributor.Service) *Services {
	// Create todos service
	todoSvc := todos.NewService(db, authSvc)

	return &Services{
		db:           db,
		systemLogger: systemLogger,
		TodoSvc:      todoSvc,
	}
}
