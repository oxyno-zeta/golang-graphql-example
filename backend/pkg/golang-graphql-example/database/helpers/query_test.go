//go:build unit

package databasehelpers

import (
	"context"
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/common"
	dbmocks "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/mocks"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/pagination"
)

func TestFind(t *testing.T) {
	now := time.Now()

	type People struct {
		database.Base
		Name       string
		FullName   string `gorm:"column:full__name"`
		LoggedOnce bool
	}
	type Filter struct {
		Name       *common.GenericFilter `dbfield:"name"`
		FullName   *common.GenericFilter `dbfield:"full__name"`
		LoggedOnce *common.GenericFilter `dbfield:"logged_once"`
		AND        []*Filter
		OR         []*Filter
	}
	type SortOrder struct {
		Name     *common.SortOrderEnum `dbfield:"name"`
		FullName *common.SortOrderEnum `dbfield:"full__name"`
	}
	type Projection struct {
		Name       bool `dbfield:"name"`
		FullName   bool `dbfield:"full__name"`
		LoggedOnce bool `dbfield:"logged_once"`
	}
	type args struct {
		filter     *Filter
		sorts      []*SortOrder
		projection *Projection
	}
	tests := []struct {
		name             string
		args             args
		wantErr          bool
		errorString      string
		expectedSQLQuery string
		expectedSQLArgs  []driver.Value
	}{
		{
			name: "1 filter",
			args: args{
				filter: &Filter{
					OR: []*Filter{
						{Name: &common.GenericFilter{Eq: "fake"}},
						{LoggedOnce: &common.GenericFilter{Eq: true}},
					},
				},
			},
			expectedSQLQuery: `SELECT * FROM "peoples" WHERE name = $1 OR logged_once = $2 ORDER BY created_at DESC`,
			expectedSQLArgs:  []driver.Value{"fake", true},
		},
		{
			name: "1 sort order",
			args: args{
				sorts: []*SortOrder{{Name: &common.SortOrderEnumAsc}},
			},
			expectedSQLQuery: `SELECT * FROM "peoples" ORDER BY name ASC`,
			expectedSQLArgs:  []driver.Value{},
		},
		{
			name: "1 projection",
			args: args{
				projection: &Projection{Name: true},
			},
			expectedSQLQuery: `SELECT "name" FROM "peoples" ORDER BY created_at DESC`,
			expectedSQLArgs:  []driver.Value{},
		},
		{
			name: "1 filter and 1 sort and 1 projection",
			args: args{
				filter: &Filter{
					OR: []*Filter{
						{Name: &common.GenericFilter{Eq: "fake"}},
						{LoggedOnce: &common.GenericFilter{Eq: true}},
					},
				},
				sorts:      []*SortOrder{{Name: &common.SortOrderEnumAsc}},
				projection: &Projection{Name: true},
			},
			expectedSQLQuery: `SELECT "name" FROM "peoples" WHERE name = $1 OR logged_once = $2 ORDER BY name ASC`,
			expectedSQLArgs:  []driver.Value{"fake", true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Error(err)

				return
			}
			defer sqlDB.Close()

			db, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{Logger: logger.Discard, NowFunc: func() time.Time {
				return now
			}})
			if err != nil {
				t.Error(err)

				return
			}

			ctrl := gomock.NewController(t)
			dbSvc := dbmocks.NewMockDB(ctrl)
			dbSvc.EXPECT().GetTransactionalOrDefaultGormDB(gomock.Any()).AnyTimes().Return(db)

			mock.ExpectQuery(tt.expectedSQLQuery).
				WithArgs(tt.expectedSQLArgs...).
				WillReturnRows(
					sqlmock.NewRows([]string{}),
				)

			ctx := context.TODO()
			res, err := Find(ctx, make([]*People, 0), dbSvc, tt.args.sorts, tt.args.filter, tt.args.projection)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if err != nil && err.Error() != tt.errorString {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.errorString)

				return
			}
			assert.Len(t, res, 0)
		})
	}
}

