package config

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"emperror.dev/errors"
	"github.com/fsnotify/fsnotify"
	"github.com/go-playground/validator/v10"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"github.com/spf13/viper"
	"github.com/thoas/go-funk"
)

// TemplateErrLoadingEnvCredentialEmpty Template Error when Loading Environment variable Credentials.
var TemplateErrLoadingEnvCredentialEmpty = "error loading credentials, environment variable %s is empty" //nolint: gosec // False positive

var validate = validator.New()

type managerimpl struct {
	cfg                       *Config
	configs                   []*viper.Viper
	onChangeHooks             []func()
	logger                    log.Logger
	internalFileWatchChannels []chan bool
}

func (ctx *managerimpl) AddOnChangeHook(hook func()) {
	ctx.onChangeHooks = append(ctx.onChangeHooks, hook)
}

func (ctx *managerimpl) Load(inputConfigFilePath string) error {
	// Initialize config file folder path
	configFolderPath := DefaultMainConfigFolderPath
	// Check if input is set to change for this one
	if inputConfigFilePath != "" {
		configFolderPath = inputConfigFilePath
	}

	// List files
	files, err := os.ReadDir(configFolderPath)
	if err != nil {
		return errors.WithStack(err)
	}

	// Generate viper instances for static configs
	ctx.configs = generateViperInstances(files, configFolderPath)

	// Load configuration
	err = ctx.loadConfiguration()
	if err != nil {
		return err
	}

	// Loop over config files
	funk.ForEach(ctx.configs, func(vip *viper.Viper) {
		// Add hooks for on change events
		vip.OnConfigChange(func(in fsnotify.Event) {
			ctx.logger.Infof("Reload configuration detected for file %s", vip.ConfigFileUsed())

			// Reload config
			err2 := ctx.loadConfiguration()
			if err2 != nil {
				ctx.logger.Error(err2)
				// Stop here and do not call hooks => configuration is unstable
				return
			}
			// Call all hooks in sequence in order to manage correctly reload database and after
			// services that depends on it
			funk.ForEach(ctx.onChangeHooks, func(hook func()) { hook() })
		})
		// Watch for configuration changes
		vip.WatchConfig()
	})

	return nil
}

// Imported and modified from viper v1.7.0.
func (ctx *managerimpl) watchInternalFile(filePath string, forceStop chan bool, onChange func()) {
	initWG := sync.WaitGroup{}
	initWG.Add(1)

	go func() {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			ctx.logger.Fatal(errors.WithStack(err))
		}
		defer watcher.Close()

		configFile := filepath.Clean(filePath)
		configDir, _ := filepath.Split(configFile)
		realConfigFile, _ := filepath.EvalSymlinks(filePath)

		eventsWG := sync.WaitGroup{}
		eventsWG.Add(1)

		go func() {
			for {
				select {
				case <-forceStop:
					eventsWG.Done()

					return
				case event, ok := <-watcher.Events:
					if !ok { // 'Events' channel is closed
						eventsWG.Done()

						return
					}

					currentConfigFile, _ := filepath.EvalSymlinks(filePath)
					// we only care about the config file with the following cases:
					// 1 - if the config file was modified or created
					// 2 - if the real path to the config file changed (eg: k8s ConfigMap replacement)
					const writeOrCreateMask = fsnotify.Write | fsnotify.Create
					if (filepath.Clean(event.Name) == configFile &&
						event.Op&writeOrCreateMask != 0) ||
						(currentConfigFile != "" && currentConfigFile != realConfigFile) {
						realConfigFile = currentConfigFile

						// Call on change
						onChange()
					} else if filepath.Clean(event.Name) == configFile && event.Op&fsnotify.Remove&fsnotify.Remove != 0 {
						eventsWG.Done()

						return
					}

				case err, ok := <-watcher.Errors:
					if ok { // 'Errors' channel is not closed
						ctx.logger.Errorf("watcher error: %v\n", err)
					}

					eventsWG.Done()

					return
				}
			}
		}()

		_ = watcher.Add(configDir)

		initWG.Done()   // done initializing the watch in this go routine, so the parent routine can move on...
		eventsWG.Wait() // now, wait for event loop to end in this go-routine...
	}()
	initWG.Wait() // make sure that the go routine above fully ended before returning
}

