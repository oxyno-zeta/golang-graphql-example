package errors

import (
	"net/http"
)

const UnauthorizedErrorCode = "UNAUTHORIZED"

func NewUnauthorizedError(msg string, options ...GenericErrorOption) Error {
	return NewUnauthorizedErrorWithOptions(
		append([]GenericErrorOption{WithErrorMessage(msg)}, options...)...)
}

func NewUnauthorizedErrorWithError(err error, options ...GenericErrorOption) Error {
	return NewUnauthorizedErrorWithOptions(
		append([]GenericErrorOption{WithError(err)}, options...)...)
}

func NewUnauthorizedErrorWithOptions(options ...GenericErrorOption) Error {
	return NewGenericError(append([]GenericErrorOption{
		WithCode(UnauthorizedErrorCode),
		WithPublicErrorMessage("unauthorized"),
		WithStatusCode(http.StatusUnauthorized),
	}, options...)...)
}
