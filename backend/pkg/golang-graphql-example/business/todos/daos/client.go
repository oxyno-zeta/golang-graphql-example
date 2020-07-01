package daos

import "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"

type Dao interface{}

func NewDao(db database.DB) Dao {
	return &dao{db: db}
}
