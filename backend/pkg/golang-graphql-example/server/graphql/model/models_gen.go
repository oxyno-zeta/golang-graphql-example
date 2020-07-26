// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
)

type NewTodo struct {
	Text string `json:"text"`
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor"`
	EndCursor       *string `json:"endCursor"`
}

type TodoConnection struct {
	Edges    []*TodoEdge `json:"edges"`
	PageInfo *PageInfo   `json:"pageInfo"`
}

type TodoEdge struct {
	Cursor string       `json:"cursor"`
	Node   *models.Todo `json:"node"`
}

type UpdateTodo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}
