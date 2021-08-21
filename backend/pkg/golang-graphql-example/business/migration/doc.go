package migration

// This package is here to manage migration of database objects with
// a migration tool and a dedicated database to store migration status.
// This service should be the only service knowing that database exists
// and how to request it. Otherwise, this should be done in a dao.
