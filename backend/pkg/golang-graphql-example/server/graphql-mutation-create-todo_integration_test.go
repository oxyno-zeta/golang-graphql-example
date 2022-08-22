//go:build integration

package server

import (
	"context"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/mappers"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/utils"
)

func (suite *GraphQLTestSuite) TestMutationCreateTodo() {
	var m struct {
		Todo struct {
			ID   string
			Text string
			Done bool
		} `graphql:"createTodo(input: $input)"`
	}
	type NewTodo struct {
		Text string `json:"text"`
	}
	variables := map[string]interface{}{
		"input": NewTodo{Text: "Fake !"},
	}

	err := suite.graphqlClient.Mutate(context.TODO(), &m, variables)

	suite.NoError(err)
	suite.Equal(m.Todo.Done, false)
	suite.Equal(m.Todo.Text, "Fake !")
	suite.NotEmpty(m.Todo.ID)
	uuid, err := utils.FromIDRelay(string(m.Todo.ID), mappers.TodoIDPrefix)
	suite.NoError(err)
	suite.NotEmpty(uuid)

	var res models.Todo
	dbRes := suite.db.GetGormDB().Where("id", uuid).First(&res)
	suite.NoError(dbRes.Error)
	suite.Equal(res.ID, uuid)
	suite.Equal(res.Text, m.Todo.Text)
	suite.NotEmpty(res.CreatedAt)
	suite.NotEmpty(res.UpdatedAt)
}
