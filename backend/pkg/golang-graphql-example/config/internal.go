package config

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"emperror.dev/errors"
	"github.com/fsnotify/fsnotify"
	"github.com/go-playground/validator/v10"
	"github.com/samber/lo"
	"github.com/spf13/viper"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
)

// TemplateErrLoadingEnvCredentialEmpty Template Error when Loading Environment variable Credentials.
var TemplateErrLoadingEnvCredentialEmpty = "error loading credentials, environment variable %s is empty" //nolint: gosec // False positive

var validate = validator.New()

type managerimpl struct {
	cfg                       *Config
	configs                   []*viper.Viper
	onChangeHooks             []*HookDefinition
	logger                    log.Logger
	metricsSvc                MetricsService
	internalFileWatchChannels []chan bool
	credentialConfigPathList  [][]string
}

func (impl *managerimpl) AddOnChangeHook(input *HookDefinition) {
	// Check if retry count is set to set a default value
	if input.RetryCount == 0 {
		input.RetryCount = 1
	}

	// Save
	impl.onChangeHooks = append(impl.onChangeHooks, input)
}

func (impl *managerimpl) Load(inputConfigFilePath string) error {
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
	impl.configs = impl.generateViperInstances(files, configFolderPath)

	// Load configuration
	err = impl.loadConfiguration()
	if err != nil {
		return err
	}

	// Loop over config files
	lo.ForEach(impl.configs, func(vip *viper.Viper, _ int) {
		// Add hooks for on change events
		vip.OnConfigChange(func(_ fsnotify.Event) {
			impl.logger.Infof("Reload configuration detected for file %s", vip.ConfigFileUsed())

			// Reload config
			err2 := impl.loadConfiguration()
			if err2 != nil {
				impl.logger.Error(err2)

				// Check if metrics service is set
				if impl.metricsSvc != nil {
					// Up metrics
					impl.metricsSvc.UpFailedConfigReload()
				}

				// Stop here and do not call hooks => configuration is unstable
				return
			}
			// Call all hooks in sequence in order to manage correctly reload database and after
			// services that depends on it
			impl.runAllHooks()
		})
		// Watch for configuration changes
		vip.WatchConfig()
	})

	return nil
}

// Imported and modified from viper v1.7.0.
func (impl *managerimpl) watchInternalFile(filePath string, forceStop chan bool, onChange func()) {
	initWG := sync.WaitGroup{}
	initWG.Add(1)

	go func() {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			impl.logger.Fatal(errors.WithStack(err))
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
						impl.logger.Errorf("watcher error: %v\n", err)
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

func (impl *managerimpl) generateViperInstances(files []os.DirEntry, configFolderPath string) []*viper.Viper {
	list := make([]*viper.Viper, 0)
	// Loop over static files to create viper instance for them
	lo.ForEach(files, func(file os.DirEntry, _ int) {
		filename := file.Name()
		// Create config file name
		cfgFileName := strings.TrimSuffix(filename, path.Ext(filename))
		// Test if config file name is compliant (ignore hidden files like .keep or directory)
		if !strings.HasPrefix(filename, ".") && cfgFileName != "" && !file.IsDir() {
			// Create new viper instance
			vip := viper.NewWithOptions(viper.WithLogger(impl.logger.GetSlogInstance()))
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

func (impl *managerimpl) loadConfiguration() error {
	// Load must start by flushing all existing watcher on internal files
	for i := 0; i < len(impl.internalFileWatchChannels); i++ {
		ch := impl.internalFileWatchChannels[i]
		// Send the force stop
		ch <- true
	}

	// Create a viper instance for default value and merging
	globalViper := viper.NewWithOptions(viper.WithLogger(impl.logger.GetSlogInstance()))

	// Put default values
	impl.loadDefaultConfigurationValues(globalViper)

	// Loop over configs
	for _, vip := range impl.configs {
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

	// Parse time, duration, regex, ...
	err = parseValues(&out)
	// Check error
	if err != nil {
		return err
	}

	// Load all credentials
	credentials, err := impl.loadAllCredentials(&out)
	if err != nil {
		return err
	}
	// Initialize or flush watch internal file channels
	internalFileWatchChannels := make([]chan bool, 0)
	impl.internalFileWatchChannels = internalFileWatchChannels
	// Loop over all credentials in order to watch file change
	lo.ForEach(credentials, func(cred *CredentialConfig, _ int) {
		// Check if credential is about a path
		if cred != nil && cred.Path != "" {
			// Create channel
			ch := make(chan bool)
			// Run the watch file
			impl.watchInternalFile(cred.Path, ch, func() {
				// File change detected
				impl.logger.Infof("Reload credential file detected for path %s", cred.Path)

				// Reload config
				err2 := loadCredential(cred)
				if err2 != nil {
					impl.logger.Error(err2)
					// Stop here and do not call hooks => configuration is unstable
					return
				}
				// Call all hooks in sequence in order to manage correctly reload database and after
				// services that depends on it
				impl.runAllHooks()
			})
			// Add channel to list of channels
			impl.internalFileWatchChannels = append(impl.internalFileWatchChannels, ch)
		}
	})

	// Configuration validation
	err = validate.Struct(out)
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	err = validateBusinessConfig(&out)
	if err != nil {
		return err
	}

	impl.cfg = &out

	return nil
}

func loadCredential(credCfg *CredentialConfig) error {
	// Check nil
	if credCfg == nil {
		return nil
	}

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
func (impl *managerimpl) GetConfig() *Config {
	return impl.cfg
}

func (impl *managerimpl) InitializeOnce() error {
	cl, err := getCredentialConfigPathList()
	// Check error
	if err != nil {
		return err
	}

	// Save
	impl.credentialConfigPathList = cl

	// Default
	return nil
}

func (impl *managerimpl) SetExtraServices(metricsSvc MetricsService) {
	impl.metricsSvc = metricsSvc
}

func (impl *managerimpl) runAllHooks() {
	// Init failed
	failed := false
	// Run all hooks
	lo.ForEach(impl.onChangeHooks, func(h *HookDefinition, _ int) {
		// Run hook
		r := impl.runReloadHook(h)
		// Save failed
		failed = failed || r
	})

	// Check if metrics services is set
	if impl.metricsSvc != nil {
		// If failed, up metric. Otherwise, down it
		if failed {
			impl.metricsSvc.UpFailedConfigReload()
		} else {
			impl.metricsSvc.DownFailedConfigReload()
		}
	}
}

func (impl *managerimpl) runReloadHook(h *HookDefinition) (res bool) {
	// Loop on retry count
	for i := range h.RetryCount {
		// Run hook
		err := h.Hook()
		// Check if error is empty
		if err == nil {
			// Stop here
			return res
		}

		// Save last fail
		res = i == (h.RetryCount - 1)

		// Check if we are on max try and fatal error
		if res && h.FatalOnMaxTry {
			// Log fatal error
			impl.logger.Fatal(err)
		} else {
			// Log error
			impl.logger.Error(err)
		}

		// Sleep the wait duration
		time.Sleep(h.RetryWaitDuration)
	}

	return res
}
