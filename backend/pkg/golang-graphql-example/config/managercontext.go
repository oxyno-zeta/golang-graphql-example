package config

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/go-playground/validator/v10"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"github.com/spf13/viper"
	"github.com/thoas/go-funk"
)

// Main configuration folder path
const mainConfigFolderPath = "conf/"

var validate = validator.New()

type managercontext struct {
	cfg           *Config
	configs       []*viper.Viper
	onChangeHooks []func()
	logger        log.Logger
}

func (ctx *managercontext) AddOnChangeHook(hook func()) {
	ctx.onChangeHooks = append(ctx.onChangeHooks, hook)
}

func (ctx *managercontext) Load() error {
	// List files
	files, err := ioutil.ReadDir(mainConfigFolderPath)
	if err != nil {
		return err
	}

	// Generate viper instances for static configs
	ctx.configs = generateViperInstances(files)

	// Put default values
	ctx.loadDefaultConfigurationValues()

	err = ctx.loadConfiguration()

	// Loop over reloadable configs
	for _, vip := range ctx.configs {
		// Watch for configuration changes
		vip.WatchConfig()
		// Add hooks for on change events
		vip.OnConfigChange(func(in fsnotify.Event) {
			ctx.logger.Infof("Reload configuration detected for file %s", in.Name)
			// Reload config
			err2 := ctx.loadConfiguration()
			if err2 != nil {
				ctx.logger.Error(err2)
				// Stop here and do not call hooks => configuration is unstable
				return
			}
			// Call all hooks
			funk.ForEach(ctx.onChangeHooks, func(hook func()) { hook() })
		})
	}

	return err
}

func (ctx *managercontext) loadDefaultConfigurationValues() {
	// Load default configuration
	viper.SetDefault("log.level", DefaultLogLevel)
	viper.SetDefault("log.format", DefaultLogFormat)
	viper.SetDefault("server.port", DefaultPort)
	viper.SetDefault("internalServer.port", DefaultInternalPort)
}

func generateViperInstances(files []os.FileInfo) []*viper.Viper {
	list := make([]*viper.Viper, 0)
	// Loop over static files to create viper instance for them
	funk.ForEach(files, func(file os.FileInfo) {
		filename := file.Name()
		// Create config file name
		cfgFileName := strings.TrimSuffix(filename, path.Ext(filename))
		// Test if config file name is compliant (ignore hidden files like .keep)
		if cfgFileName != "" {
			// Create new viper instance
			vip := viper.New()
			// Set config name
			vip.SetConfigName(cfgFileName)
			// Add configuration path
			vip.AddConfigPath(mainConfigFolderPath)
			// Append it
			list = append(list, vip)
		}
	})

	return list
}

func (ctx *managercontext) loadConfiguration() error {
	// Loop over configs
	for _, vip := range ctx.configs {
		err := vip.ReadInConfig()
		if err != nil {
			return err
		}

		err = viper.MergeConfigMap(vip.AllSettings())
		if err != nil {
			return err
		}
	}

	// Prepare configuration object
	var out Config
	// Quick unmarshal.
	err := viper.Unmarshal(&out)
	if err != nil {
		return err
	}

	// Configuration validation
	err = validate.Struct(out)
	if err != nil {
		return err
	}

	ctx.cfg = &out

	return nil
}

// GetConfig allow to get configuration object
func (ctx *managercontext) GetConfig() *Config {
	return ctx.cfg
}
