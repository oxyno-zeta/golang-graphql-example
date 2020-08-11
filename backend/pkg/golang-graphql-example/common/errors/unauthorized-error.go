package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

func NewUnauthorizedError(msg string) Error {
	return &defaultError{
		err:        errors.New(msg),
		ext:        map[string]interface{}{"code": "UNAUTHORIZED"},
		statusCode: http.StatusUnauthorized,
	}
}
