package main

var migrateDBTarget = &targetDefinition{
	Run:         migrateDBTargetRun,
	Primary:     true,
	InAllTarget: true,
}

func migrateDBTargetRun(targets []string, sv *services) {
	sv.logger.Info("Starting database migration")
	// Migrate database
	err := sv.busServices.MigrateDB()
	if err != nil {
		sv.logger.Fatal(err)
	}
}
