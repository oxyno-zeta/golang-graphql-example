package dev

import "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"

//go:generate go run github.com/oxyno-zeta/golang-graphql-example/tools/generator/modeltagsgen github.com/oxyno-zeta/golang-graphql-example/tools/generator/modeltagsgen/gen/dev Dev1
type Dev1 struct {
	database.Base
	Field1 string `gorm:"-"             json:"field1"`
	Field2 string `gorm:"column:fake2"  json:"-"`
	Field3 string `gorm:"column:fake3"  json:"field3"`
	Field4 string `gorm:"column:field4" json:"field4"`
}

//go:generate go run github.com/oxyno-zeta/golang-graphql-example/tools/generator/modeltagsgen github.com/oxyno-zeta/golang-graphql-example/tools/generator/modeltagsgen/gen/dev Object1
type Object1 struct {
	database.Base
	Name string
}

//go:generate go run github.com/oxyno-zeta/golang-graphql-example/tools/generator/modeltagsgen github.com/oxyno-zeta/golang-graphql-example/tools/generator/modeltagsgen/gen/dev Object2
type Object2 struct {
	database.Base
	Name string `json:"-"`
}

//go:generate go run github.com/oxyno-zeta/golang-graphql-example/tools/generator/modeltagsgen github.com/oxyno-zeta/golang-graphql-example/tools/generator/modeltagsgen/gen/dev Object3 --disable-gorm
type Object3 struct {
	database.Base
	Name string `gorm:"column:full__name" json:"name"`
}

//go:generate go run github.com/oxyno-zeta/golang-graphql-example/tools/generator/modeltagsgen github.com/oxyno-zeta/golang-graphql-example/tools/generator/modeltagsgen/gen/dev Object4
type Object4 struct {
	database.Base
	Name string `gorm:"-" json:"name"`
}

//go:generate go run github.com/oxyno-zeta/golang-graphql-example/tools/generator/modeltagsgen github.com/oxyno-zeta/golang-graphql-example/tools/generator/modeltagsgen/gen/dev Object5
type Object5 struct {
	database.Base
	Name string `json:"name,omitempty"`
}
