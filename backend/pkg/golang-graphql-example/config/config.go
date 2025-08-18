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
const (
	DefaultTracingType  = TracingOtelHTTPType
	TracingOtelHTTPType = "OTEL_HTTP"
)

// Default tracing sampler type.
const (
	TracingSamplerAlwaysOff = "ALWAYS_OFF"
	TracingSamplerAlwaysOn  = "ALWAYS_ON"
	TracingSamplerRatio     = "RATIO"
)

// Config Configuration object.
type Config struct {
	Log                    *LogConfig              `mapstructure:"log"                    json:"log,omitempty"`
	Tracing                *TracingConfig          `mapstructure:"tracing"                json:"tracing,omitempty"`
	Server                 *ServerConfig           `mapstructure:"server"                 json:"server,omitempty"`
	InternalServer         *ServerConfig           `mapstructure:"internalServer"         json:"internalServer,omitempty"`
	Database               *DatabaseConfig         `mapstructure:"database"               json:"database,omitempty"               validate:"required"`
	LockDistributor        *LockDistributorConfig  `mapstructure:"lockDistributor"        json:"lockDistributor,omitempty"        validate:"required"`
	OIDCAuthentication     *OIDCAuthConfig         `mapstructure:"oidcAuthentication"     json:"oidcAuthentication,omitempty"`
	OPAServerAuthorization *OPAServerAuthorization `mapstructure:"opaServerAuthorization" json:"opaServerAuthorization,omitempty"`
	SMTP                   *SMTPConfig             `mapstructure:"smtp"                   json:"smtp,omitempty"                   validate:"omitempty"`
	AMQP                   *AMQPConfig             `mapstructure:"amqp"                   json:"amqp,omitempty"                   validate:"omitempty"`
}

// AMQPConfig AMQP Message Bus configuration.
type AMQPConfig struct {
	Connection *AMQPConnectionConfig  `mapstructure:"connection" validate:"required"               json:"connection,omitempty"`
	ChannelQos *AMQPChannelQosConfig  `mapstructure:"channelQos" validate:"omitempty"              json:"channelQos,omitempty"`
	Exchanges  []*AMQPExchangeConfig  `mapstructure:"exchanges"  validate:"required,dive,required" json:"exchanges,omitempty"`
	Queues     []*AMQPQueueConfig     `mapstructure:"queues"     validate:"omitempty,dive"         json:"queues,omitempty"`
	QueueBinds []*AMQPQueueBindConfig `mapstructure:"queueBinds" validate:"omitempty,dive"         json:"queueBinds,omitempty"`
}

// AMQPChannelQosConfig AMQP Channel Qos Configuration.
type AMQPChannelQosConfig struct {
	PrefetchCount int  `mapstructure:"prefetchCount" json:"prefetchCount,omitempty"`
	PrefetchSize  int  `mapstructure:"prefetchSize"  json:"prefetchSize,omitempty"`
	Global        bool `mapstructure:"global"        json:"global,omitempty"`
}

// AMQPConnectionConfig AMQP Connection Configuration.
type AMQPConnectionConfig struct {
	URL               *CredentialConfig `mapstructure:"url"               validate:"required"               json:"url,omitempty"`
	Username          *CredentialConfig `mapstructure:"username"          validate:"required_with=Password" json:"username,omitempty"`
	Password          *CredentialConfig `mapstructure:"password"          validate:"required_with=Username" json:"password,omitempty"`
	ExtraArgs         map[string]any    `mapstructure:"extraArgs"                                           json:"extraArgs,omitempty"`
	HeartbeatDuration string            `mapstructure:"heartbeatDuration"                                   json:"heartbeatDuration,omitempty"`
	ChannelMax        uint16            `mapstructure:"channelMax"                                          json:"channelMax,omitempty"`
	FrameSize         int               `mapstructure:"frameSize"                                           json:"frameSize,omitempty"`
}

