package lockdistributor

import (
	"context"

	"github.com/pkg/errors"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
)

// ErrAcquireTransactionSerialize will be thrown when a transaction serialize is caught
// and retry number have been reached.
var ErrAcquireTransactionSerialize = errors.New("transaction error during lock acquire")

//go:generate mockgen -destination=./mocks/mock_Service.go -package=mocks github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/lockdistributor Service
type Service interface {
	// Get a lock object (semaphore on string) that can be acquired and release
	GetLock(name string) Lock
	// Initialize service
	Initialize(logger log.Logger) error
}

//go:generate mockgen -destination=./mocks/mock_Lock.go -package=mocks github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/lockdistributor Lock
type Lock interface {
	// Acquire lock
	Acquire() error
	// Acquire lock with context
	AcquireWithContext(ctx context.Context) error
	// Release lock
	Release() error
	// Check if a lock with this name is already taken
	IsAlreadyTaken() (bool, error)
	// Check if the lock is released or lost because of missing heartbeat
	IsReleased() (bool, error)
}

func NewService(cfgManager config.Manager, db database.DB) Service {
	return &service{
		cfgManager: cfgManager,
		db:         db,
	}
}
