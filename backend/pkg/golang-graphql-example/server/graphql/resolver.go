package graphql

import "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/model"

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos []*model.Todo
}
