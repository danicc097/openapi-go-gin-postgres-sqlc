// Code adapted from:
// https://github.com/MarioCarrion/todo-api-microservice-example

package internal

import (
	"errors"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
)

// Error represents an error that could be wrapping another error,
// It includes a code for determining what triggered the error.
type Error struct {
	orig error
	msg  string
	code models.ErrorCode
	loc  []string
}

// WrapErrorf returns a wrapped error.
func WrapErrorf(orig error, code models.ErrorCode, format string, a ...any) error {
	return &Error{
		code: code,
		orig: orig,
		msg:  fmt.Sprintf(format, a...),
	}
}

// WrapErrorWithLocf accumulates a given path in loc.
func WrapErrorWithLocf(orig error, code models.ErrorCode, loc []string, format string, a ...interface{}) error {
	var previousCode models.ErrorCode
	var previousLoc []string

	var ierr *Error
	if errors.As(orig, &ierr) {
		previousLoc = ierr.loc // accumulate
		previousCode = ierr.code
	}

	if previousCode == "" {
		previousCode = models.ErrorCodeUnknown
	}

	if code == "" {
		code = previousCode
	}

	return &Error{
		orig: orig,
		code: code,
		msg:  fmt.Sprintf(format, a...),
		loc:  append(loc, previousLoc...),
	}
}

// NewErrorWithLocf instantiates a new error with error location.
func NewErrorWithLocf(code models.ErrorCode, loc []string, format string, a ...any) error {
	return WrapErrorWithLocf(nil, code, loc, format, a...)
}

// NewErrorf instantiates a new error.
func NewErrorf(code models.ErrorCode, format string, a ...any) error {
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
func (e *Error) Code() models.ErrorCode {
	return e.code
}

// Loc returns the accumulated error location.
func (e *Error) Loc() []string {
	return e.loc
}

// Cause returns the root error cause in the chain.
func (e *Error) Cause() error {
	var err error
	err = e
	for {
		_err := errors.Unwrap(err)
		if _err == nil {
			return err
		}
		err = _err
	}
}
