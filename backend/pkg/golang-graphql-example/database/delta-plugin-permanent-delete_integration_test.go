//go:build integration

package database_test

import (
	"context"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/deltaplugin"
	databasehelpers "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/helpers"
	"github.com/stretchr/testify/assert"
)

func (suite *DeltaPluginTestSuite) TestPermanentDeleteWithNotFoundItem() {
	ctx := context.TODO()

	_, err := databasehelpers.PermanentDelete(ctx, &People{Base: database.Base{ID: "fake-id"}}, suite.db)
	suite.NoError(err)

	// Check channel length
	if len(suite.deltaNotificationChan) != 0 {
		suite.Fail("delta received from channel")
	}
}

func (suite *DeltaPluginTestSuite) TestPermanentDeleteWithFoundItem() {
	// Save item
	suite.setupGenericDataset([]interface{}{
		&People{
			Base:       database.Base{ID: "fake-id"},
			Name:       "name",
			FullName:   "full-name",
			LoggedOnce: true,
		},
	})
	suite.cleanDeltaNotificationChannel()

	ctx := context.TODO()
	expectedV := &deltaplugin.Delta{
		Table:     "peoples",
		Action:    deltaplugin.DELETE,
		Result:    &People{Base: database.Base{ID: "fake-id"}},
		EventDate: deltaplugin.NanoDateTime(suite.now),
	}

	_, err := databasehelpers.PermanentDelete(ctx, &People{Base: database.Base{ID: "fake-id"}}, suite.db)
	suite.NoError(err)

	suite.EventuallyWithT(func(collect *assert.CollectT) {
		d := <-suite.deltaNotificationChan
		suite.Equal(expectedV, d)
	}, EventuallyWaitFor, EventuallyTick)
}

func (suite *DeltaPluginTestSuite) TestPermanentDeleteWithAnError() {
	ctx := context.TODO()

	_, err := databasehelpers.PermanentDelete(ctx, &PeopleNotDBCreated{Base: database.Base{ID: "fake-id"}}, suite.db)
	suite.Error(err)

	// Check channel length
	if len(suite.deltaNotificationChan) != 0 {
		suite.Fail("delta received from channel")
	}
}