func TestFindWithPagination(t *testing.T) {
	now := time.Now()

	type People struct {
		database.Base
		Name       string
		FullName   string `gorm:"column:full__name"`
		LoggedOnce bool
	}
	type Filter struct {
		Name       *common.GenericFilter `dbfield:"name"`
		FullName   *common.GenericFilter `dbfield:"full__name"`
		LoggedOnce *common.GenericFilter `dbfield:"logged_once"`
		AND        []*Filter
		OR         []*Filter
	}
	type SortOrder struct {
		Name     *common.SortOrderEnum `dbfield:"name"`
		FullName *common.SortOrderEnum `dbfield:"full__name"`
	}
	type Projection struct {
		Name       bool `dbfield:"name"`
		FullName   bool `dbfield:"full__name"`
		LoggedOnce bool `dbfield:"logged_once"`
	}
	type args struct {
		filter     *Filter
		sorts      []*SortOrder
		projection *Projection
		page       *pagination.PageInput
	}
	tests := []struct {
		name             string
		args             args
		wantErr          bool
		errorString      string
		expectedSQLQuery string
		expectedSQLArgs  []driver.Value
	}{
		{
			name: "1 custom pagination",
			args: args{
				page: &pagination.PageInput{Limit: 100, Skip: 50},
			},
			expectedSQLQuery: `SELECT * FROM "peoples" ORDER BY created_at DESC LIMIT $1 OFFSET $2`,
			expectedSQLArgs:  []driver.Value{100, 50},
		},
		{
			name: "1 filter",
			args: args{
				filter: &Filter{
					OR: []*Filter{
						{Name: &common.GenericFilter{Eq: "fake"}},
						{LoggedOnce: &common.GenericFilter{Eq: true}},
					},
				},
				page: &pagination.PageInput{Limit: 100},
			},
			expectedSQLQuery: `SELECT * FROM "peoples" WHERE name = $1 OR logged_once = $2 ORDER BY created_at DESC LIMIT $3`,
			expectedSQLArgs:  []driver.Value{"fake", true, 100},
		},
		{
			name: "1 sort order",
			args: args{
				sorts: []*SortOrder{{Name: &common.SortOrderEnumAsc}},
				page:  &pagination.PageInput{Limit: 100},
			},
			expectedSQLQuery: `SELECT * FROM "peoples" ORDER BY name ASC LIMIT $1`,
			expectedSQLArgs:  []driver.Value{100},
		},
		{
			name: "1 projection",
			args: args{
				projection: &Projection{Name: true},
				page:       &pagination.PageInput{Limit: 100},
			},
			expectedSQLQuery: `SELECT "name" FROM "peoples" ORDER BY created_at DESC LIMIT $1`,
			expectedSQLArgs:  []driver.Value{100},
		},
		{
			name: "1 filter and 1 sort and 1 projection",
			args: args{
				filter: &Filter{
					OR: []*Filter{
						{Name: &common.GenericFilter{Eq: "fake"}},
						{LoggedOnce: &common.GenericFilter{Eq: true}},
					},
				},
				sorts:      []*SortOrder{{Name: &common.SortOrderEnumAsc}},
				projection: &Projection{Name: true},
				page:       &pagination.PageInput{Limit: 100},
			},
			expectedSQLQuery: `SELECT "name" FROM "peoples" WHERE name = $1 OR logged_once = $2 ORDER BY name ASC LIMIT $3`,
			expectedSQLArgs:  []driver.Value{"fake", true, 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Error(err)

				return
			}
			defer sqlDB.Close()

			db, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{Logger: logger.Discard, NowFunc: func() time.Time {
				return now
			}})
			if err != nil {
				t.Error(err)

				return
			}

			ctrl := gomock.NewController(t)
			dbSvc := dbmocks.NewMockDB(ctrl)
			dbSvc.EXPECT().GetTransactionalOrDefaultGormDB(gomock.Any()).AnyTimes().Return(db)

			mock.ExpectQuery(tt.expectedSQLQuery).
				WithArgs(tt.expectedSQLArgs...).
				WillReturnRows(
					sqlmock.NewRows([]string{}),
				)

			ctx := context.TODO()
			res, err := FindWithPagination(ctx, make([]*People, 0), dbSvc, tt.args.page, tt.args.sorts, tt.args.filter, tt.args.projection)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindWithPagination() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if err != nil && err.Error() != tt.errorString {
				t.Errorf("FindWithPagination() error = %v, wantErr %v", err, tt.errorString)

				return
			}
			assert.Len(t, res, 0)
		})
	}
}

