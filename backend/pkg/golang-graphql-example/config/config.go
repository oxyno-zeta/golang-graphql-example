package config

// DefaultLogLevel Default log level
const DefaultLogLevel = "info"

// DefaultLogFormat Default Log format
const DefaultLogFormat = "json"

// DefaultPort Default port
const DefaultPort = 8080

// DefaultInternalPort Default internal port
const DefaultInternalPort = 9090

// Config Configuration object
type Config struct {
	Log            *LogConfig      `mapstructure:"log"`
	Tracing        *TracingConfig  `mapstructure:"tracing"`
	Server         *ServerConfig   `mapstructure:"server"`
	InternalServer *ServerConfig   `mapstructure:"internalServer"`
	Database       *DatabaseConfig `mapstructure:"database" validate:"required"`
}

// TracingConfig represents the Tracing configuration structure
type TracingConfig struct {
	Enabled       bool                   `mapstructure:"enabled"`
	LogSpan       bool                   `mapstructure:"logSpan"`
	FlushInterval string                 `mapstructure:"flushInterval"`
	UDPHost       string                 `mapstructure:"udpHost"`
	QueueSize     int                    `mapstructure:"queueSize"`
	FixedTags     map[string]interface{} `mapstructure:"fixedTags"`
}

// LogConfig Log configuration
type LogConfig struct {
	Level    string `mapstructure:"level" validate:"required"`
	Format   string `mapstructure:"format" validate:"required"`
	FilePath string `mapstructure:"filePath"`
}

// ServerConfig Server configuration
type ServerConfig struct {
	ListenAddr string `mapstructure:"listenAddr"`
	Port       int    `mapstructure:"port" validate:"required"`
}

// DatabaseConfig Database configuration
type DatabaseConfig struct {
	ConnectionURL string `mapstructure:"connectionUrl" validate:"required"`
	Dialect       string `mapstructure:"dialect" validate:"required"`
}

// CredentialConfig Credential Configurations
type CredentialConfig struct {
	Path  string `mapstructure:"path" validate:"required_without_all=Env Value"`
	Env   string `mapstructure:"env" validate:"required_without_all=Path Value"`
	Value string `mapstructure:"value" validate:"required_without_all=Path Env"`
}
