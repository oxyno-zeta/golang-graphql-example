//go:build unit

package config

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_internalLoadAllCredentials(t *testing.T) {
	type obj1 struct {
		C1 *CredentialConfig
	}
	type obj2 struct {
		P1 *struct {
			C1 *CredentialConfig
		}
		C2 *CredentialConfig
	}
	type obj3 struct {
		P1 *struct {
			C1 []*CredentialConfig
		}
		C2 []*CredentialConfig
	}
	type obj4 struct {
		P1 *struct {
			C1 []*CredentialConfig
		}
		C2 *CredentialConfig
	}

	type args struct {
		out interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []*CredentialConfig
		wantErr bool
	}{
		{
			name: "simple object without anything",
			args: args{
				out: &obj1{},
			},
			want: []*CredentialConfig{},
		},
		{
			name: "simple object with value",
			args: args{
				out: &obj1{C1: &CredentialConfig{Value: "c1"}},
			},
			want: []*CredentialConfig{{Value: "c1"}},
		},
		{
			name: "second object without anything",
			args: args{
				out: &obj2{},
			},
			want: []*CredentialConfig{},
		},
		{
			name: "second object with partial content",
			args: args{
				out: &obj2{
					C2: &CredentialConfig{Value: "c2"},
				},
			},
			want: []*CredentialConfig{{Value: "c2"}},
		},
		{
			name: "second object with full content",
			args: args{
				out: &obj2{
					P1: &struct{ C1 *CredentialConfig }{
						C1: &CredentialConfig{Value: "c1"},
					},
					C2: &CredentialConfig{Value: "c2"},
				},
			},
			want: []*CredentialConfig{{Value: "c1"}, {Value: "c2"}},
		},
		{
			name: "third object without anything",
			args: args{
				out: &obj3{},
			},
			want: []*CredentialConfig{},
		},
		{
			name: "third object with partial content",
			args: args{
				out: &obj3{
					C2: []*CredentialConfig{{Value: "c2"}, {Value: "c1"}},
				},
			},
			want: []*CredentialConfig{{Value: "c2"}, {Value: "c1"}},
		},
		{
			name: "third object with full content",
			args: args{
				out: &obj3{
					P1: &struct{ C1 []*CredentialConfig }{
						C1: []*CredentialConfig{{Value: "c1"}, {Value: "c2"}},
					},
					C2: []*CredentialConfig{{Value: "c3"}, {Value: "c4"}},
				},
			},
			want: []*CredentialConfig{{Value: "c1"}, {Value: "c2"}, {Value: "c3"}, {Value: "c4"}},
		},
		{
			name: "fourth object with full content",
			args: args{
				out: &obj4{
					P1: &struct{ C1 []*CredentialConfig }{
						C1: []*CredentialConfig{{Value: "c1"}, {Value: "c2"}},
					},
					C2: &CredentialConfig{Value: "c3"},
				},
			},
			want: []*CredentialConfig{{Value: "c1"}, {Value: "c2"}, {Value: "c3"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			credentialConfigPathList, err := getRecursivelyCredentialConfigPathList([]string{}, reflect.TypeOf(tt.args.out).Elem())
			assert.NoError(t, err)

			got, err := internalLoadAllCredentials(tt.args.out, credentialConfigPathList)
			if (err != nil) != tt.wantErr {
				t.Errorf("internalLoadAllCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getRecursivelyCredentialConfigPathList(t *testing.T) {
	type obj1 struct {
		C1 *CredentialConfig
		S1 string
	}
	type obj2 struct {
		P1 *struct {
			C1 *CredentialConfig
			S1 string
		}
		C2 *CredentialConfig
		S1 string
	}
	type obj3 struct {
		P1 *struct {
			C1 *CredentialConfig
			S1 string
		}
		P2 *struct {
			P3 *struct {
				C3 *CredentialConfig
				C4 *CredentialConfig
			}
		}
		C2 *CredentialConfig
		S1 string
	}
	type obj4 struct {
		P1 *struct {
			C1 *CredentialConfig
			S1 string
		}
		P2 *struct {
			P3 *struct {
				C3 *CredentialConfig
				C4 *CredentialConfig
			}
		}
		P4 *struct {
			P5 *struct {
				S1 string
				P6 *struct {
					C3 *CredentialConfig
					C4 *CredentialConfig
					S1 string
				}
			}
		}
		C2 *CredentialConfig
		S1 string
	}
	type obj5 struct {
		C1 []*CredentialConfig
		S1 string
	}
	type obj6 struct {
		C1 []CredentialConfig
		S1 string
	}
	type obj7 struct {
		P1 *struct {
			C1 []*CredentialConfig
			C2 []CredentialConfig
			S1 string
		}
		C1 []*CredentialConfig
		C2 []CredentialConfig
		S1 string
	}
	type obj8 struct {
		P1 *struct {
			C1 *CredentialConfig
			S1 string
		}
		P2 *struct {
			P3 *struct {
				C1 []*CredentialConfig
				C2 []CredentialConfig
				C3 *CredentialConfig
				C4 *CredentialConfig
			}
		}
		P4 *struct {
			P5 *struct {
				S1 string
				P6 *struct {
					C1 []*CredentialConfig
					C2 []CredentialConfig
					C3 *CredentialConfig
					C4 *CredentialConfig
					S1 string
				}
			}
		}
		C2 *CredentialConfig
		S1 string
	}

	tests := []struct {
		name    string
		input   interface{}
		want    [][]string
		wantErr bool
	}{
		{
			name:  "case 1",
			input: obj1{},
			want:  [][]string{{"C1"}},
		},
		{
			name:  "case 2",
			input: obj2{},
			want:  [][]string{{"P1", "C1"}, {"C2"}},
		},
		{
			name:  "case 3",
			input: obj3{},
			want:  [][]string{{"P1", "C1"}, {"P2", "P3", "C3"}, {"P2", "P3", "C4"}, {"C2"}},
		},
		{
			name:  "case 4",
			input: obj4{},
			want:  [][]string{{"P1", "C1"}, {"P2", "P3", "C3"}, {"P2", "P3", "C4"}, {"P4", "P5", "P6", "C3"}, {"P4", "P5", "P6", "C4"}, {"C2"}},
		},
		{
			name:  "case 5",
			input: obj5{},
			want:  [][]string{{"C1"}},
		},
		{
			name:  "case 6",
			input: obj6{},
			want:  [][]string{{"C1"}},
		},
		{
			name:  "case 7",
			input: obj7{},
			want: [][]string{
				{"P1", "C1"},
				{"P1", "C2"},
				{"C1"},
				{"C2"},
			},
		},
		{
			name:  "case 8",
			input: obj8{},
			want: [][]string{
				{"P1", "C1"},
				{"P2", "P3", "C1"},
				{"P2", "P3", "C2"},
				{"P2", "P3", "C3"},
				{"P2", "P3", "C4"},
				{"P4", "P5", "P6", "C1"},
				{"P4", "P5", "P6", "C2"},
				{"P4", "P5", "P6", "C3"},
				{"P4", "P5", "P6", "C4"},
				{"C2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getRecursivelyCredentialConfigPathList([]string{}, reflect.TypeOf(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("getRecursivelyCredentialConfigPathList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRecursivelyCredentialConfigPathList() = %v, want %v", got, tt.want)
			}
		})
	}
}
