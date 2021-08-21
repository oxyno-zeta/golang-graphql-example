package migration

import "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"

// Client will represent the service that will migrate database.
type Client interface {
	// Migrate.
	Migrate() error
}

func New(dbSvc database.DB) Client {
	return &service{dbSvc: dbSvc}
}
