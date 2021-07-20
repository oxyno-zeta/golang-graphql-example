package main

import (
	"time"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/authentication"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/authorization"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/email"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/lockdistributor"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/metrics"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/tracing"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/version"
	"golang.org/x/sync/errgroup"
)

func main() {
	// Create new logger
	logger := log.NewLogger()

	// Create configuration manager
	cfgManager := config.NewManager(logger)

	// Load configuration
	err := cfgManager.Load()
	// Check error
	if err != nil {
		logger.Fatal(err)
	}

	// Get configuration
	cfg := cfgManager.GetConfig()
	// Configure logger
	err = logger.Configure(cfg.Log.Level, cfg.Log.Format, cfg.Log.FilePath)
	// Check error
	if err != nil {
		logger.Fatal(err)
	}

	// Watch change for logger (special case)
	cfgManager.AddOnChangeHook(func() {
		// Get configuration
		cfg := cfgManager.GetConfig()
		// Configure logger
		err = logger.Configure(cfg.Log.Level, cfg.Log.Format, cfg.Log.FilePath)
		// Check error
		if err != nil {
			logger.Fatal(err)
		}
	})

	// Getting version
	v := version.GetVersion()

	logger.Infof("Starting version: %s (git commit: %s) built on %s", v.Version, v.GitCommit, v.BuildDate)

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
		err = tracingSvc.Reload()
		// Check error
		if err != nil {
			logger.Fatal(err)
		}
	})

	// Create database service
	db := database.NewDatabase("main", cfgManager, logger, metricsCl)
	// Connect to engine
	err = db.Connect()
	// Check error
	if err != nil {
		logger.Fatal(err)
	}
	// Add configuration reload hook
	cfgManager.AddOnChangeHook(func() {
		err = db.Reconnect()
		// Check error
		if err != nil {
			logger.Fatal(err)
		}
	})

	// Create new mail service
	mailSvc := email.NewService(cfgManager, logger)
	// Try to connect
	err = mailSvc.Initialize()
	// Check error
	if err != nil {
		logger.Fatal(err)
	}
	// Add configuration reload hook
	cfgManager.AddOnChangeHook(func() {
		err = mailSvc.Initialize()
		// Check error
		if err != nil {
			logger.Fatal(err)
		}
	})

	// Create lock distributor service
	ld := lockdistributor.NewService(cfgManager, db)
	// Initialize lock distributor
	err = ld.Initialize(logger)
	// Check error
	if err != nil {
		logger.Fatal(err)
	}
	// Add configuration reload hook
	cfgManager.AddOnChangeHook(func() {
		err = ld.Initialize(logger)
		// Check error
		if err != nil {
			logger.Fatal(err)
		}
	})

	// Create authentication service
	authoSvc := authorization.NewService(cfgManager)

	// Create business services
	busServices := business.NewServices(logger, db, authoSvc, ld)

	// Migrate database
	err = busServices.MigrateDB()
	if err != nil {
		logger.Fatal(err)
	}

	// Create authentication service
	authenticationSvc := authentication.NewService(cfgManager)

	// Create servers
	svr := server.NewServer(logger, cfgManager, metricsCl, tracingSvc, busServices, authenticationSvc, authoSvc)
	intSvr := server.NewInternalServer(logger, cfgManager, metricsCl)

	// Add checker for database
	intSvr.AddChecker(&server.CheckerInput{
		Name:     "database",
		CheckFn:  db.Ping,
		Interval: 2 * time.Second, //nolint:gomnd // Won't do a const for that
		Timeout:  time.Second,
	})
	// Add checker for email service
	intSvr.AddChecker(&server.CheckerInput{
		Name:    "email",
		CheckFn: mailSvc.Check,
		// Interval is long because it takes a lot of time to connect SMTP server (can be 1 second).
		// Moreover, connect 6 time per minute should be ok.
		Interval: 10 * time.Second, //nolint:gomnd // Won't do a const for that
		Timeout:  3 * time.Second,  //nolint:gomnd // Won't do a const for that
	})

	// Generate server
	err = svr.GenerateServer()
	if err != nil {
		logger.Fatal(err)
	}
	// Generate internal server
	err = intSvr.GenerateServer()
	if err != nil {
		logger.Fatal(err)
	}

	var g errgroup.Group

	g.Go(svr.Listen)
	g.Go(intSvr.Listen)

	if err := g.Wait(); err != nil {
		logger.Fatal(err)
	}
}
