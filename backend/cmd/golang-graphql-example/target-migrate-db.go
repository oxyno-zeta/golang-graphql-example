package main

import (
	"context"
	"time"
)

var migrateDBTarget = &targetDefinition{
	Run:         migrateDBTargetRun,
	Primary:     true,
	InAllTarget: true,
}

func migrateDBTargetRun(targets []string, sv *services) {
	// Add trace
	ctx, trace := sv.tracingSvc.StartTrace(context.TODO(), "migrate-db")
	// Defer
	defer func() {
		trace.Finish()
		// Check targets
		if len(targets) == 1 && targets[0] != "all" {
			// Wait
			time.Sleep(5 * time.Second) //nolint:mnd
		}
	}()

	sv.logger.Info("Starting database migration")
	// Migrate database
	err := sv.busServices.MigrateDB(ctx)
	if err != nil {
		// trace
		trace.AddAndMarkError(err)

		sv.logger.Fatal(err)
	}
}
