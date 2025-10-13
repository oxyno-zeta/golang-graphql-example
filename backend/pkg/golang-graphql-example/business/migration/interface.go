package migration

import (
	"context"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
)

// Service will represent the service that will migrate database.
type Service interface {
	// Migrate.
	Migrate(ctx context.Context) error
}

func New(dbSvc database.DB) Service {
	return &service{dbSvc: dbSvc}
}
