package errors

import (
	"net/http"
)

const ForbiddenErrorCode = "FORBIDDEN"

func NewForbiddenError(msg string, options ...GenericErrorOption) Error {
	return NewForbiddenErrorWithOptions(append([]GenericErrorOption{WithErrorMessage(msg)}, options...)...)
}

func NewForbiddenErrorWithError(err error, options ...GenericErrorOption) Error {
	return NewForbiddenErrorWithOptions(append([]GenericErrorOption{WithError(err)}, options...)...)
}

func NewForbiddenErrorWithOptions(options ...GenericErrorOption) Error {
	return NewGenericError(append([]GenericErrorOption{
		WithCode(ForbiddenErrorCode),
		WithPublicErrorMessage("forbidden"),
		WithStatusCode(http.StatusForbidden),
	}, options...)...)
}
