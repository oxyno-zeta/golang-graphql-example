package lockdistributor

import (
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
)

type Service interface {
	GetLock(name string) Lock
	Initialize(logger log.Logger) error
}

type Lock interface {
	Acquire() error
	Release() error
}

func NewService(cfgManager config.Manager, db database.DB) Service {
	return &service{
		cfgManager: cfgManager,
		db:         db,
	}
}
