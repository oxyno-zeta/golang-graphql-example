package errors

import (
	"net/http"

	"emperror.dev/errors"
)

const NotFoundErrorCode = "NOT_FOUND"

func NewNotFoundError(msg string) Error {
	return NewNotFoundErrorWithExtensionsPublicErrorAndError(errors.New(msg), nil, nil)
}

func NewNotFoundErrorWithPublicMessage(msg, pubMsg string) Error {
	return NewNotFoundErrorWithExtensionsPublicErrorAndError(errors.New(msg), errors.New(pubMsg), nil)
}

func NewNotFoundErrorWithError(err error) Error {
	return NewNotFoundErrorWithExtensionsPublicErrorAndError(err, nil, nil)
}

func NewNotFoundErrorWithErrorAndPublicMessage(err error, pubMsg string) Error {
	return NewNotFoundErrorWithExtensionsPublicErrorAndError(err, errors.New(pubMsg), nil)
}

func NewNotFoundErrorWithExtensions(msg string, customExtensions map[string]interface{}) Error {
	return NewNotFoundErrorWithExtensionsPublicErrorAndError(errors.New(msg), nil, customExtensions)
}

func NewNotFoundErrorWithExtensionsAndError(err error, customExtensions map[string]interface{}) Error {
	return NewNotFoundErrorWithExtensionsPublicErrorAndError(err, nil, customExtensions)
}

func NewNotFoundErrorWithExtensionsPublicErrorAndError(err, publicError error, customExtensions map[string]interface{}) Error {
	// Check if custom extensions exists
	if customExtensions == nil {
		customExtensions = map[string]interface{}{}
	}
	// Add code in custom extensions
	customExtensions["code"] = NotFoundErrorCode

	pubErr := errors.New("not found")
	// Check if public error is set
	if publicError != nil {
		pubErr = publicError
	}

	// Return new error
	return &GenericError{
		err:         errors.WithStack(err),
		ext:         customExtensions,
		publicError: pubErr,
		statusCode:  http.StatusNotFound,
		code:        NotFoundErrorCode,
	}
}
