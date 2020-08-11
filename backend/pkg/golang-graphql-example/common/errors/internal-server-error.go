package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

func NewInternalServerError(msg string) Error {
	return &defaultError{
		err:        errors.New(msg),
		ext:        map[string]interface{}{"code": "INTERNAL_SERVER_ERROR"},
		statusCode: http.StatusInternalServerError,
	}
}
