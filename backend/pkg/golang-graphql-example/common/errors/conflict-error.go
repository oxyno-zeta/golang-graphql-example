package errors

import (
	"net/http"

	"emperror.dev/errors"
)

const ConflictErrorCode = "CONFLICT"

func NewConflictError(msg string) Error {
	return NewConflictErrorWithExtensionsPublicErrorAndError(errors.New(msg), nil, nil)
}

func NewConflictErrorWithPublicMessage(msg, pubMsg string) Error {
	return NewConflictErrorWithExtensionsPublicErrorAndError(errors.New(msg), errors.New(pubMsg), nil)
}

func NewConflictErrorWithError(err error) Error {
	return NewConflictErrorWithExtensionsPublicErrorAndError(err, nil, nil)
}

func NewConflictErrorWithErrorAndPublicMessage(err error, pubMsg string) Error {
	return NewConflictErrorWithExtensionsPublicErrorAndError(err, errors.New(pubMsg), nil)
}

func NewConflictErrorWithExtensions(msg string, customExtensions map[string]interface{}) Error {
	return NewConflictErrorWithExtensionsPublicErrorAndError(errors.New(msg), nil, customExtensions)
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
	customExtensions["code"] = ConflictErrorCode

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
		code:        ConflictErrorCode,
	}
}
