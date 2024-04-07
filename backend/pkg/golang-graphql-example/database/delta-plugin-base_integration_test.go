//go:build integration

package database_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	cmocks "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config/mocks"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/deltaplugin"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/tracing"
	"github.com/stretchr/testify/suite"
	"go.uber.org/goleak"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	EventuallyTick    = 200 * time.Millisecond
	EventuallyWaitFor = 2 * time.Second
)

type DeltaPluginTestSuite struct {
	suite.Suite

	db                    database.DB
	deltaNotificationChan chan *deltaplugin.Delta
	now                   time.Time
}

type People struct {
	database.Base
	Name       string
	FullName   string
	LoggedOnce bool
}

// Override before create to avoid uuid generation
func (base *People) BeforeCreate(_ *gorm.DB) error {
	base.Base.ID = "init-fake-id"
	return nil
}

type PeopleNotDBCreated struct {
	database.Base
	Name       string
	FullName   string
	LoggedOnce bool
}

// this function executes before the test suite begins execution
func (suite *DeltaPluginTestSuite) SetupSuite() {
	fmt.Println("SetupSuite phase")

	// Create mock configuration
	ctrl := gomock.NewController(suite.T())
	cfgManagerMock := cmocks.NewMockManager(ctrl)
	cfgManagerMock.EXPECT().GetConfig().AnyTimes().Return(integrationTestsCfg)

	// Create logger
	logger := log.NewLogger()
	err := logger.Configure(integrationTestsCfg.Log.Level, integrationTestsCfg.Log.Format, "")
	suite.NoError(err)

	// Create tracing service
	tracingSvc := tracing.New(cfgManagerMock, logger)
	err = tracingSvc.InitializeAndReload()
	suite.NoError(err)

	now := time.Now()
	deltaNotificationChan := make(chan *deltaplugin.Delta, 100)
	// Create db service
	db := database.NewDatabase("main", cfgManagerMock, logger, metricsCtx, tracingSvc, deltaNotificationChan)
	// Connect
	err = db.Connect()
	suite.NoError(err)

	// Override now function
	db.GetGormDB().NowFunc = func() time.Time {
		return now
	}

	// Migrate
	err = db.GetGormDB().AutoMigrate(&People{})
	suite.NoError(err)

	// Save data
	suite.db = db
	suite.deltaNotificationChan = deltaNotificationChan
	suite.now = now
}

// this function executes after all tests executed
func (suite *DeltaPluginTestSuite) TearDownSuite() {
	fmt.Println("TearDownSuite phase")
}

func (suite *DeltaPluginTestSuite) AfterTest(suiteName, testName string) {
	fmt.Println("AfterTest phase")
	suite.cleanDB()
	suite.cleanDeltaNotificationChannel()
}

func (suite *DeltaPluginTestSuite) cleanDeltaNotificationChannel() {
	for len(suite.deltaNotificationChan) > 0 {
		<-suite.deltaNotificationChan
	}
}

func (suite *DeltaPluginTestSuite) cleanDB() {
	modelList := []interface{}{
		&People{},
	}

	for _, item := range modelList {
		sch, err := schema.Parse(item, &sync.Map{}, suite.db.GetGormDB().NamingStrategy)
		suite.NoError(err)

		gdb := suite.db.GetGormDB().Exec(fmt.Sprintf("TRUNCATE %s;", sch.Table))
		suite.NoError(gdb.Error)
	}
}

func (suite *DeltaPluginTestSuite) setupGenericDataset(dataset []interface{}) {
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
	suite.Run(t, new(DeltaPluginTestSuite))
}
