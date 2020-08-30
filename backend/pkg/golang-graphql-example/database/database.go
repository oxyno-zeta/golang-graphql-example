package database

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
)

//go:generate mockgen -destination=./mocks/mock_DB.go -package=mocks github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database DB
type DB interface {
	// Get Gorm db object.
	GetGormDB() *gorm.DB
	// Get SQL db object.
	GetSQLDB() *sql.DB
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
func NewDatabase(cfgManager config.Manager, logger log.Logger) DB {
	return &postresdb{logger: logger, cfgManager: cfgManager}
}
