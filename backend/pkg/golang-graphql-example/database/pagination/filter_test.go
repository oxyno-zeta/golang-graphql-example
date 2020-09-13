package pagination

import (
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

// func Test_manageFilter(t *testing.T) {
// 	starInterface := func(s interface{}) *interface{} { return &s }
// 	type Person struct {
// 		Name string
// 	}
// 	type FilterSt1 struct {
// 		Field1 *GenericFilter `db_col:"field_1"`
// 	}
// 	type args struct {
// 		filter interface{}
// 	}
// 	tests := []struct {
// 		name                      string
// 		args                      args
// 		expectedIntermediateQuery string
// 		expectedArgs              []driver.Value
// 		wantErr                   bool
// 		errorString               string
// 	}{
// 		// {
// 		// 	name:        "wrong input",
// 		// 	args:        args{filter: false},
// 		// 	wantErr:     true,
// 		// 	errorString: "filter must be an object",
// 		// },
// 		// {
// 		// 	name: "nil sort object",
// 		// 	args: args{
// 		// 		filter: nil,
// 		// 	},
// 		// 	expectedIntermediateQuery: "",
// 		// },
// 		// {
// 		// 	name: "",
// 		// 	args: args{
// 		// 		filter: &FilterSt1{
// 		// 			Field1: &GenericFilter{
// 		// 				Eq:     starInterface("fake"),
// 		// 				NotGte: "dkk",
// 		// 				NotEq:  1,
// 		// 			},
// 		// 		},
// 		// 	},
// 		// 	expectedIntermediateQuery: "WHERE (field_1 = $1) AND (dazd = $2)",
// 		// 	expectedArgs:              []driver.Value{"fake", "oazdko"},
// 		// },
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			sqlDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
// 			if err != nil {
// 				t.Error(err)
// 				return
// 			}
// 			defer sqlDB.Close()

// 			db, err := gorm.Open("postgres", sqlDB)
// 			if err != nil {
// 				t.Error(err)
// 				return
// 			}
// 			db.LogMode(false)

// 			got, err := manageFilter(tt.args.filter, db)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("manageFilter() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 			if err != nil && err.Error() != tt.errorString {
// 				t.Errorf("manageFilter() error = %v, wantErr %v", err, tt.errorString)
// 				return
// 			}
// 			if err != nil {
// 				return
// 			}

// 			// Create expected query
// 			expectedQuery := `SELECT * FROM "people" ` + tt.expectedIntermediateQuery
// 			if tt.expectedIntermediateQuery != "" {
// 				expectedQuery += " "
// 			}
// 			expectedQuery += "LIMIT 1"

// 			mock.ExpectQuery(expectedQuery).
// 				WithArgs(tt.expectedArgs...).
// 				WillReturnRows(
// 					sqlmock.NewRows([]string{"name"}).AddRow("fake"),
// 				)

// 			// Run fake find to force query to be run
// 			res := got.First(&Person{})
// 			// Test error
// 			if res.Error != nil {
// 				t.Error(res.Error)
// 			}
// 		})
// 	}
// }

type StringTestEnum string

const FakeStringTestEnum StringTestEnum = "FAKE"

type IntTestEnum int

const FakeIntTestEum IntTestEnum = 1

func Test_manageGenericFilter(t *testing.T) {
	starInterface := func(s interface{}) interface{} { return &s }
	starString := func(s string) *string { return &s }
	now := time.Now()

	type Person struct {
		Name string
	}
	type args struct {
		v  *GenericFilter
		db *gorm.DB
	}
	tests := []struct {
		name                      string
		args                      args
		expectedIntermediateQuery string
		expectedArgs              []driver.Value
		wantErr                   bool
		errorString               string
	}{
		// EQ
		{
			name: "eq case with string",
			args: args{
				v: &GenericFilter{Eq: "fake"},
			},
			expectedIntermediateQuery: "WHERE (field_1 = $1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "eq case with *string",
			args: args{
				v: &GenericFilter{Eq: starInterface("fake")},
			},
			expectedIntermediateQuery: "WHERE (field_1 = $1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "eq case with int",
			args: args{
				v: &GenericFilter{Eq: 1},
			},
			expectedIntermediateQuery: "WHERE (field_1 = $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "eq case with *int",
			args: args{
				v: &GenericFilter{Eq: starInterface(1)},
			},
			expectedIntermediateQuery: "WHERE (field_1 = $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "eq case with bool",
			args: args{
				v: &GenericFilter{Eq: true},
			},
			expectedIntermediateQuery: "WHERE (field_1 = $1)",
			expectedArgs:              []driver.Value{true},
		},
		{
			name: "eq case with *bool",
			args: args{
				v: &GenericFilter{Eq: starInterface(true)},
			},
			expectedIntermediateQuery: "WHERE (field_1 = $1)",
			expectedArgs:              []driver.Value{true},
		},
		{
			name: "eq case with date",
			args: args{
				v: &GenericFilter{Eq: now},
			},
			expectedIntermediateQuery: "WHERE (field_1 = $1)",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "eq case with *date",
			args: args{
				v: &GenericFilter{Eq: starInterface(now)},
			},
			expectedIntermediateQuery: "WHERE (field_1 = $1)",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "eq case with Enum struct",
			args: args{
				v: &GenericFilter{Eq: FakeStringTestEnum},
			},
			expectedIntermediateQuery: "WHERE (field_1 = $1)",
			expectedArgs:              []driver.Value{"FAKE"},
		},
		{
			name: "eq case with *Enum struct",
			args: args{
				v: &GenericFilter{Eq: starInterface(FakeStringTestEnum)},
			},
			expectedIntermediateQuery: "WHERE (field_1 = $1)",
			expectedArgs:              []driver.Value{"FAKE"},
		},
		// NOT EQ
		{
			name: "not eq case with string",
			args: args{
				v: &GenericFilter{NotEq: "fake"},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 = $1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not eq case with *string",
			args: args{
				v: &GenericFilter{NotEq: starInterface("fake")},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 = $1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not eq case with int",
			args: args{
				v: &GenericFilter{NotEq: 1},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 = $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not eq case with *int",
			args: args{
				v: &GenericFilter{NotEq: starInterface(1)},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 = $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not eq case with bool",
			args: args{
				v: &GenericFilter{NotEq: true},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 = $1)",
			expectedArgs:              []driver.Value{true},
		},
		{
			name: "not eq case with *bool",
			args: args{
				v: &GenericFilter{NotEq: starInterface(true)},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 = $1)",
			expectedArgs:              []driver.Value{true},
		},
		{
			name: "not eq case with date",
			args: args{
				v: &GenericFilter{NotEq: now},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 = $1)",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "not eq case with *date",
			args: args{
				v: &GenericFilter{NotEq: starInterface(now)},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 = $1)",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "not eq case with Enum struct",
			args: args{
				v: &GenericFilter{NotEq: FakeStringTestEnum},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 = $1)",
			expectedArgs:              []driver.Value{"FAKE"},
		},
		{
			name: "not eq case with *Enum struct",
			args: args{
				v: &GenericFilter{NotEq: starInterface(FakeStringTestEnum)},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 = $1)",
			expectedArgs:              []driver.Value{"FAKE"},
		},
		// GTE
		{
			name: "gte case with string",
			args: args{
				v: &GenericFilter{Gte: 1},
			},
			expectedIntermediateQuery: "WHERE (field_1 >= $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "gte case with *string",
			args: args{
				v: &GenericFilter{Gte: starInterface("fake")},
			},
			expectedIntermediateQuery: "WHERE (field_1 >= $1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "gte case with int",
			args: args{
				v: &GenericFilter{Gte: 1},
			},
			expectedIntermediateQuery: "WHERE (field_1 >= $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "gte case with *int",
			args: args{
				v: &GenericFilter{Gte: starInterface(1)},
			},
			expectedIntermediateQuery: "WHERE (field_1 >= $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "gte case with date",
			args: args{
				v: &GenericFilter{Gte: now},
			},
			expectedIntermediateQuery: "WHERE (field_1 >= $1)",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "gte case with *date",
			args: args{
				v: &GenericFilter{Gte: starInterface(now)},
			},
			expectedIntermediateQuery: "WHERE (field_1 >= $1)",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "gte case with Enum struct",
			args: args{
				v: &GenericFilter{Gte: FakeIntTestEum},
			},
			expectedIntermediateQuery: "WHERE (field_1 >= $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "gte case with *Enum struct",
			args: args{
				v: &GenericFilter{Gte: starInterface(FakeIntTestEum)},
			},
			expectedIntermediateQuery: "WHERE (field_1 >= $1)",
			expectedArgs:              []driver.Value{1},
		},
		// NOT GTE
		{
			name: "not gte case with string",
			args: args{
				v: &GenericFilter{NotGte: 1},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 >= $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not gte case with *string",
			args: args{
				v: &GenericFilter{NotGte: starInterface("fake")},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 >= $1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not gte case with int",
			args: args{
				v: &GenericFilter{NotGte: 1},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 >= $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not gte case with *int",
			args: args{
				v: &GenericFilter{NotGte: starInterface(1)},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 >= $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not gte case with date",
			args: args{
				v: &GenericFilter{NotGte: now},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 >= $1)",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "not gte case with *date",
			args: args{
				v: &GenericFilter{NotGte: starInterface(now)},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 >= $1)",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "not gte case with Enum struct",
			args: args{
				v: &GenericFilter{NotGte: FakeIntTestEum},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 >= $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not gte case with *Enum struct",
			args: args{
				v: &GenericFilter{NotGte: starInterface(FakeIntTestEum)},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 >= $1)",
			expectedArgs:              []driver.Value{1},
		},
		// GT
		{
			name: "gt case with string",
			args: args{
				v: &GenericFilter{Gt: 1},
			},
			expectedIntermediateQuery: "WHERE (field_1 > $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "gt case with *string",
			args: args{
				v: &GenericFilter{Gt: starInterface("fake")},
			},
			expectedIntermediateQuery: "WHERE (field_1 > $1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "gt case with int",
			args: args{
				v: &GenericFilter{Gt: 1},
			},
			expectedIntermediateQuery: "WHERE (field_1 > $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "gt case with *int",
			args: args{
				v: &GenericFilter{Gt: starInterface(1)},
			},
			expectedIntermediateQuery: "WHERE (field_1 > $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "gt case with date",
			args: args{
				v: &GenericFilter{Gt: now},
			},
			expectedIntermediateQuery: "WHERE (field_1 > $1)",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "gt case with *date",
			args: args{
				v: &GenericFilter{Gt: starInterface(now)},
			},
			expectedIntermediateQuery: "WHERE (field_1 > $1)",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "gt case with Enum struct",
			args: args{
				v: &GenericFilter{Gt: FakeIntTestEum},
			},
			expectedIntermediateQuery: "WHERE (field_1 > $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "gt case with *Enum struct",
			args: args{
				v: &GenericFilter{Gt: starInterface(FakeIntTestEum)},
			},
			expectedIntermediateQuery: "WHERE (field_1 > $1)",
			expectedArgs:              []driver.Value{1},
		},
		// NOT GT
		{
			name: "not gt case with string",
			args: args{
				v: &GenericFilter{NotGt: 1},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 > $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not gt case with *string",
			args: args{
				v: &GenericFilter{NotGt: starInterface("fake")},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 > $1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not gt case with int",
			args: args{
				v: &GenericFilter{NotGt: 1},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 > $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not gt case with *int",
			args: args{
				v: &GenericFilter{NotGt: starInterface(1)},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 > $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not gt case with date",
			args: args{
				v: &GenericFilter{NotGt: now},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 > $1)",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "not gt case with *date",
			args: args{
				v: &GenericFilter{NotGt: starInterface(now)},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 > $1)",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "not gt case with Enum struct",
			args: args{
				v: &GenericFilter{NotGt: FakeIntTestEum},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 > $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not gt case with *Enum struct",
			args: args{
				v: &GenericFilter{NotGt: starInterface(FakeIntTestEum)},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 > $1)",
			expectedArgs:              []driver.Value{1},
		},
		// LTE
		{
			name: "lte case with string",
			args: args{
				v: &GenericFilter{Lte: 1},
			},
			expectedIntermediateQuery: "WHERE (field_1 <= $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "lte case with *string",
			args: args{
				v: &GenericFilter{Lte: starInterface("fake")},
			},
			expectedIntermediateQuery: "WHERE (field_1 <= $1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "lte case with int",
			args: args{
				v: &GenericFilter{Lte: 1},
			},
			expectedIntermediateQuery: "WHERE (field_1 <= $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "lte case with *int",
			args: args{
				v: &GenericFilter{Lte: starInterface(1)},
			},
			expectedIntermediateQuery: "WHERE (field_1 <= $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "lte case with date",
			args: args{
				v: &GenericFilter{Lte: now},
			},
			expectedIntermediateQuery: "WHERE (field_1 <= $1)",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "lte case with *date",
			args: args{
				v: &GenericFilter{Lte: starInterface(now)},
			},
			expectedIntermediateQuery: "WHERE (field_1 <= $1)",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "lte case with Enum struct",
			args: args{
				v: &GenericFilter{Lte: FakeIntTestEum},
			},
			expectedIntermediateQuery: "WHERE (field_1 <= $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "lte case with *Enum struct",
			args: args{
				v: &GenericFilter{Lte: starInterface(FakeIntTestEum)},
			},
			expectedIntermediateQuery: "WHERE (field_1 <= $1)",
			expectedArgs:              []driver.Value{1},
		},
		// NOT LTE
		{
			name: "not lte case with string",
			args: args{
				v: &GenericFilter{NotLte: 1},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 <= $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not lte case with *string",
			args: args{
				v: &GenericFilter{NotLte: starInterface("fake")},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 <= $1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not lte case with int",
			args: args{
				v: &GenericFilter{NotLte: 1},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 <= $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not lte case with *int",
			args: args{
				v: &GenericFilter{NotLte: starInterface(1)},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 <= $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not lte case with date",
			args: args{
				v: &GenericFilter{NotLte: now},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 <= $1)",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "not lte case with *date",
			args: args{
				v: &GenericFilter{NotLte: starInterface(now)},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 <= $1)",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "not lte case with Enum struct",
			args: args{
				v: &GenericFilter{NotLte: FakeIntTestEum},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 <= $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not lte case with *Enum struct",
			args: args{
				v: &GenericFilter{NotLte: starInterface(FakeIntTestEum)},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 <= $1)",
			expectedArgs:              []driver.Value{1},
		},
		// LT
		{
			name: "lt case with string",
			args: args{
				v: &GenericFilter{Lt: 1},
			},
			expectedIntermediateQuery: "WHERE (field_1 < $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "lt case with *string",
			args: args{
				v: &GenericFilter{Lt: starInterface("fake")},
			},
			expectedIntermediateQuery: "WHERE (field_1 < $1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "lt case with int",
			args: args{
				v: &GenericFilter{Lt: 1},
			},
			expectedIntermediateQuery: "WHERE (field_1 < $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "lt case with *int",
			args: args{
				v: &GenericFilter{Lt: starInterface(1)},
			},
			expectedIntermediateQuery: "WHERE (field_1 < $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "lt case with date",
			args: args{
				v: &GenericFilter{Lt: now},
			},
			expectedIntermediateQuery: "WHERE (field_1 < $1)",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "lt case with *date",
			args: args{
				v: &GenericFilter{Lt: starInterface(now)},
			},
			expectedIntermediateQuery: "WHERE (field_1 < $1)",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "lt case with Enum struct",
			args: args{
				v: &GenericFilter{Lt: FakeIntTestEum},
			},
			expectedIntermediateQuery: "WHERE (field_1 < $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "lt case with *Enum struct",
			args: args{
				v: &GenericFilter{Lt: starInterface(FakeIntTestEum)},
			},
			expectedIntermediateQuery: "WHERE (field_1 < $1)",
			expectedArgs:              []driver.Value{1},
		},
		// NOT LT
		{
			name: "not lt case with string",
			args: args{
				v: &GenericFilter{NotLt: 1},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 < $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not lt case with *string",
			args: args{
				v: &GenericFilter{NotLt: starInterface("fake")},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 < $1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not lt case with int",
			args: args{
				v: &GenericFilter{NotLt: 1},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 < $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not lt case with *int",
			args: args{
				v: &GenericFilter{NotLt: starInterface(1)},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 < $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not lt case with date",
			args: args{
				v: &GenericFilter{NotLt: now},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 < $1)",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "not lt case with *date",
			args: args{
				v: &GenericFilter{NotLt: starInterface(now)},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 < $1)",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "not lt case with Enum struct",
			args: args{
				v: &GenericFilter{NotLt: FakeIntTestEum},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 < $1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not lt case with *Enum struct",
			args: args{
				v: &GenericFilter{NotLt: starInterface(FakeIntTestEum)},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 < $1)",
			expectedArgs:              []driver.Value{1},
		},
		// CONTAINS
		{
			name: "contains case with *string",
			args: args{
				v: &GenericFilter{Contains: starString("fake")},
			},
			expectedIntermediateQuery: "WHERE (field_1 LIKE $1)",
			expectedArgs:              []driver.Value{"%fake%"},
		},
		{
			name: "contains case with *string",
			args: args{
				v: &GenericFilter{Contains: starString("fake")},
			},
			expectedIntermediateQuery: "WHERE (field_1 LIKE $1)",
			expectedArgs:              []driver.Value{"%fake%"},
		},
		{
			name: "contains case with Enum struct",
			args: args{
				v: &GenericFilter{Contains: FakeStringTestEnum},
			},
			expectedIntermediateQuery: "WHERE (field_1 LIKE $1)",
			expectedArgs:              []driver.Value{"%FAKE%"},
		},
		{
			name: "contains case with *Enum struct",
			args: args{
				v: &GenericFilter{Contains: starInterface(FakeStringTestEnum)},
			},
			wantErr:     true,
			errorString: "value must be a string or *string",
		},
		{
			name: "contains case with int",
			args: args{
				v: &GenericFilter{Contains: 1},
			},
			wantErr:     true,
			errorString: "value must be a string or *string",
		},
		{
			name: "contains case with *int",
			args: args{
				v: &GenericFilter{Contains: starInterface(1)},
			},
			wantErr:     true,
			errorString: "value must be a string or *string",
		},
		{
			name: "contains case with date",
			args: args{
				v: &GenericFilter{Contains: now},
			},
			wantErr:     true,
			errorString: "value must be a string or *string",
		},
		{
			name: "contains case with *date",
			args: args{
				v: &GenericFilter{Contains: starInterface(now)},
			},
			wantErr:     true,
			errorString: "value must be a string or *string",
		},
		{
			name: "contains case with bool",
			args: args{
				v: &GenericFilter{Contains: true},
			},
			wantErr:     true,
			errorString: "value must be a string or *string",
		},
		{
			name: "contains case with *bool",
			args: args{
				v: &GenericFilter{Contains: starInterface(true)},
			},
			wantErr:     true,
			errorString: "value must be a string or *string",
		},
		// NOT CONTAINS
		{
			name: "not contains case with *string",
			args: args{
				v: &GenericFilter{NotContains: starString("fake")},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 LIKE $1)",
			expectedArgs:              []driver.Value{"%fake%"},
		},
		{
			name: "not contains case with *string",
			args: args{
				v: &GenericFilter{NotContains: starString("fake")},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 LIKE $1)",
			expectedArgs:              []driver.Value{"%fake%"},
		},
		{
			name: "not contains case with Enum struct",
			args: args{
				v: &GenericFilter{NotContains: FakeStringTestEnum},
			},
			expectedIntermediateQuery: "WHERE NOT (field_1 LIKE $1)",
			expectedArgs:              []driver.Value{"%FAKE%"},
		},
		{
			name: "not contains case with *Enum struct",
			args: args{
				v: &GenericFilter{NotContains: starInterface(FakeStringTestEnum)},
			},
			wantErr:     true,
			errorString: "value must be a string or *string",
		},
		{
			name: "not contains case with int",
			args: args{
				v: &GenericFilter{NotContains: 1},
			},
			wantErr:     true,
			errorString: "value must be a string or *string",
		},
		{
			name: "not contains case with *int",
			args: args{
				v: &GenericFilter{NotContains: starInterface(1)},
			},
			wantErr:     true,
			errorString: "value must be a string or *string",
		},
		{
			name: "not contains case with date",
			args: args{
				v: &GenericFilter{NotContains: now},
			},
			wantErr:     true,
			errorString: "value must be a string or *string",
		},
		{
			name: "not contains case with *date",
			args: args{
				v: &GenericFilter{NotContains: starInterface(now)},
			},
			wantErr:     true,
			errorString: "value must be a string or *string",
		},
		{
			name: "not contains case with bool",
			args: args{
				v: &GenericFilter{NotContains: true},
			},
			wantErr:     true,
			errorString: "value must be a string or *string",
		},
		{
			name: "not contains case with *bool",
			args: args{
				v: &GenericFilter{NotContains: starInterface(true)},
			},
			wantErr:     true,
			errorString: "value must be a string or *string",
		},
		// IN
		{
			name: "in case with string",
			args: args{
				v: &GenericFilter{In: "fake"},
			},
			expectedIntermediateQuery: "WHERE (field_1 IN ($1))",
			expectedArgs:              []driver.Value{"fake"},
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

			db, err := gorm.Open("postgres", sqlDB)
			if err != nil {
				t.Error(err)
				return
			}
			db.LogMode(false)

			got, err := manageGenericFilter("field_1", tt.args.v, db)
			if (err != nil) != tt.wantErr {
				t.Errorf("manageGenericFilter() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && err.Error() != tt.errorString {
				t.Errorf("manageGenericFilter() error = %v, wantErr %v", err, tt.errorString)
				return
			}
			if err != nil {
				return
			}

			// Create expected query
			expectedQuery := `SELECT * FROM "people" ` + tt.expectedIntermediateQuery
			if tt.expectedIntermediateQuery != "" {
				expectedQuery += " "
			}
			expectedQuery += "LIMIT 1"

			mock.ExpectQuery(expectedQuery).
				WithArgs(tt.expectedArgs...).
				WillReturnRows(
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
