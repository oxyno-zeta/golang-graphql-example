package errors

import (
	"net/http"
)

const LockedErrorCode = "LOCKED"

func NewLockedError(msg string, options ...GenericErrorOption) Error {
	return NewLockedErrorWithOptions(append([]GenericErrorOption{WithErrorMessage(msg)}, options...)...)
}

func NewLockedErrorWithError(err error, options ...GenericErrorOption) Error {
	return NewLockedErrorWithOptions(append([]GenericErrorOption{WithError(err)}, options...)...)
}

func NewLockedErrorWithOptions(options ...GenericErrorOption) Error {
	return NewGenericError(append([]GenericErrorOption{
		WithCode(LockedErrorCode),
		WithPublicErrorMessage("locked"),
		WithStatusCode(http.StatusLocked),
	}, options...)...)
}
