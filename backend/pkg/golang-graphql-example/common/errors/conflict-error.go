package errors

import "net/http"

const ConflictErrorCode = "CONFLICT"

func NewConflictError(msg string, options ...GenericErrorOption) Error {
	return NewConflictErrorWithOptions(
		append([]GenericErrorOption{WithErrorMessage(msg)}, options...)...)
}

func NewConflictErrorWithError(err error, options ...GenericErrorOption) Error {
	return NewConflictErrorWithOptions(append([]GenericErrorOption{WithError(err)}, options...)...)
}

func NewConflictErrorWithOptions(options ...GenericErrorOption) Error {
	return NewGenericError(append([]GenericErrorOption{
		WithCode(ConflictErrorCode),
		WithPublicErrorMessage("conflict"),
		WithStatusCode(http.StatusConflict),
	}, options...)...)
}
