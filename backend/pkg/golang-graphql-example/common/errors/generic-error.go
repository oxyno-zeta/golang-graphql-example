package errors

import "emperror.dev/errors"

type GenericError struct {
	err         error
	ext         map[string]any
	publicError error
	code        string
	statusCode  int
}

func (e *GenericError) Error() string {
	return e.err.Error()
}

func (e *GenericError) StackTrace() errors.StackTrace {
	// Cast internal error as a stack tracer error
	//nolint: errorlint // Ignore this because the aim is to catch stack trace error at first level
	if err2, ok := e.err.(stackTracerError); ok {
		return err2.StackTrace()
	}
	// Return nothing
	return nil
}

func (e *GenericError) Extensions() map[string]any {
	return e.ext
}

func (e *GenericError) StatusCode() int {
	return e.statusCode
}

func (e *GenericError) PublicMessage() string {
	return e.publicError.Error()
}

func (e *GenericError) PublicError() error {
	return e.publicError
}

func (e *GenericError) Code() string {
	return e.code
}
