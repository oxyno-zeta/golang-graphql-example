package main

import (
	"time"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server"
)

func GenerateInternalServer(sv *services) (*server.InternalServer, error) {
	intSvr := server.NewInternalServer(sv.logger, sv.cfgManager, sv.metricsCl, sv.signalHandlerSvc)

	// Add checker for database
	intSvr.AddChecker(&server.CheckerInput{
		Name:     "database",
		CheckFn:  sv.db.Ping,
		Interval: 2 * time.Second, //nolint:gomnd // Won't do a const for that
		Timeout:  time.Second,
	})
	// Add checker for email service
	intSvr.AddChecker(&server.CheckerInput{
		Name:    "email",
		CheckFn: sv.mailSvc.Check,
		// Interval is long because it takes a lot of time to connect SMTP server (can be 1 second).
		// Moreover, connect 6 time per minute should be ok.
		Interval: 10 * time.Second, //nolint:gomnd // Won't do a const for that
		Timeout:  3 * time.Second,  //nolint:gomnd // Won't do a const for that
	})

	// Check if amqp service exists
	if sv.amqpSvc != nil {
		// Add checker for amqp service
		intSvr.AddChecker(&server.CheckerInput{
			Name:     "amqp",
			CheckFn:  sv.amqpSvc.Ping,
			Interval: 2 * time.Second, //nolint:gomnd // Won't do a const for that
			Timeout:  time.Second,
		})
	}

	// Generate internal server
	err := intSvr.GenerateServer()
	if err != nil {
		return nil, err
	}

	return intSvr, nil
}
