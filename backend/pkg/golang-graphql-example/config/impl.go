package config

import (
	"github.com/spf13/viper"
)

func (ctx *managerimpl) loadDefaultConfigurationValues(vip *viper.Viper) {
	// Load default configuration
	vip.SetDefault("log.level", DefaultLogLevel)
	vip.SetDefault("log.format", DefaultLogFormat)
	vip.SetDefault("server.port", DefaultPort)
	vip.SetDefault("internalServer.port", DefaultInternalPort)
	vip.SetDefault("database.driver", DefaultDatabaseDriver)
	vip.SetDefault("lockDistributor.tableName", DefaultLockDistributorTableName)
	vip.SetDefault("lockDistributor.leaseDuration", DefaultLockDistributorLeaseDuration)
	vip.SetDefault("lockDistributor.heartbeatFrequency", DefaultLockDistributionHeartbeatFrequency)
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

// Load credential configs here.
func loadAllCredentials(out *Config) ([]*CredentialConfig, error) {
	// Initialize answer
	result := make([]*CredentialConfig, 0)

	// Load database credential
	err := loadCredential(out.Database.ConnectionURL)
	if err != nil {
		return nil, err
	}
	// Append result
	result = append(result, out.Database.ConnectionURL)

	// Load credential for OIDC configuration
	if out.OIDCAuthentication != nil && out.OIDCAuthentication.ClientSecret != nil {
		err := loadCredential(out.OIDCAuthentication.ClientSecret)
		if err != nil {
			return nil, err
		}
		// Append result
		result = append(result, out.OIDCAuthentication.ClientSecret)
	}

	// SMTP configuration
	if out.SMTP != nil {
		// Load credential for SMTP username
		if out.SMTP.Username != nil {
			err := loadCredential(out.SMTP.Username)
			if err != nil {
				return nil, err
			}
			// Append result
			result = append(result, out.SMTP.Username)
		}

		// Load credential for SMTP password
		if out.SMTP.Password != nil {
			err := loadCredential(out.SMTP.Password)
			if err != nil {
				return nil, err
			}
			// Append result
			result = append(result, out.SMTP.Password)
		}
	}

	// Load credential for AMQP configuration
	if out.AMQP != nil && out.AMQP.Connection != nil && out.AMQP.Connection.URL != nil {
		err := loadCredential(out.AMQP.Connection.URL)
		if err != nil {
			return nil, err
		}
		// Append result
		result = append(result, out.AMQP.Connection.URL)
	}

	// TODO Load credential configs here

	return result, nil
}

func parseValues(out *Config) error {
	// TODO make any parsing here
	// Default
	return nil
}
