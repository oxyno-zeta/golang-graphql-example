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
)

func TestPatchUpdate(t *testing.T) {
	now := time.Now()

	type People struct {
		database.Base
		Name       string
		FullName   string `gorm:"column:full__name"`
		LoggedOnce bool
	}
	type args struct {
		model any
		input map[string]any
	}
	tests := []struct {
		name             string
		args             args
		want             any
		wantErr          bool
		errorString      string
		expectedSQLQuery string
		expectedSQLArgs  []driver.Value
	}{
		{
			name: "1 simple field",
			args: args{
				model: &People{
					Base:       database.Base{ID: "id1", UpdatedAt: now.Add(-time.Second)},
					Name:       "original",
					LoggedOnce: true,
				},
				input: map[string]any{"name": "updated"},
			},
			expectedSQLQuery: `UPDATE "peoples" SET "name"=$1,"updated_at"=$2 WHERE "id" = $3`,
			expectedSQLArgs:  []driver.Value{"updated", now, "id1"},
			want: &People{
				Base:       database.Base{ID: "id1", UpdatedAt: now},
				Name:       "updated",
				LoggedOnce: true,
			},
		},
		{
			name: "1 custom field",
			args: args{
				model: &People{
					Base:       database.Base{ID: "id1", UpdatedAt: now.Add(-time.Second)},
					Name:       "original",
					LoggedOnce: true,
				},
				input: map[string]any{"full__name": "updated"},
			},
			expectedSQLQuery: `UPDATE "peoples" SET "full__name"=$1,"updated_at"=$2 WHERE "id" = $3`,
			expectedSQLArgs:  []driver.Value{"updated", now, "id1"},
			want: &People{
				Base:       database.Base{ID: "id1", UpdatedAt: now},
				Name:       "original",
				FullName:   "updated",
				LoggedOnce: true,
			},
		},
		{
			name: "2 simple fields",
			args: args{
				model: &People{
					Base:       database.Base{ID: "id1", UpdatedAt: now.Add(-time.Second)},
					Name:       "original",
					LoggedOnce: true,
				},
				input: map[string]any{"name": "updated", "logged_once": false},
			},
			expectedSQLQuery: `UPDATE "peoples" SET "logged_once"=$1,"name"=$2,"updated_at"=$3 WHERE "id" = $4`,
			expectedSQLArgs:  []driver.Value{false, "updated", now, "id1"},
			want: &People{
				Base:       database.Base{ID: "id1", UpdatedAt: now},
				Name:       "updated",
				LoggedOnce: false,
			},
		},
		{
			name: "1 custom and 1 simple field",
			args: args{
				model: &People{
					Base:       database.Base{ID: "id1", UpdatedAt: now.Add(-time.Second)},
					Name:       "original",
					LoggedOnce: true,
				},
				input: map[string]any{"full__name": "updated", "name": "updated"},
			},
			expectedSQLQuery: `UPDATE "peoples" SET "full__name"=$1,"name"=$2,"updated_at"=$3 WHERE "id" = $4`,
			expectedSQLArgs:  []driver.Value{"updated", "updated", now, "id1"},
			want: &People{
				Base:       database.Base{ID: "id1", UpdatedAt: now},
				Name:       "updated",
				FullName:   "updated",
				LoggedOnce: true,
			},
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

			mock.ExpectBegin()
			mock.ExpectExec(tt.expectedSQLQuery).
				WithArgs(tt.expectedSQLArgs...).
				WillReturnResult(sqlmock.NewResult(0, 1))
			mock.ExpectCommit()

			ctx := context.TODO()
			got, err := PatchUpdate(ctx, tt.args.model, tt.args.input, dbSvc)
			if (err != nil) != tt.wantErr {
				t.Errorf("PatchUpdate() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if err != nil && err.Error() != tt.errorString {
				t.Errorf("PatchUpdate() error = %v, wantErr %v", err, tt.errorString)

				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPatchUpdateAllFiltered(t *testing.T) {
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
	type args struct {
		model  any
		input  map[string]any
		filter any
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
			name: "1 simple field",
			args: args{
				model: &People{},
				input: map[string]any{"name": "updated"},
				filter: &Filter{
					OR: []*Filter{
						{Name: &common.GenericFilter{Eq: "fake"}},
						{LoggedOnce: &common.GenericFilter{Eq: true}},
					},
				},
			},
			expectedSQLQuery: `UPDATE "peoples" SET "name"=$1,"updated_at"=$2 WHERE name = $3 OR logged_once = $4`,
			expectedSQLArgs:  []driver.Value{"updated", now, "fake", true},
		},
		{
			name: "1 custom field",
			args: args{
				model: &People{},
				input: map[string]any{"full__name": "updated"},
				filter: &Filter{
					OR: []*Filter{
						{Name: &common.GenericFilter{Eq: "fake"}},
						{LoggedOnce: &common.GenericFilter{Eq: true}},
					},
				},
			},
			expectedSQLQuery: `UPDATE "peoples" SET "full__name"=$1,"updated_at"=$2 WHERE name = $3 OR logged_once = $4`,
			expectedSQLArgs:  []driver.Value{"updated", now, "fake", true},
		},
		{
			name: "2 simple fields",
			args: args{
				model: &People{},
				input: map[string]any{"name": "updated", "logged_once": false},
				filter: &Filter{
					OR: []*Filter{
						{Name: &common.GenericFilter{Eq: "fake"}},
						{LoggedOnce: &common.GenericFilter{Eq: true}},
					},
				},
			},
			expectedSQLQuery: `UPDATE "peoples" SET "logged_once"=$1,"name"=$2,"updated_at"=$3 WHERE name = $4 OR logged_once = $5`,
			expectedSQLArgs:  []driver.Value{false, "updated", now, "fake", true},
		},
		{
			name: "1 custom and 1 simple field",
			args: args{
				model: &People{},
				input: map[string]any{"full__name": "updated", "name": "updated"},
				filter: &Filter{
					OR: []*Filter{
						{Name: &common.GenericFilter{Eq: "fake"}},
						{LoggedOnce: &common.GenericFilter{Eq: true}},
					},
				},
			},
			expectedSQLQuery: `UPDATE "peoples" SET "full__name"=$1,"name"=$2,"updated_at"=$3 WHERE name = $4 OR logged_once = $5`,
			expectedSQLArgs:  []driver.Value{"updated", "updated", now, "fake", true},
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

			mock.ExpectBegin()
			mock.ExpectExec(tt.expectedSQLQuery).
				WithArgs(tt.expectedSQLArgs...).
				WillReturnResult(sqlmock.NewResult(0, 1))
			mock.ExpectCommit()

			ctx := context.TODO()
			err = PatchUpdateAllFiltered(ctx, tt.args.model, tt.args.input, tt.args.filter, dbSvc)
			if (err != nil) != tt.wantErr {
				t.Errorf("PatchUpdateAllFiltered() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if err != nil && err.Error() != tt.errorString {
				t.Errorf("PatchUpdateAllFiltered() error = %v, wantErr %v", err, tt.errorString)

				return
			}
		})
	}
}

func TestSoftDelete(t *testing.T) {
	now := time.Now()

	type People struct {
		database.Base
		Name       string
		FullName   string `gorm:"column:full__name"`
		LoggedOnce bool
	}
	type args struct {
		input any
	}
	tests := []struct {
		name             string
		args             args
		want             any
		wantErr          bool
		errorString      string
		expectedSQLQuery string
		expectedSQLArgs  []driver.Value
	}{
		{
			name: "simple case",
			args: args{
				input: &People{
					Base:       database.Base{ID: "id1", UpdatedAt: now.Add(-time.Second)},
					Name:       "original",
					LoggedOnce: true,
				},
			},
			expectedSQLQuery: `DELETE FROM "peoples" WHERE "peoples"."id" = $1`,
			expectedSQLArgs:  []driver.Value{"id1"},
			want: &People{
				Base:       database.Base{ID: "id1", UpdatedAt: now.Add(-time.Second)},
				Name:       "original",
				LoggedOnce: true,
			},
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

			mock.ExpectBegin()
			mock.ExpectExec(tt.expectedSQLQuery).
				WithArgs(tt.expectedSQLArgs...).
				WillReturnResult(sqlmock.NewResult(0, 1))
			mock.ExpectCommit()

			ctx := context.TODO()
			got, err := SoftDelete(ctx, tt.args.input, dbSvc)
			if (err != nil) != tt.wantErr {
				t.Errorf("SoftDelete() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if err != nil && err.Error() != tt.errorString {
				t.Errorf("SoftDelete() error = %v, wantErr %v", err, tt.errorString)

				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSoftDeleteFiltered(t *testing.T) {
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
	type args struct {
		model  any
		filter any
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
			name: "1 simple field",
			args: args{
				model: &People{},
				filter: &Filter{
					OR: []*Filter{
						{Name: &common.GenericFilter{Eq: "fake"}},
						{LoggedOnce: &common.GenericFilter{Eq: true}},
					},
				},
			},
			expectedSQLQuery: `DELETE FROM "peoples" WHERE name = $1 OR logged_once = $2`,
			expectedSQLArgs:  []driver.Value{"fake", true},
		},
		{
			name: "1 custom field",
			args: args{
				model: &People{},
				filter: &Filter{
					OR: []*Filter{
						{Name: &common.GenericFilter{Eq: "fake"}},
						{LoggedOnce: &common.GenericFilter{Eq: true}},
					},
				},
			},
			expectedSQLQuery: `DELETE FROM "peoples" WHERE name = $1 OR logged_once = $2`,
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

			mock.ExpectBegin()
			mock.ExpectExec(tt.expectedSQLQuery).
				WithArgs(tt.expectedSQLArgs...).
				WillReturnResult(sqlmock.NewResult(0, 1))
			mock.ExpectCommit()

			ctx := context.TODO()
			err = SoftDeleteFiltered(ctx, tt.args.model, tt.args.filter, dbSvc)
			if (err != nil) != tt.wantErr {
				t.Errorf("SoftDeleteFiltered() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if err != nil && err.Error() != tt.errorString {
				t.Errorf("SoftDeleteFiltered() error = %v, wantErr %v", err, tt.errorString)

				return
			}
		})
	}
}

func TestPermanentDelete(t *testing.T) {
	now := time.Now()

	type People struct {
		database.Base
		Name       string
		FullName   string `gorm:"column:full__name"`
		LoggedOnce bool
	}
	type args struct {
		input any
	}
	tests := []struct {
		name             string
		args             args
		want             any
		wantErr          bool
		errorString      string
		expectedSQLQuery string
		expectedSQLArgs  []driver.Value
	}{
		{
			name: "simple case",
			args: args{
				input: &People{
					Base:       database.Base{ID: "id1", UpdatedAt: now.Add(-time.Second)},
					Name:       "original",
					LoggedOnce: true,
				},
			},
			expectedSQLQuery: `DELETE FROM "peoples" WHERE "peoples"."id" = $1`,
			expectedSQLArgs:  []driver.Value{"id1"},
			want: &People{
				Base:       database.Base{ID: "id1", UpdatedAt: now.Add(-time.Second)},
				Name:       "original",
				LoggedOnce: true,
			},
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

			mock.ExpectBegin()
			mock.ExpectExec(tt.expectedSQLQuery).
				WithArgs(tt.expectedSQLArgs...).
				WillReturnResult(sqlmock.NewResult(0, 1))
			mock.ExpectCommit()

			ctx := context.TODO()
			got, err := PermanentDelete(ctx, tt.args.input, dbSvc)
			if (err != nil) != tt.wantErr {
				t.Errorf("PermanentDelete() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if err != nil && err.Error() != tt.errorString {
				t.Errorf("PermanentDelete() error = %v, wantErr %v", err, tt.errorString)

				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPermanentDeleteFiltered(t *testing.T) {
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
	type args struct {
		model  any
		filter any
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
			name: "1 simple field",
			args: args{
				model: &People{},
				filter: &Filter{
					OR: []*Filter{
						{Name: &common.GenericFilter{Eq: "fake"}},
						{LoggedOnce: &common.GenericFilter{Eq: true}},
					},
				},
			},
			expectedSQLQuery: `DELETE FROM "peoples" WHERE name = $1 OR logged_once = $2`,
			expectedSQLArgs:  []driver.Value{"fake", true},
		},
		{
			name: "1 custom field",
			args: args{
				model: &People{},
				filter: &Filter{
					OR: []*Filter{
						{Name: &common.GenericFilter{Eq: "fake"}},
						{LoggedOnce: &common.GenericFilter{Eq: true}},
					},
				},
			},
			expectedSQLQuery: `DELETE FROM "peoples" WHERE name = $1 OR logged_once = $2`,
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

			mock.ExpectBegin()
			mock.ExpectExec(tt.expectedSQLQuery).
				WithArgs(tt.expectedSQLArgs...).
				WillReturnResult(sqlmock.NewResult(0, 1))
			mock.ExpectCommit()

			ctx := context.TODO()
			err = PermanentDeleteFiltered(ctx, tt.args.model, tt.args.filter, dbSvc)
			if (err != nil) != tt.wantErr {
				t.Errorf("PermanentDeleteFiltered() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if err != nil && err.Error() != tt.errorString {
				t.Errorf("PermanentDeleteFiltered() error = %v, wantErr %v", err, tt.errorString)

				return
			}
		})
	}
}
