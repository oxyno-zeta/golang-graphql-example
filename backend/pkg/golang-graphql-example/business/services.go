package business

import (
	"context"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/authorization"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/migration"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/metrics"
)

type Services struct {
	db           database.DB
	systemLogger log.Logger
	TodoSvc      todos.Service
}

func (s *Services) MigrateDB(ctx context.Context) error {
	// Create migration service
	migrationSvc := migration.New(s.db)

	return migrationSvc.Migrate(ctx)
}

func (s *Services) GetBusinessMetricDefinitions() []*metrics.BusinessMetricDefinition {
	return s.TodoSvc.GetBusinessMetricDefinitions()
}

func NewServices(systemLogger log.Logger, db database.DB, authSvc authorization.Service) *Services {
	// Create todos service
	todoSvc := todos.NewService(db, authSvc)

	return &Services{
		db:           db,
		systemLogger: systemLogger,
		TodoSvc:      todoSvc,
	}
}
