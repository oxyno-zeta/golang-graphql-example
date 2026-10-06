package main

import (
	"context"

	"github.com/thoas/go-funk"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
)

var businessMetricsExporterTarget = &targetDefinition{
	Run:         businessMetricsExporterTargetRun,
	Primary:     false,
	InAllTarget: true,
}

func businessMetricsExporterTargetRun(targets []string, sv *services) {
	// Build context
	ctx := log.SetLoggerToContext(context.TODO(), sv.logger)

	// Loop to add business metric definitions
	for _, def := range sv.busServices.GetBusinessMetricDefinitions() {
		sv.metricsSvc.AddBusinessMetricDefinition(def)
	}

	// Register metrics
	err := sv.metricsSvc.RegisterBusinessMetricDefinitions()
	// Check error
	if err != nil {
		sv.logger.Fatal(err)

		return
	}

	// Initial fetch of metrics
	err = sv.metricsSvc.InitialFetchBusinessMetricDefinitions(ctx)
	// Check error
	if err != nil {
		sv.logger.Fatal(err)

		return
	}

	// Start updater
	sv.metricsSvc.StartUpdaterBusinessMetricDefinitions(sv.logger)

	// If worker target is launched without any other target including an internal server
	// Then we want to launch an internal server
	if !funk.ContainsString(targets, "all") && !funk.ContainsString(targets, "server") {
		// Generate internal server
		intSvr, err := GenerateInternalServer(sv)
		if err != nil {
			sv.logger.Fatal(err)
		}

		// Start internal server
		err = intSvr.Listen()
		// Check error
		if err != nil {
			sv.logger.Fatal(err)
		}
	}
}
