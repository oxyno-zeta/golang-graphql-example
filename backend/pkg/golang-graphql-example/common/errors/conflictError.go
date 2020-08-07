package errors

import "github.com/pkg/errors"

func NewConflictError(msg string) Error {
	return &defaultError{err: errors.New(msg), ext: map[string]interface{}{"code": "CONFLICT"}}
}
