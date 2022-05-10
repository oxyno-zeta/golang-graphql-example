package main

import (
	"os"
	"syscall"
	"time"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/authentication"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/authorization"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/email"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/lockdistributor"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	amqpbusmessage "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/messagebus/amqp"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/metrics"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/signalhandler"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/tracing"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/version"
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
		newCfg := cfgManager.GetConfig()
		// Configure logger
		err = logger.Configure(newCfg.Log.Level, newCfg.Log.Format, newCfg.Log.FilePath)
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
	db := database.NewDatabase("main", cfgManager, logger, metricsCl, tracingSvc)
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

	// Create signal handler service
	signalHandlerSvc := signalhandler.NewClient(logger, true, []os.Signal{syscall.SIGTERM, syscall.SIGINT})
	// Initialize service
	err = signalHandlerSvc.Initialize()
	// Check error
	if err != nil {
		logger.Fatal(err)
	}
	// Register closing database connections on system stop
	signalHandlerSvc.OnExit(func() {
		err = db.Close()
		// Check error
		if err != nil {
			logger.Fatal(err)
		}
	})

	// Initialize
	var amqpSvc amqpbusmessage.Client
	// Check if amqp have configuration set
	if cfg.AMQP != nil {
		// Create amqp bus message service
		amqpSvc = amqpbusmessage.New(logger, cfgManager, tracingSvc, signalHandlerSvc, metricsCl)
		// Connect
		err = amqpSvc.Connect()
		// Check error
		if err != nil {
			logger.Fatal(err)
		}
		// Add configuration reload hook
		cfgManager.AddOnChangeHook(func() {
			err = amqpSvc.Reconnect()
			// Check error
			if err != nil {
				logger.Fatal(err)
			}
		})
		// Register closing connections on system stop
		signalHandlerSvc.OnExit(func() {
			err2 := amqpSvc.Close()
			// Check error
			if err2 != nil {
				logger.Fatal(err2)
			}
		})
		// Register canceling consumers on SIGTERM or SIGINT
		signalHandlerSvc.OnSignal(syscall.SIGTERM, func() {
			err2 := amqpSvc.CancelAllConsumers()
			// Check error
			if err2 != nil {
				logger.Fatal(err2)
			}
		})
		signalHandlerSvc.OnSignal(syscall.SIGINT, func() {
			err2 := amqpSvc.CancelAllConsumers()
			// Check error
			if err2 != nil {
				logger.Fatal(err2)
			}
		})
	}

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
	svr := server.NewServer(logger, cfgManager, metricsCl, tracingSvc, busServices, authenticationSvc, authoSvc, signalHandlerSvc)
	intSvr := server.NewInternalServer(logger, cfgManager, metricsCl, signalHandlerSvc)

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
	// Check if amqp service exists
	if amqpSvc != nil {
		// Add checker for amqp service
		intSvr.AddChecker(&server.CheckerInput{
			Name:     "amqp",
			CheckFn:  amqpSvc.Ping,
			Interval: 2 * time.Second, //nolint:gomnd // Won't do a const for that
			Timeout:  time.Second,
		})
	}

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

	// Start server in routine
	go func() {
		err2 := svr.Listen()
		// Check error
		if err2 != nil {
			logger.Fatal(err2)
		}
	}()

	// Start internal server
	err = intSvr.Listen()
	// Check error
	if err != nil {
		logger.Fatal(err)
	}
}
