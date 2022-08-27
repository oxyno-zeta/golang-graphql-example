package main

import (
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server"
)

var serverTarget = &targetDefinition{
	Run:         serverTargetRun,
	Primary:     false,
	InAllTarget: true,
}

func serverTargetRun(targets []string, sv *services) {
	// Create servers
	svr := server.NewServer(
		sv.logger,
		sv.cfgManager,
		sv.metricsCl,
		sv.tracingSvc,
		sv.busServices,
		sv.authenticationSvc,
		sv.authorizationSvc,
		sv.signalHandlerSvc,
	)

	// Generate server
	err := svr.GenerateServer()
	if err != nil {
		sv.logger.Fatal(err)
	}
	// Generate internal server
	intSvr, err := GenerateInternalServer(sv)
	if err != nil {
		sv.logger.Fatal(err)
	}

	// Start server in routine
	go func() {
		err2 := svr.Listen()
		// Check error
		if err2 != nil {
			sv.logger.Fatal(err2)
		}
	}()

	// Start internal server
	err = intSvr.Listen()
	// Check error
	if err != nil {
		sv.logger.Fatal(err)
	}
}
