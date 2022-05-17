package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

func NewConflictError(msg string) Error {
	return NewConflictErrorWithExtensions(msg, nil)
}

func NewConflictErrorWithPublicMessage(msg, pubMsg string) Error {
	return NewConflictErrorWithExtensionsPublicErrorAndError(errors.New(msg), errors.New(pubMsg), nil)
}

func NewConflictErrorWithError(err error) Error {
	return NewConflictErrorWithExtensionsAndError(err, nil)
}

func NewConflictErrorWithErrorAndPublicMessage(err error, pubMsg string) Error {
	return NewConflictErrorWithExtensionsPublicErrorAndError(err, errors.New(pubMsg), nil)
}

func NewConflictErrorWithExtensions(msg string, customExtensions map[string]interface{}) Error {
	return NewConflictErrorWithExtensionsAndError(errors.New(msg), customExtensions)
}

func NewConflictErrorWithExtensionsAndError(err error, customExtensions map[string]interface{}) Error {
	return NewConflictErrorWithExtensionsPublicErrorAndError(err, nil, customExtensions)
}

func NewConflictErrorWithExtensionsPublicErrorAndError(err, publicError error, customExtensions map[string]interface{}) Error {
	// Check if custom extensions exists
	if customExtensions == nil {
		customExtensions = map[string]interface{}{}
	}
	// Add code in custom extensions
	customExtensions["code"] = "CONFLICT"

	pubErr := errors.New("conflict")
	// Check if public error is set
	if publicError != nil {
		pubErr = publicError
	}

	// Return new error
	return &GenericError{
		err:         errors.WithStack(err),
		ext:         customExtensions,
		publicError: pubErr,
		statusCode:  http.StatusConflict,
	}
}
