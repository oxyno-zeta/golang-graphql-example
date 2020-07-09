package server

import (
	"net/http"
	"strconv"

	"github.com/99designs/gqlgen-contrib/gqlopentracing"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/apollotracing"
	"github.com/99designs/gqlgen/graphql/playground"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/authentication"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/authorization"
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
	logger            log.Logger
	cfgManager        config.Manager
	metricsCl         metrics.Client
	tracingSvc        tracing.Service
	busiServices      *business.Services
	authenticationSvc authentication.Client
	server            *http.Server
}

// nolint:whitespace
func NewServer(
	logger log.Logger, cfgManager config.Manager, metricsCl metrics.Client,
	tracingSvc tracing.Service, busiServices *business.Services,
	authenticationSvc authentication.Client,
) *Server {
	return &Server{
		logger:            logger,
		cfgManager:        cfgManager,
		metricsCl:         metricsCl,
		tracingSvc:        tracingSvc,
		busiServices:      busiServices,
		authenticationSvc: authenticationSvc,
	}
}

func (svr *Server) GenerateServer() error {
	// Get configuration
	cfg := svr.cfgManager.GetConfig()
	// Generate router
	r, err := svr.generateRouter()
	if err != nil {
		return err
	}

	// Create server
	addr := cfg.Server.ListenAddr + ":" + strconv.Itoa(cfg.Server.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	// Prepare for configuration onChange
	svr.cfgManager.AddOnChangeHook(func() {
		// Generate router
		r, err2 := svr.generateRouter()
		if err2 != nil {
			svr.logger.Fatal(err2)
		}
		// Change server handler
		server.Handler = r
		svr.logger.Info("Server handler reloaded")
	})

	// Store server
	svr.server = server

	return nil
}

func (svr *Server) generateRouter() (http.Handler, error) {
	// Get configuration
	cfg := svr.cfgManager.GetConfig()
	// Set release mod
	gin.SetMode(gin.ReleaseMode)
	// Create router
	router := gin.New()
	// Manage no route
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "404 not found"})
	})
	// Add middlewares
	router.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithDecompressFn(gzip.DefaultDecompressHandle)))
	router.Use(gin.Recovery())
	router.Use(helmet.Default())
	router.Use(middlewares.RequestID(svr.logger))
	router.Use(svr.tracingSvc.Middleware(middlewares.GetRequestIDFromContext))
	router.Use(log.Middleware(svr.logger, middlewares.GetRequestIDFromGin, tracing.GetSpanIDFromContext))
	router.Use(svr.metricsCl.Instrument())

	// Add authentication middleware if configuration exists
	if cfg.OIDCAuthentication != nil {
		// Add endpoints
		err := svr.authenticationSvc.OIDCEndpoints(router)
		// Check error
		if err != nil {
			return nil, err
		}

		// Add authentication middleware
		router.Use(svr.authenticationSvc.Middleware())

		// Add authorization middleware is configuration exists
		if cfg.OPAServerAuthorization != nil {
			router.Use(authorization.Middleware(cfg.OPAServerAuthorization))
		}
	}

	// Add static files
	router.Use(static.Serve("/", static.LocalFile("static/", true)))

	// Add graphql endpoints
	router.POST("/api/graphql", graphqlHandler(svr.busiServices))
	router.GET("/api/graphql", playgroundHandler())

	return router, nil
}

func (svr *Server) Listen() error {
	svr.logger.Infof("Server listening on %s", svr.server.Addr)
	err := svr.server.ListenAndServe()

	return err
}

// Defining the Graphql handler
func graphqlHandler(busiServices *business.Services) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graphql.Resolver{BusiServices: busiServices},
	}))
	h.Use(apollotracing.Tracer{})
	h.Use(gqlopentracing.Tracer{})

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
