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
	Server         *ServerConfig   `mapstructure:"server"`
	InternalServer *ServerConfig   `mapstructure:"internalServer"`
	Database       *DatabaseConfig `mapstructure:"database" validate:"required"`
}

// LogConfig Log configuration
type LogConfig struct {
	Level  string `mapstructure:"level" validate:"required"`
	Format string `mapstructure:"format"`
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
