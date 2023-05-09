package graphql

import "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	BusiServices *business.Services
}
