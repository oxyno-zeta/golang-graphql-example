//go:build integration

package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hasura/go-graphql-client"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/authentication"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/authorization"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
	cmocks "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config/mocks"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	lockdistributor "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/lockdistributor/sql"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/signalhandler"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/tracing"
	"github.com/stretchr/testify/suite"
	"go.uber.org/goleak"
	"gorm.io/gorm/schema"
)

type GraphQLTestSuite struct {
	suite.Suite

	testServer    *httptest.Server
	graphqlClient *graphql.Client
	db            database.DB
	busiServices  *business.Services
}

// this function executes before the test suite begins execution
func (suite *GraphQLTestSuite) SetupSuite() {
	fmt.Println("SetupSuite phase")
	// Override static files
	StaticFiles = "../../../static/*.html"

	// Create mock configuration
	ctrl := gomock.NewController(suite.T())
	cfgManagerMock := cmocks.NewMockManager(ctrl)
	cfgManagerMock.EXPECT().GetConfig().AnyTimes().Return(integrationTestsCfg)

	// Create logger
	logger := log.NewLogger()
	err := logger.Configure(integrationTestsCfg.Log.Level, integrationTestsCfg.Log.Format, "")
	suite.NoError(err)

	// Create tracing service
	tracingSvc, err := tracing.New(cfgManagerMock, logger)
	suite.NoError(err)
	// Create signalhandler service
	signalHandlerSvc := signalhandler.NewService(logger, false, []os.Signal{syscall.SIGTERM, syscall.SIGINT})
	// Create db service
	db := database.NewDatabase("main", cfgManagerMock, logger, metricsCtx, tracingSvc)
	// Connect
	err = db.Connect()
	suite.NoError(err)
	// Create lockdistributor
	ld := lockdistributor.NewService(cfgManagerMock, db)
	err = ld.InitializeAndReload(logger)
	suite.NoError(err)
	// Create authentication service
	authCl := authentication.NewService(cfgManagerMock)
	// Create authorization service
	authoCl := authorization.NewService(cfgManagerMock)
	// Create services
	bSvc := business.NewServices(logger, db, authoCl)
	// Migrate
	err = bSvc.MigrateDB()
	suite.NoError(err)
	// Create server
	s := NewServer(logger, cfgManagerMock, metricsCtx, tracingSvc, bSvc, authCl, authoCl, signalHandlerSvc)
	// Create handler
	got, err := s.generateRouter()
	suite.NoError(err)

	// Create server
	svr := httptest.NewUnstartedServer(got)

	// Start server
	svr.Start()

	// Create graphql client
	gcl := graphql.NewClient(svr.URL+"/api/graphql", svr.Client())
	// Add request modifier to add authentication
	gcl = gcl.WithRequestModifier(func(req *http.Request) {
		data := url.Values{}
		data.Set("username", "user")
		data.Set("password", "password")
		data.Set("client_id", integrationTestsCfg.OIDCAuthentication.ClientID)
		data.Set("grant_type", "password")
		data.Set("scope", "openid profile")

		authentUrlStr := integrationTestsCfg.OIDCAuthentication.IssuerURL + "/protocol/openid-connect/token"

		clientAuth := &http.Client{}
		r, err := http.NewRequest("POST", authentUrlStr, strings.NewReader(data.Encode())) // URL-encoded payload
		// Check err
		suite.NoError(err)

		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

		resp, err := clientAuth.Do(r)
		// Check err
		suite.NoError(err)

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			suite.NoError(err)
			return
		}
		body := string(bodyBytes)

		// Check response
		if resp.StatusCode != 200 {
			suite.Fail(fmt.Sprintf("%d - %s", resp.StatusCode, body))
			return
		}

		type tokensResponseBody struct {
			IDToken string `json:"id_token"`
		}

		var to tokensResponseBody
		// Parse token
		err = json.Unmarshal(bodyBytes, &to)
		suite.NoError(err)

		// Add header to request
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", to.IDToken))
	})

	// Save data
	suite.testServer = svr
	suite.graphqlClient = gcl
	suite.db = db
	suite.busiServices = bSvc
}

// this function executes after all tests executed
func (suite *GraphQLTestSuite) TearDownSuite() {
	fmt.Println("TearDownSuite phase")
	// Close server if it exists
	if suite.testServer != nil {
		suite.testServer.Close()
	}
}

func (suite *GraphQLTestSuite) AfterTest(suiteName, testName string) {
	fmt.Println("AfterTest phase")
	suite.cleanDB()
}

func (suite *GraphQLTestSuite) cleanDB() {
	modelList := []interface{}{
		&models.Todo{},
	}

	for _, item := range modelList {
		sch, err := schema.Parse(item, &sync.Map{}, suite.db.GetGormDB().NamingStrategy)
		suite.NoError(err)

		gdb := suite.db.GetGormDB().Exec(fmt.Sprintf("TRUNCATE %s;", sch.Table))
		suite.NoError(gdb.Error)
	}
}

func (suite *GraphQLTestSuite) setupGenericDataset(dataset []interface{}) {
	for _, it := range dataset {
		suite.NoError(suite.db.GetGormDB().Save(it).Error)
	}
}

func TestGraphQLTestSuite(t *testing.T) {
	// Verify there isn't any go routine leak
	defer goleak.VerifyNone(
		t,
		// Ignore database specific
		goleak.IgnoreTopFunction("database/sql.(*DB).connectionOpener"),
		// Ignore gorm prometheus plugin because it creates an infinite go routine
		goleak.IgnoreTopFunction("gorm.io/plugin/prometheus.(*Prometheus).Initialize.func1.1"),
	)
	suite.Run(t, new(GraphQLTestSuite))
}
