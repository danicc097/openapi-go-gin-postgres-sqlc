package internal_test

import (
	"errors"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/stretchr/testify/assert"
)

func TestErrorCause(t *testing.T) {
	var ierr *internal.Error

	err := internal.NewErrorf(internal.ErrorCodeUnknown, "root")
	err = internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "wrapped 1")
	err = internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "wrapped 2")
	errors.As(err, &ierr)

	assert.Equal(t, "root", ierr.Cause().Error() )
}
