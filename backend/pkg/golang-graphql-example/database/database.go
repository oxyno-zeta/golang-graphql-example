package database

import (
	"context"
	"database/sql"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/metrics"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/tracing"
	"gorm.io/gorm"
)

const (
	PostgresDriverSelector = "POSTGRES"
	SqliteDriverSelector   = "SQLITE"
)

//go:generate mockgen -destination=./mocks/mock_DB.go -package=mocks github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database DB
type DB interface {
	// ExecuteTransaction will execute a transaction.
	ExecuteTransaction(ctx context.Context, cb func(context.Context) error) error
	// GetTransactionalOrDefaultGormDB will return a transactional gorm db if it exists in context, otherwise the default db.
	GetTransactionalOrDefaultGormDB(ctx context.Context) *gorm.DB
	// Get Gorm db object.
	GetGormDB() *gorm.DB
	// Get SQL db object.
	GetSQLDB() (*sql.DB, error)
	// Connect to database.
	Connect() error
	// Close database connection.
	Close() error
	// Ping database.
	Ping() error
	// Reconnect to database.
	Reconnect() error
}

// NewDatabase will generate a new DB object.
func NewDatabase(
	connectionName string,
	cfgManager config.Manager,
	logger log.Logger,
	metricsCl metrics.Client,
	tracingSvc tracing.Service,
) DB {
	return &sqldb{
		logger:         logger,
		cfgManager:     cfgManager,
		metricsCl:      metricsCl,
		tracingSvc:     tracingSvc,
		connectionName: connectionName,
	}
}
