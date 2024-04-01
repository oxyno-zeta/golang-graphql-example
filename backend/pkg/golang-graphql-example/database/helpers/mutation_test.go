//go:build unit

package databasehelpers

import (
	"context"
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	dbmocks "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestUpdate(t *testing.T) {
	now := time.Now()

	type People struct {
		database.Base
		Name       string
		FullName   string `gorm:"column:full__name"`
		LoggedOnce bool
	}
	type args struct {
		model interface{}
		input map[string]interface{}
	}
	tests := []struct {
		name             string
		args             args
		want             interface{}
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
				input: map[string]interface{}{"name": "updated"},
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
				input: map[string]interface{}{"full__name": "updated"},
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
				input: map[string]interface{}{"name": "updated", "logged_once": false},
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
				input: map[string]interface{}{"full__name": "updated", "name": "updated"},
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
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.errorString {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.errorString)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
