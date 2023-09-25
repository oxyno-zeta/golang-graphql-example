package config

import "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"

// Main configuration folder path.
const DefaultMainConfigFolderPath = "conf/"

// Manager.
//
//go:generate mockgen -destination=./mocks/mock_Manager.go -package=mocks github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config Manager
type Manager interface {
	// Initialize Once
	InitializeOnce() error
	// Load configuration
	Load(inputConfigFilePath string) error
	// Get configuration object
	GetConfig() *Config
	// Add on change hook for configuration change
	AddOnChangeHook(hook func())
}

func NewManager(logger log.Logger) Manager {
	return &managerimpl{logger: logger}
}
