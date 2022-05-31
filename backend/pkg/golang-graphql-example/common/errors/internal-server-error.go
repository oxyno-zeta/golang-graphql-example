package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

const InternalServerErrorCode = "INTERNAL_SERVER_ERROR"

func NewInternalServerError(msg string) Error {
	return NewInternalServerErrorWithExtensions(msg, nil)
}

func NewInternalServerErrorWithPublicMessage(msg, pubMsg string) Error {
	return NewInternalServerErrorWithExtensionsPublicErrorAndError(errors.New(msg), errors.New(pubMsg), nil)
}

func NewInternalServerErrorWithError(err error) Error {
	return NewInternalServerErrorWithExtensionsAndError(err, nil)
}

func NewInternalServerErrorWithErrorAndPublicMessage(err error, pubMsg string) Error {
	return NewInternalServerErrorWithExtensionsPublicErrorAndError(err, errors.New(pubMsg), nil)
}

func NewInternalServerErrorWithExtensions(msg string, customExtensions map[string]interface{}) Error {
	return NewInternalServerErrorWithExtensionsAndError(errors.New(msg), customExtensions)
}

func NewInternalServerErrorWithExtensionsAndError(err error, customExtensions map[string]interface{}) Error {
	return NewInternalServerErrorWithExtensionsPublicErrorAndError(err, nil, customExtensions)
}

func NewInternalServerErrorWithExtensionsPublicErrorAndError(err, publicError error, customExtensions map[string]interface{}) Error {
	// Check if custom extensions exists
	if customExtensions == nil {
		customExtensions = map[string]interface{}{}
	}
	// Add code in custom extensions
	customExtensions["code"] = InternalServerErrorCode

	pubErr := errors.New("internal server error")
	// Check if public error is set
	if publicError != nil {
		pubErr = publicError
	}

	// Return new error
	return &GenericError{
		err:         errors.WithStack(err),
		ext:         customExtensions,
		publicError: pubErr,
		statusCode:  http.StatusInternalServerError,
		code:        InternalServerErrorCode,
	}
}
