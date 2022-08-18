package errors

import (
	"net/http"

	"emperror.dev/errors"
)

const TooManyRequestsErrorCode = "TOO_MANY_REQUESTS"

func NewTooManyRequestsError(msg string) Error {
	return NewTooManyRequestsErrorWithExtensions(msg, nil)
}

func NewTooManyRequestsErrorWithPublicMessage(msg, pubMsg string) Error {
	return NewTooManyRequestsErrorWithExtensionsPublicErrorAndError(errors.New(msg), errors.New(pubMsg), nil)
}

func NewTooManyRequestsErrorWithError(err error) Error {
	return NewTooManyRequestsErrorWithExtensionsAndError(err, nil)
}

func NewTooManyRequestsErrorWithErrorAndPublicMessage(err error, pubMsg string) Error {
	return NewTooManyRequestsErrorWithExtensionsPublicErrorAndError(err, errors.New(pubMsg), nil)
}

func NewTooManyRequestsErrorWithExtensions(msg string, customExtensions map[string]interface{}) Error {
	return NewTooManyRequestsErrorWithExtensionsAndError(errors.New(msg), customExtensions)
}

func NewTooManyRequestsErrorWithExtensionsAndError(err error, customExtensions map[string]interface{}) Error {
	return NewTooManyRequestsErrorWithExtensionsPublicErrorAndError(err, nil, customExtensions)
}

func NewTooManyRequestsErrorWithExtensionsPublicErrorAndError(err, publicError error, customExtensions map[string]interface{}) Error {
	// Check if custom extensions exists
	if customExtensions == nil {
		customExtensions = map[string]interface{}{}
	}
	// Add code in custom extensions
	customExtensions["code"] = TooManyRequestsErrorCode

	pubErr := errors.New("too many requests")
	// Check if public error is set
	if publicError != nil {
		pubErr = publicError
	}

	// Return new error
	return &GenericError{
		err:         errors.WithStack(err),
		ext:         customExtensions,
		publicError: pubErr,
		statusCode:  http.StatusTooManyRequests,
		code:        TooManyRequestsErrorCode,
	}
}
