package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

func NewNotFoundError(msg string) Error {
	return &defaultError{
		err:        errors.New(msg),
		ext:        map[string]interface{}{"code": "NOT_FOUND"},
		statusCode: http.StatusNotFound,
	}
}
