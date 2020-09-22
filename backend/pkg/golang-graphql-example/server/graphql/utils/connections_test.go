package utils

import (
	"reflect"
	"testing"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/pagination"
)

func TestMapConnection(t *testing.T) {
	// starString := func(s string) *string { return &s }
	type Person struct{ Name string }
	type PersonEdge struct {
		Cursor string
		Node   *Person
	}
	type WrongEdge1 struct {
	}
	type WrongEdge2 struct {
		Cursor int
		Node   *Person
	}
	type WrongEdge3 struct {
		Cursor string
		Node   string
	}
	type WrongEdge4 struct {
		Cursor string
		Node   Person
	}
	type PersonConnection struct {
		Edges    []*PersonEdge
		PageInfo *PageInfo
	}
	type WrongConnection1 struct {
	}
	type WrongConnection2 struct {
		Edges    []string
		PageInfo *PageInfo
	}
	type WrongConnection3 struct {
		Edges    []PersonEdge
		PageInfo *PageInfo
	}
	type WrongConnection4 struct {
		Edges    []*PersonEdge
		PageInfo string
	}
	type WrongConnection5 struct {
		Edges    []*PersonEdge
		PageInfo PageInfo
	}
	type WrongConnection6 struct {
		Edges []*PersonEdge
	}
	type WrongConnection7 struct {
		Edges    []*WrongEdge1
		PageInfo *PageInfo
	}
	type WrongConnection8 struct {
		Edges    []*WrongEdge2
		PageInfo *PageInfo
	}
	type WrongConnection9 struct {
		Edges    []*WrongEdge3
		PageInfo *PageInfo
	}
	type WrongConnection10 struct {
		Edges    []*WrongEdge4
		PageInfo *PageInfo
	}
	type args struct {
		result  interface{}
		list    interface{}
		pageOut *pagination.PageOutput
	}
	tests := []struct {
		name           string
		args           args
		wantErr        bool
		errorString    string
		expectedResult interface{}
		testResult     bool
	}{
		{
			name: "nil connection result",
			args: args{
				result:  nil,
				list:    []*Person{{Name: "fake1"}, {Name: "fake2"}},
				pageOut: &pagination.PageOutput{},
			},
			wantErr:     true,
			errorString: "connection result argument mustn't be nil",
		},
		{
			name: "nil input list",
			args: args{
				result:  &PersonConnection{},
				list:    nil,
				pageOut: &pagination.PageOutput{},
			},
			wantErr:     true,
			errorString: "list argument mustn't be nil",
		},
		{
			name: "wrong type as connection result",
			args: args{
				result:  "fake",
				list:    []*Person{{Name: "fake1"}, {Name: "fake2"}},
				pageOut: &pagination.PageOutput{},
			},
			wantErr:     true,
			errorString: "connection result argument must be a pointer to a connection object",
		},
		{
			name: "wrong type as input list",
			args: args{
				result:  &PersonConnection{},
				list:    "fake",
				pageOut: &pagination.PageOutput{},
			},
			wantErr:     true,
			errorString: "list argument must be a slice",
		},
		{
			name: "no field Edges in connection",
			args: args{
				result:  &WrongConnection1{},
				list:    []*Person{{Name: "fake1"}, {Name: "fake2"}},
				pageOut: &pagination.PageOutput{},
			},
			wantErr:     true,
			errorString: "field Edges not found in connection object",
		},
		{
			name: "empty list array",
			args: args{
				result:  &PersonConnection{},
				list:    []*Person{},
				pageOut: &pagination.PageOutput{},
			},
			expectedResult: &PersonConnection{
				Edges: nil,
				PageInfo: &PageInfo{
					HasNextPage:     false,
					HasPreviousPage: false,
					EndCursor:       nil,
					StartCursor:     nil,
				},
			},
			testResult: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MapConnection(tt.args.result, tt.args.list, tt.args.pageOut)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapConnection() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && err.Error() != tt.errorString {
				t.Errorf("MapConnection() error = %v, wantErr %v", err, tt.errorString)
				return
			}
			if tt.testResult && !reflect.DeepEqual(tt.args.result, tt.expectedResult) {
				t.Errorf("MapConnection() result = %v, want %v", tt.args.result, tt.expectedResult)
			}
		})
	}
}
