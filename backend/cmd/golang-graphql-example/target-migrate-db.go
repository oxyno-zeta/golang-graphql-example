package main

var migrateDBTarget = &targetDefinition{
	Run:         migrateDBTargetRun,
	Primary:     true,
	InAllTarget: true,
}

func migrateDBTargetRun(sv *services) {
	// Migrate database
	err := sv.busServices.MigrateDB()
	if err != nil {
		sv.logger.Fatal(err)
	}
}