func TestCount(t *testing.T) {
	now := time.Now()

	type People struct {
		database.Base
		Name       string
		FullName   string `gorm:"column:full__name"`
		LoggedOnce bool
	}
	type Filter struct {
		Name       *common.GenericFilter `dbfield:"name"`
		FullName   *common.GenericFilter `dbfield:"full__name"`
		LoggedOnce *common.GenericFilter `dbfield:"logged_once"`
		AND        []*Filter
		OR         []*Filter
	}
	type SortOrder struct {
		Name     *common.SortOrderEnum `dbfield:"name"`
		FullName *common.SortOrderEnum `dbfield:"full__name"`
	}
	type Projection struct {
		Name       bool `dbfield:"name"`
		FullName   bool `dbfield:"full__name"`
		LoggedOnce bool `dbfield:"logged_once"`
	}
	type args struct {
		filter *Filter
	}
	tests := []struct {
		name             string
		args             args
		wantErr          bool
		errorString      string
		expectedSQLQuery string
		expectedSQLArgs  []driver.Value
	}{
		{
			name: "1 filter",
			args: args{
				filter: &Filter{
					OR: []*Filter{
						{Name: &common.GenericFilter{Eq: "fake"}},
						{LoggedOnce: &common.GenericFilter{Eq: true}},
					},
				},
			},
			expectedSQLQuery: `SELECT count(*) FROM "peoples" WHERE name = $1 OR logged_once = $2`,
			expectedSQLArgs:  []driver.Value{"fake", true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Error(err)

				return
			}
			defer sqlDB.Close()

			db, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{Logger: logger.Discard, NowFunc: func() time.Time {
				return now
			}})
			if err != nil {
				t.Error(err)

				return
			}

			ctrl := gomock.NewController(t)
			dbSvc := dbmocks.NewMockDB(ctrl)
			dbSvc.EXPECT().GetTransactionalOrDefaultGormDB(gomock.Any()).AnyTimes().Return(db)

			mock.ExpectQuery(tt.expectedSQLQuery).
				WithArgs(tt.expectedSQLArgs...).
				WillReturnRows(
					sqlmock.NewRows([]string{}),
				)

			ctx := context.TODO()
			res, err := Count(ctx, dbSvc, &People{}, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("Count() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if err != nil && err.Error() != tt.errorString {
				t.Errorf("Count() error = %v, wantErr %v", err, tt.errorString)

				return
			}
			assert.Equal(t, int64(0), res)
		})
	}
}