func generateViperInstances(files []os.DirEntry, configFolderPath string) []*viper.Viper {
	list := make([]*viper.Viper, 0)
	// Loop over static files to create viper instance for them
	funk.ForEach(files, func(file os.DirEntry) {
		filename := file.Name()
		// Create config file name
		cfgFileName := strings.TrimSuffix(filename, path.Ext(filename))
		// Test if config file name is compliant (ignore hidden files like .keep or directory)
		if !strings.HasPrefix(filename, ".") && cfgFileName != "" && !file.IsDir() {
			// Create new viper instance
			vip := viper.New()
			// Set config name
			vip.SetConfigName(cfgFileName)
			// Add configuration path
			vip.AddConfigPath(configFolderPath)
			// Append it
			list = append(list, vip)
		}
	})

	return list
}

func (ctx *managerimpl) loadConfiguration() error {
	// Load must start by flushing all existing watcher on internal files
	for i := 0; i < len(ctx.internalFileWatchChannels); i++ {
		ch := ctx.internalFileWatchChannels[i]
		// Send the force stop
		ch <- true
	}

	// Create a viper instance for default value and merging
	globalViper := viper.New()

	// Put default values
	ctx.loadDefaultConfigurationValues(globalViper)

	// Loop over configs
	for _, vip := range ctx.configs {
		// Read configuration
		err := vip.ReadInConfig()
		// Check error
		if err != nil {
			return errors.WithStack(err)
		}

		// Merge all configurations
		err = globalViper.MergeConfigMap(vip.AllSettings())
		// Check error
		if err != nil {
			return errors.WithStack(err)
		}
	}

	// Prepare configuration object
	var out Config
	// Quick unmarshal.
	err := globalViper.Unmarshal(&out)
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	// Load default values
	err = loadBusinessDefaultValues(&out)
	if err != nil {
		return err
	}

	// Configuration validation
	err = validate.Struct(out)
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	// Parse time, duration, regex, ...
	err = parseValues(&out)
	// Check error
	if err != nil {
		return err
	}

	// Load all credentials
	credentials, err := loadAllCredentials(&out)
	if err != nil {
		return err
	}
	// Initialize or flush watch internal file channels
	internalFileWatchChannels := make([]chan bool, 0)
	ctx.internalFileWatchChannels = internalFileWatchChannels
	// Loop over all credentials in order to watch file change
	funk.ForEach(credentials, func(cred *CredentialConfig) {
		// Check if credential is about a path
		if cred.Path != "" {
			// Create channel
			ch := make(chan bool)
			// Run the watch file
			ctx.watchInternalFile(cred.Path, ch, func() {
				// File change detected
				ctx.logger.Infof("Reload credential file detected for path %s", cred.Path)

				// Reload config
				err2 := loadCredential(cred)
				if err2 != nil {
					ctx.logger.Error(err2)
					// Stop here and do not call hooks => configuration is unstable
					return
				}
				// Call all hooks in sequence in order to manage correctly reload database and after
				// services that depends on it
				funk.ForEach(ctx.onChangeHooks, func(hook func()) { hook() })
			})
			// Add channel to list of channels
			ctx.internalFileWatchChannels = append(ctx.internalFileWatchChannels, ch)
		}
	})

	err = validateBusinessConfig(&out)
	if err != nil {
		return err
	}

	ctx.cfg = &out

	return nil
}

func loadCredential(credCfg *CredentialConfig) error {
	if credCfg.Path != "" {
		// Secret file
		databytes, err := os.ReadFile(credCfg.Path)
		// Check error
		if err != nil {
			return errors.WithStack(err)
		}
		// Store value
		credCfg.Value = string(databytes)
	} else if credCfg.Env != "" {
		// Environment variable
		envValue := os.Getenv(credCfg.Env)
		if envValue == "" {
			return errors.Errorf(TemplateErrLoadingEnvCredentialEmpty, credCfg.Env)
		}
		// Store value
		credCfg.Value = envValue
	}
	// Default value
	return nil
}

// GetConfig allow to get configuration object.
func (ctx *managerimpl) GetConfig() *Config {
	return ctx.cfg
}
