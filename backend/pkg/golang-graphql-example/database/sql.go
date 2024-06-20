package database

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"emperror.dev/errors"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/deltaplugin"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/metrics"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/tracing"
	"github.com/thoas/go-funk"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var (
	transactionContextKey = &contextKey{name: "TRANSACTION"}
	transactionTraceName  = "database:execute-transaction"
)

type sqldb struct {
	logger                log.Logger
	cfgManager            config.Manager
	metricsSvc            metrics.Service
	tracingSvc            tracing.Service
	db                    *gorm.DB
	deltaNotificationChan chan *deltaplugin.Delta
	connectionName        string
}

func SetTransactionalGormDBToContext(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, transactionContextKey, db)
}

func GetTransactionalGormDBFromContext(ctx context.Context) *gorm.DB {
	// Get transactional db from context
	resInt := ctx.Value(transactionContextKey)

	// Check if exists
	if resInt == nil {
		return nil
	}

	// Cast it
	res, _ := resInt.(*gorm.DB)

	return res
}

func (sdb *sqldb) ExecuteTransaction(ctx context.Context, cb func(context.Context) error, opts ...TransactionOption) error {
	// Create transaction callback
	txCb := func(tx *gorm.DB) (err error) {
		// Get parent trace
		parentTrace := tracing.GetTraceFromContext(ctx)
		// Create child trace
		cctx, childTrace := parentTrace.GetChildTrace(ctx, transactionTraceName)
		// Defer close
		defer func() {
			// Check error
			if err != nil {
				childTrace.MarkAsError()
			}
			// Finish trace
			childTrace.Finish()
		}()

		// Inject transactional db in context
		newCtx := SetTransactionalGormDBToContext(cctx, tx)

		// Callback
		return cb(newCtx)
	}

	// Create options
	optCfg := &TransactionOptionsConfig{}
	// Apply options
	for _, fn := range opts {
		fn(optCfg)
	}

	db := sdb.db
	sqlOpts := &sql.TxOptions{}
	// Apply clauses for read transaction
	if optCfg.ReadTransaction {
		db = db.Clauses(dbresolver.Read)
		// Mark sql options too
		sqlOpts.ReadOnly = true
	} else {
		db = db.Clauses(dbresolver.Write)
	}

	// Add isolation
	sqlOpts.Isolation = optCfg.IsolationLevel

	return db.Transaction(txCb, sqlOpts)
}

func (sdb *sqldb) GetTransactionalOrDefaultGormDB(ctx context.Context) *gorm.DB {
	// Get transactional gorm db in context
	tx := GetTransactionalGormDBFromContext(ctx)

	// Check if last exists
	if tx != nil {
		// Return transaction and add context to gorm
		return tx.WithContext(ctx)
	}

	// Default
	return sdb.GetGormDB().WithContext(ctx)
}

func (sdb *sqldb) GetSQLDB() (*sql.DB, error) {
	return sdb.db.DB()
}

// GetGormDB will return a gorm database object.
func (sdb *sqldb) GetGormDB() *gorm.DB {
	return sdb.db
}

