package errors

import (
	"net/http"

	"emperror.dev/errors"
)

const ForbiddenErrorCode = "FORBIDDEN"

func NewForbiddenError(msg string) Error {
	return NewForbiddenErrorWithExtensionsPublicErrorAndError(errors.New(msg), nil, nil)
}

func NewForbiddenErrorWithPublicMessage(msg, pubMsg string) Error {
	return NewForbiddenErrorWithExtensionsPublicErrorAndError(errors.New(msg), errors.New(pubMsg), nil)
}

func NewForbiddenErrorWithError(err error) Error {
	return NewForbiddenErrorWithExtensionsPublicErrorAndError(err, nil, nil)
}

func NewForbiddenErrorWithErrorAndPublicMessage(err error, pubMsg string) Error {
	return NewForbiddenErrorWithExtensionsPublicErrorAndError(err, errors.New(pubMsg), nil)
}

func NewForbiddenErrorWithExtensions(msg string, customExtensions map[string]interface{}) Error {
	return NewForbiddenErrorWithExtensionsPublicErrorAndError(errors.New(msg), nil, customExtensions)
}

func NewForbiddenErrorWithExtensionsAndError(err error, customExtensions map[string]interface{}) Error {
	return NewForbiddenErrorWithExtensionsPublicErrorAndError(err, nil, customExtensions)
}

func NewForbiddenErrorWithExtensionsPublicErrorAndError(err, publicError error, customExtensions map[string]interface{}) Error {
	// Check if custom extensions exists
	if customExtensions == nil {
		customExtensions = map[string]interface{}{}
	}
	// Add code in custom extensions
	customExtensions["code"] = ForbiddenErrorCode

	pubErr := errors.New("forbidden")
	// Check if public error is set
	if publicError != nil {
		pubErr = publicError
	}

	// Return new error
	return &GenericError{
		err:         errors.WithStack(err),
		ext:         customExtensions,
		publicError: pubErr,
		statusCode:  http.StatusForbidden,
		code:        ForbiddenErrorCode,
	}
}
