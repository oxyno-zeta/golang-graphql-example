//go:build integration

package database_test

import (
	"context"

	databasehelpers "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/helpers"
)

func (suite *TransactionalOutboxPluginTestSuite) TestCreateOrUpdateWithNotFoundItem() {
	ctx := context.TODO()
	err := suite.db.ExecuteTransaction(ctx, func(ctx context.Context) error {
		_, err := databasehelpers.CreateOrUpdate(
			ctx,
			&People{Name: "fake"},
			suite.db,
		)
		return err
	})
	suite.NoError(err)
}