// AMQPQueueBindConfig AMQP Message Bus QueueBind Configuration.
type AMQPQueueBindConfig struct {
	ExtraArgs map[string]any `mapstructure:"extraArgs" json:"extraArgs,omitempty"`
	Name      string         `mapstructure:"name"      json:"name,omitempty"      validate:"required"`
	Key       string         `mapstructure:"key"       json:"key,omitempty"       validate:"required"`
	Exchange  string         `mapstructure:"exchange"  json:"exchange,omitempty"  validate:"required"`
	NoWait    bool           `mapstructure:"noWait"    json:"noWait,omitempty"`
}

// AMQPQueueConfig AMQP Message Bus Queue configuration.
type AMQPQueueConfig struct {
	ExtraArgs  map[string]any `mapstructure:"extraArgs"  json:"extraArgs,omitempty"`
	Name       string         `mapstructure:"name"       json:"name,omitempty"       validate:"required"`
	Durable    bool           `mapstructure:"durable"    json:"durable,omitempty"`
	AutoDelete bool           `mapstructure:"autoDelete" json:"autoDelete,omitempty"`
	Exclusive  bool           `mapstructure:"exclusive"  json:"exclusive,omitempty"`
	NoWait     bool           `mapstructure:"noWait"     json:"noWait,omitempty"`
}

// AMQPExchangeConfig AMQP Message Bus Exchange configuration.
type AMQPExchangeConfig struct {
	ExtraArgs  map[string]any `mapstructure:"extraArgs"  json:"extraArgs,omitempty"`
	Name       string         `mapstructure:"name"       json:"name,omitempty"       validate:"required"`
	Type       string         `mapstructure:"type"       json:"type,omitempty"       validate:"required"`
	Durable    bool           `mapstructure:"durable"    json:"durable,omitempty"`
	AutoDelete bool           `mapstructure:"autoDelete" json:"autoDelete,omitempty"`
	Internal   bool           `mapstructure:"internal"   json:"internal,omitempty"`
	NoWait     bool           `mapstructure:"noWait"     json:"noWait,omitempty"`
}

// LockDistributorConfig Lock distributor configuration.
type LockDistributorConfig struct {
	TableName          string `mapstructure:"tableName"          validate:"required" json:"tableName,omitempty"`
	LeaseDuration      string `mapstructure:"leaseDuration"      validate:"required" json:"leaseDuration,omitempty"`
	HeartbeatFrequency string `mapstructure:"heartbeatFrequency" validate:"required" json:"heartbeatFrequency,omitempty"`
}

// OIDCAuthConfig OpenID Connect authentication configurations.
type OIDCAuthConfig struct {
	ClientSecret      *CredentialConfig `mapstructure:"clientSecret"      validate:"omitempty"     json:"clientSecret,omitempty"`
	ClientID          string            `mapstructure:"clientId"          validate:"required"      json:"clientId,omitempty"`
	IssuerURL         string            `mapstructure:"issuerUrl"         validate:"required,url"  json:"issuerUrl,omitempty"`
	RedirectURL       string            `mapstructure:"redirectUrl"       validate:"required,url"  json:"redirectUrl,omitempty"`
	LogoutRedirectURL string            `mapstructure:"logoutRedirectUrl" validate:"omitempty,url" json:"logoutRedirectUrl,omitempty"`
	State             string            `mapstructure:"state"             validate:"required"      json:"state,omitempty"`
	CookieName        string            `mapstructure:"cookieName"                                 json:"cookieName,omitempty"`
	Scopes            []string          `mapstructure:"scopes"                                     json:"scopes,omitempty"`
	EmailVerified     bool              `mapstructure:"emailVerified"                              json:"emailVerified,omitempty"`
	CookieSecure      bool              `mapstructure:"cookieSecure"                               json:"cookieSecure,omitempty"`
}

// OPAServerAuthorization OPA Server authorization.
type OPAServerAuthorization struct {
	Tags map[string]string `mapstructure:"tags" json:"tags,omitempty"`
	URL  string            `mapstructure:"url"  json:"url,omitempty"  validate:"required,url"`
}

