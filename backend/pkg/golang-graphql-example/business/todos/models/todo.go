package models

import "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"

//go:generate go run ../../../../../tools/generator/modeltags github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models Todo
type Todo struct {
	database.Base
	Text string `gorm:"type:varchar(2000)"`
	Done bool
}
