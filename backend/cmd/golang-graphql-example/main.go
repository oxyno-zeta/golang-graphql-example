package main

import (
	"flag"
	"strings"
	"sync"

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
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/signalhandler"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/tracing"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/version"
	"github.com/thoas/go-funk"
)

type services struct {
	// Mandatory
	logger     log.Logger
	cfgManager config.Manager
	version    *version.AppVersion
	// Basics
	metricsCl         metrics.Client
	tracingSvc        tracing.Service
	db                database.DB
	mailSvc           email.Service
	ldSvc             lockdistributor.Service
	signalHandlerSvc  signalhandler.Client
	amqpSvc           amqpbusmessage.Client
	authorizationSvc  authorization.Service
	authenticationSvc authentication.Client
	// Extra
	// Business
	busServices *business.Services
}

var targetDefinitionsMap = map[string]*targetDefinition{
	// Basics
	"migrate-db": migrateDBTarget,
	"server":     serverTarget,
	// Extra
}

// WaitGroup is used to wait for the program to finish goroutines.
var wg sync.WaitGroup

func main() {
	// Compute possible targets
	possibleTargetValues, _ := funk.Keys(targetDefinitionsMap).([]string)
	// Add "all" in those cases
	possibleTargetValues = append(possibleTargetValues, "all")

	// Init flags
	var targets arrayFlags

	// Create target flag
	flag.Var(
		&targets,
		"target",
		"Represents the application target to be launched (possible values:"+strings.Join(possibleTargetValues, ",")+")",
	)
	// Parse flags
	flag.Parse()

	// Init services
	sv := &services{}

	// Setup mandatory services
	setupMandatoryServices(sv)

	// Catch any panic
	defer func() {
		// Catch panic
		if errI := recover(); errI != nil {
			// Panic caught => Log and exit
			// Try to cast error
			err, ok := errI.(error)
			// Check if cast wasn't ok
			if !ok {
				// Transform it
				err = errors.New(fmt.Sprintf("%+v", errI))
			} else {
				// Map introduce stack trace
				err = errors.WithStack(err)
			}

			// Log
			sv.logger.Fatal(err)
		}
	}()

	sv.logger.Infof("Application version: %s (git commit: %s) built on %s", sv.version.Version, sv.version.GitCommit, sv.version.BuildDate)

	// Check if list is empty
	if len(targets) == 0 {
		// Add "all" for default values
		targets = append(targets, "all")
	}

	// Check if "all" is present with other things
	if funk.ContainsString(targets, "all") && len(targets) != 1 {
		// Reset to "all"
		targets = []string{"all"}
	}
	// Uniq targets
	targets = funk.UniqString(targets)
	// Check if target list have only accepted values
	for _, targetFlag := range targets {
		if !funk.ContainsString(possibleTargetValues, targetFlag) {
			sv.logger.Fatalf("target %s not supported", targetFlag)
		}
	}

	sv.logger.Infof("Starting application with targets: %s", targets)

	// Setup services
	setupBasicsServices(targets, sv)

	// Setup extra services
	setupExtraServices(targets, sv)

	// Setup business services
	setupBusinessServices(targets, sv)

	// Select targets and filter them by primary or not
	// Initialize target definitions lists
	primaryList := []*targetDefinition{}
	otherList := []*targetDefinition{}

	// Check if this is a "all" target
	if len(targets) == 1 && targets[0] == "all" {
		// Loop over all possible targets
		for _, tDef := range targetDefinitionsMap {
			// Check if acceptable in "all"
			if tDef.InAllTarget {
				// Check if primary
				if tDef.Primary {
					primaryList = append(primaryList, tDef)
				} else {
					otherList = append(otherList, tDef)
				}
			}
		}
	} else {
		// Loop over targets
		for _, target := range targets {
			// Get target definition
			tDef := targetDefinitionsMap[target]
			// Check if primary
			if tDef.Primary {
				primaryList = append(primaryList, tDef)
			} else {
				otherList = append(otherList, tDef)
			}
		}
	}

	// Start all primary targets
	for _, tDef := range primaryList {
		// Run
		tDef.Run(sv)
	}

	// Add count of other targets for waiting group
	wg.Add(len(otherList))

	// Start all other targets
	for _, tDef := range otherList {
		// Start routine
		go func(tDef *targetDefinition) {
			// Inform routine is completed
			defer wg.Done()

			// Run target
			tDef.Run(sv)
		}(tDef)
	}

	// Wait
	wg.Wait()
}