// TracingConfig represents the Tracing configuration structure.
type TracingConfig struct {
	FixedTags    map[string]string      `mapstructure:"fixedTags"    json:"fixedTags,omitempty"`
	OtelHTTP     *TracingOtelHTTPConfig `mapstructure:"otelHttp"     json:"otelHttp,omitempty"     validate:"required_if=Type OTEL_HTTP Enabled true"`
	SamplerCfg   *TracingSamplerConfig  `mapstructure:"samplerCfg"   json:"samplerCfg,omitempty"`
	SamplerType  string                 `mapstructure:"samplerType"  json:"samplerType,omitempty"  validate:"omitempty,oneof=ALWAYS_OFF ALWAYS_ON RATIO"`
	Type         string                 `mapstructure:"type"         json:"type,omitempty"         validate:"oneof=OTEL_HTTP"`
	MaxQueueSize int                    `mapstructure:"maxQueueSize" json:"maxQueueSize,omitempty" validate:"omitempty,gte=0"`
	MaxBatchSize int                    `mapstructure:"maxBatchSize" json:"maxBatchSize,omitempty" validate:"omitempty,gte=0"`
	Enabled      bool                   `mapstructure:"enabled"      json:"enabled,omitempty"`
}

// TracingSamplerConfig Tracing Sampler configuration.
type TracingSamplerConfig struct {
	RatioCfg *TracingRatioSamplerConfig `mapstructure:"ratioCfg" json:"ratioCfg,omitempty"`
}

// TracingRatioSamplerConfig Tracing Ratio Sampler configuration.
type TracingRatioSamplerConfig struct {
	Ratio float64 `mapstructure:"ratio" validate:"required,gte=0" json:"ratio,omitempty"`
}

// TracingOtelHTTPConfig represents the OTEL HTTP configuration structure.
type TracingOtelHTTPConfig struct {
	ServerURL     string            `mapstructure:"serverUrl" validate:"required,http_url" json:"serverUrl,omitempty"`
	Headers       map[string]string `mapstructure:"headers"                                json:"headers,omitempty"`
	TimeoutString string            `mapstructure:"timeout"                                json:"timeoutString,omitempty"`
}

// LogConfig Log configuration.
type LogConfig struct {
	Level    string `mapstructure:"level"    validate:"required" json:"level,omitempty"`
	Format   string `mapstructure:"format"   validate:"required" json:"format,omitempty"`
	FilePath string `mapstructure:"filePath"                     json:"filePath,omitempty"`
}

// ServerConfig Server configuration.
type ServerConfig struct {
	CORS       *ServerCorsConfig     `mapstructure:"cors"       validate:"omitempty" json:"cors,omitempty"`
	Compress   *ServerCompressConfig `mapstructure:"compress"                        json:"compress,omitempty"`
	ListenAddr string                `mapstructure:"listenAddr"                      json:"listenAddr,omitempty"`
	Port       int                   `mapstructure:"port"       validate:"required"  json:"port,omitempty"`
}

// ServerCompressConfig Server compress configuration.
type ServerCompressConfig struct {
	Enabled bool `mapstructure:"enabled" json:"enabled,omitempty"`
}

