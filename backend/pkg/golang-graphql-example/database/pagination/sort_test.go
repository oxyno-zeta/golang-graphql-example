// +build unit

package pagination

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_manageSortOrder(t *testing.T) {
	type Person struct {
		Name string
	}
	type Sort1 struct {
		Fake1 *SortOrderEnum `db_col:"fake_1"`
		Fake2 *SortOrderEnum `db_col:"fake_2"`
	}
	type Sort2 struct {
		Fake1 *SortOrderEnum `db_col:"fake_1"`
		Fake2 *SortOrderEnum `db_col:"-"`
	}
	type Sort3 struct {
		Fake1 *SortOrderEnum `db_col:"fake_1"`
		Fake2 *SortOrderEnum
	}
	type Sort4 struct {
		Fake1 *SortOrderEnum `db_col:"fake_1"`
		Fake2 string         `db_col:"fake_2"`
	}
	type Sort5 struct {
		Fake1 *SortOrderEnum `db_col:"fake_1"`
		Fake2 string
	}
	type Sort6 struct {
		Fake1 SortOrderEnum `db_col:"fake_1"`
	}
	type args struct {
		sort interface{}
	}
	tests := []struct {
		name              string
		args              args
		expectedSortQuery string
		wantErr           bool
		errorString       string
	}{
		{
			name:        "wrong input",
			args:        args{sort: false},
			wantErr:     true,
			errorString: "sort must be an object",
		},
		{
			name: "nil sort object",
			args: args{
				sort: nil,
			},
			expectedSortQuery: "",
		},
		{
			name: "empty sort object",
			args: args{
				sort: &Sort1{},
			},
		},
		{
			name: "full set sort pointer object",
			args: args{
				sort: &Sort1{Fake1: &SortOrderEnumAsc, Fake2: &SortOrderEnumDesc},
			},
			expectedSortQuery: "ORDER BY fake_1 ASC,fake_2 DESC",
		},
		{
			name: "full set sort object",
			args: args{
				sort: Sort1{Fake1: &SortOrderEnumAsc, Fake2: &SortOrderEnumDesc},
			},
			expectedSortQuery: "ORDER BY fake_1 ASC,fake_2 DESC",
		},
		{
			name: "ignored filed",
			args: args{
				sort: &Sort2{Fake1: &SortOrderEnumAsc, Fake2: &SortOrderEnumDesc},
			},
			expectedSortQuery: "ORDER BY fake_1 ASC",
		},
		{
			name: "no tag",
			args: args{
				sort: &Sort3{Fake1: &SortOrderEnumAsc, Fake2: &SortOrderEnumDesc},
			},
			expectedSortQuery: "ORDER BY fake_1 ASC",
		},
		{
			name: "tag but not on right type",
			args: args{
				sort: &Sort4{Fake1: &SortOrderEnumAsc, Fake2: "fake"},
			},
			wantErr:     true,
			errorString: "field with sort tag must be a *SortOrderEnum",
		},
		{
			name: "wrong type without field must be ignored",
			args: args{
				sort: &Sort5{Fake1: &SortOrderEnumAsc, Fake2: "fake"},
			},
			expectedSortQuery: "ORDER BY fake_1 ASC",
		},
		{
			name: "wrong enum type used (no pointer)",
			args: args{
				sort: &Sort6{Fake1: SortOrderEnumAsc},
			},
			wantErr:     true,
			errorString: "field with sort tag must be a *SortOrderEnum",
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

			db, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})
			if err != nil {
				t.Error(err)
				return
			}

			got, err := manageSortOrder(tt.args.sort, db)
			if (err != nil) != tt.wantErr {
				t.Errorf("manageSortOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && err.Error() != tt.errorString {
				t.Errorf("manageSortOrder() error = %v, wantErr %v", err, tt.errorString)
				return
			}
			if err != nil {
				return
			}

			// Create expected query
			expectedQuery := `SELECT * FROM "people" ` + tt.expectedSortQuery
			if tt.expectedSortQuery != "" {
				expectedQuery += " "
			}
			expectedQuery += "LIMIT 1"
			mock.ExpectQuery(expectedQuery).WithArgs().WillReturnRows(
				sqlmock.NewRows([]string{"name"}).AddRow("fake"),
			)

			// Run fake find to force query to be run
			res := got.First(&Person{})
			// Test error
			if res.Error != nil {
				t.Error(res.Error)
			}
		})
	}
}
