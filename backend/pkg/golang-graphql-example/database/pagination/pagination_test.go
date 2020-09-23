// +build unit

package pagination

import (
	"database/sql/driver"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/common"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestPaging(t *testing.T) {
	type Person struct{ Name string }
	type Sort struct {
		Name *common.SortOrderEnum `dbfield:"name"`
	}
	type Filter struct {
		Name *common.GenericFilter `dbfield:"name"`
	}
	type args struct {
		p         *PageInput
		sort      interface{}
		filter    interface{}
		extraFunc func(db *gorm.DB) (*gorm.DB, error)
	}
	tests := []struct {
		name                            string
		args                            args
		countExpectedIntermediateQuery  string
		countExpectedArgs               []driver.Value
		selectExpectedIntermediateQuery string
		selectExpectedArgs              []driver.Value
		countResult                     int
		want                            *PageOutput
		wantErr                         bool
	}{
		{
			name: "no sort, no filter, no extra function, no limit",
			args: args{
				p: &PageInput{},
			},
			countExpectedIntermediateQuery:  "",
			countExpectedArgs:               []driver.Value{},
			countResult:                     3,
			selectExpectedIntermediateQuery: "ORDER BY created_at DESC LIMIT 10",
			selectExpectedArgs:              []driver.Value{},
			want:                            &PageOutput{TotalRecord: 3, Limit: 10},
		},
		{
			name: "no sort, no filter, no extra function",
			args: args{
				p: &PageInput{Limit: 5},
			},
			countExpectedIntermediateQuery:  "",
			countExpectedArgs:               []driver.Value{},
			countResult:                     3,
			selectExpectedIntermediateQuery: "ORDER BY created_at DESC LIMIT 5",
			selectExpectedArgs:              []driver.Value{},
			want:                            &PageOutput{TotalRecord: 3, Limit: 5},
		},
		{
			name: "no sort, no filter, no extra function with next page",
			args: args{
				p: &PageInput{Limit: 5},
			},
			countExpectedIntermediateQuery:  "",
			countExpectedArgs:               []driver.Value{},
			countResult:                     30,
			selectExpectedIntermediateQuery: "ORDER BY created_at DESC LIMIT 5",
			selectExpectedArgs:              []driver.Value{},
			want:                            &PageOutput{TotalRecord: 30, Limit: 5, HasNext: true},
		},
		{
			name: "no sort, no filter, no extra function with next and previous page and skip",
			args: args{
				p: &PageInput{Limit: 5, Skip: 20},
			},
			countExpectedIntermediateQuery:  "",
			countExpectedArgs:               []driver.Value{},
			countResult:                     30,
			selectExpectedIntermediateQuery: "ORDER BY created_at DESC LIMIT 5 OFFSET 20",
			selectExpectedArgs:              []driver.Value{},
			want:                            &PageOutput{TotalRecord: 30, Limit: 5, Skip: 20, HasNext: true, HasPrevious: true},
		},
		{
			name: "sort, filter, no extra function with next and previous page and skip",
			args: args{
				p:      &PageInput{Limit: 5, Skip: 20},
				sort:   &Sort{Name: &common.SortOrderEnumDesc},
				filter: &Filter{Name: &common.GenericFilter{Eq: "fake"}},
			},
			countExpectedIntermediateQuery:  "WHERE name = $1",
			countExpectedArgs:               []driver.Value{"fake"},
			countResult:                     30,
			selectExpectedIntermediateQuery: "WHERE name = $1 ORDER BY name DESC LIMIT 5 OFFSET 20",
			selectExpectedArgs:              []driver.Value{"fake"},
			want:                            &PageOutput{TotalRecord: 30, Limit: 5, Skip: 20, HasNext: true, HasPrevious: true},
		},
		{
			name: "extra function throwing error",
			args: args{
				p: &PageInput{Limit: 5, Skip: 20},
				extraFunc: func(db *gorm.DB) (*gorm.DB, error) {
					return nil, errors.New("fake")
				},
			},
			wantErr: true,
		},
		{
			name: "no sort, no filter, extra function with next and previous page and skip",
			args: args{
				p: &PageInput{Limit: 5, Skip: 20},
				extraFunc: func(db *gorm.DB) (*gorm.DB, error) {
					return db.Where("fake = ?", "fake1"), nil
				},
			},
			countExpectedIntermediateQuery:  "WHERE fake = $1",
			countExpectedArgs:               []driver.Value{"fake1"},
			countResult:                     30,
			selectExpectedIntermediateQuery: "WHERE fake = $1 ORDER BY created_at DESC LIMIT 5 OFFSET 20",
			selectExpectedArgs:              []driver.Value{"fake1"},
			want:                            &PageOutput{TotalRecord: 30, Limit: 5, Skip: 20, HasNext: true, HasPrevious: true},
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

			db, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{Logger: logger.Discard})
			if err != nil {
				t.Error(err)
				return
			}

			// Create expected query
			countExpectedQuery := `SELECT count(1) FROM "people" ` + tt.countExpectedIntermediateQuery
			// Create expected query
			selectExpectedQuery := `SELECT * FROM "people" ` + tt.selectExpectedIntermediateQuery

			mock.ExpectQuery(countExpectedQuery).
				WithArgs(tt.countExpectedArgs...).
				WillReturnRows(
					sqlmock.NewRows([]string{"count"}).AddRow(tt.countResult),
				)
			mock.ExpectQuery(selectExpectedQuery).
				WithArgs(tt.selectExpectedArgs...).
				WillReturnRows(
					sqlmock.NewRows([]string{"name"}).AddRow("fake"),
				)

			res := make([]*Person, 0)

			got, err := Paging(&res, db, tt.args.p, tt.args.sort, tt.args.filter, tt.args.extraFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("Paging() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Paging() = %v, want %v", got, tt.want)
			}
		})
	}
}
