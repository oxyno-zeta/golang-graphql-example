package dataloaders

import (
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business"
	todosdataloaders "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/dataloaders/todos"
)

type Dataloaders struct {
	Todos *todosdataloaders.TodosDataloaders
}

func newDataloaders(busiSvr *business.Services) *Dataloaders {
	return &Dataloaders{
		Todos: todosdataloaders.New(busiSvr),
	}
}
