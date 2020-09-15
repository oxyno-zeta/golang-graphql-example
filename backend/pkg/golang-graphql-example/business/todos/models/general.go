package models

import "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/pagination"

type SortOrder struct {
	CreatedAt *pagination.SortOrderEnum `db_col:"created_at"`
	UpdatedAt *pagination.SortOrderEnum `db_col:"updated_at"`
	Text      *pagination.SortOrderEnum `db_col:"text"`
	Done      *pagination.SortOrderEnum `db_col:"done"`
}

type Filter struct {
	Text *pagination.GenericFilter
}