// Connect will connect to database engine.
func (sdb *sqldb) Connect() error {
	// Get configuration
	cfg := sdb.cfgManager.GetConfig()

	var sqlConnectionMaxLifetimeDuration time.Duration
	// Try to parse sql connection max lifetime duration
	if cfg.Database.SQLConnectionMaxLifetimeDuration != "" {
		// Initialize
		var err error
		// Parse time
		sqlConnectionMaxLifetimeDuration, err = time.ParseDuration(cfg.Database.SQLConnectionMaxLifetimeDuration)
		// Check error
		if err != nil {
			return errors.WithStack(err)
		}
	}

	// Create gorm configuration
	gcfg := &gorm.Config{
		// Insert now function to be sure that automatic dates are in UTC
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		// Add logger
		Logger: sdb.logger.GetGormLogger(),
		// Disable foreign key constraint when migrating
		DisableForeignKeyConstraintWhenMigrating: cfg.Database.DisableForeignKeyWhenMigrating,
		// Allow global update
		AllowGlobalUpdate: cfg.Database.AllowGlobalUpdate,
		// Prepare statement for caching
		PrepareStmt: cfg.Database.PrepareStatement,
	}

	// Select postgres driver
	openFunction := postgres.Open
	// Check if sqlite driver is selected
	if cfg.Database.Driver == SqliteDriverSelector {
		openFunction = sqlite.Open
	}

	// Trim url
	sURL := strings.TrimSpace(cfg.Database.ConnectionURL.Value)

	sdb.logger.Debugf("Trying to connect to database engine of type %s", cfg.Database.Driver)
	// Connect to database
	dbResult, err := gorm.Open(openFunction(sURL), gcfg)
	// Check if error exists
	if err != nil {
		return errors.WithStack(err)
	}

	// Check if there are replica in configuration
	if len(cfg.Database.ReplicaConnectionURLs) != 0 {
		// Create connections list
		conns, _ := funk.Map(cfg.Database.ReplicaConnectionURLs, func(sc *config.CredentialConfig) gorm.Dialector {
			return openFunction(strings.TrimSpace(sc.Value))
		}).([]gorm.Dialector)

		// Inject db resolver configuration
		err = dbResult.Use(dbresolver.Register(dbresolver.Config{
			Replicas: conns,
			// sources/replicas load balancing policy
			Policy: dbresolver.RandomPolicy{},
			// print sources/replicas mode in logger
			TraceResolverMode: true,
		}))
		// Check error
		if err != nil {
			return errors.WithStack(err)
		}
	}

	// Check if delta plugin is wanted
	if sdb.deltaNotificationChan != nil {
		err = dbResult.Use(deltaplugin.New(sdb.deltaNotificationChan))
		// Check if error exists
		if err != nil {
			return errors.WithStack(err)
		}
	}

	// Get prometheus gorm middleware
	md := sdb.metricsSvc.DatabaseMiddleware(sdb.connectionName)
	// Apply middleware
	err = dbResult.Use(md)
	// Check if error exists
	if err != nil {
		return errors.WithStack(err)
	}

	// Apply tracing middleware
	err = dbResult.Use(sdb.tracingSvc.DatabaseMiddleware())
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

	// Check if max idle connections exists
	if cfg.Database.SQLMaxIdleConnections != 0 {
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxIdleConns(cfg.Database.SQLMaxIdleConnections)
	}

	// Check if max opened connections exists
	if cfg.Database.SQLMaxOpenConnections != 0 {
		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(cfg.Database.SQLMaxOpenConnections)
	}

	// Check if connection max lifetime exists
	if cfg.Database.SQLConnectionMaxLifetimeDuration != "" {
		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDB.SetConnMaxLifetime(sqlConnectionMaxLifetimeDuration)
	}

	// Save gorm db object
	sdb.db = dbResult

	sdb.logger.Infof("Successfully connected to database engine of type %s", cfg.Database.Driver)

	// Return
	return nil
}

// Close will close connection to database.
func (sdb *sqldb) Close() error {
	sdb.logger.Info("Closing database connection")
	// Get sql database
	sqlDB, err := sdb.db.DB()
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	// Close database connection
	err = sqlDB.Close()
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	// Default
	return nil
}

// Ping will ping database engine in order to test connection to engine.
func (sdb *sqldb) Ping() error {
	// Get sql database
	sqlDB, err := sdb.db.DB()
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

func (sdb *sqldb) Reconnect() error {
	// Get old sql db object
	sqlDB, err := sdb.db.DB()
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}
	// Connect to new database
	err = sdb.Connect()
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	// Wait for 1 sec before closing connection to old db
	// Here, we suppose that waiting 1 second is enough for reload
	time.Sleep(time.Second)

	// Closing old database connection
	err = sqlDB.Close()
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
