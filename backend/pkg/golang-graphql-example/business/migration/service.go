package migration

import (
	"emperror.dev/errors"
	"github.com/go-gormigrate/gormigrate/v2"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/migration/sequences"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
)

type service struct {
	dbSvc database.DB
}

func (s *service) Migrate() error {
	// Get gorm database
	db := s.dbSvc.GetGormDB()

	// Create array of sequences
	sequencesList := [][]*gormigrate.Migration{
		sequences.Seq201608List,
		sequences.Seq202108List,
	}

	// Create migrationSequences
	migrationSequences := []*gormigrate.Migration{}
	// Loot of list
	for _, k := range sequencesList {
		migrationSequences = append(migrationSequences, k...)
	}

	// Create migration sequence
	m := gormigrate.New(
		db,
		// Due to #76, force transaction to have a rollback
		// https://github.com/go-gormigrate/gormigrate/issues/76
		&gormigrate.Options{UseTransaction: true},
		migrationSequences,
	)

	// Start migration
	return errors.WithStack(m.Migrate())
}
