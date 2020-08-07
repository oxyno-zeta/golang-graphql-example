package errors

import "github.com/pkg/errors"

func NewInvalidInputError(msg string) Error {
	return &defaultError{err: errors.New(msg), ext: map[string]interface{}{"code": "INVALID_INPUT"}}
}
