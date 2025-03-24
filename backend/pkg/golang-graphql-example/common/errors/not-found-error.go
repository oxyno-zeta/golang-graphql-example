package errors

import (
	"net/http"
)

const NotFoundErrorCode = "NOT_FOUND"

func NewNotFoundError(msg string, options ...GenericErrorOption) Error {
	return NewNotFoundErrorWithOptions(
		append([]GenericErrorOption{WithErrorMessage(msg)}, options...)...)
}

func NewNotFoundErrorWithError(err error, options ...GenericErrorOption) Error {
	return NewNotFoundErrorWithOptions(append([]GenericErrorOption{WithError(err)}, options...)...)
}

func NewNotFoundErrorWithOptions(options ...GenericErrorOption) Error {
	return NewGenericError(append([]GenericErrorOption{
		WithCode(NotFoundErrorCode),
		WithPublicErrorMessage("not found"),
		WithStatusCode(http.StatusNotFound),
	}, options...)...)
}
