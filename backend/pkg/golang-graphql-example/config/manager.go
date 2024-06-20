package config

import (
	"time"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
)

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
	AddOnChangeHook(input *HookDefinition)
	// SetExtraServices will set extra services
	SetExtraServices(metricsSvc MetricsService)
}

//go:generate mockgen -destination=./mocks/mock_MetricsService.go -package=mocks github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config MetricsService
type MetricsService interface {
	// UpFailedConfigReload will raise the failed configuration reload gauge.
	UpFailedConfigReload()
	// DownFailedConfigReload will down the failed configuration reload gauge.
	DownFailedConfigReload()
}

type HookDefinition struct {
	// Hook to run
	// @mandatory
	Hook func() error
	// Maximum Retry count to try to run hook
	// @optional
	RetryCount int
	// Retry Wait Duration to wait between 2 try
	// @optional
	RetryWaitDuration time.Duration
	// Log fatal on maximum try with an error
	// @optional
	FatalOnMaxTry bool
}

func NewManager(logger log.Logger) Manager {
	return &managerimpl{logger: logger}
}
