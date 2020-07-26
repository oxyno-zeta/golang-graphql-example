package mappers

import (
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/pagination"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/model"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/utils"
)

const TodoIDPrefix = "todos"

func MapTodoConnection(allTodos []*models.Todo, pageOut *pagination.PageOutput) *model.TodoConnection {
	var res model.TodoConnection

	// Loop over all todos
	for i := 0; i < len(allTodos); i++ {
		// Get todo
		todo := allTodos[i]
		// Create new edge
		res.Edges = append(res.Edges, &model.TodoEdge{
			Cursor: utils.GetPaginateCursor(i, pageOut.Skip),
			Node:   todo,
		})
	}

	// Create page info cursors
	startCursor := ""
	endCursor := ""

	edgesLen := len(res.Edges)

	// Check if edges exist in order to map start and end cursor
	if edgesLen != 0 {
		startCursor = res.Edges[0].Cursor
		endCursor = res.Edges[edgesLen-1].Cursor
	}

	// Create page info object
	res.PageInfo = utils.GetPageInfo(startCursor, endCursor, pageOut)

	return &res
}