func TestCountPaginated(t *testing.T) {
	now := time.Now()

	type People struct {
		database.Base
		Name       string
		FullName   string `gorm:"column:full__name"`
		LoggedOnce bool
	}
	type Filter struct {
		Name       *common.GenericFilter `dbfield:"name"`
		FullName   *common.GenericFilter `dbfield:"full__name"`
		LoggedOnce *common.GenericFilter `dbfield:"logged_once"`
		AND        []*Filter
		OR         []*Filter
	}
	type SortOrder struct {
		Name     *common.SortOrderEnum `dbfield:"name"`
		FullName *common.SortOrderEnum `dbfield:"full__name"`
	}
	type Projection struct {
		Name       bool `dbfield:"name"`
		FullName   bool `dbfield:"full__name"`
		LoggedOnce bool `dbfield:"logged_once"`
	}
	type args struct {
		filter     *Filter
		sorts      []*SortOrder
		projection *Projection
		page       *pagination.PageInput
	}
	tests := []struct {
		name             string
		args             args
		wantErr          bool
		errorString      string
		expectedSQLQuery string
		expectedSQLArgs  []driver.Value
	}{
		{
			name: "1 custom pagination",
			args: args{
				page: &pagination.PageInput{Limit: 100, Skip: 50},
			},
			expectedSQLQuery: `SELECT count(*) FROM "peoples" LIMIT $1 OFFSET $2`,
			expectedSQLArgs:  []driver.Value{100, 50},
		},
		{
			name: "1 filter",
			args: args{
				filter: &Filter{
					OR: []*Filter{
						{Name: &common.GenericFilter{Eq: "fake"}},
						{LoggedOnce: &common.GenericFilter{Eq: true}},
					},
				},
				page: &pagination.PageInput{Limit: 100},
			},
			expectedSQLQuery: `SELECT count(*) FROM "peoples" WHERE name = $1 OR logged_once = $2 LIMIT $3`,
			expectedSQLArgs:  []driver.Value{"fake", true, 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Error(err)

				return
			}
			defer sqlDB.Close()

			db, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{Logger: logger.Discard, NowFunc: func() time.Time {
				return now
			}})
			if err != nil {
				t.Error(err)

				return
			}

			ctrl := gomock.NewController(t)
			dbSvc := dbmocks.NewMockDB(ctrl)
			dbSvc.EXPECT().GetTransactionalOrDefaultGormDB(gomock.Any()).AnyTimes().Return(db)

			mock.ExpectQuery(tt.expectedSQLQuery).
				WithArgs(tt.expectedSQLArgs...).
				WillReturnRows(
					sqlmock.NewRows([]string{}),
				)

			ctx := context.TODO()
			res, err := CountPaginated(ctx, dbSvc, &People{}, tt.args.page, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountPaginated() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if err != nil && err.Error() != tt.errorString {
				t.Errorf("CountPaginated() error = %v, wantErr %v", err, tt.errorString)

				return
			}
			assert.Equal(t, int64(0), res)
		})
	}
}

func TestFindByID(t *testing.T) {
	now := time.Now()

	type People struct {
		database.Base
		Name       string
		FullName   string `gorm:"column:full__name"`
		LoggedOnce bool
	}
	type Filter struct {
		Name       *common.GenericFilter `dbfield:"name"`
		FullName   *common.GenericFilter `dbfield:"full__name"`
		LoggedOnce *common.GenericFilter `dbfield:"logged_once"`
		AND        []*Filter
		OR         []*Filter
	}
	type SortOrder struct {
		Name     *common.SortOrderEnum `dbfield:"name"`
		FullName *common.SortOrderEnum `dbfield:"full__name"`
	}
	type Projection struct {
		Name       bool `dbfield:"name"`
		FullName   bool `dbfield:"full__name"`
		LoggedOnce bool `dbfield:"logged_once"`
	}
	type args struct {
		id         string
		projection *Projection
	}
	tests := []struct {
		name             string
		args             args
		wantErr          bool
		errorString      string
		expectedSQLQuery string
		expectedSQLArgs  []driver.Value
	}{
		{
			name: "classic",
			args: args{
				id: "fake",
			},
			expectedSQLQuery: `SELECT * FROM "peoples" WHERE id = $1 ORDER BY "peoples"."id" LIMIT $2`,
			expectedSQLArgs:  []driver.Value{"fake", 1},
		},
		{
			name: "1 projection",
			args: args{
				id:         "fake",
				projection: &Projection{Name: true},
			},
			expectedSQLQuery: `SELECT "name" FROM "peoples" WHERE id = $1 ORDER BY "peoples"."id" LIMIT $2`,
			expectedSQLArgs:  []driver.Value{"fake", 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Error(err)

				return
			}
			defer sqlDB.Close()

			db, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{Logger: logger.Discard, NowFunc: func() time.Time {
				return now
			}})
			if err != nil {
				t.Error(err)

				return
			}

			ctrl := gomock.NewController(t)
			dbSvc := dbmocks.NewMockDB(ctrl)
			dbSvc.EXPECT().GetTransactionalOrDefaultGormDB(gomock.Any()).AnyTimes().Return(db)

			mock.ExpectQuery(tt.expectedSQLQuery).
				WithArgs(tt.expectedSQLArgs...).
				WillReturnRows(
					sqlmock.NewRows([]string{}),
				)

			ctx := context.TODO()
			res, err := FindByID(ctx, &People{}, dbSvc, tt.args.id, tt.args.projection)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByID() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if err != nil && err.Error() != tt.errorString {
				t.Errorf("FindByID() error = %v, wantErr %v", err, tt.errorString)

				return
			}
			assert.Nil(t, res)
		})
	}
}

