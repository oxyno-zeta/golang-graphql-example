//go:build integration

package database_test

import (
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/metrics"
)

// Generate metrics instance
var metricsCtx = metrics.NewService()

var integrationTestsCfg *config.Config = &config.Config{
	Server:  &config.ServerConfig{},
	Log:     &config.LogConfig{Level: "debug", Format: "human"},
	Tracing: &config.TracingConfig{Enabled: false},
	Database: &config.DatabaseConfig{
		Driver: config.DefaultDatabaseDriver,
		ConnectionURL: &config.CredentialConfig{
			Value: "host=localhost port=5432 user=postgres dbname=postgres-integration password=postgres sslmode=disable",
		},
	},
}
