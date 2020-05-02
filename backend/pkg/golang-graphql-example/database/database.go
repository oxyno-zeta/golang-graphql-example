package database

import (
	"github.com/jinzhu/gorm"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
)

type DB interface {
	GetGormDB() *gorm.DB
	Connect() error
	Close() error
	Ping() error
	Reconnect() error
}

// NewDatabase will generate a new DB object
func NewDatabase(cfgManager config.Manager, logger log.Logger) DB {
	return &postresdb{logger: logger, cfgManager: cfgManager}
}
