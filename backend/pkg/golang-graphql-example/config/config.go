package config

// DefaultLogLevel Default log level.
const DefaultLogLevel = "info"

// DefaultLogFormat Default Log format.
const DefaultLogFormat = "json"

// DefaultPort Default port.
const DefaultPort = 8080

// DefaultInternalPort Default internal port.
const DefaultInternalPort = 9090

// Default lock distributor table name.
const DefaultLockDistributorTableName = "locks"

// Default lock distribution lease duration.
const DefaultLockDistributorLeaseDuration = "3s"

// Default lock distributor heartbeat frequency.
const DefaultLockDistributionHeartbeatFrequency = "1s"

// DefaultOIDCScopes Default OIDC scopes.
var DefaultOIDCScopes = []string{"openid", "email", "profile"}

// Default cookie name.
const DefaultCookieName = "oidc"

// Config Configuration object.
type Config struct {
	Log                    *LogConfig              `mapstructure:"log"`
	Tracing                *TracingConfig          `mapstructure:"tracing"`
	Server                 *ServerConfig           `mapstructure:"server"`
	InternalServer         *ServerConfig           `mapstructure:"internalServer"`
	Database               *DatabaseConfig         `mapstructure:"database" validate:"required"`
	LockDistributor        *LockDistributorConfig  `mapstructure:"lockDistributor" validate:"required"`
	OIDCAuthentication     *OIDCAuthConfig         `mapstructure:"oidcAuthentication"`
	OPAServerAuthorization *OPAServerAuthorization `mapstructure:"opaServerAuthorization"`
	SMTP                   *SMTPConfig             `mapstructure:"smtp" validate:"omitempty"`
}

// LockDistributorConfig Lock distributor configuration.
type LockDistributorConfig struct {
	TableName          string `mapstructure:"tableName" validate:"required"`
	LeaseDuration      string `mapstructure:"leaseDuration" validate:"required"`
	HeartbeatFrequency string `mapstructure:"heartbeatFrequency" validate:"required"`
}

// OIDCAuthConfig OpenID Connect authentication configurations.
type OIDCAuthConfig struct {
	ClientID          string            `mapstructure:"clientId" validate:"required"`
	ClientSecret      *CredentialConfig `mapstructure:"clientSecret" validate:"omitempty,dive"`
	IssuerURL         string            `mapstructure:"issuerUrl" validate:"required,url"`
	RedirectURL       string            `mapstructure:"redirectUrl" validate:"required,url"`
	LogoutRedirectURL string            `mapstructure:"logoutRedirectUrl" validate:"omitempty,url"`
	Scopes            []string          `mapstructure:"scopes"`
	State             string            `mapstructure:"state" validate:"required"`
	CookieName        string            `mapstructure:"cookieName"`
	EmailVerified     bool              `mapstructure:"emailVerified"`
	CookieSecure      bool              `mapstructure:"cookieSecure"`
}

// OPAServerAuthorization OPA Server authorization.
type OPAServerAuthorization struct {
	URL  string            `mapstructure:"url" validate:"required,url"`
	Tags map[string]string `mapstructure:"tags"`
}

// TracingConfig represents the Tracing configuration structure.
type TracingConfig struct {
	Enabled       bool                   `mapstructure:"enabled"`
	LogSpan       bool                   `mapstructure:"logSpan"`
	FlushInterval string                 `mapstructure:"flushInterval"`
	UDPHost       string                 `mapstructure:"udpHost"`
	QueueSize     int                    `mapstructure:"queueSize"`
	FixedTags     map[string]interface{} `mapstructure:"fixedTags"`
}

// LogConfig Log configuration.
type LogConfig struct {
	Level    string `mapstructure:"level" validate:"required"`
	Format   string `mapstructure:"format" validate:"required"`
	FilePath string `mapstructure:"filePath"`
}

// ServerConfig Server configuration.
type ServerConfig struct {
	ListenAddr string            `mapstructure:"listenAddr"`
	Port       int               `mapstructure:"port" validate:"required"`
	CORS       *ServerCorsConfig `mapstructure:"cors" validate:"omitempty"`
}

// ServerCorsConfig Server CORS configuration.
type ServerCorsConfig struct {
	AllowOrigins            []string `mapstructure:"allowOrigins"`
	AllowMethods            []string `mapstructure:"allowMethods"`
	AllowHeaders            []string `mapstructure:"allowHeaders"`
	ExposeHeaders           []string `mapstructure:"exposeHeaders"`
	MaxAgeDuration          string   `mapstructure:"maxAgeDuration"`
	AllowCredentials        *bool    `mapstructure:"allowCredentials"`
	AllowWildcard           *bool    `mapstructure:"allowWildcard"`
	AllowBrowserExtensions  *bool    `mapstructure:"allowBrowserExtensions"`
	AllowWebSockets         *bool    `mapstructure:"allowWebSockets"`
	AllowFiles              *bool    `mapstructure:"allowFiles"`
	AllowAllOrigins         *bool    `mapstructure:"allowAllOrigins"`
	UseDefaultConfiguration bool     `mapstructure:"useDefaultConfiguration"`
}

// DatabaseConfig Database configuration.
type DatabaseConfig struct {
	ConnectionURL                    *CredentialConfig `mapstructure:"connectionUrl" validate:"required"`
	DisableForeignKeyWhenMigrating   bool              `mapstructure:"disableForeignKeyWhenMigrating"`
	AllowGlobalUpdate                bool              `mapstructure:"allowGlobalUpdate"`
	PrepareStatement                 bool              `mapstructure:"prepareStatement"`
	SQLMaxIdleConnections            int               `mapstructure:"sqlMaxIdleConnections"`
	SQLMaxOpenConnections            int               `mapstructure:"sqlMaxOpenConnections"`
	SQLConnectionMaxLifetimeDuration string            `mapstructure:"sqlConnectionMaxLifetimeDuration"`
}

// SMTPConfig SMTP Configuration.
type SMTPConfig struct {
	Host               string            `mapstructure:"host" validation:"fqdn,required"`
	Port               int               `mapstructure:"port" validation:"gt=0,required"`
	Username           *CredentialConfig `mapstructure:"username"`
	Password           *CredentialConfig `mapstructure:"password"`
	Encryption         string            `mapstructure:"encryption" validation:"omitempty,oneof=NONE TLS SSL"`
	AuthenticationType string            `mapstructure:"authenticationType" validation:"omitempty,oneof=PLAIN LOGIN CRAM-MD5"`
	ConnectTimeout     string            `mapstructure:"connectTimeout"`
	SendTimeout        string            `mapstructure:"sendTimeout"`
	KeepAlive          bool              `mapstructure:"keepAlive"`
	TLSSkipVerify      bool              `mapstructure:"tlsSkipVerify"`
}

// CredentialConfig Credential Configurations.
type CredentialConfig struct {
	Path  string `mapstructure:"path" validate:"required_without_all=Env Value"`
	Env   string `mapstructure:"env" validate:"required_without_all=Path Value"`
	Value string `mapstructure:"value" validate:"required_without_all=Path Env"`
}
