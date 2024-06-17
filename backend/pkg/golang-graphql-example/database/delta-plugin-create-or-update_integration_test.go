//go:build integration

package database_test

import (
	"context"
	"time"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/deltaplugin"
	databasehelpers "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/helpers"
	"github.com/stretchr/testify/assert"
)

func (suite *DeltaPluginTestSuite) TestCreateOrUpdateWithNotFoundItem() {
	ctx := context.TODO()
	expectedV := &deltaplugin.Delta{
		Table:     "peoples",
		Action:    deltaplugin.CREATE,
		Result:    &People{Base: database.Base{ID: "init-fake-id", CreatedAt: suite.now, UpdatedAt: suite.now}, Name: "fake"},
		EventDate: deltaplugin.NanoDateTime(suite.now),
	}

	_, err := databasehelpers.CreateOrUpdate(
		ctx,
		&People{Name: "fake"},
		suite.db,
	)
	suite.NoError(err)

	suite.EventuallyWithT(func(collect *assert.CollectT) {
		d := <-suite.deltaNotificationChan
		suite.Equal(expectedV, d)
	}, EventuallyWaitFor, EventuallyTick)
}

func (suite *DeltaPluginTestSuite) TestCreateOrUpdateWithNotFoundItemAndFromID() {
	ctx := context.TODO()
	expectedV := &deltaplugin.Delta{
		Table:     "peoples",
		Action:    deltaplugin.CREATE,
		Result:    &People{Base: database.Base{ID: "fake-id", CreatedAt: suite.now, UpdatedAt: suite.now}, Name: "fake"},
		EventDate: deltaplugin.NanoDateTime(suite.now),
	}

	_, err := databasehelpers.CreateOrUpdate(
		ctx,
		&People{Base: database.Base{ID: "fake-id"}, Name: "fake"},
		suite.db,
	)
	suite.NoError(err)

	suite.EventuallyWithT(func(collect *assert.CollectT) {
		d := <-suite.deltaNotificationChan
		suite.Equal(expectedV, d)
	}, EventuallyWaitFor, EventuallyTick)
}

func (suite *DeltaPluginTestSuite) TestCreateOrUpdateWithFoundItem() {
	// Save item
	suite.setupGenericDataset([]interface{}{
		&People{
			Base:       database.Base{ID: "fake-id", CreatedAt: suite.now.Add(-time.Second), UpdatedAt: suite.now.Add(-time.Second)},
			Name:       "name",
			FullName:   "full-name",
			LoggedOnce: true,
		},
	})
	suite.cleanDeltaNotificationChannel()

	ctx := context.TODO()
	expectedV := &deltaplugin.Delta{
		Table:     "peoples",
		Action:    deltaplugin.UPDATE,
		Result:    &People{Base: database.Base{ID: "fake-id", UpdatedAt: suite.now}, Name: "fake"},
		EventDate: deltaplugin.NanoDateTime(suite.now),
	}

	_, err := databasehelpers.CreateOrUpdate(
		ctx,
		&People{Base: database.Base{ID: "fake-id"}, Name: "fake"},
		suite.db,
	)
	suite.NoError(err)

	suite.EventuallyWithT(func(collect *assert.CollectT) {
		d := <-suite.deltaNotificationChan
		suite.Equal(expectedV, d)
	}, EventuallyWaitFor, EventuallyTick)
}

func (suite *DeltaPluginTestSuite) TestCreateOrUpdateWithAnError() {
	ctx := context.TODO()

	_, err := databasehelpers.CreateOrUpdate(
		ctx,
		&PeopleNotDBCreated{Base: database.Base{ID: "fake-id"}},
		suite.db,
	)
	suite.Error(err)

	// Check channel length
	if len(suite.deltaNotificationChan) != 0 {
		suite.Fail("delta received from channel")
	}
}
