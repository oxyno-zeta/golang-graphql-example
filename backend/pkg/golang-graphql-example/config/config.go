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

// Default Database driver.
const DefaultDatabaseDriver = "POSTGRES"

// Default tracing type.
const DefaultTracingType = TracingOtelHTTPType
const TracingJaegerHTTPType = "JAEGER_HTTP"
const TracingOtelHTTPType = "OTEL_HTTP"

// Config Configuration object.
type Config struct {
	Log                    *LogConfig              `mapstructure:"log"`
	Tracing                *TracingConfig          `mapstructure:"tracing"`
	Server                 *ServerConfig           `mapstructure:"server"`
	InternalServer         *ServerConfig           `mapstructure:"internalServer"`
	Database               *DatabaseConfig         `mapstructure:"database"               validate:"required"`
	LockDistributor        *LockDistributorConfig  `mapstructure:"lockDistributor"        validate:"required"`
	OIDCAuthentication     *OIDCAuthConfig         `mapstructure:"oidcAuthentication"`
	OPAServerAuthorization *OPAServerAuthorization `mapstructure:"opaServerAuthorization"`
	SMTP                   *SMTPConfig             `mapstructure:"smtp"                   validate:"omitempty"`
	AMQP                   *AMQPConfig             `mapstructure:"amqp"                   validate:"omitempty,dive"`
}

// AMQPConfig AMQP Message Bus configuration.
type AMQPConfig struct {
	Connection *AMQPConnectionConfig  `mapstructure:"connection" validate:"required"`
	ChannelQos *AMQPChannelQosConfig  `mapstructure:"channelQos" validate:"omitempty,dive"`
	Exchanges  []*AMQPExchangeConfig  `mapstructure:"exchanges"  validate:"required,dive,required"`
	Queues     []*AMQPQueueConfig     `mapstructure:"queues"     validate:"omitempty,dive"`
	QueueBinds []*AMQPQueueBindConfig `mapstructure:"queueBinds" validate:"omitempty,dive"`
}

// AMQPChannelQosConfig AMQP Channel Qos Configuration.
type AMQPChannelQosConfig struct {
	PrefetchCount int  `mapstructure:"prefetchCount"`
	PrefetchSize  int  `mapstructure:"prefetchSize"`
	Global        bool `mapstructure:"global"`
}

// AMQPConnectionConfig AMQP Connection Configuration.
type AMQPConnectionConfig struct {
	URL               *CredentialConfig      `mapstructure:"url"               validate:"required"`
	ExtraArgs         map[string]interface{} `mapstructure:"extraArgs"`
	HeartbeatDuration string                 `mapstructure:"heartbeatDuration"`
	ChannelMax        int                    `mapstructure:"channelMax"`
	FrameSize         int                    `mapstructure:"frameSize"`
}

// AMQPQueueBindConfig AMQP Message Bus QueueBind Configuration.
type AMQPQueueBindConfig struct {
	ExtraArgs map[string]interface{} `mapstructure:"extraArgs"`
	Name      string                 `mapstructure:"name"      validate:"required"`
	Key       string                 `mapstructure:"key"       validate:"required"`
	Exchange  string                 `mapstructure:"exchange"  validate:"required"`
	NoWait    bool                   `mapstructure:"noWait"`
}

// AMQPQueueConfig AMQP Message Bus Queue configuration.
type AMQPQueueConfig struct {
	ExtraArgs  map[string]interface{} `mapstructure:"extraArgs"`
	Name       string                 `mapstructure:"name"       validate:"required"`
	Durable    bool                   `mapstructure:"durable"`
	AutoDelete bool                   `mapstructure:"autoDelete"`
	Exclusive  bool                   `mapstructure:"exclusive"`
	NoWait     bool                   `mapstructure:"noWait"`
}

// AMQPExchangeConfig AMQP Message Bus Exchange configuration.
type AMQPExchangeConfig struct {
	ExtraArgs  map[string]interface{} `mapstructure:"extraArgs"`
	Name       string                 `mapstructure:"name"       validate:"required"`
	Type       string                 `mapstructure:"type"       validate:"required"`
	Durable    bool                   `mapstructure:"durable"`
	AutoDelete bool                   `mapstructure:"autoDelete"`
	Internal   bool                   `mapstructure:"internal"`
	NoWait     bool                   `mapstructure:"noWait"`
}

// LockDistributorConfig Lock distributor configuration.
type LockDistributorConfig struct {
	TableName          string `mapstructure:"tableName"          validate:"required"`
	LeaseDuration      string `mapstructure:"leaseDuration"      validate:"required"`
	HeartbeatFrequency string `mapstructure:"heartbeatFrequency" validate:"required"`
}

// OIDCAuthConfig OpenID Connect authentication configurations.
type OIDCAuthConfig struct {
	ClientSecret      *CredentialConfig `mapstructure:"clientSecret"      validate:"omitempty,dive"`
	ClientID          string            `mapstructure:"clientId"          validate:"required"`
	IssuerURL         string            `mapstructure:"issuerUrl"         validate:"required,url"`
	RedirectURL       string            `mapstructure:"redirectUrl"       validate:"required,url"`
	LogoutRedirectURL string            `mapstructure:"logoutRedirectUrl" validate:"omitempty,url"`
	State             string            `mapstructure:"state"             validate:"required"`
	CookieName        string            `mapstructure:"cookieName"`
	Scopes            []string          `mapstructure:"scopes"`
	EmailVerified     bool              `mapstructure:"emailVerified"`
	CookieSecure      bool              `mapstructure:"cookieSecure"`
}

