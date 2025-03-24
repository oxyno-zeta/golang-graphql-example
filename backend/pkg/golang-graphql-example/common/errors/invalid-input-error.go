package errors

import (
	"net/http"
)

const InvalidInputErrorCode = "INVALID_INPUT"

func NewInvalidInputError(msg string, options ...GenericErrorOption) Error {
	return NewInvalidInputErrorWithOptions(
		append([]GenericErrorOption{WithErrorMessage(msg)}, options...)...)
}

func NewInvalidInputErrorWithError(err error, options ...GenericErrorOption) Error {
	return NewInvalidInputErrorWithOptions(
		append([]GenericErrorOption{WithError(err)}, options...)...)
}

func NewInvalidInputErrorWithOptions(options ...GenericErrorOption) Error {
	return NewGenericError(append([]GenericErrorOption{
		WithCode(InvalidInputErrorCode),
		WithPublicErrorMessage("invalid input"),
		WithStatusCode(http.StatusBadRequest),
	}, options...)...)
}
