package errors

import (
	"net/http"

	"emperror.dev/errors"
)

const LockedErrorCode = "LOCKED"

func NewLockedError(msg string) Error {
	return NewLockedErrorWithExtensionsPublicErrorAndError(errors.New(msg), nil, nil)
}

func NewLockedErrorWithPublicMessage(msg, pubMsg string) Error {
	return NewLockedErrorWithExtensionsPublicErrorAndError(errors.New(msg), errors.New(pubMsg), nil)
}

func NewLockedErrorWithError(err error) Error {
	return NewLockedErrorWithExtensionsPublicErrorAndError(err, nil, nil)
}

func NewLockedErrorWithErrorAndPublicMessage(err error, pubMsg string) Error {
	return NewLockedErrorWithExtensionsPublicErrorAndError(err, errors.New(pubMsg), nil)
}

func NewLockedErrorWithExtensions(msg string, customExtensions map[string]interface{}) Error {
	return NewLockedErrorWithExtensionsPublicErrorAndError(errors.New(msg), nil, customExtensions)
}

func NewLockedErrorWithExtensionsAndError(err error, customExtensions map[string]interface{}) Error {
	return NewLockedErrorWithExtensionsPublicErrorAndError(err, nil, customExtensions)
}

func NewLockedErrorWithExtensionsPublicErrorAndError(err, publicError error, customExtensions map[string]interface{}) Error {
	// Check if custom extensions exists
	if customExtensions == nil {
		customExtensions = map[string]interface{}{}
	}
	// Add code in custom extensions
	customExtensions["code"] = LockedErrorCode

	pubErr := errors.New("locked")
	// Check if public error is set
	if publicError != nil {
		pubErr = publicError
	}

	// Return new error
	return &GenericError{
		err:         errors.WithStack(err),
		ext:         customExtensions,
		publicError: pubErr,
		statusCode:  http.StatusLocked,
		code:        LockedErrorCode,
	}
}
