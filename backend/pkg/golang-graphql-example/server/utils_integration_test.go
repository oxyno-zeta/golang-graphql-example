//go:build integration

package server

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
	OIDCAuthentication: &config.OIDCAuthConfig{
		ClientID:          "client-without-secret",
		State:             "my-secret-state-key",
		IssuerURL:         "http://localhost:8088/auth/realms/integration",
		RedirectURL:       "http://localhost:8080/",
		LogoutRedirectURL: "http://localhost:8080/",
		EmailVerified:     true,
		Scopes:            config.DefaultOIDCScopes,
		CookieName:        config.DefaultCookieName,
	},
	LockDistributor: &config.LockDistributorConfig{
		TableName:          config.DefaultLockDistributorTableName,
		LeaseDuration:      config.DefaultLockDistributorLeaseDuration,
		HeartbeatFrequency: config.DefaultLockDistributionHeartbeatFrequency,
	},
	Database: &config.DatabaseConfig{
		Driver: config.DefaultDatabaseDriver,
		ConnectionURL: &config.CredentialConfig{
			Value: "host=localhost port=5432 user=postgres dbname=postgres-integration password=postgres sslmode=disable",
		},
	},
}
