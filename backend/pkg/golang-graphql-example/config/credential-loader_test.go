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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("internalLoadAllCredentials() = %v, want %v", got, tt.want)
			}
		})
	}
}
