package database

import (
	"database/sql"
	"time"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

type postresdb struct {
	logger     log.Logger
	db         *gorm.DB
	cfgManager config.Manager
}

func (ctx *postresdb) GetSQLDB() (*sql.DB, error) {
	return ctx.db.DB()
}

// GetGormDB will return a gorm database object.
func (ctx *postresdb) GetGormDB() *gorm.DB {
	return ctx.db
}

// Connect will connect to database engine.
func (ctx *postresdb) Connect() error {
	// Get configuration
	cfg := ctx.cfgManager.GetConfig()

	// Create gorm configuration
	gcfg := &gorm.Config{
		// Insert now function to be sure that automatic dates are in UTC
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		// Add logger
		Logger: ctx.logger.GetGormLogger(),
	}

	ctx.logger.Debug("Trying to connect to database engine of type PostgreSQL")
	// Connect to database
	dbResult, err := gorm.Open(postgres.Open(cfg.Database.ConnectionURL.Value), gcfg)
	// Check if error exists
	if err != nil {
		return errors.WithStack(err)
	}

	// Trying to ping database
	sqlDB, err := dbResult.DB()
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	// Ping db
	err = sqlDB.Ping()
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	// Save gorm db object
	ctx.db = dbResult

	ctx.logger.Info("Successfully connected to database engine of type PostgreSQL")

	// Return
	return nil
}

// Close will close connection to database.
func (ctx *postresdb) Close() error {
	ctx.logger.Info("Closing database connection")
	// Get sql database
	sqlDB, err := ctx.db.DB()
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	return sqlDB.Close()
}

// Ping will ping database engine in order to test connection to engine.
func (ctx *postresdb) Ping() error {
	// Get sql database
	sqlDB, err := ctx.db.DB()
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}
	// Ping database to test connection
	pingErr := sqlDB.Ping()
	// Check error
	if pingErr != nil {
		return errors.WithStack(pingErr)
	}

	return nil
}

func (ctx *postresdb) Reconnect() error {
	// Get old sql db object
	sqlDB, err := ctx.db.DB()
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}
	// Defer closing old database connection
	defer sqlDB.Close()
	// Connect to new database
	err = ctx.Connect()
	if err != nil {
		return errors.WithStack(err)
	}
	// Wait for 1 sec before closing connection to old db
	// Here, we suppose that waiting 1 second is enough for reload
	time.Sleep(time.Second)

	return nil
}
