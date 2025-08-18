package databasehelpers

import (
	"context"

	"gorm.io/gorm"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/common"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/pagination"
)

type GormOpt = func(ctx context.Context, gdb *gorm.DB) (*gorm.DB, error)

func WithFilterGormOpt(filter any) GormOpt {
	return func(_ context.Context, gdb *gorm.DB) (*gorm.DB, error) {
		return common.ManageFilter(filter, gdb)
	}
}

func WithSortOrderGormOpt(sorts any) GormOpt {
	return func(_ context.Context, gdb *gorm.DB) (*gorm.DB, error) {
		return common.ManageSortOrder(sorts, gdb)
	}
}

func WithProjectionGormOpt(projection any) GormOpt {
	return func(_ context.Context, gdb *gorm.DB) (*gorm.DB, error) {
		return common.ManageProjection(projection, gdb)
	}
}

func WithPaginationGormOpt(page *pagination.PageInput) GormOpt {
	return func(_ context.Context, gdb *gorm.DB) (*gorm.DB, error) {
		return gdb.Offset(page.Skip).Limit(page.Limit), nil
	}
}
