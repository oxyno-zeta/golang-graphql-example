package models

import "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"

type Todo struct {
	database.Base
	Text string `gorm:"type:varchar(2000)"`
	Done bool
}
