package errors

import "net/http"

const InternalServerErrorCode = "INTERNAL_SERVER_ERROR"

func NewInternalServerError(msg string, options ...GenericErrorOption) Error {
	return NewInternalServerErrorWithOptions(append([]GenericErrorOption{WithErrorMessage(msg)}, options...)...)
}

func NewInternalServerErrorWithError(err error, options ...GenericErrorOption) Error {
	return NewGenericError(append([]GenericErrorOption{WithError(err)}, options...)...)
}

func NewInternalServerErrorWithOptions(options ...GenericErrorOption) Error {
	return NewGenericError(append([]GenericErrorOption{
		WithCode(InternalServerErrorCode),
		WithPublicErrorMessage("internal server error"),
		WithStatusCode(http.StatusInternalServerError),
	}, options...)...)
}