// ServerCorsConfig Server CORS configuration.
type ServerCorsConfig struct {
	AllowCredentials        *bool    `mapstructure:"allowCredentials"        json:"allowCredentials,omitempty"`
	AllowWildcard           *bool    `mapstructure:"allowWildcard"           json:"allowWildcard,omitempty"`
	AllowBrowserExtensions  *bool    `mapstructure:"allowBrowserExtensions"  json:"allowBrowserExtensions,omitempty"`
	AllowWebSockets         *bool    `mapstructure:"allowWebSockets"         json:"allowWebSockets,omitempty"`
	AllowFiles              *bool    `mapstructure:"allowFiles"              json:"allowFiles,omitempty"`
	AllowAllOrigins         *bool    `mapstructure:"allowAllOrigins"         json:"allowAllOrigins,omitempty"`
	MaxAgeDuration          string   `mapstructure:"maxAgeDuration"          json:"maxAgeDuration,omitempty"`
	AllowOrigins            []string `mapstructure:"allowOrigins"            json:"allowOrigins,omitempty"`
	AllowMethods            []string `mapstructure:"allowMethods"            json:"allowMethods,omitempty"`
	AllowHeaders            []string `mapstructure:"allowHeaders"            json:"allowHeaders,omitempty"`
	ExposeHeaders           []string `mapstructure:"exposeHeaders"           json:"exposeHeaders,omitempty"`
	UseDefaultConfiguration bool     `mapstructure:"useDefaultConfiguration" json:"useDefaultConfiguration,omitempty"`
}

// DatabaseConfig Database configuration.
type DatabaseConfig struct {
	ConnectionURL                    *CredentialConfig   `mapstructure:"connectionUrl"                    validate:"required"                       json:"connectionUrl,omitempty"`
	Driver                           string              `mapstructure:"driver"                           validate:"required,oneof=POSTGRES SQLITE" json:"driver,omitempty"`
	SQLConnectionMaxLifetimeDuration string              `mapstructure:"sqlConnectionMaxLifetimeDuration"                                           json:"sqlConnectionMaxLifetimeDuration,omitempty"`
	ReplicaConnectionURLs            []*CredentialConfig `mapstructure:"replicaConnectionUrls"                                                      json:"replicaConnectionUrls,omitempty"`
	SQLMaxIdleConnections            int                 `mapstructure:"sqlMaxIdleConnections"                                                      json:"sqlMaxIdleConnections,omitempty"`
	SQLMaxOpenConnections            int                 `mapstructure:"sqlMaxOpenConnections"                                                      json:"sqlMaxOpenConnections,omitempty"`
	DisableForeignKeyWhenMigrating   bool                `mapstructure:"disableForeignKeyWhenMigrating"                                             json:"disableForeignKeyWhenMigrating,omitempty"`
	AllowGlobalUpdate                bool                `mapstructure:"allowGlobalUpdate"                                                          json:"allowGlobalUpdate,omitempty"`
	PrepareStatement                 bool                `mapstructure:"prepareStatement"                                                           json:"prepareStatement,omitempty"`
}

// SMTPConfig SMTP Configuration.
type SMTPConfig struct {
	Username           *CredentialConfig `mapstructure:"username"           json:"username,omitempty"`
	Password           *CredentialConfig `mapstructure:"password"           json:"password,omitempty"`
	Host               string            `mapstructure:"host"               json:"host,omitempty"               validation:"fqdn,required"`
	Encryption         string            `mapstructure:"encryption"         json:"encryption,omitempty"         validation:"omitempty,oneof=NONE TLS SSL"`
	AuthenticationType string            `mapstructure:"authenticationType" json:"authenticationType,omitempty" validation:"omitempty,oneof=PLAIN LOGIN CRAM-MD5"`
	ConnectTimeout     string            `mapstructure:"connectTimeout"     json:"connectTimeout,omitempty"`
	SendTimeout        string            `mapstructure:"sendTimeout"        json:"sendTimeout,omitempty"`
	Port               int               `mapstructure:"port"               json:"port,omitempty"               validation:"gt=0,required"`
	KeepAlive          bool              `mapstructure:"keepAlive"          json:"keepAlive,omitempty"`
	TLSSkipVerify      bool              `mapstructure:"tlsSkipVerify"      json:"tlsSkipVerify,omitempty"`
}

// CredentialConfig Credential Configurations.
type CredentialConfig struct {
	Path  string `mapstructure:"path"  validate:"required_without_all=Env Value"  json:"path,omitempty"`
	Env   string `mapstructure:"env"   validate:"required_without_all=Path Value" json:"env,omitempty"`
	Value string `mapstructure:"value" validate:"required"                        json:"-"` // Ignore this key in json marshal
}
