package lockdistributor

import (
	"strings"
	"time"

	"cirello.io/pglock"
	"emperror.dev/errors"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
)

type service struct {
	cfgManager config.Manager
	db         database.DB
	cl         *pglock.Client
}

func (s *service) Initialize(logger log.Logger) error {
	// Get configuration
	cfg := s.cfgManager.GetConfig()

	// Parse durations
	ld, err := time.ParseDuration(cfg.LockDistributor.LeaseDuration)
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	hf, err := time.ParseDuration(cfg.LockDistributor.HeartbeatFrequency)
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	// Get sql database
	sqlDB, err := s.db.GetSQLDB()
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	// Log
	logger.Debug("Trying to create lock distributor client")

	// Create pglock client
	c, err := pglock.UnsafeNew(
		sqlDB,
		pglock.WithLeaseDuration(ld),
		pglock.WithHeartbeatFrequency(hf),
		pglock.WithCustomTable(cfg.LockDistributor.TableName),
		pglock.WithLogger(logger.GetLockDistributorLogger()),
	)
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	// Create lock table
	err = c.CreateTable()
	// Check error
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return errors.WithStack(err)
	}

	// Save client
	s.cl = c

	// Log
	logger.Info("Successfully created lock distributor client")

	return nil
}

func (s *service) GetLock(name string) Lock {
	return &lock{
		name: name,
		s:    s,
	}
}
