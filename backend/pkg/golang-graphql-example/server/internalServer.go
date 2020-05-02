package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/InVisionApp/go-health/v2"
	"github.com/InVisionApp/go-health/v2/handlers"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/metrics"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/middlewares"
)

type InternalServer struct {
	logger     log.Logger
	cfgManager config.Manager
	metricsCl  metrics.Client
	checkers   []*health.Config
}

type CheckerInput struct {
	Name     string
	CheckFn  func() error
	Interval time.Duration
}

func NewInternalServer(logger log.Logger, cfgManager config.Manager, metricsCl metrics.Client) *InternalServer {
	return &InternalServer{
		logger:     logger,
		cfgManager: cfgManager,
		metricsCl:  metricsCl,
		checkers:   make([]*health.Config, 0),
	}
}

func (svr *InternalServer) AddChecker(chI *CheckerInput) {
	svr.checkers = append(svr.checkers, &health.Config{
		Name:     chI.Name,
		Checker:  &customHealthChecker{fn: chI.CheckFn},
		Fatal:    true,
		Interval: chI.Interval,
	})
}

func (svr *InternalServer) Listen() error {
	// Set release mod
	gin.SetMode(gin.ReleaseMode)
	// Create router
	router := gin.New()
	// Add middlewares
	router.Use(gin.Recovery())
	router.Use(helmet.Default())
	router.Use(middlewares.RequestID(svr.logger))
	router.Use(middlewares.LogMiddleware(svr.logger))
	router.Use(svr.metricsCl.Instrument())

	// Create a new health instance
	h := health.New()

	// Disable logging
	h.DisableLogging()

	// Add checkers
	err := h.AddChecks(svr.checkers)
	if err != nil {
		return err
	}

	// Start health checker
	err = h.Start()
	if err != nil {
		return err
	}

	// Add metrics path
	router.GET("/metrics", gin.WrapH(svr.metricsCl.GetPrometheusHTTPHandler()))
	router.GET("/health", gin.WrapH(handlers.NewJSONHandlerFunc(h, nil)))

	// Get configuration
	cfg := svr.cfgManager.GetConfig()

	// Create server
	addr := cfg.InternalServer.ListenAddr + ":" + strconv.Itoa(cfg.InternalServer.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	// Listen
	svr.logger.Infof("Internal Server listening on %s", addr)

	return server.ListenAndServe()
}
