package migration

import (
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

	// Create sequences
	sequences := []*gormigrate.Migration{}
	// Loot of list
	for _, k := range sequencesList {
		sequences = append(sequences, k...)
	}

	// Create migration sequence
	m := gormigrate.New(
		db,
		gormigrate.DefaultOptions,
		sequences,
	)

	// Start migration
	return m.Migrate()
}
