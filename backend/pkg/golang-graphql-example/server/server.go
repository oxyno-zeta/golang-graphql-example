package server

import (
	"context"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/pkg/errors"

	gqlgraphql "github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/authentication"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/authorization"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
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
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/middlewares"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/signalhandler"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/tracing"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const GraphqlComplexityLimit = 2000

type Server struct {
	logger            log.Logger
	cfgManager        config.Manager
	metricsCl         metrics.Client
	tracingSvc        tracing.Service
	busiServices      *business.Services
	authenticationSvc authentication.Client
	authorizationSvc  authorization.Service
	signalHandlerSvc  signalhandler.Client
	server            *http.Server
}

func NewServer(
	logger log.Logger, cfgManager config.Manager, metricsCl metrics.Client,
	tracingSvc tracing.Service, busiServices *business.Services,
	authenticationSvc authentication.Client, authoSvc authorization.Service,
	signalHandlerSvc signalhandler.Client,
) *Server {
	return &Server{
		logger:            logger,
		cfgManager:        cfgManager,
		metricsCl:         metricsCl,
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
		ReadHeaderTimeout: 10 * time.Second, //nolint: gomnd // Ignored to see it clearly
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
		// Create not found error
		err := cerrors.NewNotFoundError("404 not found")

		// Answer
		utils.AnswerWithError(c, err)
	})
	// Add middlewares
	router.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithDecompressFn(gzip.DefaultDecompressHandle)))
	router.Use(gin.Recovery())
	router.Use(svr.signalHandlerSvc.ActiveRequestCounterMiddleware())
	router.Use(middlewares.RequestID(svr.logger))
	router.Use(svr.tracingSvc.HTTPMiddleware(middlewares.GetRequestIDFromContext))
	router.Use(log.Middleware(svr.logger, middlewares.GetRequestIDFromGin, tracing.GetTraceIDFromContext))
	router.Use(svr.metricsCl.Instrument("business", true))
	// Add helmet for security
	router.Use(helmet.Default())
	// Add cors if configured
	err := manageCORS(router, cfg.Server)
	// Check error
	if err != nil {
		return nil, err
	}

	// Create api prefix path regexp
	apiReg, err := regexp.Compile("^/api")
	// Check error
	if err != nil {
		return nil, errors.WithStack(err)
	}

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
	router.GET("/api/graphql", gin.WrapH(graphiqlHandler("GraphQL", "/api/graphql")))

	// Add gin html files for answer
	router.LoadHTMLGlob("static/*.html")
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
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graphql.Resolver{BusiServices: busiServices},
		Complexity: generated.ComplexityRoot{
			Mutation: struct {
				CloseTodo  func(childComplexity int, todoID string) int
				CreateTodo func(childComplexity int, input model.NewTodo) int
				UpdateTodo func(childComplexity int, input *model.UpdateTodo) int
			}{
				CloseTodo: func(childComplexity int, todoID string) int {
					return gutils.CalculateMutationComplexity(childComplexity)
				},
				CreateTodo: func(childComplexity int, input model.NewTodo) int {
					return gutils.CalculateMutationComplexity(childComplexity)
				},
				UpdateTodo: func(childComplexity int, input *model.UpdateTodo) int {
					return gutils.CalculateMutationComplexity(childComplexity)
				},
			},
			Query: struct {
				Todo  func(childComplexity int, id string) int
				Todos func(childComplexity int, after *string, before *string, first *int, last *int, sort *models.SortOrder, filter *models.Filter) int
			}{
				Todo: func(childComplexity int, id string) int {
					return gutils.CalculateQuerySimpleStructComplexity(childComplexity)
				},
				Todos: func(childComplexity int, after, before *string, first, last *int, sort *models.SortOrder, filter *models.Filter) int {
					return gutils.CalculateQueryConnectionComplexity(childComplexity, after, before, first, last)
				},
			},
		},
	}))
	h.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		// Get logger
		logger := log.GetLoggerFromContext(ctx)
		// Log error
		logger.Error(err)
		// Initialize potential generic error
		var err2 *cerrors.GenericError
		// Get generic error if available
		if errors.As(err, &err2) {
			// Return graphql error
			return &gqlerror.Error{
				Path:       gqlgraphql.GetPath(ctx),
				Extensions: err2.Extensions(),
				Message:    err2.PublicMessage(),
			}
		}
		// Return
		return gqlgraphql.DefaultErrorPresenter(ctx, cerrors.NewInternalServerError("internal server error"))
	})
	h.Use(svr.tracingSvc.GraphqlMiddleware())
	h.Use(svr.metricsCl.GraphqlMiddleware())
	h.Use(extension.FixedComplexityLimit(GraphqlComplexityLimit))

	return gin.WrapH(h)
}
