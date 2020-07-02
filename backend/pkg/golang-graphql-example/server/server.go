package server

import (
	"net/http"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/metrics"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/generated"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/middlewares"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/tracing"
)

type Server struct {
	logger       log.Logger
	cfgManager   config.Manager
	metricsCl    metrics.Client
	tracingSvc   tracing.Service
	busiServices *business.Services
}

func NewServer(logger log.Logger, cfgManager config.Manager, metricsCl metrics.Client, tracingSvc tracing.Service, busiServices *business.Services) *Server {
	return &Server{
		logger:       logger,
		cfgManager:   cfgManager,
		metricsCl:    metricsCl,
		tracingSvc:   tracingSvc,
		busiServices: busiServices,
	}
}

func (svr *Server) Listen() error {
	// Set release mod
	gin.SetMode(gin.ReleaseMode)
	// Create router
	router := gin.New()
	// Add middlewares
	router.Use(gin.Recovery())
	router.Use(helmet.Default())
	router.Use(middlewares.RequestID(svr.logger))
	router.Use(svr.tracingSvc.Middleware(middlewares.GetRequestIDFromContext))
	router.Use(log.Middleware(svr.logger, middlewares.GetRequestIDFromGin, tracing.GetSpanIDFromContext))
	router.Use(svr.metricsCl.Instrument())

	// Add static files
	router.Use(static.Serve("/", static.LocalFile("static/", true)))

	// Add graphql endpoints
	router.POST("/api/graphql", graphqlHandler(svr.busiServices))
	router.GET("/api/graphql", playgroundHandler())

	// Get configuration
	cfg := svr.cfgManager.GetConfig()

	// Create server
	addr := cfg.Server.ListenAddr + ":" + strconv.Itoa(cfg.Server.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	// Listen
	svr.logger.Infof("Server listening on %s", addr)

	return server.ListenAndServe()
}

// Defining the Graphql handler
func graphqlHandler(busiServices *business.Services) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graphql.Resolver{BusiServices: busiServices},
	}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/api/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
