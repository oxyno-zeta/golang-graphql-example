package errors

import "github.com/pkg/errors"

type stackTracerError interface {
	StackTrace() errors.StackTrace
}

// Error is the interface all common errors must implement.
type Error interface {
	error
	stackTracerError
	// Extensions will return extensions.
	Extensions() map[string]interface{}
	// StatusCode will return status code linked to error.
	StatusCode() int
	// PublicMessage will return a public message to be displayed externally.
	PublicMessage() string
	// PublicError will return an error with public message inside.
	PublicError() error
	// Code will return the error code.
	Code() string
}
