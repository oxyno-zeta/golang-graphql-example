//go:build unit

package common

import (
	"database/sql/driver"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Test_ManageFilter(t *testing.T) {
	starInterface := func(s interface{}) *interface{} { return &s }
	dateStr := "2020-09-19T23:10:35+02:00"
	date, err := time.Parse(time.RFC3339Nano, dateStr)
	if err != nil {
		t.Error(err)
		return
	}
	date = date.UTC()

	type Person struct {
		Name string
	}
	type Filter1 struct {
		Field1 *GenericFilter `dbfield:"field_1"`
	}
	type Filter2 struct {
		Field1 *GenericFilter `dbfield:"field_1"`
		Field2 *GenericFilter `dbfield:"-"`
	}
	type Filter3 struct {
		Field1 *GenericFilter `dbfield:"field_1"`
		Field2 *GenericFilter
	}
	type Filter4 struct {
		Field1 *GenericFilter `dbfield:"field_1"`
		Field2 string         `dbfield:"field_2"`
	}
	type Filter5 struct {
		Field1 *GenericFilter `dbfield:"field_1"`
		Field2 GenericFilter  `dbfield:"field_2"`
	}
	type Filter6 struct {
		OR     []*Filter6
		Field1 *GenericFilter `dbfield:"field_1"`
		Field2 *GenericFilter `dbfield:"field_2"`
	}
	type Filter7 struct {
		AND    []*Filter7
		Field1 *GenericFilter `dbfield:"field_1"`
		Field2 *GenericFilter `dbfield:"field_2"`
	}
	type Filter8 struct {
		AND    []*Filter8
		OR     []*Filter8
		Field1 *GenericFilter `dbfield:"field_1"`
		Field2 *GenericFilter `dbfield:"field_2"`
	}
	type Filter9 struct {
		OR     string
		Field1 *GenericFilter `dbfield:"field_1"`
		Field2 *GenericFilter `dbfield:"field_2"`
	}
	type Filter10 struct {
		OR     []string
		Field1 *GenericFilter `dbfield:"field_1"`
		Field2 *GenericFilter `dbfield:"field_2"`
	}
	type Filter11 struct {
		OR     []*Person
		Field1 *GenericFilter `dbfield:"field_1"`
		Field2 *GenericFilter `dbfield:"field_2"`
	}
	type Filter12 struct {
		AND    string
		Field1 *GenericFilter `dbfield:"field_1"`
		Field2 *GenericFilter `dbfield:"field_2"`
	}
	type Filter13 struct {
		AND    []string
		Field1 *GenericFilter `dbfield:"field_1"`
		Field2 *GenericFilter `dbfield:"field_2"`
	}
	type Filter14 struct {
		AND    []*Person
		Field1 *GenericFilter `dbfield:"field_1"`
		Field2 *GenericFilter `dbfield:"field_2"`
	}
	type Filter15 struct {
		Field1 *GenericFilter `dbfield:"field_1"`
		Field2 *Person        `dbfield:"field_2"`
	}
	type Filter16 struct {
		Field1 *DateFilter `dbfield:"field_1"`
	}
	type args struct {
		filter interface{}
	}
	tests := []struct {
		name                      string
		args                      args
		expectedIntermediateQuery string
		expectedArgs              []driver.Value
		wantErr                   bool
		errorString               string
	}{
		{
			name:        "wrong input",
			args:        args{filter: false},
			wantErr:     true,
			errorString: "filter must be an object",
		},
		{
			name: "nil sort object",
			args: args{
				filter: nil,
			},
			expectedIntermediateQuery: "",
		},
		{
			name: "date filter",
			args: args{
				filter: &Filter16{
					Field1: &DateFilter{Eq: &dateStr},
				},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1",
			expectedArgs:              []driver.Value{date},
		},
		{
			name: "one field",
			args: args{
				filter: &Filter1{
					Field1: &GenericFilter{
						Eq: starInterface("fake"),
					},
				},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "one field with upper value",
			args: args{
				filter: &Filter1{
					Field1: &GenericFilter{
						Eq:             starInterface("fake"),
						ValueUppercase: true,
					},
				},
			},
			expectedIntermediateQuery: "WHERE field_1 = upper($1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "one field with lower value",
			args: args{
				filter: &Filter1{
					Field1: &GenericFilter{
						Eq:             starInterface("fake"),
						ValueLowercase: true,
					},
				},
			},
			expectedIntermediateQuery: "WHERE field_1 = lower($1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "one field with lower field",
			args: args{
				filter: &Filter1{
					Field1: &GenericFilter{
						Eq:             starInterface("fake"),
						FieldLowercase: true,
					},
				},
			},
			expectedIntermediateQuery: "WHERE lower(field_1) = $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "one field with upper field",
			args: args{
				filter: &Filter1{
					Field1: &GenericFilter{
						Eq:             starInterface("fake"),
						FieldUppercase: true,
					},
				},
			},
			expectedIntermediateQuery: "WHERE upper(field_1) = $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "one field with lower and upper field but lower must be prioritized",
			args: args{
				filter: &Filter1{
					Field1: &GenericFilter{
						Eq:             starInterface("fake"),
						FieldLowercase: true,
						FieldUppercase: true,
					},
				},
			},
			expectedIntermediateQuery: "WHERE lower(field_1) = $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "one field with upper field and lower value",
			args: args{
				filter: &Filter1{
					Field1: &GenericFilter{
						Eq:             starInterface("fake"),
						FieldUppercase: true,
						ValueLowercase: true,
					},
				},
			},
			expectedIntermediateQuery: "WHERE upper(field_1) = lower($1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "one field with case insensitive",
			args: args{
				filter: &Filter1{
					Field1: &GenericFilter{
						In:              []string{"FAKE"},
						CaseInsensitive: true,
					},
				},
			},
			expectedIntermediateQuery: "WHERE lower(field_1) IN ($1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "one field IN with value lowercase",
			args: args{
				filter: &Filter1{
					Field1: &GenericFilter{
						In:             []string{"FAKE"},
						ValueLowercase: true,
					},
				},
			},
			expectedIntermediateQuery: "WHERE field_1 IN ($1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "one field IN with value uppercase",
			args: args{
				filter: &Filter1{
					Field1: &GenericFilter{
						In:             []string{"fake"},
						ValueUppercase: true,
					},
				},
			},
			expectedIntermediateQuery: "WHERE field_1 IN ($1)",
			expectedArgs:              []driver.Value{"FAKE"},
		},
		{
			name: "two fields and root level",
			args: args{
				filter: &Filter6{
					Field1: &GenericFilter{
						Eq: starInterface("fake"),
					},
					Field2: &GenericFilter{
						Eq: starInterface("fake2"),
					},
				},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1 AND field_2 = $2",
			expectedArgs:              []driver.Value{"fake", "fake2"},
		},
		{
			name: "two fields and root level with first one with upper value and second with lower field",
			args: args{
				filter: &Filter6{
					Field1: &GenericFilter{
						Eq:             starInterface("fake"),
						ValueUppercase: true,
					},
					Field2: &GenericFilter{
						Eq:             starInterface("fake2"),
						FieldLowercase: true,
					},
				},
			},
			expectedIntermediateQuery: "WHERE field_1 = upper($1) AND lower(field_2) = $2",
			expectedArgs:              []driver.Value{"fake", "fake2"},
		},
		{
			name: "two fields and root level with first one contains with upper value and second with lower field",
			args: args{
				filter: &Filter6{
					Field1: &GenericFilter{
						Contains:       "fake",
						ValueUppercase: true,
					},
					Field2: &GenericFilter{
						Eq:             starInterface("fake2"),
						FieldLowercase: true,
					},
				},
			},
			expectedIntermediateQuery: "WHERE field_1 LIKE upper($1) AND lower(field_2) = $2",
			expectedArgs:              []driver.Value{"%fake%", "fake2"},
		},
		{
			name: "one field with nil",
			args: args{
				filter: &Filter1{
					Field1: nil,
				},
			},
			expectedIntermediateQuery: "",
			expectedArgs:              []driver.Value{},
		},
		{
			name: "2 fields with one ignored (tag ignore)",
			args: args{
				filter: &Filter2{
					Field1: &GenericFilter{
						Eq: "fake",
					},
					Field2: &GenericFilter{
						Contains: "fak",
					},
				},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "2 fields with one ignored (no tag)",
			args: args{
				filter: &Filter3{
					Field1: &GenericFilter{
						Eq: "fake",
					},
					Field2: &GenericFilter{
						Contains: "fak",
					},
				},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "tag with wrong type",
			args: args{
				filter: &Filter4{
					Field1: &GenericFilter{
						Eq: "fake",
					},
					Field2: "fake",
				},
			},
			wantErr:     true,
			errorString: "field Field2 with filter tag must be a *GenericFilter or implement GenericFilterBuilder interface",
		},
		{
			name: "tag with wrong type 2",
			args: args{
				filter: &Filter5{
					Field1: &GenericFilter{
						Eq: "fake",
					},
					Field2: GenericFilter{
						Contains: "fak",
					},
				},
			},
			wantErr:     true,
			errorString: "field Field2 with filter tag must be a *GenericFilter or implement GenericFilterBuilder interface",
		},
		{
			name: "tag with wrong type 3 (struct pointer)",
			args: args{
				filter: &Filter15{
					Field1: &GenericFilter{Eq: "fake"},
					Field2: &Person{Name: "fake"},
				},
			},
			wantErr:     true,
			errorString: "field Field2 with filter tag must be a *GenericFilter or implement GenericFilterBuilder interface",
		},
		{
			name: "OR list without any root fields",
			args: args{
				filter: &Filter6{
					OR: []*Filter6{
						{
							Field1: &GenericFilter{Eq: "fake2"},
						},
						{
							Field1: &GenericFilter{Eq: "fake3"},
						},
					},
				},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1 OR field_1 = $2",
			expectedArgs:              []driver.Value{"fake2", "fake3"},
		},
		{
			name: "OR and root fields",
			args: args{
				filter: &Filter6{
					OR: []*Filter6{
						{
							Field1: &GenericFilter{Eq: "fake2"},
						},
						{
							Field1: &GenericFilter{Eq: "fake3"},
						},
					},
					Field1: &GenericFilter{
						Eq: "fake",
					},
					Field2: &GenericFilter{
						Contains: "fak",
					},
				},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1 AND field_2 LIKE $2 AND (field_1 = $3 OR field_1 = $4)",
			expectedArgs:              []driver.Value{"fake", "%fak%", "fake2", "fake3"},
		},
		{
			name: "OR with AND at root level for 1 side only",
			args: args{
				filter: &Filter8{
					OR: []*Filter8{
						{Field1: &GenericFilter{Eq: "fake1"}},
						{Field2: &GenericFilter{Eq: "fake2"}, Field1: &GenericFilter{Contains: "fake3"}},
					},
				},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1 OR (field_1 LIKE $2 AND field_2 = $3)",
			expectedArgs:              []driver.Value{"fake1", "%fake3%", "fake2"},
		},
		{
			name: "real AND with OR on right",
			args: args{
				filter: &Filter8{
					AND: []*Filter8{
						{
							OR: []*Filter8{
								{Field1: &GenericFilter{Eq: "fake1"}},
								{Field2: &GenericFilter{Eq: "fake2"}},
							},
						},
						{Field1: &GenericFilter{Contains: "fake3"}},
					},
				},
			},
			expectedIntermediateQuery: "WHERE (field_1 = $1 OR field_2 = $2) AND field_1 LIKE $3",
			expectedArgs:              []driver.Value{"fake1", "fake2", "%fake3%"},
		},
		{
			name: "OR cascade list with root fields on second level",
			args: args{
				filter: &Filter6{
					OR: []*Filter6{
						{
							Field1: &GenericFilter{Eq: "fake2"},
							OR: []*Filter6{
								{
									Field2: &GenericFilter{Eq: "fake"},
								},
								{
									Field2: &GenericFilter{Eq: "fake4"},
								},
							},
						},
						{
							Field1: &GenericFilter{Eq: "fake3"},
						},
					},
				},
			},
			expectedIntermediateQuery: "WHERE (field_1 = $1 AND (field_2 = $2 OR field_2 = $3)) OR field_1 = $4",
			expectedArgs:              []driver.Value{"fake2", "fake", "fake4", "fake3"},
		},
		{
			name: "OR cascade list without root fields on second level",
			args: args{
				filter: &Filter6{
					OR: []*Filter6{
						{
							OR: []*Filter6{
								{
									Field2: &GenericFilter{Eq: "fake"},
								},
								{
									Field2: &GenericFilter{Eq: "fake4"},
								},
							},
						},
						{
							Field1: &GenericFilter{Eq: "fake3"},
						},
					},
				},
			},
			expectedIntermediateQuery: "WHERE (field_2 = $1 OR field_2 = $2) OR field_1 = $3",
			expectedArgs:              []driver.Value{"fake", "fake4", "fake3"},
		},
		{
			name: "OR with an unsupported type (string) should be ignored",
			args: args{
				filter: &Filter9{
					OR:     "fake1",
					Field1: &GenericFilter{Eq: "fake"},
				},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "OR with an unsupported type ([]string) should be ignored",
			args: args{
				filter: &Filter10{
					OR:     []string{"fake1"},
					Field1: &GenericFilter{Eq: "fake"},
				},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "OR with an unsupported type ([]*Person) should be ignored",
			args: args{
				filter: &Filter11{
					OR:     []*Person{{Name: "fake1"}},
					Field1: &GenericFilter{Eq: "fake"},
				},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "AND and root fields",
			args: args{
				filter: &Filter7{
					AND: []*Filter7{
						{
							Field1: &GenericFilter{Eq: "fake2"},
						},
						{
							Field1: &GenericFilter{Eq: "fake3"},
						},
					},
					Field1: &GenericFilter{
						Eq: "fake",
					},
					Field2: &GenericFilter{
						Contains: "fak",
					},
				},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1 AND field_2 LIKE $2 AND (field_1 = $3 AND field_1 = $4)",
			expectedArgs:              []driver.Value{"fake", "%fak%", "fake2", "fake3"},
		},
		{
			name: "AND without root fields",
			args: args{
				filter: &Filter7{
					AND: []*Filter7{
						{
							Field1: &GenericFilter{Eq: "fake2"},
						},
						{
							Field1: &GenericFilter{Eq: "fake3"},
						},
					},
				},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1 AND field_1 = $2",
			expectedArgs:              []driver.Value{"fake2", "fake3"},
		},
		{
			name: "AND cascade list with root fields on second level",
			args: args{
				filter: &Filter7{
					AND: []*Filter7{
						{
							Field1: &GenericFilter{Eq: "fake2"},
							AND: []*Filter7{
								{
									Field2: &GenericFilter{Eq: "fake"},
								},
								{
									Field2: &GenericFilter{Eq: "fake4"},
								},
							},
						},
						{
							Field1: &GenericFilter{Eq: "fake3"},
						},
					},
				},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1 AND (field_2 = $2 AND field_2 = $3) AND field_1 = $4",
			expectedArgs:              []driver.Value{"fake2", "fake", "fake4", "fake3"},
		},
		{
			name: "AND cascade list without root fields on second level",
			args: args{
				filter: &Filter7{
					AND: []*Filter7{
						{
							AND: []*Filter7{
								{
									Field2: &GenericFilter{Eq: "fake"},
								},
								{
									Field2: &GenericFilter{Eq: "fake4"},
								},
							},
						},
						{
							Field1: &GenericFilter{Eq: "fake3"},
						},
					},
				},
			},
			expectedIntermediateQuery: "WHERE (field_2 = $1 AND field_2 = $2) AND field_1 = $3",
			expectedArgs:              []driver.Value{"fake", "fake4", "fake3"},
		},
		{
			name: "AND with an unsupported type (string) should be ignored",
			args: args{
				filter: &Filter12{
					AND:    "fake1",
					Field1: &GenericFilter{Eq: "fake"},
				},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "AND with an unsupported type ([]string) should be ignored",
			args: args{
				filter: &Filter13{
					AND:    []string{"fake1"},
					Field1: &GenericFilter{Eq: "fake"},
				},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "AND with an unsupported type ([]*Person) should be ignored",
			args: args{
				filter: &Filter14{
					AND:    []*Person{{Name: "fake1"}},
					Field1: &GenericFilter{Eq: "fake"},
				},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "AND and OR with root fields and OR with AND inside",
			args: args{
				filter: &Filter8{
					Field1: &GenericFilter{Eq: "fake11"},
					AND: []*Filter8{
						{
							AND: []*Filter8{
								{
									Field2: &GenericFilter{Eq: "fake1"},
								},
								{
									Field2: &GenericFilter{Eq: "fake2"},
								},
							},
						},
						{
							Field1: &GenericFilter{Eq: "fake3"},
						},
					},
					OR: []*Filter8{
						{
							Field2: &GenericFilter{Eq: "fake10"},
							OR: []*Filter8{
								{
									Field2: &GenericFilter{Eq: "fake4"},
								},
								{
									Field2: &GenericFilter{Eq: "fake5"},
								},
							},
						},
						{
							Field1: &GenericFilter{Eq: "fake6"},
						},
						{
							AND: []*Filter8{
								{
									Field1: &GenericFilter{Eq: "fake7"},
								},
								{
									OR: []*Filter8{
										{
											Field2: &GenericFilter{Eq: "fake8"},
										},
										{
											Field2: &GenericFilter{Eq: "fake9"},
											Field1: &GenericFilter{Eq: "fake10"},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1 AND ((field_2 = $2 AND field_2 = $3) AND field_1 = $4)" +
				" AND ((field_2 = $5 AND (field_2 = $6 OR field_2 = $7)) OR field_1 = $8" +
				" OR (field_1 = $9 AND (field_2 = $10 OR (field_1 = $11 AND field_2 = $12))))",
			expectedArgs: []driver.Value{
				"fake11",
				"fake1",
				"fake2",
				"fake3",
				"fake10",
				"fake4",
				"fake5",
				"fake6",
				"fake7",
				"fake8",
				"fake10",
				"fake9",
			},
		},
		{
			name: "AND and OR with root fields and OR with AND inside without any init root field",
			args: args{
				filter: &Filter8{
					AND: []*Filter8{
						{
							AND: []*Filter8{
								{
									Field2: &GenericFilter{Eq: "fake1"},
								},
								{
									Field2: &GenericFilter{Eq: "fake2"},
								},
							},
						},
						{
							Field1: &GenericFilter{Eq: "fake3"},
						},
					},
					OR: []*Filter8{
						{
							Field2: &GenericFilter{Eq: "fake10"},
							OR: []*Filter8{
								{
									Field2: &GenericFilter{Eq: "fake4"},
								},
								{
									Field2: &GenericFilter{Eq: "fake5"},
								},
							},
						},
						{
							Field1: &GenericFilter{Eq: "fake6"},
						},
						{
							AND: []*Filter8{
								{
									Field1: &GenericFilter{Eq: "fake7"},
								},
								{
									OR: []*Filter8{
										{
											Field2: &GenericFilter{Eq: "fake8"},
										},
										{
											Field2: &GenericFilter{Eq: "fake9"},
											Field1: &GenericFilter{Eq: "fake10"},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedIntermediateQuery: "WHERE ((field_2 = $1 AND field_2 = $2) AND field_1 = $3)" +
				" AND ((field_2 = $4 AND (field_2 = $5 OR field_2 = $6)) OR field_1 = $7" +
				" OR (field_1 = $8 AND (field_2 = $9 OR (field_1 = $10 AND field_2 = $11))))",
			expectedArgs: []driver.Value{
				"fake1",
				"fake2",
				"fake3",
				"fake10",
				"fake4",
				"fake5",
				"fake6",
				"fake7",
				"fake8",
				"fake10",
				"fake9",
			},
		},
		{
			name: "OR AND cascade",
			args: args{
				filter: &Filter8{
					OR: []*Filter8{
						{
							AND: []*Filter8{
								{Field1: &GenericFilter{Eq: "fake1"}},
								{Field1: &GenericFilter{Eq: "fake2"}},
							},
						},
						{Field2: &GenericFilter{Eq: "fake3"}},
						{
							AND: []*Filter8{
								{Field1: &GenericFilter{Eq: "fake4"}},
								{Field1: &GenericFilter{Eq: "fake5"}},
							},
						},
					},
				},
			},
			expectedIntermediateQuery: "WHERE (field_1 = $1 AND field_1 = $2) OR field_2 = $3 OR (field_1 = $4 AND field_1 = $5)",
			expectedArgs:              []driver.Value{"fake1", "fake2", "fake3", "fake4", "fake5"},
		},
		{
			name: "AND OR cascade",
			args: args{
				filter: &Filter8{
					AND: []*Filter8{
						{
							OR: []*Filter8{
								{Field1: &GenericFilter{Eq: "fake1"}},
								{Field1: &GenericFilter{Eq: "fake2"}},
							},
						},
						{Field2: &GenericFilter{Eq: "fake3"}},
						{
							OR: []*Filter8{
								{Field1: &GenericFilter{Eq: "fake4"}},
								{Field1: &GenericFilter{Eq: "fake5"}},
							},
						},
					},
				},
			},
			expectedIntermediateQuery: "WHERE (field_1 = $1 OR field_1 = $2) AND field_2 = $3 AND (field_1 = $4 OR field_1 = $5)",
			expectedArgs:              []driver.Value{"fake1", "fake2", "fake3", "fake4", "fake5"},
		},
		{
			name: "AND with only one AND inside",
			args: args{
				filter: &Filter8{
					AND: []*Filter8{
						{
							AND: []*Filter8{
								{Field1: &GenericFilter{Eq: "fake1"}},
								{Field1: &GenericFilter{Eq: "fake2"}},
							},
						},
					},
				},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1 AND field_1 = $2",
			expectedArgs:              []driver.Value{"fake1", "fake2"},
		},
		{
			name: "OR with only one AND inside",
			args: args{
				filter: &Filter8{
					OR: []*Filter8{
						{
							AND: []*Filter8{
								{Field1: &GenericFilter{Eq: "fake1"}},
								{Field1: &GenericFilter{Eq: "fake2"}},
							},
						},
					},
				},
			},
			expectedIntermediateQuery: "WHERE (field_1 = $1 AND field_1 = $2)",
			expectedArgs:              []driver.Value{"fake1", "fake2"},
		},
		{
			name: "AND with only one OR inside",
			args: args{
				filter: &Filter8{
					AND: []*Filter8{
						{
							OR: []*Filter8{
								{Field1: &GenericFilter{Eq: "fake1"}},
								{Field1: &GenericFilter{Eq: "fake2"}},
							},
						},
					},
				},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1 OR field_1 = $2",
			expectedArgs:              []driver.Value{"fake1", "fake2"},
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

			got, err := ManageFilter(tt.args.filter, db)
			if (err != nil) != tt.wantErr {
				t.Errorf("ManageFilter() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && err.Error() != tt.errorString {
				t.Errorf("ManageFilter() error = %v, wantErr %v", err, tt.errorString)
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
			expectedQuery += `ORDER BY "people"."name" LIMIT $` + strconv.Itoa(len(tt.expectedArgs)+1)
			// Add limit to expected args
			tt.expectedArgs = append(tt.expectedArgs, 1)

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

type StringTestEnum string

const FakeStringTestEnum StringTestEnum = "FAKE"

type IntTestEnum int

const FakeIntTestEum IntTestEnum = 1

func Test_manageFilterRequest(t *testing.T) {
	starInterface := func(s interface{}) interface{} { return &s }
	starString := func(s string) *string { return &s }
	now := time.Now()

	type Person struct {
		Name string
	}
	type args struct {
		v *GenericFilter
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
			expectedIntermediateQuery: "WHERE field_1 = $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "eq case with *string",
			args: args{
				v: &GenericFilter{Eq: starInterface("fake")},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "eq case with int",
			args: args{
				v: &GenericFilter{Eq: 1},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "eq case with *int",
			args: args{
				v: &GenericFilter{Eq: starInterface(1)},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "eq case with bool",
			args: args{
				v: &GenericFilter{Eq: true},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1",
			expectedArgs:              []driver.Value{true},
		},
		{
			name: "eq case with *bool",
			args: args{
				v: &GenericFilter{Eq: starInterface(true)},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1",
			expectedArgs:              []driver.Value{true},
		},
		{
			name: "eq case with date",
			args: args{
				v: &GenericFilter{Eq: now},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "eq case with *date",
			args: args{
				v: &GenericFilter{Eq: starInterface(now)},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "eq case with Enum struct",
			args: args{
				v: &GenericFilter{Eq: FakeStringTestEnum},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1",
			expectedArgs:              []driver.Value{"FAKE"},
		},
		{
			name: "eq case with *Enum struct",
			args: args{
				v: &GenericFilter{Eq: starInterface(FakeStringTestEnum)},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1",
			expectedArgs:              []driver.Value{"FAKE"},
		},
		{
			name: "eq case with upper value",
			args: args{
				v: &GenericFilter{Eq: "fake", ValueUppercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 = upper($1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "eq case with lower value",
			args: args{
				v: &GenericFilter{Eq: "fake", ValueLowercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 = lower($1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "eq case with upper field",
			args: args{
				v: &GenericFilter{Eq: "fake", FieldUppercase: true},
			},
			expectedIntermediateQuery: "WHERE upper(field_1) = $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "eq case with lower field",
			args: args{
				v: &GenericFilter{Eq: "fake", FieldLowercase: true},
			},
			expectedIntermediateQuery: "WHERE lower(field_1) = $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "eq case with case insensitive enabled",
			args: args{
				v: &GenericFilter{Eq: "fake", CaseInsensitive: true},
			},
			expectedIntermediateQuery: "WHERE lower(field_1) = lower($1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		// NOT EQ
		{
			name: "not eq case with string",
			args: args{
				v: &GenericFilter{NotEq: "fake"},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 = $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not eq case with *string",
			args: args{
				v: &GenericFilter{NotEq: starInterface("fake")},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 = $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not eq case with int",
			args: args{
				v: &GenericFilter{NotEq: 1},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 = $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not eq case with *int",
			args: args{
				v: &GenericFilter{NotEq: starInterface(1)},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 = $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not eq case with bool",
			args: args{
				v: &GenericFilter{NotEq: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 = $1",
			expectedArgs:              []driver.Value{true},
		},
		{
			name: "not eq case with *bool",
			args: args{
				v: &GenericFilter{NotEq: starInterface(true)},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 = $1",
			expectedArgs:              []driver.Value{true},
		},
		{
			name: "not eq case with date",
			args: args{
				v: &GenericFilter{NotEq: now},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 = $1",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "not eq case with *date",
			args: args{
				v: &GenericFilter{NotEq: starInterface(now)},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 = $1",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "not eq case with Enum struct",
			args: args{
				v: &GenericFilter{NotEq: FakeStringTestEnum},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 = $1",
			expectedArgs:              []driver.Value{"FAKE"},
		},
		{
			name: "not eq case with *Enum struct",
			args: args{
				v: &GenericFilter{NotEq: starInterface(FakeStringTestEnum)},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 = $1",
			expectedArgs:              []driver.Value{"FAKE"},
		},
		{
			name: "not eq case with upper value",
			args: args{
				v: &GenericFilter{NotEq: "fake", ValueUppercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 = upper($1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not eq case with lower value",
			args: args{
				v: &GenericFilter{NotEq: "fake", ValueLowercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 = lower($1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not eq case with upper field",
			args: args{
				v: &GenericFilter{NotEq: "fake", FieldUppercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT upper(field_1) = $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not eq case with lower field",
			args: args{
				v: &GenericFilter{NotEq: "fake", FieldLowercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT lower(field_1) = $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not eq case with case insensitive enabled",
			args: args{
				v: &GenericFilter{NotEq: "fake", CaseInsensitive: true},
			},
			expectedIntermediateQuery: "WHERE NOT lower(field_1) = lower($1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		// GTE
		{
			name: "gte case with string",
			args: args{
				v: &GenericFilter{Gte: 1},
			},
			expectedIntermediateQuery: "WHERE field_1 >= $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "gte case with *string",
			args: args{
				v: &GenericFilter{Gte: starInterface("fake")},
			},
			expectedIntermediateQuery: "WHERE field_1 >= $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "gte case with int",
			args: args{
				v: &GenericFilter{Gte: 1},
			},
			expectedIntermediateQuery: "WHERE field_1 >= $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "gte case with *int",
			args: args{
				v: &GenericFilter{Gte: starInterface(1)},
			},
			expectedIntermediateQuery: "WHERE field_1 >= $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "gte case with date",
			args: args{
				v: &GenericFilter{Gte: now},
			},
			expectedIntermediateQuery: "WHERE field_1 >= $1",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "gte case with *date",
			args: args{
				v: &GenericFilter{Gte: starInterface(now)},
			},
			expectedIntermediateQuery: "WHERE field_1 >= $1",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "gte case with Enum struct",
			args: args{
				v: &GenericFilter{Gte: FakeIntTestEum},
			},
			expectedIntermediateQuery: "WHERE field_1 >= $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "gte case with *Enum struct",
			args: args{
				v: &GenericFilter{Gte: starInterface(FakeIntTestEum)},
			},
			expectedIntermediateQuery: "WHERE field_1 >= $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "gte case with upper value must be ignored",
			args: args{
				v: &GenericFilter{Gte: "fake", ValueUppercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 >= $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "gte case with lower value must be ignored",
			args: args{
				v: &GenericFilter{Gte: "fake", ValueLowercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 >= $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "gte case with upper field must be ignored",
			args: args{
				v: &GenericFilter{Gte: "fake", FieldUppercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 >= $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "gte case with lower field must be ignored",
			args: args{
				v: &GenericFilter{Gte: "fake", FieldLowercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 >= $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "gte case with case insensitive must be ignored",
			args: args{
				v: &GenericFilter{Gte: "fake", CaseInsensitive: true},
			},
			expectedIntermediateQuery: "WHERE field_1 >= $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		// NOT GTE
		{
			name: "not gte case with string",
			args: args{
				v: &GenericFilter{NotGte: 1},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 >= $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not gte case with *string",
			args: args{
				v: &GenericFilter{NotGte: starInterface("fake")},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 >= $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not gte case with int",
			args: args{
				v: &GenericFilter{NotGte: 1},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 >= $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not gte case with *int",
			args: args{
				v: &GenericFilter{NotGte: starInterface(1)},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 >= $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not gte case with date",
			args: args{
				v: &GenericFilter{NotGte: now},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 >= $1",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "not gte case with *date",
			args: args{
				v: &GenericFilter{NotGte: starInterface(now)},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 >= $1",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "not gte case with Enum struct",
			args: args{
				v: &GenericFilter{NotGte: FakeIntTestEum},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 >= $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not gte case with *Enum struct",
			args: args{
				v: &GenericFilter{NotGte: starInterface(FakeIntTestEum)},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 >= $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not gte case with upper value must be ignored",
			args: args{
				v: &GenericFilter{NotGte: "fake", ValueUppercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 >= $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not gte case with lower value must be ignored",
			args: args{
				v: &GenericFilter{NotGte: "fake", ValueLowercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 >= $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not gte case with upper field must be ignored",
			args: args{
				v: &GenericFilter{NotGte: "fake", FieldUppercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 >= $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not gte case with lower field must be ignored",
			args: args{
				v: &GenericFilter{NotGte: "fake", FieldLowercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 >= $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not gte case with case insensitive must be ignored",
			args: args{
				v: &GenericFilter{NotGte: "fake", CaseInsensitive: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 >= $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		// GT
		{
			name: "gt case with string",
			args: args{
				v: &GenericFilter{Gt: 1},
			},
			expectedIntermediateQuery: "WHERE field_1 > $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "gt case with *string",
			args: args{
				v: &GenericFilter{Gt: starInterface("fake")},
			},
			expectedIntermediateQuery: "WHERE field_1 > $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "gt case with int",
			args: args{
				v: &GenericFilter{Gt: 1},
			},
			expectedIntermediateQuery: "WHERE field_1 > $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "gt case with *int",
			args: args{
				v: &GenericFilter{Gt: starInterface(1)},
			},
			expectedIntermediateQuery: "WHERE field_1 > $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "gt case with date",
			args: args{
				v: &GenericFilter{Gt: now},
			},
			expectedIntermediateQuery: "WHERE field_1 > $1",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "gt case with *date",
			args: args{
				v: &GenericFilter{Gt: starInterface(now)},
			},
			expectedIntermediateQuery: "WHERE field_1 > $1",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "gt case with Enum struct",
			args: args{
				v: &GenericFilter{Gt: FakeIntTestEum},
			},
			expectedIntermediateQuery: "WHERE field_1 > $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "gt case with *Enum struct",
			args: args{
				v: &GenericFilter{Gt: starInterface(FakeIntTestEum)},
			},
			expectedIntermediateQuery: "WHERE field_1 > $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "gt case with upper value must be ignored",
			args: args{
				v: &GenericFilter{Gt: "fake", ValueUppercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 > $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "gt case with lower value must be ignored",
			args: args{
				v: &GenericFilter{Gt: "fake", ValueLowercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 > $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "gt case with upper field must be ignored",
			args: args{
				v: &GenericFilter{Gt: "fake", FieldUppercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 > $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "gt case with lower field must be ignored",
			args: args{
				v: &GenericFilter{Gt: "fake", FieldLowercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 > $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "gt case with case insensitive must be ignored",
			args: args{
				v: &GenericFilter{Gt: "fake", CaseInsensitive: true},
			},
			expectedIntermediateQuery: "WHERE field_1 > $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		// NOT GT
		{
			name: "not gt case with string",
			args: args{
				v: &GenericFilter{NotGt: 1},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 > $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not gt case with *string",
			args: args{
				v: &GenericFilter{NotGt: starInterface("fake")},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 > $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not gt case with int",
			args: args{
				v: &GenericFilter{NotGt: 1},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 > $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not gt case with *int",
			args: args{
				v: &GenericFilter{NotGt: starInterface(1)},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 > $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not gt case with date",
			args: args{
				v: &GenericFilter{NotGt: now},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 > $1",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "not gt case with *date",
			args: args{
				v: &GenericFilter{NotGt: starInterface(now)},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 > $1",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "not gt case with Enum struct",
			args: args{
				v: &GenericFilter{NotGt: FakeIntTestEum},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 > $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not gt case with *Enum struct",
			args: args{
				v: &GenericFilter{NotGt: starInterface(FakeIntTestEum)},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 > $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not gt case with upper value must be ignored",
			args: args{
				v: &GenericFilter{NotGt: "fake", ValueUppercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 > $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not gt case with lower value must be ignored",
			args: args{
				v: &GenericFilter{NotGt: "fake", ValueLowercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 > $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not gt case with upper field must be ignored",
			args: args{
				v: &GenericFilter{NotGt: "fake", FieldUppercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 > $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not gt case with lower field must be ignored",
			args: args{
				v: &GenericFilter{NotGt: "fake", FieldLowercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 > $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not gt case with case insensitive must be ignored",
			args: args{
				v: &GenericFilter{NotGt: "fake", CaseInsensitive: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 > $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		// LTE
		{
			name: "lte case with string",
			args: args{
				v: &GenericFilter{Lte: 1},
			},
			expectedIntermediateQuery: "WHERE field_1 <= $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "lte case with *string",
			args: args{
				v: &GenericFilter{Lte: starInterface("fake")},
			},
			expectedIntermediateQuery: "WHERE field_1 <= $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "lte case with int",
			args: args{
				v: &GenericFilter{Lte: 1},
			},
			expectedIntermediateQuery: "WHERE field_1 <= $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "lte case with *int",
			args: args{
				v: &GenericFilter{Lte: starInterface(1)},
			},
			expectedIntermediateQuery: "WHERE field_1 <= $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "lte case with date",
			args: args{
				v: &GenericFilter{Lte: now},
			},
			expectedIntermediateQuery: "WHERE field_1 <= $1",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "lte case with *date",
			args: args{
				v: &GenericFilter{Lte: starInterface(now)},
			},
			expectedIntermediateQuery: "WHERE field_1 <= $1",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "lte case with Enum struct",
			args: args{
				v: &GenericFilter{Lte: FakeIntTestEum},
			},
			expectedIntermediateQuery: "WHERE field_1 <= $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "lte case with *Enum struct",
			args: args{
				v: &GenericFilter{Lte: starInterface(FakeIntTestEum)},
			},
			expectedIntermediateQuery: "WHERE field_1 <= $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "lte case with upper value must be ignored",
			args: args{
				v: &GenericFilter{Lte: "fake", ValueUppercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 <= $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "lte case with lower value must be ignored",
			args: args{
				v: &GenericFilter{Lte: "fake", ValueLowercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 <= $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "lte case with upper field must be ignored",
			args: args{
				v: &GenericFilter{Lte: "fake", FieldUppercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 <= $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "lte case with lower field must be ignored",
			args: args{
				v: &GenericFilter{Lte: "fake", FieldLowercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 <= $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "lte case with case insensitive must be ignored",
			args: args{
				v: &GenericFilter{Lte: "fake", CaseInsensitive: true},
			},
			expectedIntermediateQuery: "WHERE field_1 <= $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		// NOT LTE
		{
			name: "not lte case with string",
			args: args{
				v: &GenericFilter{NotLte: 1},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 <= $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not lte case with *string",
			args: args{
				v: &GenericFilter{NotLte: starInterface("fake")},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 <= $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not lte case with int",
			args: args{
				v: &GenericFilter{NotLte: 1},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 <= $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not lte case with *int",
			args: args{
				v: &GenericFilter{NotLte: starInterface(1)},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 <= $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not lte case with date",
			args: args{
				v: &GenericFilter{NotLte: now},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 <= $1",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "not lte case with *date",
			args: args{
				v: &GenericFilter{NotLte: starInterface(now)},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 <= $1",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "not lte case with Enum struct",
			args: args{
				v: &GenericFilter{NotLte: FakeIntTestEum},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 <= $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not lte case with *Enum struct",
			args: args{
				v: &GenericFilter{NotLte: starInterface(FakeIntTestEum)},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 <= $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not lte case with upper value must be ignored",
			args: args{
				v: &GenericFilter{NotLte: "fake", ValueUppercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 <= $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not lte case with lower value must be ignored",
			args: args{
				v: &GenericFilter{NotLte: "fake", ValueLowercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 <= $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not lte case with upper field must be ignored",
			args: args{
				v: &GenericFilter{NotLte: "fake", FieldUppercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 <= $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not lte case with lower field must be ignored",
			args: args{
				v: &GenericFilter{NotLte: "fake", FieldLowercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 <= $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not lte case with case insensitive must be ignored",
			args: args{
				v: &GenericFilter{NotLte: "fake", CaseInsensitive: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 <= $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		// LT
		{
			name: "lt case with string",
			args: args{
				v: &GenericFilter{Lt: 1},
			},
			expectedIntermediateQuery: "WHERE field_1 < $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "lt case with *string",
			args: args{
				v: &GenericFilter{Lt: starInterface("fake")},
			},
			expectedIntermediateQuery: "WHERE field_1 < $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "lt case with int",
			args: args{
				v: &GenericFilter{Lt: 1},
			},
			expectedIntermediateQuery: "WHERE field_1 < $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "lt case with *int",
			args: args{
				v: &GenericFilter{Lt: starInterface(1)},
			},
			expectedIntermediateQuery: "WHERE field_1 < $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "lt case with date",
			args: args{
				v: &GenericFilter{Lt: now},
			},
			expectedIntermediateQuery: "WHERE field_1 < $1",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "lt case with *date",
			args: args{
				v: &GenericFilter{Lt: starInterface(now)},
			},
			expectedIntermediateQuery: "WHERE field_1 < $1",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "lt case with Enum struct",
			args: args{
				v: &GenericFilter{Lt: FakeIntTestEum},
			},
			expectedIntermediateQuery: "WHERE field_1 < $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "lt case with *Enum struct",
			args: args{
				v: &GenericFilter{Lt: starInterface(FakeIntTestEum)},
			},
			expectedIntermediateQuery: "WHERE field_1 < $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "lt case with upper value must be ignored",
			args: args{
				v: &GenericFilter{Lt: "fake", ValueUppercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 < $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "lt case with lower value must be ignored",
			args: args{
				v: &GenericFilter{Lt: "fake", ValueLowercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 < $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "lt case with upper field must be ignored",
			args: args{
				v: &GenericFilter{Lt: "fake", FieldUppercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 < $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "lt case with lower field must be ignored",
			args: args{
				v: &GenericFilter{Lt: "fake", FieldLowercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 < $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "lt case with case insensitive must be ignored",
			args: args{
				v: &GenericFilter{Lt: "fake", CaseInsensitive: true},
			},
			expectedIntermediateQuery: "WHERE field_1 < $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		// NOT LT
		{
			name: "not lt case with string",
			args: args{
				v: &GenericFilter{NotLt: 1},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 < $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not lt case with *string",
			args: args{
				v: &GenericFilter{NotLt: starInterface("fake")},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 < $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not lt case with int",
			args: args{
				v: &GenericFilter{NotLt: 1},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 < $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not lt case with *int",
			args: args{
				v: &GenericFilter{NotLt: starInterface(1)},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 < $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not lt case with date",
			args: args{
				v: &GenericFilter{NotLt: now},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 < $1",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "not lt case with *date",
			args: args{
				v: &GenericFilter{NotLt: starInterface(now)},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 < $1",
			expectedArgs:              []driver.Value{now},
		},
		{
			name: "not lt case with Enum struct",
			args: args{
				v: &GenericFilter{NotLt: FakeIntTestEum},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 < $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not lt case with *Enum struct",
			args: args{
				v: &GenericFilter{NotLt: starInterface(FakeIntTestEum)},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 < $1",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not lt case with upper value must be ignored",
			args: args{
				v: &GenericFilter{NotLt: "fake", ValueUppercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 < $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not lt case with lower value must be ignored",
			args: args{
				v: &GenericFilter{NotLt: "fake", ValueLowercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 < $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not lt case with upper field must be ignored",
			args: args{
				v: &GenericFilter{NotLt: "fake", FieldUppercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 < $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not lt case with lower field must be ignored",
			args: args{
				v: &GenericFilter{NotLt: "fake", FieldLowercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 < $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not lt case with case insensitive must be ignored",
			args: args{
				v: &GenericFilter{NotLt: "fake", CaseInsensitive: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 < $1",
			expectedArgs:              []driver.Value{"fake"},
		},
		// CONTAINS
		{
			name: "contains case with *string",
			args: args{
				v: &GenericFilter{Contains: starString("fake")},
			},
			expectedIntermediateQuery: "WHERE field_1 LIKE $1",
			expectedArgs:              []driver.Value{"%fake%"},
		},
		{
			name: "contains case with *string",
			args: args{
				v: &GenericFilter{Contains: starString("fake")},
			},
			expectedIntermediateQuery: "WHERE field_1 LIKE $1",
			expectedArgs:              []driver.Value{"%fake%"},
		},
		{
			name: "contains case with Enum struct",
			args: args{
				v: &GenericFilter{Contains: FakeStringTestEnum},
			},
			expectedIntermediateQuery: "WHERE field_1 LIKE $1",
			expectedArgs:              []driver.Value{"%FAKE%"},
		},
		{
			name: "contains case with *Enum struct",
			args: args{
				v: &GenericFilter{Contains: starInterface(FakeStringTestEnum)},
			},
			wantErr:     true,
			errorString: "contains value must be a string or *string",
		},
		{
			name: "contains case with int",
			args: args{
				v: &GenericFilter{Contains: 1},
			},
			wantErr:     true,
			errorString: "contains value must be a string or *string",
		},
		{
			name: "contains case with *int",
			args: args{
				v: &GenericFilter{Contains: starInterface(1)},
			},
			wantErr:     true,
			errorString: "contains value must be a string or *string",
		},
		{
			name: "contains case with date",
			args: args{
				v: &GenericFilter{Contains: now},
			},
			wantErr:     true,
			errorString: "contains value must be a string or *string",
		},
		{
			name: "contains case with *date",
			args: args{
				v: &GenericFilter{Contains: starInterface(now)},
			},
			wantErr:     true,
			errorString: "contains value must be a string or *string",
		},
		{
			name: "contains case with bool",
			args: args{
				v: &GenericFilter{Contains: true},
			},
			wantErr:     true,
			errorString: "contains value must be a string or *string",
		},
		{
			name: "contains case with *bool",
			args: args{
				v: &GenericFilter{Contains: starInterface(true)},
			},
			wantErr:     true,
			errorString: "contains value must be a string or *string",
		},
		{
			name: "contains case with upper value",
			args: args{
				v: &GenericFilter{Contains: "fake", ValueUppercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 LIKE upper($1)",
			expectedArgs:              []driver.Value{"%fake%"},
		},
		{
			name: "contains case with lower value",
			args: args{
				v: &GenericFilter{Contains: "fake", ValueLowercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 LIKE lower($1)",
			expectedArgs:              []driver.Value{"%fake%"},
		},
		{
			name: "contains case with upper field",
			args: args{
				v: &GenericFilter{Contains: "fake", FieldUppercase: true},
			},
			expectedIntermediateQuery: "WHERE upper(field_1) LIKE $1",
			expectedArgs:              []driver.Value{"%fake%"},
		},
		{
			name: "contains case with lower field",
			args: args{
				v: &GenericFilter{Contains: "fake", FieldLowercase: true},
			},
			expectedIntermediateQuery: "WHERE lower(field_1) LIKE $1",
			expectedArgs:              []driver.Value{"%fake%"},
		},
		{
			name: "contains case with case insensitive",
			args: args{
				v: &GenericFilter{Contains: "fake", CaseInsensitive: true},
			},
			expectedIntermediateQuery: "WHERE lower(field_1) LIKE lower($1)",
			expectedArgs:              []driver.Value{"%fake%"},
		},
		// NOT CONTAINS
		{
			name: "not contains case with *string",
			args: args{
				v: &GenericFilter{NotContains: starString("fake")},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 LIKE $1",
			expectedArgs:              []driver.Value{"%fake%"},
		},
		{
			name: "not contains case with *string",
			args: args{
				v: &GenericFilter{NotContains: starString("fake")},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 LIKE $1",
			expectedArgs:              []driver.Value{"%fake%"},
		},
		{
			name: "not contains case with Enum struct",
			args: args{
				v: &GenericFilter{NotContains: FakeStringTestEnum},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 LIKE $1",
			expectedArgs:              []driver.Value{"%FAKE%"},
		},
		{
			name: "not contains case with *Enum struct",
			args: args{
				v: &GenericFilter{NotContains: starInterface(FakeStringTestEnum)},
			},
			wantErr:     true,
			errorString: "notContains value must be a string or *string",
		},
		{
			name: "not contains case with int",
			args: args{
				v: &GenericFilter{NotContains: 1},
			},
			wantErr:     true,
			errorString: "notContains value must be a string or *string",
		},
		{
			name: "not contains case with *int",
			args: args{
				v: &GenericFilter{NotContains: starInterface(1)},
			},
			wantErr:     true,
			errorString: "notContains value must be a string or *string",
		},
		{
			name: "not contains case with date",
			args: args{
				v: &GenericFilter{NotContains: now},
			},
			wantErr:     true,
			errorString: "notContains value must be a string or *string",
		},
		{
			name: "not contains case with *date",
			args: args{
				v: &GenericFilter{NotContains: starInterface(now)},
			},
			wantErr:     true,
			errorString: "notContains value must be a string or *string",
		},
		{
			name: "not contains case with bool",
			args: args{
				v: &GenericFilter{NotContains: true},
			},
			wantErr:     true,
			errorString: "notContains value must be a string or *string",
		},
		{
			name: "not contains case with *bool",
			args: args{
				v: &GenericFilter{NotContains: starInterface(true)},
			},
			wantErr:     true,
			errorString: "notContains value must be a string or *string",
		},
		{
			name: "not contains case with upper value",
			args: args{
				v: &GenericFilter{NotContains: "fake", ValueUppercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 LIKE upper($1)",
			expectedArgs:              []driver.Value{"%fake%"},
		},
		{
			name: "not contains case with lower value",
			args: args{
				v: &GenericFilter{NotContains: "fake", ValueLowercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 LIKE lower($1)",
			expectedArgs:              []driver.Value{"%fake%"},
		},
		{
			name: "not contains case with upper field",
			args: args{
				v: &GenericFilter{NotContains: "fake", FieldUppercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT upper(field_1) LIKE $1",
			expectedArgs:              []driver.Value{"%fake%"},
		},
		{
			name: "not contains case with lower field",
			args: args{
				v: &GenericFilter{NotContains: "fake", FieldLowercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT lower(field_1) LIKE $1",
			expectedArgs:              []driver.Value{"%fake%"},
		},
		{
			name: "not contains case with case insensitive",
			args: args{
				v: &GenericFilter{NotContains: "fake", CaseInsensitive: true},
			},
			expectedIntermediateQuery: "WHERE NOT lower(field_1) LIKE lower($1)",
			expectedArgs:              []driver.Value{"%fake%"},
		},
		// STARTS WITH
		{
			name: "starts with case with *string",
			args: args{
				v: &GenericFilter{StartsWith: starString("fake")},
			},
			expectedIntermediateQuery: "WHERE field_1 LIKE $1",
			expectedArgs:              []driver.Value{"fake%"},
		},
		{
			name: "starts with case with *string",
			args: args{
				v: &GenericFilter{StartsWith: starString("fake")},
			},
			expectedIntermediateQuery: "WHERE field_1 LIKE $1",
			expectedArgs:              []driver.Value{"fake%"},
		},
		{
			name: "starts with case with Enum struct",
			args: args{
				v: &GenericFilter{StartsWith: FakeStringTestEnum},
			},
			expectedIntermediateQuery: "WHERE field_1 LIKE $1",
			expectedArgs:              []driver.Value{"FAKE%"},
		},
		{
			name: "starts with case with *Enum struct",
			args: args{
				v: &GenericFilter{StartsWith: starInterface(FakeStringTestEnum)},
			},
			wantErr:     true,
			errorString: "startsWith value must be a string or *string",
		},
		{
			name: "starts with case with int",
			args: args{
				v: &GenericFilter{StartsWith: 1},
			},
			wantErr:     true,
			errorString: "startsWith value must be a string or *string",
		},
		{
			name: "starts with case with *int",
			args: args{
				v: &GenericFilter{StartsWith: starInterface(1)},
			},
			wantErr:     true,
			errorString: "startsWith value must be a string or *string",
		},
		{
			name: "starts with case with date",
			args: args{
				v: &GenericFilter{StartsWith: now},
			},
			wantErr:     true,
			errorString: "startsWith value must be a string or *string",
		},
		{
			name: "starts with case with *date",
			args: args{
				v: &GenericFilter{StartsWith: starInterface(now)},
			},
			wantErr:     true,
			errorString: "startsWith value must be a string or *string",
		},
		{
			name: "starts with case with bool",
			args: args{
				v: &GenericFilter{StartsWith: true},
			},
			wantErr:     true,
			errorString: "startsWith value must be a string or *string",
		},
		{
			name: "starts with case with *bool",
			args: args{
				v: &GenericFilter{StartsWith: starInterface(true)},
			},
			wantErr:     true,
			errorString: "startsWith value must be a string or *string",
		},
		{
			name: "starts with case with upper value",
			args: args{
				v: &GenericFilter{StartsWith: "fake", ValueUppercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 LIKE upper($1)",
			expectedArgs:              []driver.Value{"fake%"},
		},
		{
			name: "starts with case with lower value",
			args: args{
				v: &GenericFilter{StartsWith: "fake", ValueLowercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 LIKE lower($1)",
			expectedArgs:              []driver.Value{"fake%"},
		},
		{
			name: "starts with case with upper field",
			args: args{
				v: &GenericFilter{StartsWith: "fake", FieldUppercase: true},
			},
			expectedIntermediateQuery: "WHERE upper(field_1) LIKE $1",
			expectedArgs:              []driver.Value{"fake%"},
		},
		{
			name: "starts with case with lower field",
			args: args{
				v: &GenericFilter{StartsWith: "fake", FieldLowercase: true},
			},
			expectedIntermediateQuery: "WHERE lower(field_1) LIKE $1",
			expectedArgs:              []driver.Value{"fake%"},
		},
		{
			name: "starts with case with case insensitive",
			args: args{
				v: &GenericFilter{StartsWith: "fake", CaseInsensitive: true},
			},
			expectedIntermediateQuery: "WHERE lower(field_1) LIKE lower($1)",
			expectedArgs:              []driver.Value{"fake%"},
		},
		// NOT STARTS WITH
		{
			name: "not starts with case with *string",
			args: args{
				v: &GenericFilter{NotStartsWith: starString("fake")},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 LIKE $1",
			expectedArgs:              []driver.Value{"fake%"},
		},
		{
			name: "not starts with case with *string",
			args: args{
				v: &GenericFilter{NotStartsWith: starString("fake")},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 LIKE $1",
			expectedArgs:              []driver.Value{"fake%"},
		},
		{
			name: "not starts with case with Enum struct",
			args: args{
				v: &GenericFilter{NotStartsWith: FakeStringTestEnum},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 LIKE $1",
			expectedArgs:              []driver.Value{"FAKE%"},
		},
		{
			name: "not starts with case with *Enum struct",
			args: args{
				v: &GenericFilter{NotStartsWith: starInterface(FakeStringTestEnum)},
			},
			wantErr:     true,
			errorString: "notStartsWith value must be a string or *string",
		},
		{
			name: "not starts with case with int",
			args: args{
				v: &GenericFilter{NotStartsWith: 1},
			},
			wantErr:     true,
			errorString: "notStartsWith value must be a string or *string",
		},
		{
			name: "not starts with case with *int",
			args: args{
				v: &GenericFilter{NotStartsWith: starInterface(1)},
			},
			wantErr:     true,
			errorString: "notStartsWith value must be a string or *string",
		},
		{
			name: "not starts with case with date",
			args: args{
				v: &GenericFilter{NotStartsWith: now},
			},
			wantErr:     true,
			errorString: "notStartsWith value must be a string or *string",
		},
		{
			name: "not starts with case with *date",
			args: args{
				v: &GenericFilter{NotStartsWith: starInterface(now)},
			},
			wantErr:     true,
			errorString: "notStartsWith value must be a string or *string",
		},
		{
			name: "not starts with case with bool",
			args: args{
				v: &GenericFilter{NotStartsWith: true},
			},
			wantErr:     true,
			errorString: "notStartsWith value must be a string or *string",
		},
		{
			name: "not starts with case with *bool",
			args: args{
				v: &GenericFilter{NotStartsWith: starInterface(true)},
			},
			wantErr:     true,
			errorString: "notStartsWith value must be a string or *string",
		},
		{
			name: "not starts with case with upper value",
			args: args{
				v: &GenericFilter{NotStartsWith: "fake", ValueUppercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 LIKE upper($1)",
			expectedArgs:              []driver.Value{"fake%"},
		},
		{
			name: "not starts with case with lower value",
			args: args{
				v: &GenericFilter{NotStartsWith: "fake", ValueLowercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 LIKE lower($1)",
			expectedArgs:              []driver.Value{"fake%"},
		},
		{
			name: "not starts with case with upper field",
			args: args{
				v: &GenericFilter{NotStartsWith: "fake", FieldUppercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT upper(field_1) LIKE $1",
			expectedArgs:              []driver.Value{"fake%"},
		},
		{
			name: "not starts with case with lower field",
			args: args{
				v: &GenericFilter{NotStartsWith: "fake", FieldLowercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT lower(field_1) LIKE $1",
			expectedArgs:              []driver.Value{"fake%"},
		},
		{
			name: "not starts with case with case insensitive",
			args: args{
				v: &GenericFilter{NotStartsWith: "fake", CaseInsensitive: true},
			},
			expectedIntermediateQuery: "WHERE NOT lower(field_1) LIKE lower($1)",
			expectedArgs:              []driver.Value{"fake%"},
		},
		// ENDS WITH
		{
			name: "ends with case with *string",
			args: args{
				v: &GenericFilter{EndsWith: starString("fake")},
			},
			expectedIntermediateQuery: "WHERE field_1 LIKE $1",
			expectedArgs:              []driver.Value{"%fake"},
		},
		{
			name: "ends with case with *string",
			args: args{
				v: &GenericFilter{EndsWith: starString("fake")},
			},
			expectedIntermediateQuery: "WHERE field_1 LIKE $1",
			expectedArgs:              []driver.Value{"%fake"},
		},
		{
			name: "ends with case with Enum struct",
			args: args{
				v: &GenericFilter{EndsWith: FakeStringTestEnum},
			},
			expectedIntermediateQuery: "WHERE field_1 LIKE $1",
			expectedArgs:              []driver.Value{"%FAKE"},
		},
		{
			name: "ends with case with *Enum struct",
			args: args{
				v: &GenericFilter{EndsWith: starInterface(FakeStringTestEnum)},
			},
			wantErr:     true,
			errorString: "endsWith value must be a string or *string",
		},
		{
			name: "ends with case with int",
			args: args{
				v: &GenericFilter{EndsWith: 1},
			},
			wantErr:     true,
			errorString: "endsWith value must be a string or *string",
		},
		{
			name: "ends with case with *int",
			args: args{
				v: &GenericFilter{EndsWith: starInterface(1)},
			},
			wantErr:     true,
			errorString: "endsWith value must be a string or *string",
		},
		{
			name: "ends with case with date",
			args: args{
				v: &GenericFilter{EndsWith: now},
			},
			wantErr:     true,
			errorString: "endsWith value must be a string or *string",
		},
		{
			name: "ends with case with *date",
			args: args{
				v: &GenericFilter{EndsWith: starInterface(now)},
			},
			wantErr:     true,
			errorString: "endsWith value must be a string or *string",
		},
		{
			name: "ends with case with bool",
			args: args{
				v: &GenericFilter{EndsWith: true},
			},
			wantErr:     true,
			errorString: "endsWith value must be a string or *string",
		},
		{
			name: "ends with case with *bool",
			args: args{
				v: &GenericFilter{EndsWith: starInterface(true)},
			},
			wantErr:     true,
			errorString: "endsWith value must be a string or *string",
		},
		{
			name: "ends with case with upper value",
			args: args{
				v: &GenericFilter{EndsWith: "fake", ValueUppercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 LIKE upper($1)",
			expectedArgs:              []driver.Value{"%fake"},
		},
		{
			name: "ends with case with lower value",
			args: args{
				v: &GenericFilter{EndsWith: "fake", ValueLowercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 LIKE lower($1)",
			expectedArgs:              []driver.Value{"%fake"},
		},
		{
			name: "ends with case with upper field",
			args: args{
				v: &GenericFilter{EndsWith: "fake", FieldUppercase: true},
			},
			expectedIntermediateQuery: "WHERE upper(field_1) LIKE $1",
			expectedArgs:              []driver.Value{"%fake"},
		},
		{
			name: "ends with case with lower field",
			args: args{
				v: &GenericFilter{EndsWith: "fake", FieldLowercase: true},
			},
			expectedIntermediateQuery: "WHERE lower(field_1) LIKE $1",
			expectedArgs:              []driver.Value{"%fake"},
		},
		{
			name: "ends with case with case insensitive",
			args: args{
				v: &GenericFilter{EndsWith: "fake", CaseInsensitive: true},
			},
			expectedIntermediateQuery: "WHERE lower(field_1) LIKE lower($1)",
			expectedArgs:              []driver.Value{"%fake"},
		},
		// NOT ENDS WITH
		{
			name: "not ends with case with *string",
			args: args{
				v: &GenericFilter{NotEndsWith: starString("fake")},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 LIKE $1",
			expectedArgs:              []driver.Value{"%fake"},
		},
		{
			name: "not ends with case with *string",
			args: args{
				v: &GenericFilter{NotEndsWith: starString("fake")},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 LIKE $1",
			expectedArgs:              []driver.Value{"%fake"},
		},
		{
			name: "not ends with case with Enum struct",
			args: args{
				v: &GenericFilter{NotEndsWith: FakeStringTestEnum},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 LIKE $1",
			expectedArgs:              []driver.Value{"%FAKE"},
		},
		{
			name: "not ends with case with *Enum struct",
			args: args{
				v: &GenericFilter{NotEndsWith: starInterface(FakeStringTestEnum)},
			},
			wantErr:     true,
			errorString: "notEndsWith value must be a string or *string",
		},
		{
			name: "not ends with case with int",
			args: args{
				v: &GenericFilter{NotEndsWith: 1},
			},
			wantErr:     true,
			errorString: "notEndsWith value must be a string or *string",
		},
		{
			name: "not ends with case with *int",
			args: args{
				v: &GenericFilter{NotEndsWith: starInterface(1)},
			},
			wantErr:     true,
			errorString: "notEndsWith value must be a string or *string",
		},
		{
			name: "not ends with case with date",
			args: args{
				v: &GenericFilter{NotEndsWith: now},
			},
			wantErr:     true,
			errorString: "notEndsWith value must be a string or *string",
		},
		{
			name: "not ends with case with *date",
			args: args{
				v: &GenericFilter{NotEndsWith: starInterface(now)},
			},
			wantErr:     true,
			errorString: "notEndsWith value must be a string or *string",
		},
		{
			name: "not ends with case with bool",
			args: args{
				v: &GenericFilter{NotEndsWith: true},
			},
			wantErr:     true,
			errorString: "notEndsWith value must be a string or *string",
		},
		{
			name: "not ends with case with *bool",
			args: args{
				v: &GenericFilter{NotEndsWith: starInterface(true)},
			},
			wantErr:     true,
			errorString: "notEndsWith value must be a string or *string",
		},
		{
			name: "not ends with case with upper value",
			args: args{
				v: &GenericFilter{NotEndsWith: "fake", ValueUppercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 LIKE upper($1)",
			expectedArgs:              []driver.Value{"%fake"},
		},
		{
			name: "not ends with case with lower value",
			args: args{
				v: &GenericFilter{NotEndsWith: "fake", ValueLowercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT field_1 LIKE lower($1)",
			expectedArgs:              []driver.Value{"%fake"},
		},
		{
			name: "not ends with case with upper field",
			args: args{
				v: &GenericFilter{NotEndsWith: "fake", FieldUppercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT upper(field_1) LIKE $1",
			expectedArgs:              []driver.Value{"%fake"},
		},
		{
			name: "not ends with case with lower field",
			args: args{
				v: &GenericFilter{NotEndsWith: "fake", FieldLowercase: true},
			},
			expectedIntermediateQuery: "WHERE NOT lower(field_1) LIKE $1",
			expectedArgs:              []driver.Value{"%fake"},
		},
		{
			name: "not ends with case with case insensitive",
			args: args{
				v: &GenericFilter{NotEndsWith: "fake", CaseInsensitive: true},
			},
			expectedIntermediateQuery: "WHERE NOT lower(field_1) LIKE lower($1)",
			expectedArgs:              []driver.Value{"%fake"},
		},
		// IN
		{
			name: "in case with []string",
			args: args{
				v: &GenericFilter{In: []string{"fake"}},
			},
			expectedIntermediateQuery: "WHERE field_1 IN ($1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "in case with []*string",
			args: args{
				v: &GenericFilter{In: []*string{starString("fake")}},
			},
			expectedIntermediateQuery: "WHERE field_1 IN ($1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "in case with []string with 2 values",
			args: args{
				v: &GenericFilter{In: []string{"fake", "fake2"}},
			},
			expectedIntermediateQuery: "WHERE field_1 IN ($1,$2)",
			expectedArgs:              []driver.Value{"fake", "fake2"},
		},
		{
			name: "in case with []int",
			args: args{
				v: &GenericFilter{In: []int{1}},
			},
			expectedIntermediateQuery: "WHERE field_1 IN ($1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "in case with []int with 2 values",
			args: args{
				v: &GenericFilter{In: []int{1, 2}},
			},
			expectedIntermediateQuery: "WHERE field_1 IN ($1,$2)",
			expectedArgs:              []driver.Value{1, 2},
		},
		{
			name: "in case with []Enum",
			args: args{
				v: &GenericFilter{In: []StringTestEnum{FakeStringTestEnum}},
			},
			expectedIntermediateQuery: "WHERE field_1 IN ($1)",
			expectedArgs:              []driver.Value{"FAKE"},
		},
		{
			name: "in case with upper value",
			args: args{
				v: &GenericFilter{In: []string{"fake"}, ValueUppercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 IN ($1)",
			expectedArgs:              []driver.Value{"FAKE"},
		},
		{
			name: "in case with lower value",
			args: args{
				v: &GenericFilter{In: []string{"FAKE"}, ValueLowercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 IN ($1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "in case with upper field",
			args: args{
				v: &GenericFilter{In: []string{"fake"}, FieldUppercase: true},
			},
			expectedIntermediateQuery: "WHERE upper(field_1) IN ($1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "in case with lower field",
			args: args{
				v: &GenericFilter{In: []string{"fake"}, FieldLowercase: true},
			},
			expectedIntermediateQuery: "WHERE lower(field_1) IN ($1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "in case with case insensitive",
			args: args{
				v: &GenericFilter{In: []string{"FAKE"}, CaseInsensitive: true},
			},
			expectedIntermediateQuery: "WHERE lower(field_1) IN ($1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		// NOT IN
		{
			name: "not in case with []string",
			args: args{
				v: &GenericFilter{NotIn: []string{"fake"}},
			},
			expectedIntermediateQuery: "WHERE field_1 NOT IN ($1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not in case with []*string",
			args: args{
				v: &GenericFilter{NotIn: []*string{starString("fake")}},
			},
			expectedIntermediateQuery: "WHERE field_1 NOT IN ($1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not in case with []string with 2 values",
			args: args{
				v: &GenericFilter{NotIn: []string{"fake", "fake2"}},
			},
			expectedIntermediateQuery: "WHERE field_1 NOT IN ($1,$2)",
			expectedArgs:              []driver.Value{"fake", "fake2"},
		},
		{
			name: "not in case with []int",
			args: args{
				v: &GenericFilter{NotIn: []int{1}},
			},
			expectedIntermediateQuery: "WHERE field_1 NOT IN ($1)",
			expectedArgs:              []driver.Value{1},
		},
		{
			name: "not in case with []int with 2 values",
			args: args{
				v: &GenericFilter{NotIn: []int{1, 2}},
			},
			expectedIntermediateQuery: "WHERE field_1 NOT IN ($1,$2)",
			expectedArgs:              []driver.Value{1, 2},
		},
		{
			name: "not in case with []Enum",
			args: args{
				v: &GenericFilter{NotIn: []StringTestEnum{FakeStringTestEnum}},
			},
			expectedIntermediateQuery: "WHERE field_1 NOT IN ($1)",
			expectedArgs:              []driver.Value{"FAKE"},
		},
		{
			name: "not in case with upper value",
			args: args{
				v: &GenericFilter{NotIn: []string{"fake"}, ValueUppercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 NOT IN ($1)",
			expectedArgs:              []driver.Value{"FAKE"},
		},
		{
			name: "not in case with lower value",
			args: args{
				v: &GenericFilter{NotIn: []string{"FAKE"}, ValueLowercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 NOT IN ($1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not in case with upper field",
			args: args{
				v: &GenericFilter{NotIn: []string{"fake"}, FieldUppercase: true},
			},
			expectedIntermediateQuery: "WHERE upper(field_1) NOT IN ($1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		{
			name: "not in case with lower field",
			args: args{
				v: &GenericFilter{NotIn: []string{"FAKE"}, FieldLowercase: true},
			},
			expectedIntermediateQuery: "WHERE lower(field_1) NOT IN ($1)",
			expectedArgs:              []driver.Value{"FAKE"},
		},
		{
			name: "not in case with case insensitive",
			args: args{
				v: &GenericFilter{NotIn: []string{"FAKE"}, CaseInsensitive: true},
			},
			expectedIntermediateQuery: "WHERE lower(field_1) NOT IN ($1)",
			expectedArgs:              []driver.Value{"fake"},
		},
		// IS NULL
		{
			name: "is null false",
			args: args{
				v: &GenericFilter{IsNull: false},
			},
		},
		{
			name: "is null true",
			args: args{
				v: &GenericFilter{IsNull: true},
			},
			expectedIntermediateQuery: "WHERE field_1 IS NULL",
			expectedArgs:              []driver.Value{},
		},
		{
			name: "is null case with upper value",
			args: args{
				v: &GenericFilter{IsNull: true, ValueUppercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 IS NULL",
			expectedArgs:              []driver.Value{},
		},
		{
			name: "is null case with lower value",
			args: args{
				v: &GenericFilter{IsNull: true, ValueLowercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 IS NULL",
			expectedArgs:              []driver.Value{},
		},
		{
			name: "is null case with upper field",
			args: args{
				v: &GenericFilter{IsNull: true, FieldUppercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 IS NULL",
			expectedArgs:              []driver.Value{},
		},
		{
			name: "is null case with lower field",
			args: args{
				v: &GenericFilter{IsNull: true, FieldLowercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 IS NULL",
			expectedArgs:              []driver.Value{},
		},
		{
			name: "is null case with case insensitive",
			args: args{
				v: &GenericFilter{IsNull: true, CaseInsensitive: true},
			},
			expectedIntermediateQuery: "WHERE field_1 IS NULL",
			expectedArgs:              []driver.Value{},
		},
		// IS NOT NULL
		{
			name: "is not null false",
			args: args{
				v: &GenericFilter{IsNotNull: false},
			},
		},
		{
			name: "is not null true",
			args: args{
				v: &GenericFilter{IsNotNull: true},
			},
			expectedIntermediateQuery: "WHERE field_1 IS NOT NULL",
			expectedArgs:              []driver.Value{},
		},
		{
			name: "is not null case with upper value",
			args: args{
				v: &GenericFilter{IsNotNull: true, ValueUppercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 IS NOT NULL",
			expectedArgs:              []driver.Value{},
		},
		{
			name: "is not null case with lower value",
			args: args{
				v: &GenericFilter{IsNotNull: true, ValueLowercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 IS NOT NULL",
			expectedArgs:              []driver.Value{},
		},
		{
			name: "is not null case with upper field",
			args: args{
				v: &GenericFilter{IsNotNull: true, FieldUppercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 IS NOT NULL",
			expectedArgs:              []driver.Value{},
		},
		{
			name: "is not null case with lower field",
			args: args{
				v: &GenericFilter{IsNotNull: true, FieldLowercase: true},
			},
			expectedIntermediateQuery: "WHERE field_1 IS NOT NULL",
			expectedArgs:              []driver.Value{},
		},
		{
			name: "is not null case with case insensitive",
			args: args{
				v: &GenericFilter{IsNotNull: true, CaseInsensitive: true},
			},
			expectedIntermediateQuery: "WHERE field_1 IS NOT NULL",
			expectedArgs:              []driver.Value{},
		},
		// All at the same time
		{
			name: "all at the same time",
			args: args{
				v: &GenericFilter{
					Eq:            starInterface("fake-eq"),
					NotEq:         starInterface("fake-not-eq2"),
					Gte:           starInterface(10),
					Gt:            starInterface(5),
					NotGte:        starInterface(7),
					NotGt:         starInterface(3),
					Contains:      starString("fake-contains"),
					NotContains:   starString("fake-not-contains2"),
					EndsWith:      "fake-ends",
					NotEndsWith:   "fake-not-ends2",
					In:            []string{"fake-in", "fake-in2"},
					NotIn:         []string{"fake-not-in", "fake-not-in2"},
					Lt:            1,
					Lte:           2,
					NotLt:         13,
					NotLte:        4,
					StartsWith:    "fake-starts",
					NotStartsWith: "fake-not-starts2",
				},
			},
			expectedIntermediateQuery: "WHERE field_1 = $1 AND NOT field_1 = $2 AND field_1 >= $3 AND NOT field_1 >= $4 AND field_1 > $5 AND NOT field_1 > $6 AND field_1 <= $7 AND NOT field_1 <= $8 AND field_1 < $9 AND NOT field_1 < $10 AND field_1 LIKE $11 AND NOT field_1 LIKE $12 AND field_1 LIKE $13 AND NOT field_1 LIKE $14 AND field_1 LIKE $15 AND NOT field_1 LIKE $16 AND field_1 IN ($17,$18) AND field_1 NOT IN ($19,$20)",
			expectedArgs: []driver.Value{
				"fake-eq",
				"fake-not-eq2",
				10,
				7,
				5,
				3,
				2,
				4,
				1,
				13,
				"%fake-contains%",
				"%fake-not-contains2%",
				"fake-starts%",
				"fake-not-starts2%",
				"%fake-ends",
				"%fake-not-ends2",
				"fake-in",
				"fake-in2",
				"fake-not-in",
				"fake-not-in2",
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

			db, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{Logger: logger.Discard})
			if err != nil {
				t.Error(err)
				return
			}

			got, err := manageFilterRequest("field_1", tt.args.v, db)
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
			expectedQuery += `ORDER BY "people"."name" LIMIT $` + strconv.Itoa(len(tt.expectedArgs)+1)
			// Add limit to expected args
			tt.expectedArgs = append(tt.expectedArgs, 1)

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

func Test_transformToLowerOrUpperCasesList(t *testing.T) {
	type args struct {
		input       interface{}
		toLowercase bool
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "should ignore nil",
			args: args{input: nil, toLowercase: true},
			want: nil,
		},
		{
			name: "should ignore int",
			args: args{input: 1, toLowercase: true},
			want: 1,
		},
		{
			name: "should ignore *int",
			args: args{input: starAny(1), toLowercase: true},
			want: starAny(1),
		},
		{
			name: "should ignore string",
			args: args{input: "FakE", toLowercase: true},
			want: "FakE",
		},
		{
			name: "should ignore *string",
			args: args{input: starAny("FakE"), toLowercase: true},
			want: starAny("FakE"),
		},
		{
			name: "should be ok with empty []string",
			args: args{input: []string{}, toLowercase: true},
			want: []string{},
		},
		{
			name: "should be ok with values in []string",
			args: args{input: []string{"FaKe", "FAKE", "fake"}, toLowercase: true},
			want: []string{"fake", "fake", "fake"},
		},
		{
			name: "should be ok with empty []*string",
			args: args{input: []*string{}, toLowercase: true},
			want: []*string{},
		},
		{
			name: "should be ok with values in []*string",
			args: args{input: []*string{starAny("FakE"), starAny("fake"), starAny("FAKE")}, toLowercase: true},
			want: []*string{starAny("fake"), starAny("fake"), starAny("fake")},
		},
		{
			name: "should be ok with nil value in []*string",
			args: args{input: []*string{nil}, toLowercase: true},
			want: []*string{nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transformToLowerOrUpperCasesList(tt.args.input, tt.args.toLowercase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transformToLowerOrUpperCasesList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func starAny[T any](i T) *T {
	return &i
}
