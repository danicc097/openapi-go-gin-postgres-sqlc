// Code adapted from:
// https://github.com/MarioCarrion/todo-api-microservice-example

package internal

import (
	"fmt"
)

// Error represents an error that could be wrapping another error,
// It includes a code for determining what triggered the error.
type Error struct {
	orig error
	msg  string
	code ErrorCode
}

// ErrorCode defines supported error codes.
type ErrorCode uint

const (
	ErrorCodeUnknown ErrorCode = iota
	ErrorCodeNotFound
	ErrorCodeInvalidArgument
	ErrorCodeAlreadyExists
	ErrorCodeUnauthorized
	ErrorCodeUnauthenticated
	ErrorCodeValidation
	ErrorCodeResponseValidation

	ErrorCodeInvalidRole
	ErrorCodeInvalidScope

	ErrorCodeInvalidUUID
)

type HTTPErrorType string

const (
	HTTPErrorTypeResponseValidation HTTPErrorType = "response-validation-error"
)

// WrapErrorf returns a wrapped error.
func WrapErrorf(orig error, code ErrorCode, format string, a ...any) error {
	return &Error{
		code: code,
		orig: orig,
		msg:  fmt.Sprintf(format, a...),
	}
}

// NewErrorf instantiates a new error.
func NewErrorf(code ErrorCode, format string, a ...any) error {
	return WrapErrorf(nil, code, format, a...)
}

// Error returns the message, when wrapping errors the wrapped error is returned.
func (e *Error) Error() string {
	if e.orig != nil {
		return fmt.Sprintf("%s: %v", e.msg, e.orig)
	}

	return e.msg
}

// Unwrap returns the wrapped error, if any.
func (e *Error) Unwrap() error {
	return e.orig
}

// Code returns the code representing this error.
func (e *Error) Code() ErrorCode {
	return e.code
}

// Cause returns the root error cause in the chain.
func (e *Error) Cause() error {
	var err error
	root := e
	for {
		if err = root.Unwrap(); err == nil {
			return root
		}

		r, ok := err.(*Error)
		if !ok {
			return err
		}

		root = r
	}
}
