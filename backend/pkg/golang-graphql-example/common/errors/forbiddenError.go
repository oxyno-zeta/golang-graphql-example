package errors

import "github.com/pkg/errors"

func NewForbiddenError(msg string) Error {
	return &defaultError{err: errors.New(msg), ext: map[string]interface{}{"code": "FORBIDDEN"}}
}