func TestFindOne(t *testing.T) {
	now := time.Now()

	type People struct {
		database.Base
		Name       string
		FullName   string `gorm:"column:full__name"`
		LoggedOnce bool
	}
	type Filter struct {
		Name       *common.GenericFilter `dbfield:"name"`
		FullName   *common.GenericFilter `dbfield:"full__name"`
		LoggedOnce *common.GenericFilter `dbfield:"logged_once"`
		AND        []*Filter
		OR         []*Filter
	}
	type SortOrder struct {
		Name     *common.SortOrderEnum `dbfield:"name"`
		FullName *common.SortOrderEnum `dbfield:"full__name"`
	}
	type Projection struct {
		Name       bool `dbfield:"name"`
		FullName   bool `dbfield:"full__name"`
		LoggedOnce bool `dbfield:"logged_once"`
	}
	type args struct {
		filter     *Filter
		sorts      []*SortOrder
		projection *Projection
	}
	tests := []struct {
		name             string
		args             args
		wantErr          bool
		errorString      string
		expectedSQLQuery string
		expectedSQLArgs  []driver.Value
	}{
		{
			name: "1 filter",
			args: args{
				filter: &Filter{
					OR: []*Filter{
						{Name: &common.GenericFilter{Eq: "fake"}},
						{LoggedOnce: &common.GenericFilter{Eq: true}},
					},
				},
			},
			expectedSQLQuery: `SELECT * FROM "peoples" WHERE name = $1 OR logged_once = $2 ORDER BY created_at DESC,"peoples"."id" LIMIT $3`,
			expectedSQLArgs:  []driver.Value{"fake", true, 1},
		},
		{
			name: "1 sort order",
			args: args{
				sorts: []*SortOrder{{Name: &common.SortOrderEnumAsc}},
			},
			expectedSQLQuery: `SELECT * FROM "peoples" ORDER BY name ASC,"peoples"."id" LIMIT $1`,
			expectedSQLArgs:  []driver.Value{1},
		},
		{
			name: "1 projection",
			args: args{
				projection: &Projection{Name: true},
			},
			expectedSQLQuery: `SELECT "name" FROM "peoples" ORDER BY created_at DESC,"peoples"."id" LIMIT $1`,
			expectedSQLArgs:  []driver.Value{1},
		},
		{
			name: "1 filter and 1 sort and 1 projection",
			args: args{
				filter: &Filter{
					OR: []*Filter{
						{Name: &common.GenericFilter{Eq: "fake"}},
						{LoggedOnce: &common.GenericFilter{Eq: true}},
					},
				},
				sorts:      []*SortOrder{{Name: &common.SortOrderEnumAsc}},
				projection: &Projection{Name: true},
			},
			expectedSQLQuery: `SELECT "name" FROM "peoples" WHERE name = $1 OR logged_once = $2 ORDER BY name ASC,"peoples"."id" LIMIT $3`,
			expectedSQLArgs:  []driver.Value{"fake", true, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Error(err)

				return
			}
			defer sqlDB.Close()

			db, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{Logger: logger.Discard, NowFunc: func() time.Time {
				return now
			}})
			if err != nil {
				t.Error(err)

				return
			}

			ctrl := gomock.NewController(t)
			dbSvc := dbmocks.NewMockDB(ctrl)
			dbSvc.EXPECT().GetTransactionalOrDefaultGormDB(gomock.Any()).AnyTimes().Return(db)

			mock.ExpectQuery(tt.expectedSQLQuery).
				WithArgs(tt.expectedSQLArgs...).
				WillReturnRows(
					sqlmock.NewRows([]string{}),
				)

			ctx := context.TODO()
			res, err := FindOne(ctx, &People{}, dbSvc, tt.args.sorts, tt.args.filter, tt.args.projection)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindOne() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if err != nil && err.Error() != tt.errorString {
				t.Errorf("FindOne() error = %v, wantErr %v", err, tt.errorString)

				return
			}
			assert.Nil(t, res)
		})
	}
}
