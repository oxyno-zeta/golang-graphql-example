//+build unit

package utils

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/vektah/gqlparser/v2/ast"
)

func TestManageSimpleProjection(t *testing.T) {
	starString := func(s string) *string { return &s }
	type Out1 struct {
		Field1 bool
	}
	type Out2 struct {
		Field1 bool `graphql_field:"-"`
	}
	type Out3 struct {
		Field1 string `graphql_field:"field1"`
	}
	type Out4 struct {
		Field1 *string `graphql_field:"field1"`
	}
	type Out5 struct {
		Field1 bool `graphql_field:"field1"`
	}
	type InnerOut1 struct {
		InnerField1 bool `graphql_field:"innerField1"`
	}
	type Out6 struct {
		Field1 *InnerOut1 `graphql_field:"field1"`
	}
	type args struct {
		fctx          *graphql.FieldContext
		projectionOut interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		errorString string
		want        interface{}
	}{
		{
			name: "nil case",
			args: args{
				projectionOut: nil,
				fctx:          &graphql.FieldContext{},
			},
			wantErr:     true,
			errorString: "projection output cannot be nil",
		},
		{
			name: "projection object must be a pointer",
			args: args{
				projectionOut: "fake",
				fctx:          &graphql.FieldContext{},
			},
			wantErr:     true,
			errorString: "projection output must be a pointer to an object",
		},
		{
			name: "projection object must be a pointer",
			args: args{
				projectionOut: starString("fake"),
				fctx:          &graphql.FieldContext{},
			},
			wantErr:     true,
			errorString: "projection output must be a pointer to an object",
		},
		{
			name: "field ignored: no tag",
			args: args{
				projectionOut: &Out1{},
				fctx: &graphql.FieldContext{
					Field: graphql.CollectedField{
						Selections: ast.SelectionSet{
							&ast.Field{Name: "field1"},
						},
					},
				},
			},
			want: &Out1{Field1: false},
		},
		{
			name: "field ignored: ignore",
			args: args{
				projectionOut: &Out2{},
				fctx: &graphql.FieldContext{
					Field: graphql.CollectedField{
						Selections: ast.SelectionSet{
							&ast.Field{Name: "field1"},
						},
					},
				},
			},
			want: &Out2{Field1: false},
		},
		{
			name: "field ignored: field not in projection",
			args: args{
				projectionOut: &Out2{},
				fctx: &graphql.FieldContext{
					Field: graphql.CollectedField{
						Selections: ast.SelectionSet{
							&ast.Field{Name: "field2"},
						},
					},
				},
			},
			want: &Out2{Field1: false},
		},
		{
			name: "not a boolean or a struct ptr: string",
			args: args{
				projectionOut: &Out3{},
				fctx: &graphql.FieldContext{
					Field: graphql.CollectedField{
						Selections: ast.SelectionSet{
							&ast.Field{Name: "field1"},
						},
					},
				},
			},
			wantErr:     true,
			errorString: "field Field1 must be a boolean or a pointer to a struct",
		},
		{
			name: "not a boolean or a struct ptr: *string",
			args: args{
				projectionOut: &Out4{},
				fctx: &graphql.FieldContext{
					Field: graphql.CollectedField{
						Selections: ast.SelectionSet{
							&ast.Field{Name: "field1"},
						},
					},
				},
			},
			wantErr:     true,
			errorString: "field Field1 must be a boolean or a pointer to a struct",
		},
		{
			name: "simple field",
			args: args{
				projectionOut: &Out5{},
				fctx: &graphql.FieldContext{
					Field: graphql.CollectedField{
						Selections: ast.SelectionSet{
							&ast.Field{Name: "field1"},
						},
					},
				},
			},
			want: &Out5{Field1: true},
		},
		{
			name: "inner field",
			args: args{
				projectionOut: &Out6{},
				fctx: &graphql.FieldContext{
					Field: graphql.CollectedField{
						Selections: ast.SelectionSet{
							&ast.Field{
								Name: "field1",
								SelectionSet: ast.SelectionSet{
									&ast.Field{Name: "innerField1"},
								},
							},
						},
					},
				},
			},
			want: &Out6{Field1: &InnerOut1{InnerField1: true}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create context
			ctx := context.TODO()
			ctx = graphql.WithOperationContext(ctx, &graphql.OperationContext{})
			ctx = graphql.WithFieldContext(ctx, tt.args.fctx)

			err := ManageSimpleProjection(ctx, tt.args.projectionOut)
			if err != nil && err.Error() != tt.errorString {
				t.Errorf("ManageSimpleProjection() error = %v, wantErr %v", err, tt.errorString)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("ManageSimpleProjection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				assert.Equal(t, tt.want, tt.args.projectionOut)
			}
		})
	}
}

func TestManageConnectionNodeProjection(t *testing.T) {
	starString := func(s string) *string { return &s }
	type Out1 struct {
		Fake string
	}
	type InnerOut1 struct {
		InnerField1 bool `graphql_field:"innerField1"`
	}
	type Out2 struct {
		Field1 *InnerOut1 `graphql_field:"field1"`
	}
	type args struct {
		fctx          *graphql.FieldContext
		projectionOut interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		errorString string
		want        interface{}
	}{
		{
			name: "nil case",
			args: args{
				projectionOut: nil,
				fctx:          &graphql.FieldContext{},
			},
			wantErr:     true,
			errorString: "projection output cannot be nil",
		},
		{
			name: "projection object must be a pointer",
			args: args{
				projectionOut: "fake",
				fctx:          &graphql.FieldContext{},
			},
			wantErr:     true,
			errorString: "projection output must be a pointer to an object",
		},
		{
			name: "projection object must be a pointer",
			args: args{
				projectionOut: starString("fake"),
				fctx:          &graphql.FieldContext{},
			},
			wantErr:     true,
			errorString: "projection output must be a pointer to an object",
		},
		{
			name: "cannot find any edges",
			args: args{
				projectionOut: &Out1{},
				fctx: &graphql.FieldContext{
					Field: graphql.CollectedField{
						Selections: ast.SelectionSet{
							&ast.Field{
								Name: "field1",
							},
						},
					},
				},
			},
			want: &Out1{},
		},
		{
			name: "cannot find any node",
			args: args{
				projectionOut: &Out1{},
				fctx: &graphql.FieldContext{
					Field: graphql.CollectedField{
						Selections: ast.SelectionSet{
							&ast.Field{
								Name: "edges",
								SelectionSet: ast.SelectionSet{
									&ast.Field{Name: "field1"},
								},
							},
						},
					},
				},
			},
			want: &Out1{},
		},
		{
			name: "valid",
			args: args{
				projectionOut: &Out2{},
				fctx: &graphql.FieldContext{
					Field: graphql.CollectedField{
						Selections: ast.SelectionSet{
							&ast.Field{
								Name: "edges",
								SelectionSet: ast.SelectionSet{
									&ast.Field{
										Name: "node",
										SelectionSet: ast.SelectionSet{
											&ast.Field{
												Name: "field1",
												SelectionSet: ast.SelectionSet{
													&ast.Field{Name: "innerField1"},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want: &Out2{
				Field1: &InnerOut1{
					InnerField1: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create context
			ctx := context.TODO()
			ctx = graphql.WithOperationContext(ctx, &graphql.OperationContext{})
			ctx = graphql.WithFieldContext(ctx, tt.args.fctx)
			err := ManageConnectionNodeProjection(ctx, tt.args.projectionOut)
			if err != nil && err.Error() != tt.errorString {
				t.Errorf("ManageConnectionNodeProjection() error = %v, wantErr %v", err, tt.errorString)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("ManageConnectionNodeProjection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				assert.Equal(t, tt.want, tt.args.projectionOut)
			}
		})
	}
}
