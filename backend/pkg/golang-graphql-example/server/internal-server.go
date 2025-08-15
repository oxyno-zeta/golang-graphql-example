package server

import (
	"net/http"
	"strconv"
	"time"

	"emperror.dev/errors"
	"github.com/gin-gonic/gin"

	gosundheit "github.com/AppsFlyer/go-sundheit"
	healthhttp "github.com/AppsFlyer/go-sundheit/http"
	helmet "github.com/danielkov/gin-helmet/ginhelmet"

	correlationid "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/common/correlation-id"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/metrics"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/signalhandler"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/tracing"
)

const (
	DefaultHealthCheckTimeout = time.Second
)

type InternalServer struct {
	logger           log.Logger
	cfgManager       config.Manager
	metricsSvc       metrics.Service
	signalHandlerSvc signalhandler.Service
	server           *http.Server
	checkers         []*CheckerInput
}

type CheckerInput struct {
	CheckFn      func() error
	Name         string
	Interval     time.Duration
	Timeout      time.Duration
	InitialDelay time.Duration
}

func NewInternalServer(
	logger log.Logger,
	cfgManager config.Manager,
	metricsSvc metrics.Service,
	signalHandlerSvc signalhandler.Service,
) *InternalServer {
	return &InternalServer{
		logger:           logger,
		cfgManager:       cfgManager,
		metricsSvc:       metricsSvc,
		signalHandlerSvc: signalHandlerSvc,
		checkers:         make([]*CheckerInput, 0),
	}
}

// AddChecker allow to add a health checker.
func (svr *InternalServer) AddChecker(chI *CheckerInput) {
	// Append
	svr.checkers = append(svr.checkers, chI)
}

func (svr *InternalServer) generateInternalRouter() (http.Handler, error) {
	// Get configuration
	cfg := svr.cfgManager.GetConfig()

	// Set release mod
	gin.SetMode(gin.ReleaseMode)
	// Create router
	router := gin.New()
	// Add middlewares
	router.Use(gin.Recovery())
	router.Use(helmet.Default())
	router.Use(correlationid.HTTPMiddleware(svr.logger))
	router.Use(log.Middleware(svr.logger, correlationid.GetFromGin, tracing.GetTraceIDFromContext))
	router.Use(svr.metricsSvc.Instrument("internal", true))
	// Add cors if configured
	err := manageCORS(router, cfg.InternalServer)
	// Check error
	if err != nil {
		return nil, err
	}

	// create a new health instance
	h2 := gosundheit.New()

	for _, it := range svr.checkers {
		// Create logger
		logger := svr.logger.WithField("health-check-target", it.Name)

		// Initialize check options
		options := make([]gosundheit.CheckOption, 0)

		// Check if timeout is set, otherwise put a default value
		if it.Timeout == 0 {
			options = append(options, gosundheit.ExecutionTimeout(DefaultHealthCheckTimeout))
		} else {
			options = append(options, gosundheit.ExecutionTimeout(it.Timeout))
		}

		// Check if initial delay is set, otherwise ignore it
		if it.InitialDelay != 0 {
			options = append(options, gosundheit.InitialDelay(it.InitialDelay))
		}

		// Set interval
		options = append(options, gosundheit.ExecutionPeriod(it.Interval))

		// Register check
		err = h2.RegisterCheck(
			&customHealthChecker{
				logger: logger,
				name:   it.Name,
				fn:     it.CheckFn,
			},
			options...,
		)
		// Check error
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	// Add metrics path
	router.GET("/metrics", gin.WrapH(svr.metricsSvc.PrometheusHTTPHandler()))
	router.GET("/health", gin.WrapH(healthhttp.HandleHealthJSON(h2)))
	router.GET("/ready", func(c *gin.Context) {
		// Check if system is in shutdown
		if svr.signalHandlerSvc.IsStoppingSystem() {
			// Response with service unavailable flag
			c.JSON(http.StatusServiceUnavailable, gin.H{"reason": "system stopping"})

			return
		}

		// Otherwise, send health check result
		gin.WrapH(healthhttp.HandleHealthJSON(h2))(c)
	})

	return router, nil
}

func (svr *InternalServer) Listen() error {
	svr.logger.Infof("Internal server listening on %s", svr.server.Addr)
	err := svr.server.ListenAndServe()
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	// Default
	return nil
}

func (svr *InternalServer) GenerateServer() error {
	// Get configuration
	cfg := svr.cfgManager.GetConfig()
	// Generate internal router
	r, err := svr.generateInternalRouter()
	// Check error
	if err != nil {
		return err
	}
	// Create server
	addr := cfg.InternalServer.ListenAddr + ":" + strconv.Itoa(cfg.InternalServer.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: r,
		// Timeout to 10 second to limit the time to read the request headers
		// This is done to patch server against Slowloris Attack
		ReadHeaderTimeout: 10 * time.Second, //nolint: mnd // Ignored to see it clearly
	}
	// Store server
	svr.server = server

	return nil
}
