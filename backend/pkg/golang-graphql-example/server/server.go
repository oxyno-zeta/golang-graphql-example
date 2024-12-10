package server

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"emperror.dev/errors"

	gqlgraphql "github.com/99designs/gqlgen/graphql"
	gqlerrorcode "github.com/99designs/gqlgen/graphql/errcode"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	gqlplayground "github.com/99designs/gqlgen/graphql/playground"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/authentication"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/authorization"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business"
	correlationid "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/common/correlation-id"
	cerrors "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/common/errors"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/common/utils"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/metrics"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/dataloaders"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/generated"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/model"
	gutils "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/utils"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/signalhandler"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/tracing"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const GraphqlComplexityLimit = 2000

var StaticFiles = "static/*.html"

type Server struct {
	logger            log.Logger
	cfgManager        config.Manager
	metricsSvc        metrics.Service
	tracingSvc        tracing.Service
	busiServices      *business.Services
	authenticationSvc authentication.Service
	authorizationSvc  authorization.Service
	signalHandlerSvc  signalhandler.Service
	server            *http.Server
}

func NewServer(
	logger log.Logger, cfgManager config.Manager, metricsSvc metrics.Service,
	tracingSvc tracing.Service, busiServices *business.Services,
	authenticationSvc authentication.Service, authoSvc authorization.Service,
	signalHandlerSvc signalhandler.Service,
) *Server {
	return &Server{
		logger:            logger,
		cfgManager:        cfgManager,
		metricsSvc:        metricsSvc,
		tracingSvc:        tracingSvc,
		busiServices:      busiServices,
		authenticationSvc: authenticationSvc,
		authorizationSvc:  authoSvc,
		signalHandlerSvc:  signalHandlerSvc,
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
		// Timeout to 10 second to limit the time to read the request headers
		// This is done to patch server against Slowloris Attack
		ReadHeaderTimeout: 10 * time.Second, //nolint: mnd // Ignored to see it clearly
	}

	// Prepare for configuration onChange
	svr.cfgManager.AddOnChangeHook(&config.HookDefinition{
		Hook: func() error {
			// Generate router
			r, err2 := svr.generateRouter()
			// Check error
			if err2 != nil {
				return err2
			}
			// Change server handler
			server.Handler = r

			svr.logger.Info("Server handler reloaded")

			return nil
		},
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
		// Create not found error
		err := cerrors.NewNotFoundError("404 not found")

		// Answer
		utils.AnswerWithError(c, err)
	})
	// Add middlewares

	// Check if compress if enabled
	if cfg.Server.Compress != nil && cfg.Server.Compress.Enabled {
		router.Use(gzip.Gzip(gzip.DefaultCompression))
	}
	// Recover any panic in router
	// Force the first parameter to nil to avoid any log
	router.Use(gin.CustomRecoveryWithWriter(nil, func(c *gin.Context, errI any) {
		// Get logger
		logger := log.GetLoggerFromGin(c)

		var err error
		// Cast err as error
		err, ok := errI.(error)
		// Check if cast was ok to wrap it in basic error
		// If not, stringify it
		if ok {
			err = cerrors.NewInternalServerErrorWithError(err)
		} else {
			err = cerrors.NewInternalServerError(fmt.Sprintf("%+v", errI))
		}
		// Log
		logger.Error(err)
		// Answer
		utils.AnswerWithError(c, err)
	}))
	router.Use(svr.signalHandlerSvc.ActiveRequestCounterMiddleware([]string{}))
	router.Use(correlationid.HTTPMiddleware(svr.logger))
	router.Use(svr.tracingSvc.HTTPMiddlewareList(correlationid.GetFromContext)...)
	router.Use(log.Middleware(svr.logger, correlationid.GetFromGin, tracing.GetTraceIDFromContext))
	router.Use(svr.metricsSvc.Instrument("business", true))
	// Add helmet for security
	router.Use(helmet.Default())
	// Add cors if configured
	err := manageCORS(router, cfg.Server)
	// Check error
	if err != nil {
		return nil, err
	}

	// Create api prefix path regexp
	apiReg := regexp.MustCompile("^/api")

	// Add authentication middleware if configuration exists
	if cfg.OIDCAuthentication != nil {
		// Add endpoints
		err := svr.authenticationSvc.OIDCEndpoints(router)
		// Check error
		if err != nil {
			return nil, errors.WithStack(err)
		}

		// Add authentication middleware
		router.Use(svr.authenticationSvc.Middleware([]*regexp.Regexp{apiReg}))
	}

	// Integrate graphql dataloaders
	router.Use(dataloaders.Middleware(svr.busiServices))
	// Add graphql endpoints
	router.POST("/api/graphql", svr.graphqlHandler(svr.busiServices))
	router.GET("/api/graphql", gin.WrapH(gqlplayground.Handler("GraphQL", "/api/graphql")))

	// Add gin html files for answer
	router.LoadHTMLGlob(StaticFiles)
	// Add static files
	router.Use(static.Serve("/", static.LocalFile("static/", true)))
	// Add specialized support for SPA based UI
	router.Use(func(c *gin.Context) {
		// Check if patch isn't matching api based prefix
		if !apiReg.MatchString(c.Request.RequestURI) {
			// Answer with index.html to all possible path
			c.HTML(http.StatusOK, "index.html", nil)
			c.Abort()
		}
	})

	return router, nil
}

