package errors

import (
	"net/http"

	"emperror.dev/errors"
)

const UnauthorizedErrorCode = "UNAUTHORIZED"

func NewUnauthorizedError(msg string) Error {
	return NewUnauthorizedErrorWithExtensions(msg, nil)
}

func NewUnauthorizedErrorWithPublicMessage(msg, pubMsg string) Error {
	return NewUnauthorizedErrorWithExtensionsPublicErrorAndError(errors.New(msg), errors.New(pubMsg), nil)
}

func NewUnauthorizedErrorWithError(err error) Error {
	return NewUnauthorizedErrorWithExtensionsAndError(err, nil)
}

func NewUnauthorizedErrorWithErrorAndPublicMessage(err error, pubMsg string) Error {
	return NewUnauthorizedErrorWithExtensionsPublicErrorAndError(err, errors.New(pubMsg), nil)
}

func NewUnauthorizedErrorWithExtensions(msg string, customExtensions map[string]interface{}) Error {
	return NewUnauthorizedErrorWithExtensionsAndError(errors.New(msg), customExtensions)
}

func NewUnauthorizedErrorWithExtensionsAndError(err error, customExtensions map[string]interface{}) Error {
	return NewUnauthorizedErrorWithExtensionsPublicErrorAndError(err, nil, customExtensions)
}

func NewUnauthorizedErrorWithExtensionsPublicErrorAndError(err, publicError error, customExtensions map[string]interface{}) Error {
	// Check if custom extensions exists
	if customExtensions == nil {
		customExtensions = map[string]interface{}{}
	}
	// Add code in custom extensions
	customExtensions["code"] = UnauthorizedErrorCode

	pubErr := errors.New("unauthorized")
	// Check if public error is set
	if publicError != nil {
		pubErr = publicError
	}

	// Return new error
	return &GenericError{
		err:         errors.WithStack(err),
		ext:         customExtensions,
		publicError: pubErr,
		statusCode:  http.StatusUnauthorized,
		code:        UnauthorizedErrorCode,
	}
}
