package database

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"github.com/pkg/errors"

	// Import this for dialect usage in gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type postresdb struct {
	logger     log.Logger
	db         *gorm.DB
	cfgManager config.Manager
}

func (ctx *postresdb) GetSQLDB() *sql.DB {
	return ctx.db.DB()
}

// GetGormDB will return a gorm database object
func (ctx *postresdb) GetGormDB() *gorm.DB {
	return ctx.db
}

// Connect will connect to database engine
func (ctx *postresdb) Connect() error {
	// Get configuration
	cfg := ctx.cfgManager.GetConfig()

	ctx.logger.Debugf("Trying to connect to database engine of type %s", cfg.Database.Dialect)
	// Connect to database
	dbResult, err := gorm.Open(cfg.Database.Dialect, cfg.Database.ConnectionURL)
	// Check if error exists
	if err != nil {
		return errors.WithStack(err)
	}
	// Disable logger
	dbResult.LogMode(false)
	// Save gorm db object
	ctx.db = dbResult

	ctx.logger.Infof("Successfully connected to database engine of type %s", cfg.Database.Dialect)

	// Return
	return nil
}

// Close will close connection to database
func (ctx *postresdb) Close() error {
	ctx.logger.Info("Closing database connection")
	return ctx.db.Close()
}

// Ping will ping database engine in order to test connection to engine
func (ctx *postresdb) Ping() error {
	// Ping database to test connection
	pingErr := ctx.db.DB().Ping()
	// Check error
	if pingErr != nil {
		return errors.WithStack(pingErr)
	}

	return nil
}

func (ctx *postresdb) Reconnect() error {
	// Get old gorm database
	oldDB := ctx.db
	// Defer closing old database connection
	defer oldDB.Close()
	// Connect to new database
	err := ctx.Connect()
	if err != nil {
		return errors.WithStack(err)
	}
	// Wait for 1 sec before closing connection to old db
	// Here, we suppose that waiting 1 second is enough for reload
	time.Sleep(time.Second)

	return nil
}
