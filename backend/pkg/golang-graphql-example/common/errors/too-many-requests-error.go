package errors

import (
	"net/http"
)

const TooManyRequestsErrorCode = "TOO_MANY_REQUESTS"

func NewTooManyRequestsError(msg string, options ...GenericErrorOption) Error {
	return NewTooManyRequestsErrorWithOptions(append([]GenericErrorOption{WithErrorMessage(msg)}, options...)...)
}

func NewTooManyRequestsErrorWithError(err error, options ...GenericErrorOption) Error {
	return NewTooManyRequestsErrorWithOptions(append([]GenericErrorOption{WithError(err)}, options...)...)
}

func NewTooManyRequestsErrorWithOptions(options ...GenericErrorOption) Error {
	return NewGenericError(append([]GenericErrorOption{
		WithCode(TooManyRequestsErrorCode),
		WithPublicErrorMessage("too many requests"),
		WithStatusCode(http.StatusTooManyRequests),
	}, options...)...)
}
