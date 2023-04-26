//go:build unit

package errors

import (
	gerrors "errors"
	"reflect"
	"testing"

	"emperror.dev/errors"
)

func TestNewInternalServerError(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name       string
		args       args
		err        error
		ext        map[string]interface{}
		statusCode int
	}{
		{
			name:       "constructor",
			args:       args{msg: "fake"},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "INTERNAL_SERVER_ERROR"},
			statusCode: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInternalServerError(tt.args.msg)
			if !reflect.DeepEqual(got.Error(), tt.err.Error()) {
				t.Errorf("NewInternalServerError().err = %v, want %v", got.Error(), tt.err.Error())
			}
			if !reflect.DeepEqual(got.Extensions(), tt.ext) {
				t.Errorf("NewInternalServerError().ext = %v, want %v", got.Extensions(), tt.ext)
			}
			if !reflect.DeepEqual(got.StatusCode(), tt.statusCode) {
				t.Errorf("NewInternalServerError().statusCode = %v, want %v", got.StatusCode(), tt.statusCode)
			}
			if got.StackTrace() == nil {
				t.Error("NewInternalServerError().stackTrace must exists")
			}
		})
	}
}

func TestNewInternalServerErrorWithError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name       string
		args       args
		err        error
		ext        map[string]interface{}
		statusCode int
	}{
		{
			name:       "constructor",
			args:       args{err: errors.New("fake")},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "INTERNAL_SERVER_ERROR"},
			statusCode: 500,
		},
		{
			name:       "constructor with golang error",
			args:       args{err: gerrors.New("fake")},
			err:        errors.New("fake"),
			ext:        map[string]interface{}{"code": "INTERNAL_SERVER_ERROR"},
			statusCode: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInternalServerErrorWithError(tt.args.err)
			if !reflect.DeepEqual(got.Error(), tt.err.Error()) {
				t.Errorf("NewInternalServerErrorWithError().err = %v, want %v", got.Error(), tt.err.Error())
			}
			if !reflect.DeepEqual(got.Extensions(), tt.ext) {
				t.Errorf("NewInternalServerErrorWithError().ext = %v, want %v", got.Extensions(), tt.ext)
			}
			if !reflect.DeepEqual(got.StatusCode(), tt.statusCode) {
				t.Errorf("NewInternalServerErrorWithError().statusCode = %v, want %v", got.StatusCode(), tt.statusCode)
			}
			if got.StackTrace() == nil {
				t.Error("NewInternalServerErrorWithError().stackTrace must exists")
			}
		})
	}
}
