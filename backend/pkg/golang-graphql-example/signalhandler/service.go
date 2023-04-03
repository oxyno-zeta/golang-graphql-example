package signalhandler

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
)

const (
	TimeBetweenChecks = 200 * time.Millisecond
)

type service struct {
	logger                   log.Logger
	hooksStorage             map[os.Signal][]func()
	activeRequestCounterChan chan int64
	signalListToNotify       []os.Signal
	onExitHookStorage        []func()
	activeRequestCounter     int64
	serverMode               bool
	stoppingSysInProgress    bool
}

func (s *service) IncreaseActiveRequestCounter() {
	// Check if server mode isn't enabled
	// In this case, ignore the call
	if !s.serverMode {
		return
	}

	s.activeRequestCounterChan <- 1
}

func (s *service) DecreaseActiveRequestCounter() {
	// Check if server mode isn't enabled
	// In this case, ignore the call
	if !s.serverMode {
		return
	}

	s.activeRequestCounterChan <- -1
}

func (s *service) IsStoppingSystem() bool {
	return s.stoppingSysInProgress
}

func (s *service) InitializeOnce() error {
	// Create signal channel
	signalChan := make(chan os.Signal, 1)

	// Notify watcher
	signal.Notify(
		signalChan,
		s.signalListToNotify...,
	)

	go func() {
		for sig := range signalChan {
			// Log
			s.logger.Infof("Catching signal \"%s\"", sig)

			// Get hook list for signal
			hooks := s.hooksStorage[sig]

			// Run all hooks
			for _, h := range hooks {
				// Start hook
				h()
			}

			// Check if signal is SIGTERM or SIGINT
			if sig == syscall.SIGINT || sig == syscall.SIGTERM {
				// Run stop hook
				s.stoppingAppHook()
			}
		}
	}()

	// Initialize server mode
	s.initializeServerMode()

	// Default
	return nil
}

func (s *service) OnExit(hook func()) {
	s.onExitHookStorage = append(s.onExitHookStorage, hook)
}

func (s *service) OnSignal(signal os.Signal, hook func()) {
	// Check if array exist
	if s.hooksStorage[signal] != nil {
		// Create list
		s.hooksStorage[signal] = make([]func(), 0)
	}

	// Add item
	s.hooksStorage[signal] = append(s.hooksStorage[signal], hook)
}

func (s *service) stoppingAppHook() {
	// Check if application is already marked as in stopping mode
	if s.stoppingSysInProgress {
		// Avoid starting a new go routine for the same thing
		return
	}

	// Create ticker
	ticker := time.NewTicker(TimeBetweenChecks)

	// Starting the go routine
	go func() {
		// Loop
		for range ticker.C {
			// Log
			s.logger.Debug("Checking is application can be stopped")
			// Check if requests still in progress
			if s.activeRequestCounter == 0 {
				// Log
				s.logger.Info("Stopping application")
				// Run on exit all hooks
				for _, h := range s.onExitHookStorage {
					// Start hook
					h()
				}
				// Stopping application
				os.Exit(0)
			} else {
				s.logger.Infof("Cannot stop application yet, still detecting %d requests", s.activeRequestCounter)
			}
		}
	}()

	// Updating the stopping flag
	s.stoppingSysInProgress = true
}
