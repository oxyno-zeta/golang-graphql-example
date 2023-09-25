package config

import (
	"github.com/spf13/viper"
)

func (*managerimpl) loadDefaultConfigurationValues(vip *viper.Viper) {
	// Load default configuration
	vip.SetDefault("log.level", DefaultLogLevel)
	vip.SetDefault("log.format", DefaultLogFormat)
	vip.SetDefault("server.port", DefaultPort)
	vip.SetDefault("internalServer.port", DefaultInternalPort)
	vip.SetDefault("database.driver", DefaultDatabaseDriver)
	vip.SetDefault("lockDistributor.tableName", DefaultLockDistributorTableName)
	vip.SetDefault("lockDistributor.leaseDuration", DefaultLockDistributorLeaseDuration)
	vip.SetDefault("lockDistributor.heartbeatFrequency", DefaultLockDistributionHeartbeatFrequency)
	vip.SetDefault("tracing.type", DefaultTracingType)
}

// Load default values based on business rules.
func loadBusinessDefaultValues(out *Config) error {
	// Load default oidc configurations
	if out.OIDCAuthentication != nil {
		// Add default scopes
		if out.OIDCAuthentication.Scopes == nil {
			out.OIDCAuthentication.Scopes = DefaultOIDCScopes
		}
		// Add default cookie name
		if out.OIDCAuthentication.CookieName == "" {
			out.OIDCAuthentication.CookieName = DefaultCookieName
		}
	}

	// Load default tags for opa authorization
	if out.OPAServerAuthorization != nil && out.OPAServerAuthorization.Tags == nil {
		out.OPAServerAuthorization.Tags = map[string]string{}
	}

	// Load default tracing configuration
	if out.Tracing == nil {
		out.Tracing = &TracingConfig{Enabled: false}
	}

	// TODO Load default values based on business rules
	return nil
}

func parseValues(_ *Config) error {
	// TODO make any parsing here
	// Default
	return nil
}
