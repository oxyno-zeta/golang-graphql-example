package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

func NewInvalidInputError(msg string) Error {
	return NewInvalidInputErrorWithExtensions(msg, nil)
}

func NewInvalidInputErrorWithPublicMessage(msg, pubMsg string) Error {
	return NewInvalidInputErrorWithExtensionsPublicErrorAndError(errors.New(msg), errors.New(pubMsg), nil)
}

func NewInvalidInputErrorWithError(err error) Error {
	return NewInvalidInputErrorWithExtensionsAndError(err, nil)
}

func NewInvalidInputErrorWithErrorAndPublicMessage(err error, pubMsg string) Error {
	return NewInvalidInputErrorWithExtensionsPublicErrorAndError(err, errors.New(pubMsg), nil)
}

func NewInvalidInputErrorWithExtensions(msg string, customExtensions map[string]interface{}) Error {
	return NewInvalidInputErrorWithExtensionsAndError(errors.New(msg), customExtensions)
}

func NewInvalidInputErrorWithExtensionsAndError(err error, customExtensions map[string]interface{}) Error {
	return NewInvalidInputErrorWithExtensionsPublicErrorAndError(err, nil, customExtensions)
}

func NewInvalidInputErrorWithExtensionsPublicErrorAndError(err, publicError error, customExtensions map[string]interface{}) Error {
	// Check if custom extensions exists
	if customExtensions == nil {
		customExtensions = map[string]interface{}{}
	}
	// Add code in custom extensions
	customExtensions["code"] = "INVALID_INPUT"

	pubErr := errors.New("invalid input")
	// Check if public error is set
	if publicError != nil {
		pubErr = publicError
	}

	// Return new error
	return &GenericError{
		err:         errors.WithStack(err),
		ext:         customExtensions,
		publicError: pubErr,
		statusCode:  http.StatusBadRequest,
	}
}
