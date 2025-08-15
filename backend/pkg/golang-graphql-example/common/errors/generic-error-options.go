package errors

import (
	"net/http"

	"emperror.dev/errors"
)

type GenericErrorOption func(*GenericError)

func NewGenericError(options ...GenericErrorOption) Error {
	defaultErr := errors.New("internal server error")

	// Create default error
	err := &GenericError{
		err: errors.WithStack(defaultErr),
		ext: map[string]any{
			"code": InternalServerErrorCode,
		},
		publicError: defaultErr,
		statusCode:  http.StatusInternalServerError,
		code:        InternalServerErrorCode,
	}

	// Apply all options
	for _, o := range options {
		o(err)
	}

	return err
}

// WithErrorMessage will create an error from error message provided.
// Warning: This will create an error with the current package as part of the stack trace.
func WithErrorMessage(errMsg string) GenericErrorOption {
	return func(ge *GenericError) { ge.err = errors.New(errMsg) }
}

// WithError will save the error as main error and will add stack trace if not present.
func WithError(err error) GenericErrorOption {
	return func(ge *GenericError) { ge.err = errors.WithStack(err) }
}

// WithExtensions will replace custom extensions map.
// Warning: This will remove the "code" extension. It is recommended to use the WithCode option just after.
func WithExtensions(input map[string]any) GenericErrorOption {
	return func(ge *GenericError) { ge.ext = input }
}

// AddExtension will add an extension to the custom extensions map.
func AddExtension(key string, val any) GenericErrorOption {
	return func(ge *GenericError) { ge.ext[key] = val }
}

// WithPublicError will replace the public error and will add stack trace if not present.
// This error will be returned on API responses for example.
func WithPublicError(err error) GenericErrorOption {
	return func(ge *GenericError) { ge.publicError = errors.WithStack(err) }
}

// WithPublicErrorMessage will create an error from public error message provided.
// Warning: This will create an error with the current package as part of the stack trace.
func WithPublicErrorMessage(msg string) GenericErrorOption {
	return func(ge *GenericError) { ge.publicError = errors.New(msg) }
}

// WithStatusCode will replace the status code that will be used on API responses.
func WithStatusCode(st int) GenericErrorOption {
	return func(ge *GenericError) { ge.statusCode = st }
}

// WithCode will replace the error code both in code field and in extensions.
func WithCode(input string) GenericErrorOption {
	return func(ge *GenericError) {
		ge.code = input
		ge.ext["code"] = input
	}
}
