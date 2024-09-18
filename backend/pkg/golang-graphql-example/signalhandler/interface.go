package signalhandler

import (
	"context"
	"os"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"github.com/samber/lo"
)

//go:generate mockgen -destination=./mocks/mock_Service.go -package=mocks github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/signalhandler Service
type Service interface {
	// InitializeOnce will initialize service.
	// Important note: this must be called only once.
	InitializeOnce() error
	// OnSignal will add a hook on specific signal.
	// Hooks for SIGINT and SIGTERM are called here before checking that all requests are finished and calling onExit hooks.
	OnSignal(signal os.Signal, hook func())
	// OnExit will add a hook that will be called when a SIGINT or SIGTERM is caught and when the application will be closed
	// That will be launched only when all incoming requests are finished.
	OnExit(hook func())
	// Middleware to count active requests.
	ActiveRequestCounterMiddleware() gin.HandlerFunc
	// Is stopping system will return true if the application is stopping.
	IsStoppingSystem() bool
	// IncreaseActiveRequestCounter will increase active request counter by one.
	IncreaseActiveRequestCounter()
	// DecreaseActiveRequestCounter will decrease active request counter by one.
	DecreaseActiveRequestCounter()
	// GetStoppingSystemContext will return a context that will be cancelled on system stopping.
	// This context must be used by calling "ctx.Done()". This will be unblocked on system stopping.
	GetStoppingSystemContext() context.Context
}

func NewService(logger log.Logger, serverMode bool, signalListToNotify []os.Signal) Service {
	// Create signal list to notify
	signalListToNotifyInternal := []os.Signal{syscall.SIGTERM, syscall.SIGINT}
	// Append all items from input inside
	signalListToNotifyInternal = append(signalListToNotifyInternal, signalListToNotify...)
	// Filter to unique
	signalListToNotifyInternal = lo.Uniq(signalListToNotifyInternal)

	// Create ctx that will be cancelled when OS is stopping
	ctx, cancel := context.WithCancel(context.TODO())

	return &service{
		logger:                   logger,
		serverMode:               serverMode,
		signalListToNotify:       signalListToNotifyInternal,
		hooksStorage:             map[os.Signal][]func(){},
		onExitHookStorage:        []func(){},
		activeRequestCounter:     0,
		activeRequestCounterChan: make(chan int64, 100), //nolint:mnd// Ignore this
		cancelCtx:                ctx,
		cancelFunc:               cancel,
	}
}
