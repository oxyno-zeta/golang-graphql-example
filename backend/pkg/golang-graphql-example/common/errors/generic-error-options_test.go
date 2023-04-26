package errors

import (
	gerrors "errors"
	"reflect"
	"testing"
)

func TestWithErrorMessage(t *testing.T) {
	type args struct {
		errMsg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "override",
			args: args{errMsg: "fake"},
			want: "fake",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGenericError(WithErrorMessage(tt.args.errMsg))
			if got.Error() != tt.want {
				t.Errorf("WithErrorMessage() = %v, want %v", got, tt.want)
			}
			if got.StackTrace() == nil {
				t.Error("WithErrorMessage() don't have stack trace")
			}
		})
	}
}

func TestWithError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "override",
			args: args{err: gerrors.New("fake")},
			want: "fake",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGenericError(WithError(tt.args.err))
			if got.Error() != tt.want {
				t.Errorf("WithError() = %v, want %v", got, tt.want)
			}
			if got.StackTrace() == nil {
				t.Error("WithError() don't have stack trace")
			}
		})
	}
}

func TestWithPublicError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "override",
			args: args{err: gerrors.New("fake")},
			want: "fake",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGenericError(WithPublicError(tt.args.err))
			if got.PublicError().Error() != tt.want {
				t.Errorf("WithPublicError() = %v, want %v", got, tt.want)
			}
			if got.PublicMessage() != tt.want {
				t.Errorf("WithPublicError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithPublicErrorMessage(t *testing.T) {
	type args struct {
		errMsg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "override",
			args: args{errMsg: "fake"},
			want: "fake",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGenericError(WithPublicErrorMessage(tt.args.errMsg))
			if got.PublicError().Error() != tt.want {
				t.Errorf("WithPublicError() = %v, want %v", got, tt.want)
			}
			if got.PublicMessage() != tt.want {
				t.Errorf("WithPublicError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithExtensions(t *testing.T) {
	type args struct {
		input map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "override",
			args: args{input: map[string]interface{}{"fake": "option"}},
			want: map[string]interface{}{"fake": "option"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGenericError(WithExtensions(tt.args.input))
			if !reflect.DeepEqual(got.Extensions(), tt.want) {
				t.Errorf("WithExtensions() = %v, want %v", got.Extensions(), tt.want)
			}
		})
	}
}

func TestAddExtension(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "override",
			args: args{
				key:   "fake",
				value: "option",
			},
			want: map[string]interface{}{
				"fake": "option",
				"code": "INTERNAL_SERVER_ERROR",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGenericError(AddExtension(tt.args.key, tt.args.value))
			if !reflect.DeepEqual(got.Extensions(), tt.want) {
				t.Errorf("WithExtensions() = %v, want %v", got.Extensions(), tt.want)
			}
		})
	}
}

func TestWithStatusCode(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "override",
			args: args{
				input: 433,
			},
			want: 433,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGenericError(WithStatusCode(tt.args.input))
			if got.StatusCode() != tt.want {
				t.Errorf("WithStatusCode() = %v, want %v", got.StatusCode(), tt.want)
			}
		})
	}
}

func TestWithCode(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "override",
			args: args{
				input: "FAKE",
			},
			want: "FAKE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGenericError(WithCode(tt.args.input))
			if got.Code() != tt.want {
				t.Errorf("WithCode() = %v, want %v", got.Code(), tt.want)
			}
			if got.Extensions()["code"] != tt.want {
				t.Errorf("WithCode() = %v, want %v", got.Extensions()["code"], tt.want)
			}
		})
	}
}
