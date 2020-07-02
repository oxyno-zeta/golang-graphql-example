package main

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/metrics"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/tracing"
	"golang.org/x/sync/errgroup"
)

func main() {
	// Create new logger
	logger := log.NewLogger()

	// Create configuration manager
	cfgManager := config.NewManager(logger)

	// Load configuration
	err := cfgManager.Load()
	if err != nil {
		logger.Fatal(err)
	}

	// Get configuration
	cfg := cfgManager.GetConfig()
	// Configure logger
	err = logger.Configure(cfg.Log.Level, cfg.Log.Format, cfg.Log.FilePath)
	if err != nil {
		logger.Fatal(err)
	}

	// Watch change for logger (special case)
	cfgManager.AddOnChangeHook(func() {
		// Get configuration
		cfg := cfgManager.GetConfig()
		// Configure logger
		err = logger.Configure(cfg.Log.Level, cfg.Log.Format, cfg.Log.FilePath)
		if err != nil {
			logger.Error(err)
		}
	})

	// Create metrics client
	metricsCl := metrics.NewMetricsClient()

	// Generate tracing service instance
	tracingSvc, err := tracing.New(cfgManager, logger)
	// Check error
	if err != nil {
		logger.Fatal(err)
	}
	// Prepare on reload hook
	cfgManager.AddOnChangeHook(func() {
		err2 := tracingSvc.Reload()
		if err2 != nil {
			logger.Fatal(err2)
		}
	})

	// Create database service
	db := database.NewDatabase(cfgManager, logger)
	// Connect to engine
	err = db.Connect()
	if err != nil {
		logger.Fatal(err)
	}
	// Add configuration reload hook
	cfgManager.AddOnChangeHook(func() {
		err := db.Reconnect()
		if err != nil {
			logger.Fatal(err)
		}
	})

	// Create business services
	busServices := business.NewServices(db)

	// Migrate database
	err = busServices.MigrateDB()
	if err != nil {
		logger.Fatal(err)
	}

	// Create servers
	svr := server.NewServer(logger, cfgManager, metricsCl, tracingSvc, busServices)
	intSvr := server.NewInternalServer(logger, cfgManager, metricsCl)

	// Add checker for internal server
	intSvr.AddChecker(&server.CheckerInput{
		Name:     "database",
		CheckFn:  db.Ping,
		Interval: 2 * time.Second,
	})

	var g errgroup.Group

	g.Go(svr.Listen)
	g.Go(intSvr.Listen)

	if err := g.Wait(); err != nil {
		logger.Fatal(err)
	}
}