func (svr *Server) Listen() error {
	svr.logger.Infof("Server listening on %s", svr.server.Addr)
	err := svr.server.ListenAndServe()
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	// Default
	return nil
}

// Defining the Graphql handler.
func (svr *Server) graphqlHandler(busiServices *business.Services) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.New(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graphql.Resolver{BusiServices: busiServices},
		Complexity: generated.ComplexityRoot{
			Mutation: struct {
				CloseTodo  func(childComplexity int, todoID string) int
				CreateTodo func(childComplexity int, input model.NewTodo) int
				UpdateTodo func(childComplexity int, input *model.UpdateTodo) int
			}{
				CloseTodo: func(childComplexity int, _ string) int {
					return gutils.CalculateMutationComplexity(childComplexity)
				},
				CreateTodo: func(childComplexity int, _ model.NewTodo) int {
					return gutils.CalculateMutationComplexity(childComplexity)
				},
				UpdateTodo: func(childComplexity int, _ *model.UpdateTodo) int {
					return gutils.CalculateMutationComplexity(childComplexity)
				},
			},
		},
	}))

	h.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second, //nolint:mnd
	})
	h.AddTransport(transport.Options{})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})
	h.AddTransport(transport.MultipartForm{})

	h.SetQueryCache(lru.New[*ast.QueryDocument](1000)) //nolint:mnd

	h.Use(extension.Introspection{})
	h.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100), //nolint:mnd
	})
	h.Use(svr.tracingSvc.GraphqlMiddleware())
	h.Use(svr.metricsSvc.GraphqlMiddleware())
	h.Use(extension.FixedComplexityLimit(GraphqlComplexityLimit))

	h.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		// Get logger
		logger := log.GetLoggerFromContext(ctx)
		// Initialize potential generic error
		var err2 *cerrors.GenericError
		// Initialize gqlparser error
		var err3 *gqlerror.Error
		// Get generic error if available
		if errors.As(err, &err2) {
			// Log error generic error
			logger.Error(err2)
			// Return graphql error
			return &gqlerror.Error{
				Path:       gqlgraphql.GetPath(ctx),
				Extensions: err2.Extensions(),
				Message:    err2.PublicMessage(),
			}
		} else if errors.As(err, &err3) && err3.Extensions != nil && err3.Extensions["code"] == gqlerrorcode.ValidationFailed {
			// This case is for GraphQL validation failed.
			// This can arrive when a user send a request with a unknown field or if types are wrong.
			// Log error generic error
			logger.WithError(errors.WithStack(err3)).Warn(err3)
			// Return graphql error
			return err3
		}

		// Not a managed error. Manage it as internal server error
		// Wrap it as an internal server error
		err4 := cerrors.NewInternalServerErrorWithError(err)
		// Log
		logger.Error(err4)
		// Return new built error
		return &gqlerror.Error{
			Path:       gqlgraphql.GetPath(ctx),
			Extensions: err4.Extensions(),
			Message:    err4.PublicMessage(),
			Locations:  err3.Locations,
			Rule:       err3.Rule,
		}
	})
	h.SetRecoverFunc(func(ctx context.Context, errI interface{}) (userMessage error) {
		// Get logger
		logger := log.GetLoggerFromContext(ctx)

		var err error
		// Cast err as error
		err, ok := errI.(error)
		// Check if cast was ok to wrap it in basic error
		// If not, stringify it
		if ok {
			err = cerrors.NewInternalServerErrorWithError(err)
		} else {
			err = cerrors.NewInternalServerError(fmt.Sprintf("%+v", errI))
		}
		// Log
		logger.Error(err)

		return err
	})

	return gin.WrapH(h)
}
