package errors

import "github.com/pkg/errors"

type defaultError struct {
	err        error
	ext        map[string]interface{}
	statusCode int
}

func (e *defaultError) Error() string {
	return e.err.Error()
}

func (e *defaultError) StackTrace() errors.StackTrace {
	// Cast internal error as a stack tracer error
	if err2, ok := e.err.(stackTracerError); ok {
		return err2.StackTrace()
	}
	// Return nothing
	return nil
}

func (e *defaultError) Extensions() map[string]interface{} {
	return e.ext
}

func (e *defaultError) StatusCode() int {
	return e.statusCode
}
