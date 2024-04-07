//go:build integration

package database_test

import (
	"context"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/deltaplugin"
	databasehelpers "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/helpers"
	"github.com/stretchr/testify/assert"
)

func (suite *DeltaPluginTestSuite) TestPatchUpdateWithNotFoundItem() {
	ctx := context.TODO()

	_, err := databasehelpers.PatchUpdate(
		ctx,
		&People{Base: database.Base{ID: "fake-id"}},
		map[string]interface{}{"name": "fake"},
		suite.db,
	)
	suite.NoError(err)

	// Check channel length
	if len(suite.deltaNotificationChan) != 0 {
		suite.Fail("delta received from channel")
	}
}

func (suite *DeltaPluginTestSuite) TestPatchUpdateWithFoundItem() {
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
		Table:  "peoples",
		Action: deltaplugin.PATCH,
		Result: &People{Base: database.Base{ID: "fake-id", UpdatedAt: suite.now}, Name: "fake"},
		Patch: map[string]interface{}{
			"name": "fake",
		},
	}

	_, err := databasehelpers.PatchUpdate(
		ctx,
		&People{Base: database.Base{ID: "fake-id"}},
		map[string]interface{}{"name": "fake"},
		suite.db,
	)
	suite.NoError(err)

	suite.EventuallyWithT(func(collect *assert.CollectT) {
		d := <-suite.deltaNotificationChan
		suite.Equal(expectedV, d)
	}, EventuallyWaitFor, EventuallyTick)
}

func (suite *DeltaPluginTestSuite) TestPatchUpdateWithAnError() {
	ctx := context.TODO()

	_, err := databasehelpers.PatchUpdate(
		ctx,
		&PeopleNotDBCreated{Base: database.Base{ID: "fake-id"}},
		map[string]interface{}{"name": "fake"},
		suite.db,
	)
	suite.Error(err)

	// Check channel length
	if len(suite.deltaNotificationChan) != 0 {
		suite.Fail("delta received from channel")
	}
}