// OPAServerAuthorization OPA Server authorization.
type OPAServerAuthorization struct {
	Tags map[string]string `mapstructure:"tags"`
	URL  string            `mapstructure:"url"  validate:"required,url"`
}

// TracingConfig represents the Tracing configuration structure.
type TracingConfig struct {
	FixedTags    map[string]string        `mapstructure:"fixedTags"`
	JaegerHTTP   *TracingJaegerHTTPConfig `mapstructure:"jaegerHttp"   validate:"required_if=Type JAEGER_HTTP"`
	OtelHTTP     *TracingOtelHTTPConfig   `mapstructure:"otelHttp"     validate:"required_if=Type OTEL_HTTP"`
	Type         string                   `mapstructure:"type"         validate:"oneof=JAEGER_HTTP OTEL_HTTP"`
	MaxQueueSize int                      `mapstructure:"maxQueueSize" validate:"omitempty,gte=0"`
	MaxBatchSize int                      `mapstructure:"maxBatchSize" validate:"omitempty,gte=0"`
	Enabled      bool                     `mapstructure:"enabled"`
}

// TracingJaegerHTTPConfig represents the Jaeger HTTP configuration structure.
type TracingJaegerHTTPConfig struct {
	ServerURL     string `mapstructure:"serverUrl" validate:"required,http_url"`
	TimeoutString string `mapstructure:"timeout"`
}

// TracingOtelHTTPConfig represents the OTEL HTTP configuration structure.
type TracingOtelHTTPConfig struct {
	ServerURL     string            `mapstructure:"serverUrl" validate:"required,http_url"`
	Headers       map[string]string `mapstructure:"headers"`
	TimeoutString string            `mapstructure:"timeout"`
}

// LogConfig Log configuration.
type LogConfig struct {
	Level    string `mapstructure:"level"    validate:"required"`
	Format   string `mapstructure:"format"   validate:"required"`
	FilePath string `mapstructure:"filePath"`
}

// ServerConfig Server configuration.
type ServerConfig struct {
	CORS       *ServerCorsConfig     `mapstructure:"cors"       validate:"omitempty"`
	Compress   *ServerCompressConfig `mapstructure:"compress"`
	ListenAddr string                `mapstructure:"listenAddr"`
	Port       int                   `mapstructure:"port"       validate:"required"`
}

// ServerCompressConfig Server compress configuration.
type ServerCompressConfig struct {
	Enabled bool `mapstructure:"enabled"`
}

// ServerCorsConfig Server CORS configuration.
type ServerCorsConfig struct {
	AllowCredentials        *bool    `mapstructure:"allowCredentials"`
	AllowWildcard           *bool    `mapstructure:"allowWildcard"`
	AllowBrowserExtensions  *bool    `mapstructure:"allowBrowserExtensions"`
	AllowWebSockets         *bool    `mapstructure:"allowWebSockets"`
	AllowFiles              *bool    `mapstructure:"allowFiles"`
	AllowAllOrigins         *bool    `mapstructure:"allowAllOrigins"`
	MaxAgeDuration          string   `mapstructure:"maxAgeDuration"`
	AllowOrigins            []string `mapstructure:"allowOrigins"`
	AllowMethods            []string `mapstructure:"allowMethods"`
	AllowHeaders            []string `mapstructure:"allowHeaders"`
	ExposeHeaders           []string `mapstructure:"exposeHeaders"`
	UseDefaultConfiguration bool     `mapstructure:"useDefaultConfiguration"`
}

// DatabaseConfig Database configuration.
type DatabaseConfig struct {
	ConnectionURL                    *CredentialConfig `mapstructure:"connectionUrl"                    validate:"required"`
	Driver                           string            `mapstructure:"driver"                           validate:"required,oneof=POSTGRES SQLITE"`
	SQLConnectionMaxLifetimeDuration string            `mapstructure:"sqlConnectionMaxLifetimeDuration"`
	SQLMaxIdleConnections            int               `mapstructure:"sqlMaxIdleConnections"`
	SQLMaxOpenConnections            int               `mapstructure:"sqlMaxOpenConnections"`
	DisableForeignKeyWhenMigrating   bool              `mapstructure:"disableForeignKeyWhenMigrating"`
	AllowGlobalUpdate                bool              `mapstructure:"allowGlobalUpdate"`
	PrepareStatement                 bool              `mapstructure:"prepareStatement"`
}

// SMTPConfig SMTP Configuration.
type SMTPConfig struct {
	Username           *CredentialConfig `mapstructure:"username"`
	Password           *CredentialConfig `mapstructure:"password"`
	Host               string            `mapstructure:"host"               validation:"fqdn,required"`
	Encryption         string            `mapstructure:"encryption"         validation:"omitempty,oneof=NONE TLS SSL"`
	AuthenticationType string            `mapstructure:"authenticationType" validation:"omitempty,oneof=PLAIN LOGIN CRAM-MD5"`
	ConnectTimeout     string            `mapstructure:"connectTimeout"`
	SendTimeout        string            `mapstructure:"sendTimeout"`
	Port               int               `mapstructure:"port"               validation:"gt=0,required"`
	KeepAlive          bool              `mapstructure:"keepAlive"`
	TLSSkipVerify      bool              `mapstructure:"tlsSkipVerify"`
}

// CredentialConfig Credential Configurations.
type CredentialConfig struct {
	Path  string `mapstructure:"path"  validate:"required_without_all=Env Value"`
	Env   string `mapstructure:"env"   validate:"required_without_all=Path Value"`
	Value string `mapstructure:"value" validate:"required"`
}
